// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ethermint/types/v1/bitcoin_indexer.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// BitcoinTxParseResult is the value parsed bitcoin transaction
// NOTE: This struct may be extended in the future
type BitcoinTxParseResult struct {
	// tx_id is the btc transaction id
	TxId string `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	// tx_type is the type of the transaction, eg. "brc20_transfer","transfer"
	TxType string `protobuf:"bytes,2,opt,name=tx_type,json=txType,proto3" json:"tx_type,omitempty"`
	// index is the index of the transaction in the block
	Index int64 `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	// from is l2 user address, by parse bitcoin get the address
	From []*From `protobuf:"bytes,4,rep,name=from,proto3" json:"from,omitempty"`
	// to is listening address
	To string `protobuf:"bytes,5,opt,name=to,proto3" json:"to,omitempty"`
	// value is from transfer amount
	Value int64 `protobuf:"varint,6,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *BitcoinTxParseResult) Reset()         { *m = BitcoinTxParseResult{} }
func (m *BitcoinTxParseResult) String() string { return proto.CompactTextString(m) }
func (*BitcoinTxParseResult) ProtoMessage()    {}
func (*BitcoinTxParseResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_a28da31cd1817aa6, []int{0}
}
func (m *BitcoinTxParseResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BitcoinTxParseResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BitcoinTxParseResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BitcoinTxParseResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BitcoinTxParseResult.Merge(m, src)
}
func (m *BitcoinTxParseResult) XXX_Size() int {
	return m.Size()
}
func (m *BitcoinTxParseResult) XXX_DiscardUnknown() {
	xxx_messageInfo_BitcoinTxParseResult.DiscardUnknown(m)
}

var xxx_messageInfo_BitcoinTxParseResult proto.InternalMessageInfo

// From is the from transaction info
type From struct {
	// tx_id is the btc from transaction id
	TxId string `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	// from is l2 user address, by parse bitcoin get the address
	From string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
}

func (m *From) Reset()         { *m = From{} }
func (m *From) String() string { return proto.CompactTextString(m) }
func (*From) ProtoMessage()    {}
func (*From) Descriptor() ([]byte, []int) {
	return fileDescriptor_a28da31cd1817aa6, []int{1}
}
func (m *From) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *From) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_From.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *From) XXX_Merge(src proto.Message) {
	xxx_messageInfo_From.Merge(m, src)
}
func (m *From) XXX_Size() int {
	return m.Size()
}
func (m *From) XXX_DiscardUnknown() {
	xxx_messageInfo_From.DiscardUnknown(m)
}

var xxx_messageInfo_From proto.InternalMessageInfo

func (m *From) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *From) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func init() {
	proto.RegisterType((*BitcoinTxParseResult)(nil), "ethermint.types.v1.BitcoinTxParseResult")
	proto.RegisterType((*From)(nil), "ethermint.types.v1.From")
}

func init() {
	proto.RegisterFile("ethermint/types/v1/bitcoin_indexer.proto", fileDescriptor_a28da31cd1817aa6)
}

var fileDescriptor_a28da31cd1817aa6 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x4e, 0xf3, 0x30,
	0x18, 0x86, 0xe3, 0x36, 0xed, 0xaf, 0xdf, 0x48, 0x0c, 0xa6, 0x12, 0x16, 0x83, 0x89, 0x3a, 0x65,
	0x40, 0xb6, 0x0a, 0x5b, 0xc7, 0x0e, 0x48, 0x6c, 0x28, 0xea, 0xc4, 0x12, 0x35, 0x8d, 0x49, 0x2d,
	0x35, 0x71, 0xe4, 0x7c, 0x89, 0xdc, 0x1b, 0x30, 0x72, 0x04, 0x0e, 0xc1, 0x21, 0x18, 0x3b, 0x32,
	0xa2, 0xe4, 0x22, 0x28, 0x0e, 0x62, 0x00, 0xb6, 0xef, 0xb3, 0x5f, 0x3f, 0x7e, 0xf4, 0xe2, 0x50,
	0xc2, 0x4e, 0x9a, 0x5c, 0x15, 0x20, 0xe0, 0x50, 0xca, 0x4a, 0x34, 0x0b, 0x91, 0x28, 0xd8, 0x6a,
	0x55, 0xc4, 0xaa, 0x48, 0xa5, 0x95, 0x86, 0x97, 0x46, 0x83, 0x26, 0xe4, 0x3b, 0xc9, 0x5d, 0x92,
	0x37, 0x8b, 0x8b, 0x59, 0xa6, 0x33, 0xed, 0xae, 0x45, 0x3f, 0x0d, 0xc9, 0xf9, 0x2b, 0xc2, 0xb3,
	0xd5, 0xc0, 0x58, 0xdb, 0xfb, 0x8d, 0xa9, 0x64, 0x24, 0xab, 0x7a, 0x0f, 0xe4, 0x0c, 0x4f, 0xc0,
	0xc6, 0x2a, 0xa5, 0x28, 0x40, 0xe1, 0xff, 0xc8, 0x07, 0x7b, 0x97, 0x92, 0x73, 0xfc, 0x0f, 0x6c,
	0xdc, 0x23, 0xe9, 0xc8, 0x1d, 0x4f, 0xc1, 0xae, 0x0f, 0xa5, 0x24, 0x33, 0x3c, 0x71, 0x06, 0x74,
	0x1c, 0xa0, 0x70, 0x1c, 0x0d, 0x0b, 0xb9, 0xc2, 0xfe, 0xa3, 0xd1, 0x39, 0xf5, 0x83, 0x71, 0x78,
	0x72, 0x4d, 0xf9, 0x6f, 0x2b, 0x7e, 0x6b, 0x74, 0x1e, 0xb9, 0x14, 0x39, 0xc5, 0x23, 0xd0, 0x74,
	0xe2, 0xb8, 0x23, 0xd0, 0x3d, 0xb3, 0xd9, 0xec, 0x6b, 0x49, 0xa7, 0x03, 0xd3, 0x2d, 0x4b, 0xff,
	0xe9, 0xe5, 0xd2, 0x9b, 0x0b, 0xec, 0xf7, 0x2f, 0xff, 0xb6, 0x24, 0x5f, 0xdf, 0x0e, 0x8a, 0x6e,
	0x5e, 0x2d, 0xdf, 0x5a, 0x86, 0x8e, 0x2d, 0x43, 0x1f, 0x2d, 0x43, 0xcf, 0x1d, 0xf3, 0x8e, 0x1d,
	0xf3, 0xde, 0x3b, 0xe6, 0x3d, 0x04, 0x99, 0x82, 0x5d, 0x9d, 0xf0, 0xad, 0xce, 0x85, 0x6c, 0x72,
	0x5d, 0x89, 0x1f, 0x35, 0x27, 0x53, 0x57, 0xd5, 0xcd, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x05,
	0xb1, 0x8d, 0x6d, 0x80, 0x01, 0x00, 0x00,
}

func (m *BitcoinTxParseResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BitcoinTxParseResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BitcoinTxParseResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i = encodeVarintBitcoinIndexer(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x30
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintBitcoinIndexer(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.From) > 0 {
		for iNdEx := len(m.From) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.From[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBitcoinIndexer(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Index != 0 {
		i = encodeVarintBitcoinIndexer(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x18
	}
	if len(m.TxType) > 0 {
		i -= len(m.TxType)
		copy(dAtA[i:], m.TxType)
		i = encodeVarintBitcoinIndexer(dAtA, i, uint64(len(m.TxType)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TxId) > 0 {
		i -= len(m.TxId)
		copy(dAtA[i:], m.TxId)
		i = encodeVarintBitcoinIndexer(dAtA, i, uint64(len(m.TxId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *From) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *From) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *From) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintBitcoinIndexer(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TxId) > 0 {
		i -= len(m.TxId)
		copy(dAtA[i:], m.TxId)
		i = encodeVarintBitcoinIndexer(dAtA, i, uint64(len(m.TxId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBitcoinIndexer(dAtA []byte, offset int, v uint64) int {
	offset -= sovBitcoinIndexer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BitcoinTxParseResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxId)
	if l > 0 {
		n += 1 + l + sovBitcoinIndexer(uint64(l))
	}
	l = len(m.TxType)
	if l > 0 {
		n += 1 + l + sovBitcoinIndexer(uint64(l))
	}
	if m.Index != 0 {
		n += 1 + sovBitcoinIndexer(uint64(m.Index))
	}
	if len(m.From) > 0 {
		for _, e := range m.From {
			l = e.Size()
			n += 1 + l + sovBitcoinIndexer(uint64(l))
		}
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovBitcoinIndexer(uint64(l))
	}
	if m.Value != 0 {
		n += 1 + sovBitcoinIndexer(uint64(m.Value))
	}
	return n
}

func (m *From) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxId)
	if l > 0 {
		n += 1 + l + sovBitcoinIndexer(uint64(l))
	}
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovBitcoinIndexer(uint64(l))
	}
	return n
}

func sovBitcoinIndexer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBitcoinIndexer(x uint64) (n int) {
	return sovBitcoinIndexer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BitcoinTxParseResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBitcoinIndexer
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
			return fmt.Errorf("proto: BitcoinTxParseResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BitcoinTxParseResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoinIndexer
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
				return ErrInvalidLengthBitcoinIndexer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBitcoinIndexer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoinIndexer
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
				return ErrInvalidLengthBitcoinIndexer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBitcoinIndexer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoinIndexer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoinIndexer
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
				return ErrInvalidLengthBitcoinIndexer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBitcoinIndexer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = append(m.From, &From{})
			if err := m.From[len(m.From)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoinIndexer
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
				return ErrInvalidLengthBitcoinIndexer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBitcoinIndexer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoinIndexer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBitcoinIndexer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBitcoinIndexer
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
func (m *From) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBitcoinIndexer
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
			return fmt.Errorf("proto: From: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: From: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoinIndexer
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
				return ErrInvalidLengthBitcoinIndexer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBitcoinIndexer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoinIndexer
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
				return ErrInvalidLengthBitcoinIndexer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBitcoinIndexer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBitcoinIndexer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBitcoinIndexer
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
func skipBitcoinIndexer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBitcoinIndexer
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
					return 0, ErrIntOverflowBitcoinIndexer
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
					return 0, ErrIntOverflowBitcoinIndexer
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
				return 0, ErrInvalidLengthBitcoinIndexer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBitcoinIndexer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBitcoinIndexer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBitcoinIndexer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBitcoinIndexer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBitcoinIndexer = fmt.Errorf("proto: unexpected end of group")
)
