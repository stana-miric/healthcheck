#!/bin/bash
set -eu

TOTAL_COINS=100000000000stake
STAKE_COINS=100000000stake

# ---------------------------------------- Registry chain ----------------------------------------

REGISTRY_BINARY=registryd
REGISTRY_HOME=$HOME/.registry
REGISTRY_CHAIN_ID=registry
REGISTRY_MONIKER=registry
VALIDATOR=validator
RELAYER=relayer
NODE_IP="localhost"
REGISTRY_RPC_LADDR="$NODE_IP:26658"
REGISTRY_GRPC_ADDR="$NODE_IP:9091"

# Clean start
killall $REGISTRY_BINARY &> /dev/null || true
rm -rf $REGISTRY_HOME

# Init chain
./$REGISTRY_BINARY init $REGISTRY_MONIKER --home $REGISTRY_HOME --chain-id $REGISTRY_CHAIN_ID
sleep 1

# Create account keypairs
./$REGISTRY_BINARY keys add $VALIDATOR --home $REGISTRY_HOME --keyring-backend test --output json > $REGISTRY_HOME/keypair.json 2>&1
sleep 1

./$REGISTRY_BINARY keys add $RELAYER --home $REGISTRY_HOME --keyring-backend test --output json > $REGISTRY_HOME/keypair_relayer.json 2>&1
sleep 1

# Add stake to users
./$REGISTRY_BINARY add-genesis-account $(jq -r .address $REGISTRY_HOME/keypair.json) $TOTAL_COINS --home $REGISTRY_HOME --keyring-backend test
sleep 1

./$REGISTRY_BINARY add-genesis-account $(jq -r .address $REGISTRY_HOME/keypair_relayer.json) $TOTAL_COINS --home $REGISTRY_HOME --keyring-backend test
sleep 1

# Stake 1/1000 user's coins
./$REGISTRY_BINARY gentx $VALIDATOR $STAKE_COINS --chain-id $REGISTRY_CHAIN_ID --home $REGISTRY_HOME --keyring-backend test --moniker $VALIDATOR
sleep 1

./$REGISTRY_BINARY collect-gentxs --home $REGISTRY_HOME --gentx-dir $REGISTRY_HOME/config/gentx/
sleep 1

./$REGISTRY_BINARY config keyring-backend test
./$REGISTRY_BINARY config node tcp://$REGISTRY_RPC_LADDR

# Start the chain
./$REGISTRY_BINARY start \
	--home $REGISTRY_HOME \
	--rpc.laddr tcp://$REGISTRY_RPC_LADDR \
	--grpc.address $REGISTRY_GRPC_ADDR \
	--address tcp://${NODE_IP}:26655 \
	--p2p.laddr tcp://${NODE_IP}:26656 \
	--grpc-web.enable=false &> $REGISTRY_HOME/logs &
sleep 10
echo $REGISTRY_MONIKER chain started!

# ---------------------------------------- Monitored chain ----------------------------------------

MONITORED_BINARY=monitored
MONITORED_HOME=$HOME/.monitored
MONITORED_CHAIN_ID=monitored
MONITORED_MONIKER=monitored
VALIDATOR=validator
NODE_IP="localhost"
MONITORED_RPC_LADDR="$NODE_IP:26648"
MONITORED_GRPC_ADDR="$NODE_IP:9081"

# Clean start
killall $MONITORED_BINARY &> /dev/null || true
rm -rf $MONITORED_HOME

# Init chain
./$MONITORED_BINARY init $MONITORED_MONIKER --home $MONITORED_HOME --chain-id $MONITORED_CHAIN_ID
sleep 1

# Create account keypairs
./$MONITORED_BINARY keys add $VALIDATOR --home $MONITORED_HOME --keyring-backend test --output json > $MONITORED_HOME/keypair.json 2>&1
sleep 1

./$MONITORED_BINARY keys add $RELAYER --home $MONITORED_HOME --keyring-backend test --output json > $MONITORED_HOME/keypair_relayer.json 2>&1
sleep 1

# Add stake to users
./$MONITORED_BINARY add-genesis-account $(jq -r .address $MONITORED_HOME/keypair.json) $TOTAL_COINS --home $MONITORED_HOME --keyring-backend test
sleep 1

./$MONITORED_BINARY add-genesis-account $(jq -r .address $MONITORED_HOME/keypair_relayer.json) $TOTAL_COINS --home $MONITORED_HOME --keyring-backend test
sleep 1

# Stake 1/1000 user's coins
./$MONITORED_BINARY gentx $VALIDATOR $STAKE_COINS --chain-id $MONITORED_CHAIN_ID --home $MONITORED_HOME --keyring-backend test --moniker $VALIDATOR
sleep 1

./$MONITORED_BINARY collect-gentxs --home $MONITORED_HOME --gentx-dir $MONITORED_HOME/config/gentx/
sleep 1

./$MONITORED_BINARY config keyring-backend test
./$MONITORED_BINARY config node tcp://$MONITORED_RPC_LADDR

# Start the chain
./$MONITORED_BINARY start \
	--home $MONITORED_HOME \
	--rpc.laddr tcp://$MONITORED_RPC_LADDR \
	--grpc.address $MONITORED_GRPC_ADDR \
	--address tcp://${NODE_IP}:26645 \
	--p2p.laddr tcp://${NODE_IP}:26646 \
	--grpc-web.enable=false &> $MONITORED_HOME/logs &
sleep 10
echo $MONITORED_MONIKER chain started!

# ---------------------------------------- Relayer ----------------------------------------

# Setup Hermes in packet relayer mode
killall hermes 2> /dev/null || true

tee ~/.hermes/config.toml<<EOF
[global]
log_level = "trace"

[mode]

[mode.clients]
enabled = true
refresh = true
misbehaviour = true

[mode.connections]
enabled = true

[mode.channels]
enabled = true

[mode.packets]
enabled = true
clear_interval = 5

[[chains]]
id = '$MONITORED_CHAIN_ID'
type = 'CosmosSdk'
ccv_consumer_chain = false
rpc_addr = 'http://${MONITORED_RPC_LADDR}'
grpc_addr = 'tcp://${MONITORED_GRPC_ADDR}'
event_source = { mode = 'push', url = 'ws://${MONITORED_RPC_LADDR}/websocket', batch_delay = '500ms' }
rpc_timeout = '10s'
account_prefix = 'cosmos'
key_name = 'relayer'
address_type = { derivation = 'cosmos' }
store_prefix = 'ibc'
default_gas = 100000
max_gas = 3000000
gas_price = { price = 0.0025, denom = 'stake' }
gas_multiplier = 1.1
max_msg_num = 30
max_tx_size = 2097152
clock_drift = '5s'
max_block_time = '30s'
trusting_period = '14days'

[[chains]]
id = '$REGISTRY_CHAIN_ID'
type = 'CosmosSdk'
ccv_consumer_chain = false
rpc_addr = 'http://${REGISTRY_RPC_LADDR}'
grpc_addr = 'tcp://${REGISTRY_GRPC_ADDR}'
event_source = { mode = 'push', url = 'ws://${REGISTRY_RPC_LADDR}/websocket', batch_delay = '500ms' }
rpc_timeout = '10s'
account_prefix = 'cosmos'
key_name = 'relayer'
address_type = { derivation = 'cosmos' }
store_prefix = 'ibc'
default_gas = 100000
max_gas = 3000000
gas_price = { price = 0.0025, denom = 'stake' }
gas_multiplier = 1.1
max_msg_num = 30
max_tx_size = 2097152
clock_drift = '5s'
max_block_time = '30s'
trusting_period = '14days'
EOF

# Delete all previous keys in relayer
hermes keys delete --chain $MONITORED_CHAIN_ID --all
hermes keys delete --chain $REGISTRY_CHAIN_ID --all

# Restore keys to hermes relayer
echo "$(jq -r .mnemonic $MONITORED_HOME/keypair_relayer.json)" | hermes keys add --chain $MONITORED_CHAIN_ID --mnemonic-file /dev/stdin
echo "$(jq -r .mnemonic $REGISTRY_HOME/keypair_relayer.json)" | hermes keys add --chain $REGISTRY_CHAIN_ID --mnemonic-file /dev/stdin
sleep 1

hermes create connection --a-chain $MONITORED_CHAIN_ID --b-chain $REGISTRY_CHAIN_ID
sleep 1

./$REGISTRY_BINARY tx healthcheck create-chain $MONITORED_CHAIN_ID connection-0 --from $VALIDATOR -y

hermes create channel --a-chain $MONITORED_CHAIN_ID --a-connection connection-0 --a-port monitored --b-port healthcheck --channel-version 1 --order ordered
sleep 1

hermes start &> ~/.hermes/logs &