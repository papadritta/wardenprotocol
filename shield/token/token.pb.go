// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shield/token/token.proto

package token

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type Type int32

const (
	Type_ILLEGAL   Type = 0
	Type_EOF       Type = 1
	Type_IDENT     Type = 2
	Type_INT       Type = 3
	Type_STRING    Type = 4
	Type_COMMA     Type = 5
	Type_SEMICOLON Type = 6
	Type_LPAREN    Type = 7
	Type_RPAREN    Type = 8
	Type_LBRACKET  Type = 9
	Type_RBRACKET  Type = 10
	Type_AND       Type = 11
	Type_OR        Type = 12
	Type_EQ        Type = 13
	Type_NEQ       Type = 14
	Type_GT        Type = 15
	Type_LT        Type = 16
	Type_GTE       Type = 17
	Type_LTE       Type = 18
	Type_TRUE      Type = 19
	Type_FALSE     Type = 20
)

var Type_name = map[int32]string{
	0:  "ILLEGAL",
	1:  "EOF",
	2:  "IDENT",
	3:  "INT",
	4:  "STRING",
	5:  "COMMA",
	6:  "SEMICOLON",
	7:  "LPAREN",
	8:  "RPAREN",
	9:  "LBRACKET",
	10: "RBRACKET",
	11: "AND",
	12: "OR",
	13: "EQ",
	14: "NEQ",
	15: "GT",
	16: "LT",
	17: "GTE",
	18: "LTE",
	19: "TRUE",
	20: "FALSE",
}

var Type_value = map[string]int32{
	"ILLEGAL":   0,
	"EOF":       1,
	"IDENT":     2,
	"INT":       3,
	"STRING":    4,
	"COMMA":     5,
	"SEMICOLON": 6,
	"LPAREN":    7,
	"RPAREN":    8,
	"LBRACKET":  9,
	"RBRACKET":  10,
	"AND":       11,
	"OR":        12,
	"EQ":        13,
	"NEQ":       14,
	"GT":        15,
	"LT":        16,
	"GTE":       17,
	"LTE":       18,
	"TRUE":      19,
	"FALSE":     20,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fae17a9db3bdc43d, []int{0}
}

type Token struct {
	Type    Type   `protobuf:"varint,1,opt,name=type,proto3,enum=shield.token.Type" json:"type,omitempty"`
	Literal string `protobuf:"bytes,2,opt,name=literal,proto3" json:"literal,omitempty"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_fae17a9db3bdc43d, []int{0}
}
func (m *Token) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Token.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return m.Size()
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_ILLEGAL
}

func (m *Token) GetLiteral() string {
	if m != nil {
		return m.Literal
	}
	return ""
}

func init() {
	proto.RegisterEnum("shield.token.Type", Type_name, Type_value)
	proto.RegisterType((*Token)(nil), "shield.token.Token")
}

func init() { proto.RegisterFile("shield/token/token.proto", fileDescriptor_fae17a9db3bdc43d) }

var fileDescriptor_fae17a9db3bdc43d = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x4f, 0x6b, 0xea, 0x40,
	0x14, 0xc5, 0x33, 0x1a, 0x13, 0x33, 0xfe, 0x79, 0xf7, 0xcd, 0x7b, 0x8b, 0xac, 0x82, 0xbc, 0xc5,
	0x43, 0x0a, 0x8d, 0xd0, 0x42, 0xf7, 0x51, 0xc7, 0x10, 0x3a, 0x26, 0x3a, 0x4e, 0x37, 0xdd, 0xf9,
	0x27, 0x54, 0x69, 0x6a, 0xc4, 0xa6, 0x14, 0xbf, 0x45, 0x3f, 0x56, 0x97, 0x2e, 0xbb, 0x2c, 0xe6,
	0x8b, 0x94, 0x1b, 0x6b, 0x71, 0x33, 0xf7, 0x9c, 0xf3, 0x3b, 0x30, 0xc3, 0x5c, 0x6a, 0x3f, 0x2f,
	0x57, 0x71, 0xb2, 0xe8, 0x64, 0xe9, 0x63, 0xbc, 0x3e, 0x9e, 0xee, 0x66, 0x9b, 0x66, 0x29, 0xab,
	0x1f, 0x89, 0x5b, 0x64, 0xff, 0x02, 0x5a, 0x51, 0x28, 0xd8, 0x7f, 0xaa, 0x67, 0xbb, 0x4d, 0x6c,
	0x93, 0x16, 0x69, 0x37, 0xaf, 0x98, 0x7b, 0xde, 0x72, 0xd5, 0x6e, 0x13, 0xcb, 0x82, 0x33, 0x9b,
	0x9a, 0xc9, 0x2a, 0x8b, 0xb7, 0xd3, 0xc4, 0x2e, 0xb5, 0x48, 0xdb, 0x92, 0x27, 0x7b, 0x91, 0x13,
	0xaa, 0x63, 0x91, 0xd5, 0xa8, 0x19, 0x08, 0xc1, 0x7d, 0x4f, 0x80, 0xc6, 0x4c, 0x5a, 0xe6, 0xd1,
	0x00, 0x08, 0xb3, 0x68, 0x25, 0xe8, 0xf3, 0x50, 0x41, 0x09, 0xb3, 0x20, 0x54, 0x50, 0x66, 0x94,
	0x1a, 0x13, 0x25, 0x83, 0xd0, 0x07, 0x1d, 0x79, 0x2f, 0x1a, 0x0e, 0x3d, 0xa8, 0xb0, 0x06, 0xb5,
	0x26, 0x7c, 0x18, 0xf4, 0x22, 0x11, 0x85, 0x60, 0x60, 0x4b, 0x8c, 0x3c, 0xc9, 0x43, 0x30, 0x51,
	0xcb, 0xa3, 0xae, 0xb2, 0x3a, 0xad, 0x8a, 0xae, 0xf4, 0x7a, 0xb7, 0x5c, 0x81, 0x85, 0x4e, 0x9e,
	0x1c, 0xc5, 0x2b, 0xbc, 0xb0, 0x0f, 0x35, 0x66, 0xd0, 0x52, 0x24, 0xa1, 0x8e, 0x93, 0x8f, 0xa1,
	0x81, 0x20, 0xe4, 0x63, 0x68, 0x62, 0xe0, 0x2b, 0xf8, 0x85, 0x53, 0x28, 0x00, 0x04, 0xbe, 0xe2,
	0xf0, 0x1b, 0x85, 0x50, 0x1c, 0x18, 0xab, 0x52, 0x5d, 0xc9, 0x3b, 0x0e, 0x7f, 0xf0, 0x6d, 0x03,
	0x4f, 0x4c, 0x38, 0xfc, 0xed, 0x8e, 0xde, 0x0f, 0x0e, 0xd9, 0x1f, 0x1c, 0xf2, 0x79, 0x70, 0xc8,
	0x5b, 0xee, 0x68, 0xfb, 0xdc, 0xd1, 0x3e, 0x72, 0x47, 0xbb, 0xbf, 0x79, 0x58, 0x65, 0xcb, 0x97,
	0x99, 0x3b, 0x4f, 0x9f, 0x3a, 0xaf, 0xd3, 0xed, 0x22, 0x5e, 0x5f, 0x16, 0x3f, 0x3e, 0x4f, 0x93,
	0x6f, 0xff, 0x63, 0xcf, 0x97, 0x33, 0x33, 0x8a, 0xf8, 0xfa, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x60,
	0x9f, 0xf6, 0x01, 0xb3, 0x01, 0x00, 0x00,
}

func (m *Token) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Token) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Token) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Literal) > 0 {
		i -= len(m.Literal)
		copy(dAtA[i:], m.Literal)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Literal)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintToken(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintToken(dAtA []byte, offset int, v uint64) int {
	offset -= sovToken(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Token) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovToken(uint64(m.Type))
	}
	l = len(m.Literal)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	return n
}

func sovToken(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozToken(x uint64) (n int) {
	return sovToken(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Token) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowToken
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
			return fmt.Errorf("proto: Token: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Token: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Literal", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Literal = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipToken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthToken
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
func skipToken(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowToken
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
					return 0, ErrIntOverflowToken
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
					return 0, ErrIntOverflowToken
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
				return 0, ErrInvalidLengthToken
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupToken
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthToken
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthToken        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowToken          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupToken = fmt.Errorf("proto: unexpected end of group")
)
