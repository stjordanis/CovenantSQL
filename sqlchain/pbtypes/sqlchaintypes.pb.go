// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sqlchaintypes.proto

package pbtypes

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TxType int32

const (
	TxType_QUERY   TxType = 0
	TxType_STORAGE TxType = 1
)

var TxType_name = map[int32]string{
	0: "QUERY",
	1: "STORAGE",
}
var TxType_value = map[string]int32{
	"QUERY":   0,
	"STORAGE": 1,
}

func (x TxType) String() string {
	return proto.EnumName(TxType_name, int32(x))
}
func (TxType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_sqlchaintypes_12adb04c278231b0, []int{0}
}

type Signature struct {
	R                    string   `protobuf:"bytes,1,opt,name=R" json:"R,omitempty"`
	S                    string   `protobuf:"bytes,2,opt,name=S" json:"S,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Signature) Reset()         { *m = Signature{} }
func (m *Signature) String() string { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()    {}
func (*Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_sqlchaintypes_12adb04c278231b0, []int{0}
}
func (m *Signature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Signature.Unmarshal(m, b)
}
func (m *Signature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Signature.Marshal(b, m, deterministic)
}
func (dst *Signature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signature.Merge(dst, src)
}
func (m *Signature) XXX_Size() int {
	return xxx_messageInfo_Signature.Size(m)
}
func (m *Signature) XXX_DiscardUnknown() {
	xxx_messageInfo_Signature.DiscardUnknown(m)
}

var xxx_messageInfo_Signature proto.InternalMessageInfo

func (m *Signature) GetR() string {
	if m != nil {
		return m.R
	}
	return ""
}

func (m *Signature) GetS() string {
	if m != nil {
		return m.S
	}
	return ""
}

type PublicKey struct {
	PublicKey            []byte   `protobuf:"bytes,1,opt,name=PublicKey,proto3" json:"PublicKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublicKey) Reset()         { *m = PublicKey{} }
func (m *PublicKey) String() string { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()    {}
func (*PublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_sqlchaintypes_12adb04c278231b0, []int{1}
}
func (m *PublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicKey.Unmarshal(m, b)
}
func (m *PublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicKey.Marshal(b, m, deterministic)
}
func (dst *PublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicKey.Merge(dst, src)
}
func (m *PublicKey) XXX_Size() int {
	return xxx_messageInfo_PublicKey.Size(m)
}
func (m *PublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_PublicKey proto.InternalMessageInfo

func (m *PublicKey) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

type Hash struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hash) Reset()         { *m = Hash{} }
func (m *Hash) String() string { return proto.CompactTextString(m) }
func (*Hash) ProtoMessage()    {}
func (*Hash) Descriptor() ([]byte, []int) {
	return fileDescriptor_sqlchaintypes_12adb04c278231b0, []int{2}
}
func (m *Hash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hash.Unmarshal(m, b)
}
func (m *Hash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hash.Marshal(b, m, deterministic)
}
func (dst *Hash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hash.Merge(dst, src)
}
func (m *Hash) XXX_Size() int {
	return xxx_messageInfo_Hash.Size(m)
}
func (m *Hash) XXX_DiscardUnknown() {
	xxx_messageInfo_Hash.DiscardUnknown(m)
}

var xxx_messageInfo_Hash proto.InternalMessageInfo

func (m *Hash) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type UtxoEntry struct {
	IsCoinbase           bool             `protobuf:"varint,1,opt,name=IsCoinbase" json:"IsCoinbase,omitempty"`
	FromMainChain        bool             `protobuf:"varint,2,opt,name=FromMainChain" json:"FromMainChain,omitempty"`
	BlockHeight          uint32           `protobuf:"varint,3,opt,name=BlockHeight" json:"BlockHeight,omitempty"`
	SparseOutputs        map[uint32]*Utxo `protobuf:"bytes,4,rep,name=SparseOutputs" json:"SparseOutputs,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UtxoEntry) Reset()         { *m = UtxoEntry{} }
func (m *UtxoEntry) String() string { return proto.CompactTextString(m) }
func (*UtxoEntry) ProtoMessage()    {}
func (*UtxoEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_sqlchaintypes_12adb04c278231b0, []int{3}
}
func (m *UtxoEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UtxoEntry.Unmarshal(m, b)
}
func (m *UtxoEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UtxoEntry.Marshal(b, m, deterministic)
}
func (dst *UtxoEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UtxoEntry.Merge(dst, src)
}
func (m *UtxoEntry) XXX_Size() int {
	return xxx_messageInfo_UtxoEntry.Size(m)
}
func (m *UtxoEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_UtxoEntry.DiscardUnknown(m)
}

var xxx_messageInfo_UtxoEntry proto.InternalMessageInfo

func (m *UtxoEntry) GetIsCoinbase() bool {
	if m != nil {
		return m.IsCoinbase
	}
	return false
}

func (m *UtxoEntry) GetFromMainChain() bool {
	if m != nil {
		return m.FromMainChain
	}
	return false
}

func (m *UtxoEntry) GetBlockHeight() uint32 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *UtxoEntry) GetSparseOutputs() map[uint32]*Utxo {
	if m != nil {
		return m.SparseOutputs
	}
	return nil
}

type Utxo struct {
	UtxoHeader           *UtxoHeader `protobuf:"bytes,1,opt,name=UtxoHeader" json:"UtxoHeader,omitempty"`
	Spent                bool        `protobuf:"varint,2,opt,name=Spent" json:"Spent,omitempty"`
	Amount               uint64      `protobuf:"varint,3,opt,name=amount" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Utxo) Reset()         { *m = Utxo{} }
func (m *Utxo) String() string { return proto.CompactTextString(m) }
func (*Utxo) ProtoMessage()    {}
func (*Utxo) Descriptor() ([]byte, []int) {
	return fileDescriptor_sqlchaintypes_12adb04c278231b0, []int{4}
}
func (m *Utxo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Utxo.Unmarshal(m, b)
}
func (m *Utxo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Utxo.Marshal(b, m, deterministic)
}
func (dst *Utxo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Utxo.Merge(dst, src)
}
func (m *Utxo) XXX_Size() int {
	return xxx_messageInfo_Utxo.Size(m)
}
func (m *Utxo) XXX_DiscardUnknown() {
	xxx_messageInfo_Utxo.DiscardUnknown(m)
}

var xxx_messageInfo_Utxo proto.InternalMessageInfo

func (m *Utxo) GetUtxoHeader() *UtxoHeader {
	if m != nil {
		return m.UtxoHeader
	}
	return nil
}

func (m *Utxo) GetSpent() bool {
	if m != nil {
		return m.Spent
	}
	return false
}

func (m *Utxo) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type UtxoHeader struct {
	Version              int32      `protobuf:"varint,1,opt,name=Version" json:"Version,omitempty"`
	PrevTxHash           *Hash      `protobuf:"bytes,2,opt,name=PrevTxHash" json:"PrevTxHash,omitempty"`
	Signee               *PublicKey `protobuf:"bytes,3,opt,name=Signee" json:"Signee,omitempty"`
	Signature            *Signature `protobuf:"bytes,4,opt,name=Signature" json:"Signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UtxoHeader) Reset()         { *m = UtxoHeader{} }
func (m *UtxoHeader) String() string { return proto.CompactTextString(m) }
func (*UtxoHeader) ProtoMessage()    {}
func (*UtxoHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_sqlchaintypes_12adb04c278231b0, []int{5}
}
func (m *UtxoHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UtxoHeader.Unmarshal(m, b)
}
func (m *UtxoHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UtxoHeader.Marshal(b, m, deterministic)
}
func (dst *UtxoHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UtxoHeader.Merge(dst, src)
}
func (m *UtxoHeader) XXX_Size() int {
	return xxx_messageInfo_UtxoHeader.Size(m)
}
func (m *UtxoHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_UtxoHeader.DiscardUnknown(m)
}

var xxx_messageInfo_UtxoHeader proto.InternalMessageInfo

func (m *UtxoHeader) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *UtxoHeader) GetPrevTxHash() *Hash {
	if m != nil {
		return m.PrevTxHash
	}
	return nil
}

func (m *UtxoHeader) GetSignee() *PublicKey {
	if m != nil {
		return m.Signee
	}
	return nil
}

func (m *UtxoHeader) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Tx struct {
	UtxoIn               []*Utxo  `protobuf:"bytes,1,rep,name=UtxoIn" json:"UtxoIn,omitempty"`
	UtxoOut              []*Utxo  `protobuf:"bytes,2,rep,name=UtxoOut" json:"UtxoOut,omitempty"`
	Type                 TxType   `protobuf:"varint,3,opt,name=type,enum=pbtypes.TxType" json:"type,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=Content" json:"Content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tx) Reset()         { *m = Tx{} }
func (m *Tx) String() string { return proto.CompactTextString(m) }
func (*Tx) ProtoMessage()    {}
func (*Tx) Descriptor() ([]byte, []int) {
	return fileDescriptor_sqlchaintypes_12adb04c278231b0, []int{6}
}
func (m *Tx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tx.Unmarshal(m, b)
}
func (m *Tx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tx.Marshal(b, m, deterministic)
}
func (dst *Tx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tx.Merge(dst, src)
}
func (m *Tx) XXX_Size() int {
	return xxx_messageInfo_Tx.Size(m)
}
func (m *Tx) XXX_DiscardUnknown() {
	xxx_messageInfo_Tx.DiscardUnknown(m)
}

var xxx_messageInfo_Tx proto.InternalMessageInfo

func (m *Tx) GetUtxoIn() []*Utxo {
	if m != nil {
		return m.UtxoIn
	}
	return nil
}

func (m *Tx) GetUtxoOut() []*Utxo {
	if m != nil {
		return m.UtxoOut
	}
	return nil
}

func (m *Tx) GetType() TxType {
	if m != nil {
		return m.Type
	}
	return TxType_QUERY
}

func (m *Tx) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type Address struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_sqlchaintypes_12adb04c278231b0, []int{7}
}
func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (dst *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(dst, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

type Header struct {
	Version              int32    `protobuf:"varint,1,opt,name=Version" json:"Version,omitempty"`
	Producer             *Address `protobuf:"bytes,2,opt,name=Producer" json:"Producer,omitempty"`
	Root                 *Hash    `protobuf:"bytes,3,opt,name=Root" json:"Root,omitempty"`
	Parent               *Hash    `protobuf:"bytes,4,opt,name=Parent" json:"Parent,omitempty"`
	MerkleRoot           *Hash    `protobuf:"bytes,5,opt,name=MerkleRoot" json:"MerkleRoot,omitempty"`
	Timestamp            []byte   `protobuf:"bytes,6,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_sqlchaintypes_12adb04c278231b0, []int{8}
}
func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (dst *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(dst, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Header) GetProducer() *Address {
	if m != nil {
		return m.Producer
	}
	return nil
}

func (m *Header) GetRoot() *Hash {
	if m != nil {
		return m.Root
	}
	return nil
}

func (m *Header) GetParent() *Hash {
	if m != nil {
		return m.Parent
	}
	return nil
}

func (m *Header) GetMerkleRoot() *Hash {
	if m != nil {
		return m.MerkleRoot
	}
	return nil
}

func (m *Header) GetTimestamp() []byte {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type SignedHeader struct {
	Header               *Header    `protobuf:"bytes,1,opt,name=Header" json:"Header,omitempty"`
	BlockHash            *Hash      `protobuf:"bytes,2,opt,name=BlockHash" json:"BlockHash,omitempty"`
	Signee               *PublicKey `protobuf:"bytes,3,opt,name=Signee" json:"Signee,omitempty"`
	Signature            *Signature `protobuf:"bytes,4,opt,name=Signature" json:"Signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SignedHeader) Reset()         { *m = SignedHeader{} }
func (m *SignedHeader) String() string { return proto.CompactTextString(m) }
func (*SignedHeader) ProtoMessage()    {}
func (*SignedHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_sqlchaintypes_12adb04c278231b0, []int{9}
}
func (m *SignedHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedHeader.Unmarshal(m, b)
}
func (m *SignedHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedHeader.Marshal(b, m, deterministic)
}
func (dst *SignedHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedHeader.Merge(dst, src)
}
func (m *SignedHeader) XXX_Size() int {
	return xxx_messageInfo_SignedHeader.Size(m)
}
func (m *SignedHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedHeader.DiscardUnknown(m)
}

var xxx_messageInfo_SignedHeader proto.InternalMessageInfo

func (m *SignedHeader) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SignedHeader) GetBlockHash() *Hash {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *SignedHeader) GetSignee() *PublicKey {
	if m != nil {
		return m.Signee
	}
	return nil
}

func (m *SignedHeader) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

type State struct {
	Head                 *Hash    `protobuf:"bytes,1,opt,name=Head" json:"Head,omitempty"`
	Height               int32    `protobuf:"varint,2,opt,name=Height" json:"Height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_sqlchaintypes_12adb04c278231b0, []int{10}
}
func (m *State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State.Unmarshal(m, b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State.Marshal(b, m, deterministic)
}
func (dst *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(dst, src)
}
func (m *State) XXX_Size() int {
	return xxx_messageInfo_State.Size(m)
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetHead() *Hash {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *State) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterType((*Signature)(nil), "pbtypes.Signature")
	proto.RegisterType((*PublicKey)(nil), "pbtypes.PublicKey")
	proto.RegisterType((*Hash)(nil), "pbtypes.Hash")
	proto.RegisterType((*UtxoEntry)(nil), "pbtypes.UtxoEntry")
	proto.RegisterMapType((map[uint32]*Utxo)(nil), "pbtypes.UtxoEntry.SparseOutputsEntry")
	proto.RegisterType((*Utxo)(nil), "pbtypes.Utxo")
	proto.RegisterType((*UtxoHeader)(nil), "pbtypes.UtxoHeader")
	proto.RegisterType((*Tx)(nil), "pbtypes.Tx")
	proto.RegisterType((*Address)(nil), "pbtypes.Address")
	proto.RegisterType((*Header)(nil), "pbtypes.Header")
	proto.RegisterType((*SignedHeader)(nil), "pbtypes.SignedHeader")
	proto.RegisterType((*State)(nil), "pbtypes.State")
	proto.RegisterEnum("pbtypes.TxType", TxType_name, TxType_value)
}

func init() { proto.RegisterFile("sqlchaintypes.proto", fileDescriptor_sqlchaintypes_12adb04c278231b0) }

var fileDescriptor_sqlchaintypes_12adb04c278231b0 = []byte{
	// 640 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xd1, 0x6e, 0xd3, 0x3c,
	0x14, 0xfe, 0xd3, 0xa6, 0xe9, 0x7a, 0xb2, 0xfe, 0x54, 0x67, 0x08, 0x45, 0x13, 0x42, 0x25, 0x63,
	0xda, 0x18, 0x50, 0xa1, 0xee, 0x06, 0x71, 0xb7, 0x4d, 0x83, 0x4d, 0xd3, 0xd4, 0xe2, 0x64, 0x48,
	0x5c, 0xba, 0xad, 0xb5, 0x45, 0x6b, 0xe3, 0xe0, 0x38, 0x53, 0xfb, 0x18, 0x88, 0x47, 0xe1, 0x0d,
	0x78, 0x19, 0x5e, 0x03, 0xd9, 0x71, 0xd2, 0xb4, 0x1b, 0xe2, 0x8a, 0xab, 0xf8, 0x9c, 0xef, 0xb3,
	0xfd, 0xf9, 0x3b, 0xe7, 0x04, 0xb6, 0xd2, 0xaf, 0xd3, 0xf1, 0x0d, 0x8d, 0x62, 0xb9, 0x48, 0x58,
	0xda, 0x4b, 0x04, 0x97, 0x1c, 0x9b, 0xc9, 0x48, 0x87, 0xfe, 0x1e, 0xb4, 0x82, 0xe8, 0x3a, 0xa6,
	0x32, 0x13, 0x0c, 0x37, 0xc1, 0x22, 0x9e, 0xd5, 0xb5, 0xf6, 0x5b, 0xc4, 0x22, 0x2a, 0x0a, 0xbc,
	0x5a, 0x1e, 0x05, 0xfe, 0x4b, 0x68, 0x0d, 0xb3, 0xd1, 0x34, 0x1a, 0x5f, 0xb0, 0x05, 0x3e, 0xad,
	0x04, 0x7a, 0xc3, 0x26, 0x59, 0x26, 0xfc, 0x6d, 0xb0, 0xcf, 0x68, 0x7a, 0x83, 0x98, 0x7f, 0x0d,
	0x41, 0xaf, 0xfd, 0x6f, 0x35, 0x68, 0x5d, 0xc9, 0x39, 0x3f, 0x8d, 0xa5, 0x58, 0xe0, 0x33, 0x80,
	0xf3, 0xf4, 0x84, 0x47, 0xf1, 0x88, 0xa6, 0x4c, 0xf3, 0x36, 0x48, 0x25, 0x83, 0x2f, 0xa0, 0xfd,
	0x41, 0xf0, 0xd9, 0x25, 0x8d, 0xe2, 0x13, 0xf5, 0x04, 0x2d, 0x67, 0x83, 0xac, 0x26, 0xb1, 0x0b,
	0xee, 0xf1, 0x94, 0x8f, 0x6f, 0xcf, 0x58, 0x74, 0x7d, 0x23, 0xbd, 0x7a, 0xd7, 0xda, 0x6f, 0x93,
	0x6a, 0x0a, 0x2f, 0xa0, 0x1d, 0x24, 0x54, 0xa4, 0x6c, 0x90, 0xc9, 0x24, 0x93, 0xa9, 0x67, 0x77,
	0xeb, 0xfb, 0x6e, 0x7f, 0xb7, 0x67, 0x6c, 0xe8, 0x95, 0x92, 0x7a, 0x2b, 0x3c, 0x9d, 0x22, 0xab,
	0x7b, 0xb7, 0x07, 0x80, 0xf7, 0x49, 0xd8, 0x81, 0xfa, 0xad, 0x31, 0xa3, 0x4d, 0xd4, 0x12, 0x77,
	0xa0, 0x71, 0x47, 0xa7, 0x19, 0xd3, 0xa2, 0xdd, 0x7e, 0x7b, 0xe5, 0x32, 0x92, 0x63, 0xef, 0x6b,
	0xef, 0x2c, 0x3f, 0x02, 0x5b, 0xa5, 0xf0, 0x10, 0x40, 0x7d, 0xcf, 0x18, 0x9d, 0x30, 0xa1, 0x4f,
	0x72, 0xfb, 0x5b, 0x2b, 0xbb, 0x72, 0x88, 0x54, 0x68, 0xf8, 0x18, 0x1a, 0x41, 0xc2, 0x62, 0x69,
	0xac, 0xc9, 0x03, 0x7c, 0x02, 0x0e, 0x9d, 0xf1, 0x2c, 0xce, 0xdd, 0xb0, 0x89, 0x89, 0xfc, 0x1f,
	0x56, 0xf5, 0x0e, 0xf4, 0xa0, 0xf9, 0x99, 0x89, 0x34, 0xe2, 0xb1, 0xbe, 0xae, 0x41, 0x8a, 0x10,
	0xdf, 0x00, 0x0c, 0x05, 0xbb, 0x0b, 0xe7, 0xba, 0x82, 0xeb, 0x2f, 0x50, 0x49, 0x52, 0x21, 0xe0,
	0x01, 0x38, 0xaa, 0x8d, 0x18, 0xd3, 0xf7, 0xb9, 0x7d, 0x2c, 0xa9, 0x65, 0x5b, 0x10, 0xc3, 0xc0,
	0xb7, 0x95, 0x96, 0xf3, 0xec, 0x35, 0x7a, 0x89, 0x90, 0x25, 0xc9, 0xff, 0x6e, 0x41, 0x2d, 0x9c,
	0xe3, 0x2e, 0x38, 0x4a, 0xfb, 0xb9, 0x12, 0x5b, 0xbf, 0xef, 0xa8, 0x01, 0x71, 0x0f, 0x9a, 0x6a,
	0x35, 0xc8, 0x94, 0x27, 0x0f, 0xf0, 0x0a, 0x14, 0x77, 0xc0, 0x56, 0x69, 0x2d, 0xf9, 0xff, 0xfe,
	0xa3, 0x92, 0x15, 0xce, 0xc3, 0x45, 0xc2, 0x88, 0x06, 0x95, 0x45, 0x27, 0x3c, 0x96, 0xca, 0x61,
	0x5b, 0xcf, 0x42, 0x11, 0xfa, 0x3b, 0xd0, 0x3c, 0x9a, 0x4c, 0x04, 0x4b, 0x53, 0x45, 0x32, 0x4b,
	0xd3, 0xec, 0x45, 0xe8, 0xff, 0xb2, 0xc0, 0xf9, 0xab, 0xd9, 0xaf, 0x61, 0x63, 0x28, 0xf8, 0x24,
	0x1b, 0x33, 0x61, 0xac, 0xee, 0x94, 0x62, 0xcc, 0x41, 0xa4, 0x64, 0xe0, 0x73, 0xb0, 0x09, 0xe7,
	0xd2, 0x38, 0xbd, 0x56, 0x14, 0x0d, 0x29, 0xa7, 0x86, 0x54, 0x14, 0x9a, 0xef, 0x91, 0x0c, 0xa8,
	0x8a, 0x7c, 0xc9, 0xc4, 0xed, 0x94, 0xe9, 0xf3, 0x1a, 0x0f, 0x16, 0x79, 0x49, 0x50, 0x53, 0x1f,
	0x46, 0x33, 0x96, 0x4a, 0x3a, 0x4b, 0x3c, 0x27, 0x9f, 0xfa, 0x32, 0xe1, 0xff, 0xb4, 0x60, 0x53,
	0x57, 0x78, 0x62, 0xde, 0xbb, 0x57, 0xbc, 0xdc, 0xb4, 0xf2, 0xd2, 0x60, 0xd3, 0xc6, 0x85, 0x31,
	0xaf, 0xa0, 0x95, 0x0f, 0xeb, 0x1f, 0x5b, 0x6d, 0x89, 0xff, 0xe3, 0x4e, 0x3b, 0x86, 0x46, 0x20,
	0xa9, 0x64, 0xca, 0x64, 0xa5, 0xce, 0x48, 0x5f, 0x37, 0x59, 0x41, 0x6a, 0xc6, 0xcc, 0x1f, 0xa7,
	0xa6, 0xcb, 0x69, 0xa2, 0x83, 0x2e, 0x38, 0x79, 0x07, 0x61, 0x0b, 0x1a, 0x9f, 0xae, 0x4e, 0xc9,
	0x97, 0xce, 0x7f, 0xe8, 0x42, 0x33, 0x08, 0x07, 0xe4, 0xe8, 0xe3, 0x69, 0xc7, 0x1a, 0x39, 0xfa,
	0x27, 0x7c, 0xf8, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x87, 0x56, 0x6e, 0xce, 0x9b, 0x05, 0x00, 0x00,
}
