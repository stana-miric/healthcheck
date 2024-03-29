// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: healthcheck/types/packet.proto

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

type HealthcheckPacketData struct {
	// Types that are valid to be assigned to Packet:
	//	*HealthcheckPacketData_HealtcheckUpdate
	Packet isHealthcheckPacketData_Packet `protobuf_oneof:"packet"`
}

func (m *HealthcheckPacketData) Reset()         { *m = HealthcheckPacketData{} }
func (m *HealthcheckPacketData) String() string { return proto.CompactTextString(m) }
func (*HealthcheckPacketData) ProtoMessage()    {}
func (*HealthcheckPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_802e33ee4c416c7d, []int{0}
}
func (m *HealthcheckPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HealthcheckPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HealthcheckPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HealthcheckPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthcheckPacketData.Merge(m, src)
}
func (m *HealthcheckPacketData) XXX_Size() int {
	return m.Size()
}
func (m *HealthcheckPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthcheckPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_HealthcheckPacketData proto.InternalMessageInfo

type isHealthcheckPacketData_Packet interface {
	isHealthcheckPacketData_Packet()
	MarshalTo([]byte) (int, error)
	Size() int
}

type HealthcheckPacketData_HealtcheckUpdate struct {
	HealtcheckUpdate *HealthcheckUpdateData `protobuf:"bytes,1,opt,name=healtcheckUpdate,proto3,oneof" json:"healtcheckUpdate,omitempty"`
}

func (*HealthcheckPacketData_HealtcheckUpdate) isHealthcheckPacketData_Packet() {}

func (m *HealthcheckPacketData) GetPacket() isHealthcheckPacketData_Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

func (m *HealthcheckPacketData) GetHealtcheckUpdate() *HealthcheckUpdateData {
	if x, ok := m.GetPacket().(*HealthcheckPacketData_HealtcheckUpdate); ok {
		return x.HealtcheckUpdate
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthcheckPacketData) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthcheckPacketData_HealtcheckUpdate)(nil),
	}
}

type HealthcheckUpdateData struct {
	Block     uint64 `protobuf:"varint,1,opt,name=block,proto3" json:"block,omitempty"`
	Timestamp uint64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *HealthcheckUpdateData) Reset()         { *m = HealthcheckUpdateData{} }
func (m *HealthcheckUpdateData) String() string { return proto.CompactTextString(m) }
func (*HealthcheckUpdateData) ProtoMessage()    {}
func (*HealthcheckUpdateData) Descriptor() ([]byte, []int) {
	return fileDescriptor_802e33ee4c416c7d, []int{1}
}
func (m *HealthcheckUpdateData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HealthcheckUpdateData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HealthcheckUpdateData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HealthcheckUpdateData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthcheckUpdateData.Merge(m, src)
}
func (m *HealthcheckUpdateData) XXX_Size() int {
	return m.Size()
}
func (m *HealthcheckUpdateData) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthcheckUpdateData.DiscardUnknown(m)
}

var xxx_messageInfo_HealthcheckUpdateData proto.InternalMessageInfo

func (m *HealthcheckUpdateData) GetBlock() uint64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func (m *HealthcheckUpdateData) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*HealthcheckPacketData)(nil), "healthcheck.types.HealthcheckPacketData")
	proto.RegisterType((*HealthcheckUpdateData)(nil), "healthcheck.types.HealthcheckUpdateData")
}

func init() { proto.RegisterFile("healthcheck/types/packet.proto", fileDescriptor_802e33ee4c416c7d) }

var fileDescriptor_802e33ee4c416c7d = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcb, 0x48, 0x4d, 0xcc,
	0x29, 0xc9, 0x48, 0xce, 0x48, 0x4d, 0xce, 0xd6, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2f, 0x48,
	0x4c, 0xce, 0x4e, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x44, 0x92, 0xd7, 0x03,
	0xcb, 0x2b, 0x55, 0x72, 0x89, 0x7a, 0x20, 0x04, 0x03, 0xc0, 0xaa, 0x5d, 0x12, 0x4b, 0x12, 0x85,
	0xc2, 0xb8, 0x04, 0xc0, 0xaa, 0xc1, 0xe2, 0xa1, 0x05, 0x29, 0x89, 0x25, 0xa9, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0xdc, 0x46, 0x1a, 0x7a, 0x18, 0xc6, 0xe8, 0x21, 0x99, 0x01, 0x51, 0x0b, 0x32, 0xc3,
	0x83, 0x21, 0x08, 0xc3, 0x0c, 0x27, 0x0e, 0x2e, 0x36, 0x88, 0x9b, 0x94, 0xbc, 0x51, 0xac, 0x46,
	0x68, 0x13, 0x12, 0xe1, 0x62, 0x4d, 0xca, 0xc9, 0x4f, 0xce, 0x06, 0xdb, 0xc7, 0x12, 0x04, 0xe1,
	0x08, 0xc9, 0x70, 0x71, 0x96, 0x64, 0xe6, 0xa6, 0x16, 0x97, 0x24, 0xe6, 0x16, 0x48, 0x30, 0x81,
	0x65, 0x10, 0x02, 0x4e, 0xba, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91,
	0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x25,
	0x8c, 0x1c, 0x28, 0x15, 0x90, 0x60, 0x49, 0x62, 0x03, 0x07, 0x88, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x95, 0xfa, 0x87, 0x87, 0x32, 0x01, 0x00, 0x00,
}

func (m *HealthcheckPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HealthcheckPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HealthcheckPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Packet != nil {
		{
			size := m.Packet.Size()
			i -= size
			if _, err := m.Packet.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *HealthcheckPacketData_HealtcheckUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HealthcheckPacketData_HealtcheckUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.HealtcheckUpdate != nil {
		{
			size, err := m.HealtcheckUpdate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPacket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *HealthcheckUpdateData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HealthcheckUpdateData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HealthcheckUpdateData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Timestamp != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x10
	}
	if m.Block != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPacket(dAtA []byte, offset int, v uint64) int {
	offset -= sovPacket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HealthcheckPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Packet != nil {
		n += m.Packet.Size()
	}
	return n
}

func (m *HealthcheckPacketData_HealtcheckUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HealtcheckUpdate != nil {
		l = m.HealtcheckUpdate.Size()
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}
func (m *HealthcheckUpdateData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Block != 0 {
		n += 1 + sovPacket(uint64(m.Block))
	}
	if m.Timestamp != 0 {
		n += 1 + sovPacket(uint64(m.Timestamp))
	}
	return n
}

func sovPacket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPacket(x uint64) (n int) {
	return sovPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HealthcheckPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: HealthcheckPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HealthcheckPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HealtcheckUpdate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &HealthcheckUpdateData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Packet = &HealthcheckPacketData_HealtcheckUpdate{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *HealthcheckUpdateData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: HealthcheckUpdateData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HealthcheckUpdateData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func skipPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
				return 0, ErrInvalidLengthPacket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPacket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPacket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPacket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPacket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPacket = fmt.Errorf("proto: unexpected end of group")
)
