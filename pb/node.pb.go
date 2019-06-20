// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/node.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strconv "strconv"

import bytes "bytes"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SyncState int32

const (
	WAIT_FOR_SYNCING SyncState = 0
	SYNC_STARTED     SyncState = 1
	SYNC_FINISHED    SyncState = 2
	PERSIST_FINISHED SyncState = 3
)

var SyncState_name = map[int32]string{
	0: "WAIT_FOR_SYNCING",
	1: "SYNC_STARTED",
	2: "SYNC_FINISHED",
	3: "PERSIST_FINISHED",
}
var SyncState_value = map[string]int32{
	"WAIT_FOR_SYNCING": 0,
	"SYNC_STARTED":     1,
	"SYNC_FINISHED":    2,
	"PERSIST_FINISHED": 3,
}

func (SyncState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_node_90a8b35fe1774c96, []int{0}
}

type NodeData struct {
	PublicKey       []byte `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	WebsocketPort   uint32 `protobuf:"varint,2,opt,name=websocket_port,json=websocketPort,proto3" json:"websocket_port,omitempty"`
	JsonRpcPort     uint32 `protobuf:"varint,3,opt,name=json_rpc_port,json=jsonRpcPort,proto3" json:"json_rpc_port,omitempty"`
	ProtocolVersion uint32 `protobuf:"varint,4,opt,name=protocol_version,json=protocolVersion,proto3" json:"protocol_version,omitempty"`
}

func (m *NodeData) Reset()      { *m = NodeData{} }
func (*NodeData) ProtoMessage() {}
func (*NodeData) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_90a8b35fe1774c96, []int{0}
}
func (m *NodeData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NodeData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NodeData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *NodeData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeData.Merge(dst, src)
}
func (m *NodeData) XXX_Size() int {
	return m.Size()
}
func (m *NodeData) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeData.DiscardUnknown(m)
}

var xxx_messageInfo_NodeData proto.InternalMessageInfo

func (m *NodeData) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *NodeData) GetWebsocketPort() uint32 {
	if m != nil {
		return m.WebsocketPort
	}
	return 0
}

func (m *NodeData) GetJsonRpcPort() uint32 {
	if m != nil {
		return m.JsonRpcPort
	}
	return 0
}

func (m *NodeData) GetProtocolVersion() uint32 {
	if m != nil {
		return m.ProtocolVersion
	}
	return 0
}

func init() {
	proto.RegisterType((*NodeData)(nil), "pb.NodeData")
	proto.RegisterEnum("pb.SyncState", SyncState_name, SyncState_value)
}
func (x SyncState) String() string {
	s, ok := SyncState_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *NodeData) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*NodeData)
	if !ok {
		that2, ok := that.(NodeData)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.PublicKey, that1.PublicKey) {
		return false
	}
	if this.WebsocketPort != that1.WebsocketPort {
		return false
	}
	if this.JsonRpcPort != that1.JsonRpcPort {
		return false
	}
	if this.ProtocolVersion != that1.ProtocolVersion {
		return false
	}
	return true
}
func (this *NodeData) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&pb.NodeData{")
	s = append(s, "PublicKey: "+fmt.Sprintf("%#v", this.PublicKey)+",\n")
	s = append(s, "WebsocketPort: "+fmt.Sprintf("%#v", this.WebsocketPort)+",\n")
	s = append(s, "JsonRpcPort: "+fmt.Sprintf("%#v", this.JsonRpcPort)+",\n")
	s = append(s, "ProtocolVersion: "+fmt.Sprintf("%#v", this.ProtocolVersion)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringNode(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *NodeData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeData) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PublicKey) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNode(dAtA, i, uint64(len(m.PublicKey)))
		i += copy(dAtA[i:], m.PublicKey)
	}
	if m.WebsocketPort != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintNode(dAtA, i, uint64(m.WebsocketPort))
	}
	if m.JsonRpcPort != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintNode(dAtA, i, uint64(m.JsonRpcPort))
	}
	if m.ProtocolVersion != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintNode(dAtA, i, uint64(m.ProtocolVersion))
	}
	return i, nil
}

func encodeVarintNode(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedNodeData(r randyNode, easy bool) *NodeData {
	this := &NodeData{}
	v1 := r.Intn(100)
	this.PublicKey = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.PublicKey[i] = byte(r.Intn(256))
	}
	this.WebsocketPort = uint32(r.Uint32())
	this.JsonRpcPort = uint32(r.Uint32())
	this.ProtocolVersion = uint32(r.Uint32())
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyNode interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneNode(r randyNode) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringNode(r randyNode) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneNode(r)
	}
	return string(tmps)
}
func randUnrecognizedNode(r randyNode, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldNode(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldNode(dAtA []byte, r randyNode, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateNode(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateNode(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateNode(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateNode(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateNode(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateNode(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateNode(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *NodeData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PublicKey)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	if m.WebsocketPort != 0 {
		n += 1 + sovNode(uint64(m.WebsocketPort))
	}
	if m.JsonRpcPort != 0 {
		n += 1 + sovNode(uint64(m.JsonRpcPort))
	}
	if m.ProtocolVersion != 0 {
		n += 1 + sovNode(uint64(m.ProtocolVersion))
	}
	return n
}

func sovNode(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNode(x uint64) (n int) {
	return sovNode(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *NodeData) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NodeData{`,
		`PublicKey:` + fmt.Sprintf("%v", this.PublicKey) + `,`,
		`WebsocketPort:` + fmt.Sprintf("%v", this.WebsocketPort) + `,`,
		`JsonRpcPort:` + fmt.Sprintf("%v", this.JsonRpcPort) + `,`,
		`ProtocolVersion:` + fmt.Sprintf("%v", this.ProtocolVersion) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringNode(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *NodeData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNode
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NodeData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKey = append(m.PublicKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PublicKey == nil {
				m.PublicKey = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WebsocketPort", wireType)
			}
			m.WebsocketPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WebsocketPort |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field JsonRpcPort", wireType)
			}
			m.JsonRpcPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.JsonRpcPort |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtocolVersion", wireType)
			}
			m.ProtocolVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProtocolVersion |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNode
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
func skipNode(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNode
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
					return 0, ErrIntOverflowNode
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNode
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthNode
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNode
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipNode(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthNode = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNode   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("pb/node.proto", fileDescriptor_node_90a8b35fe1774c96) }

var fileDescriptor_node_90a8b35fe1774c96 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8e, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x73, 0xda, 0xcb, 0xe5, 0x76, 0x6e, 0x73, 0x6f, 0x0c, 0x2e, 0x8a, 0xe0, 0xa1, 0x14,
	0x84, 0x2a, 0xd8, 0x2e, 0xdc, 0xba, 0xa9, 0xb6, 0xd5, 0x20, 0xc4, 0x92, 0x04, 0x45, 0x10, 0x42,
	0x33, 0x1d, 0x6b, 0x6d, 0xcd, 0x0c, 0xe9, 0x54, 0xe9, 0xce, 0x47, 0xf0, 0x0d, 0xdc, 0xfa, 0x08,
	0x3e, 0x82, 0xcb, 0x2e, 0xbb, 0x34, 0xd3, 0x8d, 0xcb, 0x2e, 0x5d, 0x4a, 0x27, 0xa8, 0xbb, 0xf3,
	0x7f, 0xe7, 0xfb, 0x39, 0x87, 0x98, 0x22, 0xaa, 0xc7, 0xbc, 0xc7, 0x6a, 0x22, 0xe1, 0x92, 0xdb,
	0x39, 0x11, 0x6d, 0xec, 0xf6, 0x07, 0xf2, 0x7a, 0x12, 0xd5, 0x28, 0xbf, 0xad, 0xf7, 0x79, 0x9f,
	0xd7, 0xf5, 0x2a, 0x9a, 0x5c, 0xe9, 0xa4, 0x83, 0x9e, 0xb2, 0x4a, 0xe5, 0x09, 0xc8, 0x1f, 0x97,
	0xf7, 0x58, 0xb3, 0x2b, 0xbb, 0xf6, 0x26, 0x21, 0x62, 0x12, 0x8d, 0x06, 0x34, 0x1c, 0xb2, 0x69,
	0x09, 0xca, 0x50, 0x2d, 0x7a, 0x85, 0x8c, 0x9c, 0xb0, 0xa9, 0xbd, 0x45, 0xfe, 0xdd, 0xb3, 0x68,
	0xcc, 0xe9, 0x90, 0xc9, 0x50, 0xf0, 0x44, 0x96, 0x72, 0x65, 0xa8, 0x9a, 0x9e, 0xf9, 0x4d, 0x3b,
	0x3c, 0x91, 0x76, 0x85, 0x98, 0x37, 0x63, 0x1e, 0x87, 0x89, 0xa0, 0x99, 0x95, 0xd7, 0xd6, 0xdf,
	0x15, 0xf4, 0x04, 0xd5, 0xce, 0x36, 0xb1, 0xf4, 0x7d, 0xca, 0x47, 0xe1, 0x1d, 0x4b, 0xc6, 0x03,
	0x1e, 0x97, 0x7e, 0x69, 0xed, 0xff, 0x17, 0x3f, 0xcb, 0xf0, 0xce, 0x25, 0x29, 0xf8, 0xd3, 0x98,
	0xfa, 0xb2, 0x2b, 0x99, 0xbd, 0x4e, 0xac, 0xf3, 0x86, 0x13, 0x84, 0xed, 0x53, 0x2f, 0xf4, 0x2f,
	0xdc, 0x43, 0xc7, 0x3d, 0xb2, 0x0c, 0xdb, 0x22, 0xc5, 0x55, 0x08, 0xfd, 0xa0, 0xe1, 0x05, 0xad,
	0xa6, 0x05, 0xf6, 0x1a, 0x31, 0x35, 0x69, 0x3b, 0xae, 0xe3, 0x1f, 0xb7, 0x9a, 0x56, 0x6e, 0x55,
	0xed, 0xb4, 0x3c, 0xdf, 0xf1, 0x83, 0x1f, 0x9a, 0x3f, 0xd8, 0x9f, 0xa5, 0x68, 0xcc, 0x53, 0x34,
	0x96, 0x29, 0xc2, 0x47, 0x8a, 0xf0, 0xa0, 0x10, 0x9e, 0x15, 0xc2, 0x8b, 0x42, 0x78, 0x55, 0x08,
	0x33, 0x85, 0xf0, 0xa6, 0x10, 0xde, 0x15, 0x1a, 0x4b, 0x85, 0xf0, 0xb8, 0x40, 0x63, 0xb6, 0x40,
	0x63, 0xbe, 0x40, 0x23, 0xfa, 0xad, 0x9f, 0xdd, 0xfb, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x8e,
	0x5d, 0xf8, 0x88, 0x01, 0x00, 0x00,
}
