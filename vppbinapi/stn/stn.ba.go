// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.5.0
//  VPP:              23.02-rc0~230-g1d9780a43~b3131
// source: /usr/share/vpp/api/plugins/stn.api.json

// Package stn contains generated bindings for API file stn.api.
//
// Contents:
//
//	4 messages
package stn

import (
	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	interface_types "github.com/jalapeno/go-vpp-sr/vppbinapi/interface_types"
	ip_types "github.com/jalapeno/go-vpp-sr/vppbinapi/ip_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "stn"
	APIVersion = "2.0.0"
	VersionCrc = 0x9cfaef64
)

// StnAddDelRule defines message 'stn_add_del_rule'.
type StnAddDelRule struct {
	IPAddress ip_types.Address               `binapi:"address,name=ip_address" json:"ip_address,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	IsAdd     bool                           `binapi:"bool,name=is_add" json:"is_add,omitempty"`
}

func (m *StnAddDelRule) Reset()               { *m = StnAddDelRule{} }
func (*StnAddDelRule) GetMessageName() string { return "stn_add_del_rule" }
func (*StnAddDelRule) GetCrcString() string   { return "224c6edd" }
func (*StnAddDelRule) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *StnAddDelRule) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.IPAddress.Af
	size += 1 * 16 // m.IPAddress.Un
	size += 4      // m.SwIfIndex
	size += 1      // m.IsAdd
	return size
}
func (m *StnAddDelRule) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.IPAddress.Af))
	buf.EncodeBytes(m.IPAddress.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBool(m.IsAdd)
	return buf.Bytes(), nil
}
func (m *StnAddDelRule) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IPAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.IPAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.IsAdd = buf.DecodeBool()
	return nil
}

// StnAddDelRuleReply defines message 'stn_add_del_rule_reply'.
type StnAddDelRuleReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *StnAddDelRuleReply) Reset()               { *m = StnAddDelRuleReply{} }
func (*StnAddDelRuleReply) GetMessageName() string { return "stn_add_del_rule_reply" }
func (*StnAddDelRuleReply) GetCrcString() string   { return "e8d4e804" }
func (*StnAddDelRuleReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *StnAddDelRuleReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *StnAddDelRuleReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *StnAddDelRuleReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// StnRulesDetails defines message 'stn_rules_details'.
type StnRulesDetails struct {
	IPAddress ip_types.Address               `binapi:"address,name=ip_address" json:"ip_address,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *StnRulesDetails) Reset()               { *m = StnRulesDetails{} }
func (*StnRulesDetails) GetMessageName() string { return "stn_rules_details" }
func (*StnRulesDetails) GetCrcString() string   { return "a51935a6" }
func (*StnRulesDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *StnRulesDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.IPAddress.Af
	size += 1 * 16 // m.IPAddress.Un
	size += 4      // m.SwIfIndex
	return size
}
func (m *StnRulesDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.IPAddress.Af))
	buf.EncodeBytes(m.IPAddress.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *StnRulesDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IPAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.IPAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// StnRulesDump defines message 'stn_rules_dump'.
type StnRulesDump struct{}

func (m *StnRulesDump) Reset()               { *m = StnRulesDump{} }
func (*StnRulesDump) GetMessageName() string { return "stn_rules_dump" }
func (*StnRulesDump) GetCrcString() string   { return "51077d14" }
func (*StnRulesDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *StnRulesDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *StnRulesDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *StnRulesDump) Unmarshal(b []byte) error {
	return nil
}

func init() { file_stn_binapi_init() }
func file_stn_binapi_init() {
	api.RegisterMessage((*StnAddDelRule)(nil), "stn_add_del_rule_224c6edd")
	api.RegisterMessage((*StnAddDelRuleReply)(nil), "stn_add_del_rule_reply_e8d4e804")
	api.RegisterMessage((*StnRulesDetails)(nil), "stn_rules_details_a51935a6")
	api.RegisterMessage((*StnRulesDump)(nil), "stn_rules_dump_51077d14")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*StnAddDelRule)(nil),
		(*StnAddDelRuleReply)(nil),
		(*StnRulesDetails)(nil),
		(*StnRulesDump)(nil),
	}
}
