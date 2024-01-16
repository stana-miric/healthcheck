// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: healthcheck/healthcheck/monitored_chains.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MonitoredChain struct {
	ChainId         string `protobuf:"bytes,1,opt,name=chainId,proto3" json:"chainId,omitempty"`
	ConnectionId    string `protobuf:"bytes,2,opt,name=connectionId,proto3" json:"connectionId,omitempty"`
	Creator         string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
	TimeoutInterval uint32 `protobuf:"varint,4,opt,name=timeoutInterval,proto3" json:"timeoutInterval,omitempty"`
	UpdateInterval  uint32 `protobuf:"varint,5,opt,name=updateInterval,proto3" json:"updateInterval,omitempty"`
}

func (m *MonitoredChain) Reset()         { *m = MonitoredChain{} }
func (m *MonitoredChain) String() string { return proto.CompactTextString(m) }
func (*MonitoredChain) ProtoMessage()    {}
func (*MonitoredChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d374112fcac9bc6, []int{0}
}
func (m *MonitoredChain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MonitoredChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MonitoredChain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MonitoredChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitoredChain.Merge(m, src)
}
func (m *MonitoredChain) XXX_Size() int {
	return m.Size()
}
func (m *MonitoredChain) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitoredChain.DiscardUnknown(m)
}

var xxx_messageInfo_MonitoredChain proto.InternalMessageInfo

func (m *MonitoredChain) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *MonitoredChain) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *MonitoredChain) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MonitoredChain) GetTimeoutInterval() uint32 {
	if m != nil {
		return m.TimeoutInterval
	}
	return 0
}

func (m *MonitoredChain) GetUpdateInterval() uint32 {
	if m != nil {
		return m.UpdateInterval
	}
	return 0
}

func init() {
	proto.RegisterType((*MonitoredChain)(nil), "healthcheck.healthcheck.MonitoredChain")
}

func init() {
	proto.RegisterFile("healthcheck/healthcheck/monitored_chains.proto", fileDescriptor_9d374112fcac9bc6)
}

var fileDescriptor_9d374112fcac9bc6 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xcb, 0x48, 0x4d, 0xcc,
	0x29, 0xc9, 0x48, 0xce, 0x48, 0x4d, 0xce, 0xd6, 0x47, 0x66, 0xe7, 0xe6, 0xe7, 0x65, 0x96, 0xe4,
	0x17, 0xa5, 0xa6, 0xc4, 0x27, 0x67, 0x24, 0x66, 0xe6, 0x15, 0xeb, 0x15, 0x14, 0xe5, 0x97, 0xe4,
	0x0b, 0x89, 0x23, 0xa9, 0x41, 0xd6, 0xab, 0xb4, 0x8b, 0x91, 0x8b, 0xcf, 0x17, 0xa6, 0xc7, 0x19,
	0xa4, 0x45, 0x48, 0x82, 0x8b, 0x1d, 0xac, 0xd7, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33,
	0x08, 0xc6, 0x15, 0x52, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0xcb, 0x4b, 0x4d, 0x2e, 0xc9, 0xcc, 0x07,
	0x49, 0x33, 0x81, 0xa5, 0x51, 0xc4, 0xc0, 0xba, 0x8b, 0x52, 0x13, 0x4b, 0xf2, 0x8b, 0x24, 0x98,
	0xa1, 0xba, 0x21, 0x5c, 0x21, 0x0d, 0x2e, 0xfe, 0x92, 0xcc, 0xdc, 0xd4, 0xfc, 0xd2, 0x12, 0xcf,
	0xbc, 0x92, 0xd4, 0xa2, 0xb2, 0xc4, 0x1c, 0x09, 0x16, 0x05, 0x46, 0x0d, 0xde, 0x20, 0x74, 0x61,
	0x21, 0x35, 0x2e, 0xbe, 0xd2, 0x82, 0x94, 0xc4, 0x92, 0x54, 0xb8, 0x42, 0x56, 0xb0, 0x42, 0x34,
	0x51, 0x27, 0xcb, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71,
	0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x92, 0x47, 0x0e,
	0x93, 0x0a, 0x94, 0x10, 0x2a, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x87, 0x8b, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x71, 0x46, 0xaa, 0x5c, 0x49, 0x01, 0x00, 0x00,
}

func (m *MonitoredChain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MonitoredChain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MonitoredChain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UpdateInterval != 0 {
		i = encodeVarintMonitoredChains(dAtA, i, uint64(m.UpdateInterval))
		i--
		dAtA[i] = 0x28
	}
	if m.TimeoutInterval != 0 {
		i = encodeVarintMonitoredChains(dAtA, i, uint64(m.TimeoutInterval))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintMonitoredChains(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintMonitoredChains(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintMonitoredChains(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMonitoredChains(dAtA []byte, offset int, v uint64) int {
	offset -= sovMonitoredChains(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MonitoredChain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovMonitoredChains(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovMonitoredChains(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovMonitoredChains(uint64(l))
	}
	if m.TimeoutInterval != 0 {
		n += 1 + sovMonitoredChains(uint64(m.TimeoutInterval))
	}
	if m.UpdateInterval != 0 {
		n += 1 + sovMonitoredChains(uint64(m.UpdateInterval))
	}
	return n
}

func sovMonitoredChains(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMonitoredChains(x uint64) (n int) {
	return sovMonitoredChains(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MonitoredChain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMonitoredChains
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MonitoredChain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MonitoredChain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoredChains
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMonitoredChains
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMonitoredChains
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoredChains
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMonitoredChains
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMonitoredChains
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoredChains
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMonitoredChains
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMonitoredChains
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutInterval", wireType)
			}
			m.TimeoutInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoredChains
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutInterval |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateInterval", wireType)
			}
			m.UpdateInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoredChains
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdateInterval |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMonitoredChains(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMonitoredChains
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMonitoredChains(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMonitoredChains
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMonitoredChains
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMonitoredChains
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMonitoredChains
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMonitoredChains
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMonitoredChains
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMonitoredChains        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMonitoredChains          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMonitoredChains = fmt.Errorf("proto: unexpected end of group")
)
