// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.5.0
//  VPP:              23.02-rc0~230-g1d9780a43~b3131
// source: /usr/share/vpp/api/plugins/adl.api.json

// Package adl contains generated bindings for API file adl.api.
//
// Contents:
//
//	4 messages
package adl

import (
	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	interface_types "github.com/jalapeno/go-vpp-sr/vppbinapi/interface_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "adl"
	APIVersion = "0.0.1"
	VersionCrc = 0xb752b7a2
)

// AdlAllowlistEnableDisable defines message 'adl_allowlist_enable_disable'.
// InProgress: the message form may change in the future versions
type AdlAllowlistEnableDisable struct {
	SwIfIndex  interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	FibID      uint32                         `binapi:"u32,name=fib_id" json:"fib_id,omitempty"`
	IP4        bool                           `binapi:"bool,name=ip4" json:"ip4,omitempty"`
	IP6        bool                           `binapi:"bool,name=ip6" json:"ip6,omitempty"`
	DefaultAdl bool                           `binapi:"bool,name=default_adl" json:"default_adl,omitempty"`
}

func (m *AdlAllowlistEnableDisable) Reset()               { *m = AdlAllowlistEnableDisable{} }
func (*AdlAllowlistEnableDisable) GetMessageName() string { return "adl_allowlist_enable_disable" }
func (*AdlAllowlistEnableDisable) GetCrcString() string   { return "ea88828d" }
func (*AdlAllowlistEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *AdlAllowlistEnableDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 4 // m.FibID
	size += 1 // m.IP4
	size += 1 // m.IP6
	size += 1 // m.DefaultAdl
	return size
}
func (m *AdlAllowlistEnableDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint32(m.FibID)
	buf.EncodeBool(m.IP4)
	buf.EncodeBool(m.IP6)
	buf.EncodeBool(m.DefaultAdl)
	return buf.Bytes(), nil
}
func (m *AdlAllowlistEnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.FibID = buf.DecodeUint32()
	m.IP4 = buf.DecodeBool()
	m.IP6 = buf.DecodeBool()
	m.DefaultAdl = buf.DecodeBool()
	return nil
}

// AdlAllowlistEnableDisableReply defines message 'adl_allowlist_enable_disable_reply'.
// InProgress: the message form may change in the future versions
type AdlAllowlistEnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *AdlAllowlistEnableDisableReply) Reset() { *m = AdlAllowlistEnableDisableReply{} }
func (*AdlAllowlistEnableDisableReply) GetMessageName() string {
	return "adl_allowlist_enable_disable_reply"
}
func (*AdlAllowlistEnableDisableReply) GetCrcString() string { return "e8d4e804" }
func (*AdlAllowlistEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *AdlAllowlistEnableDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *AdlAllowlistEnableDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *AdlAllowlistEnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// AdlInterfaceEnableDisable defines message 'adl_interface_enable_disable'.
// InProgress: the message form may change in the future versions
type AdlInterfaceEnableDisable struct {
	SwIfIndex     interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	EnableDisable bool                           `binapi:"bool,name=enable_disable" json:"enable_disable,omitempty"`
}

func (m *AdlInterfaceEnableDisable) Reset()               { *m = AdlInterfaceEnableDisable{} }
func (*AdlInterfaceEnableDisable) GetMessageName() string { return "adl_interface_enable_disable" }
func (*AdlInterfaceEnableDisable) GetCrcString() string   { return "5501adee" }
func (*AdlInterfaceEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *AdlInterfaceEnableDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 1 // m.EnableDisable
	return size
}
func (m *AdlInterfaceEnableDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBool(m.EnableDisable)
	return buf.Bytes(), nil
}
func (m *AdlInterfaceEnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.EnableDisable = buf.DecodeBool()
	return nil
}

// AdlInterfaceEnableDisableReply defines message 'adl_interface_enable_disable_reply'.
// InProgress: the message form may change in the future versions
type AdlInterfaceEnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *AdlInterfaceEnableDisableReply) Reset() { *m = AdlInterfaceEnableDisableReply{} }
func (*AdlInterfaceEnableDisableReply) GetMessageName() string {
	return "adl_interface_enable_disable_reply"
}
func (*AdlInterfaceEnableDisableReply) GetCrcString() string { return "e8d4e804" }
func (*AdlInterfaceEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *AdlInterfaceEnableDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *AdlInterfaceEnableDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *AdlInterfaceEnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_adl_binapi_init() }
func file_adl_binapi_init() {
	api.RegisterMessage((*AdlAllowlistEnableDisable)(nil), "adl_allowlist_enable_disable_ea88828d")
	api.RegisterMessage((*AdlAllowlistEnableDisableReply)(nil), "adl_allowlist_enable_disable_reply_e8d4e804")
	api.RegisterMessage((*AdlInterfaceEnableDisable)(nil), "adl_interface_enable_disable_5501adee")
	api.RegisterMessage((*AdlInterfaceEnableDisableReply)(nil), "adl_interface_enable_disable_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*AdlAllowlistEnableDisable)(nil),
		(*AdlAllowlistEnableDisableReply)(nil),
		(*AdlInterfaceEnableDisable)(nil),
		(*AdlInterfaceEnableDisableReply)(nil),
	}
}
