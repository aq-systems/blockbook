// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bchain/coins/eth/ethtx.proto

/*
Package eth is a generated protocol buffer package.

It is generated from these files:
	bchain/coins/eth/ethtx.proto

It has these top-level messages:
	ProtoCompleteTransaction
*/
package iochain

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

type ProtoCompleteTransaction struct {
	BlockNumber uint32                                `protobuf:"varint,1,opt,name=BlockNumber" json:"BlockNumber,omitempty"`
	BlockTime   uint64                                `protobuf:"varint,2,opt,name=BlockTime" json:"BlockTime,omitempty"`
	Tx          *ProtoCompleteTransaction_TxType      `protobuf:"bytes,3,opt,name=Tx" json:"Tx,omitempty"`
	Receipt     *ProtoCompleteTransaction_ReceiptType `protobuf:"bytes,4,opt,name=Receipt" json:"Receipt,omitempty"`
}

func (m *ProtoCompleteTransaction) Reset()                    { *m = ProtoCompleteTransaction{} }
func (m *ProtoCompleteTransaction) String() string            { return proto.CompactTextString(m) }
func (*ProtoCompleteTransaction) ProtoMessage()               {}
func (*ProtoCompleteTransaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProtoCompleteTransaction) GetBlockNumber() uint32 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func (m *ProtoCompleteTransaction) GetBlockTime() uint64 {
	if m != nil {
		return m.BlockTime
	}
	return 0
}

func (m *ProtoCompleteTransaction) GetTx() *ProtoCompleteTransaction_TxType {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *ProtoCompleteTransaction) GetReceipt() *ProtoCompleteTransaction_ReceiptType {
	if m != nil {
		return m.Receipt
	}
	return nil
}

type ProtoCompleteTransaction_TxType struct {
	AccountNonce     uint64 `protobuf:"varint,1,opt,name=AccountNonce" json:"AccountNonce,omitempty"`
	GasPrice         []byte `protobuf:"bytes,2,opt,name=GasPrice,proto3" json:"GasPrice,omitempty"`
	GasLimit         uint64 `protobuf:"varint,3,opt,name=GasLimit" json:"GasLimit,omitempty"`
	Value            []byte `protobuf:"bytes,4,opt,name=Value,proto3" json:"Value,omitempty"`
	Payload          []byte `protobuf:"bytes,5,opt,name=Payload,proto3" json:"Payload,omitempty"`
	Hash             []byte `protobuf:"bytes,6,opt,name=Hash,proto3" json:"Hash,omitempty"`
	To               []byte `protobuf:"bytes,7,opt,name=To,proto3" json:"To,omitempty"`
	From             []byte `protobuf:"bytes,8,opt,name=From,proto3" json:"From,omitempty"`
	TransactionIndex uint32 `protobuf:"varint,9,opt,name=TransactionIndex" json:"TransactionIndex,omitempty"`
}

func (m *ProtoCompleteTransaction_TxType) Reset()         { *m = ProtoCompleteTransaction_TxType{} }
func (m *ProtoCompleteTransaction_TxType) String() string { return proto.CompactTextString(m) }
func (*ProtoCompleteTransaction_TxType) ProtoMessage()    {}
func (*ProtoCompleteTransaction_TxType) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

func (m *ProtoCompleteTransaction_TxType) GetAccountNonce() uint64 {
	if m != nil {
		return m.AccountNonce
	}
	return 0
}

func (m *ProtoCompleteTransaction_TxType) GetGasPrice() []byte {
	if m != nil {
		return m.GasPrice
	}
	return nil
}

func (m *ProtoCompleteTransaction_TxType) GetGasLimit() uint64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *ProtoCompleteTransaction_TxType) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ProtoCompleteTransaction_TxType) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ProtoCompleteTransaction_TxType) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *ProtoCompleteTransaction_TxType) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *ProtoCompleteTransaction_TxType) GetFrom() []byte {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *ProtoCompleteTransaction_TxType) GetTransactionIndex() uint32 {
	if m != nil {
		return m.TransactionIndex
	}
	return 0
}

type ProtoCompleteTransaction_ReceiptType struct {
	GasUsed []byte                                          `protobuf:"bytes,1,opt,name=GasUsed,proto3" json:"GasUsed,omitempty"`
	Status  []byte                                          `protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
	Log     []*ProtoCompleteTransaction_ReceiptType_LogType `protobuf:"bytes,3,rep,name=Log" json:"Log,omitempty"`
}

func (m *ProtoCompleteTransaction_ReceiptType) Reset()         { *m = ProtoCompleteTransaction_ReceiptType{} }
func (m *ProtoCompleteTransaction_ReceiptType) String() string { return proto.CompactTextString(m) }
func (*ProtoCompleteTransaction_ReceiptType) ProtoMessage()    {}
func (*ProtoCompleteTransaction_ReceiptType) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 1}
}

func (m *ProtoCompleteTransaction_ReceiptType) GetGasUsed() []byte {
	if m != nil {
		return m.GasUsed
	}
	return nil
}

func (m *ProtoCompleteTransaction_ReceiptType) GetStatus() []byte {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ProtoCompleteTransaction_ReceiptType) GetLog() []*ProtoCompleteTransaction_ReceiptType_LogType {
	if m != nil {
		return m.Log
	}
	return nil
}

type ProtoCompleteTransaction_ReceiptType_LogType struct {
	Address []byte   `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	Data    []byte   `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	Topics  [][]byte `protobuf:"bytes,3,rep,name=Topics,proto3" json:"Topics,omitempty"`
}

func (m *ProtoCompleteTransaction_ReceiptType_LogType) Reset() {
	*m = ProtoCompleteTransaction_ReceiptType_LogType{}
}
func (m *ProtoCompleteTransaction_ReceiptType_LogType) String() string {
	return proto.CompactTextString(m)
}
func (*ProtoCompleteTransaction_ReceiptType_LogType) ProtoMessage() {}
func (*ProtoCompleteTransaction_ReceiptType_LogType) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 1, 0}
}

func (m *ProtoCompleteTransaction_ReceiptType_LogType) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *ProtoCompleteTransaction_ReceiptType_LogType) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ProtoCompleteTransaction_ReceiptType_LogType) GetTopics() [][]byte {
	if m != nil {
		return m.Topics
	}
	return nil
}

func init() {
	proto.RegisterType((*ProtoCompleteTransaction)(nil), "eth.ProtoCompleteTransaction")
	proto.RegisterType((*ProtoCompleteTransaction_TxType)(nil), "eth.ProtoCompleteTransaction.TxType")
	proto.RegisterType((*ProtoCompleteTransaction_ReceiptType)(nil), "eth.ProtoCompleteTransaction.ReceiptType")
	proto.RegisterType((*ProtoCompleteTransaction_ReceiptType_LogType)(nil), "eth.ProtoCompleteTransaction.ReceiptType.LogType")
}

func init() { proto.RegisterFile("bchain/coins/eth/ethtx.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdf, 0x8a, 0xd4, 0x30,
	0x18, 0xc5, 0xe9, 0x9f, 0x99, 0xd9, 0xfd, 0xa6, 0x8a, 0x04, 0x91, 0x30, 0xec, 0x45, 0x59, 0xbc,
	0x18, 0xbd, 0xe8, 0xe2, 0xea, 0x0b, 0xac, 0x23, 0xae, 0xc2, 0xb0, 0x0e, 0x31, 0x7a, 0x9f, 0x49,
	0xc3, 0x36, 0x38, 0x6d, 0x4a, 0x93, 0x42, 0xf7, 0x8d, 0x7c, 0x21, 0xdf, 0xc5, 0x4b, 0xc9, 0xd7,
	0x74, 0x1d, 0x11, 0x65, 0x2f, 0x0a, 0xf9, 0x9d, 0x7e, 0xa7, 0x39, 0x27, 0x29, 0x9c, 0xed, 0x65,
	0x25, 0x74, 0x73, 0x21, 0x8d, 0x6e, 0xec, 0x85, 0x72, 0x95, 0x7f, 0xdc, 0x50, 0xb4, 0x9d, 0x71,
	0x86, 0x24, 0xca, 0x55, 0xe7, 0xdf, 0x67, 0x40, 0x77, 0x1e, 0x37, 0xa6, 0x6e, 0x0f, 0xca, 0x29,
	0xde, 0x89, 0xc6, 0x0a, 0xe9, 0xb4, 0x69, 0x48, 0x0e, 0xcb, 0xb7, 0x07, 0x23, 0xbf, 0xdd, 0xf4,
	0xf5, 0x5e, 0x75, 0x34, 0xca, 0xa3, 0xf5, 0x23, 0x76, 0x2c, 0x91, 0x33, 0x38, 0x45, 0xe4, 0xba,
	0x56, 0x34, 0xce, 0xa3, 0x75, 0xca, 0x7e, 0x0b, 0xe4, 0x0d, 0xc4, 0x7c, 0xa0, 0x49, 0x1e, 0xad,
	0x97, 0x97, 0xcf, 0x0b, 0xe5, 0xaa, 0xe2, 0x5f, 0x5b, 0x15, 0x7c, 0xe0, 0x77, 0xad, 0x62, 0x31,
	0x1f, 0xc8, 0x06, 0x16, 0x4c, 0x49, 0xa5, 0x5b, 0x47, 0x53, 0xb4, 0xbe, 0xf8, 0xbf, 0x35, 0x0c,
	0xa3, 0x7f, 0x72, 0xae, 0x7e, 0x46, 0x30, 0x1f, 0xbf, 0x49, 0xce, 0x21, 0xbb, 0x92, 0xd2, 0xf4,
	0x8d, 0xbb, 0x31, 0x8d, 0x54, 0x58, 0x23, 0x65, 0x7f, 0x68, 0x64, 0x05, 0x27, 0xd7, 0xc2, 0xee,
	0x3a, 0x2d, 0xc7, 0x1a, 0x19, 0xbb, 0xe7, 0xf0, 0x6e, 0xab, 0x6b, 0xed, 0xb0, 0x4b, 0xca, 0xee,
	0x99, 0x3c, 0x85, 0xd9, 0x57, 0x71, 0xe8, 0x15, 0x26, 0xcd, 0xd8, 0x08, 0x84, 0xc2, 0x62, 0x27,
	0xee, 0x0e, 0x46, 0x94, 0x74, 0x86, 0xfa, 0x84, 0x84, 0x40, 0xfa, 0x41, 0xd8, 0x8a, 0xce, 0x51,
	0xc6, 0x35, 0x79, 0x0c, 0x31, 0x37, 0x74, 0x81, 0x4a, 0xcc, 0x8d, 0x9f, 0x79, 0xdf, 0x99, 0x9a,
	0x9e, 0x8c, 0x33, 0x7e, 0x4d, 0x5e, 0xc2, 0x93, 0xa3, 0xca, 0x1f, 0x9b, 0x52, 0x0d, 0xf4, 0x14,
	0xaf, 0xe3, 0x2f, 0x7d, 0xf5, 0x23, 0x82, 0xe5, 0xd1, 0x99, 0xf8, 0x34, 0xd7, 0xc2, 0x7e, 0xb1,
	0xaa, 0xc4, 0xea, 0x19, 0x9b, 0x90, 0x3c, 0x83, 0xf9, 0x67, 0x27, 0x5c, 0x6f, 0x43, 0xe7, 0x40,
	0x64, 0x03, 0xc9, 0xd6, 0xdc, 0xd2, 0x24, 0x4f, 0xd6, 0xcb, 0xcb, 0x57, 0x0f, 0x3e, 0xfd, 0x62,
	0x6b, 0x6e, 0xf1, 0x16, 0xbc, 0x7b, 0xf5, 0x09, 0x16, 0x81, 0x7d, 0x82, 0xab, 0xb2, 0xec, 0x94,
	0xb5, 0x53, 0x82, 0x80, 0xbe, 0xeb, 0x3b, 0xe1, 0x44, 0xd8, 0x1f, 0xd7, 0x3e, 0x15, 0x37, 0xad,
	0x96, 0x16, 0x03, 0x64, 0x2c, 0xd0, 0x7e, 0x8e, 0xbf, 0xed, 0xeb, 0x5f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xc2, 0x69, 0x8d, 0xdf, 0xd6, 0x02, 0x00, 0x00,
}
