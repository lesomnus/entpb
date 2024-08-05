package pbgen

import (
	"github.com/lesomnus/entpb/pbgen/ident"
)

type Type = ident.Full

var (
	TypeDouble   = Type{Segments: []string{"double"}}
	TypeFloat    = Type{Segments: []string{"float"}}
	TypeInt32    = Type{Segments: []string{"int32"}}
	TypeInt64    = Type{Segments: []string{"int64"}}
	TypeUint32   = Type{Segments: []string{"uint32"}}
	TypeUint64   = Type{Segments: []string{"uint64"}}
	TypeSint32   = Type{Segments: []string{"sint32"}}
	TypeSint64   = Type{Segments: []string{"sint64"}}
	TypeFixed32  = Type{Segments: []string{"fixed32"}}
	TypeFixed64  = Type{Segments: []string{"fixed64"}}
	TypeSfixed32 = Type{Segments: []string{"sfixed32"}}
	TypeSfixed64 = Type{Segments: []string{"sfixed64"}}
	TypeBool     = Type{Segments: []string{"bool"}}
	TypeString   = Type{Segments: []string{"string"}}
	TypeBytes    = Type{Segments: []string{"bytes"}}
)

type Label string

const (
	LabelRequired = Label("required")
	LabelOptional = Label("optional")
	LabelRepeated = Label("repeated")
)

type Visibility string

const (
	VisibilityWeak   = Visibility("weak")
	VisibilityPublic = Visibility("public")
)

type Edition struct {
	Keyword string
	Value   string
}

var (
	SyntaxProto2 = Edition{"syntax", "proto2"}
	SyntaxProto3 = Edition{"syntax", "proto3"}
	Edition2023  = Edition{"edition", "2023"}
)

type ProtoFile struct {
	Edition Edition
	Package ident.Full
	Imports []Import
	Options []Option

	TopLevelDefinitions []TopLevelDef
}

func (ProtoFile) TemplateName() string {
	return "proto-file"
}

type Import struct {
	Name       string
	Visibility Visibility
}

type Enum struct {
	Name    ident.Ident
	Options []Option
	Body    []EnumBody

	topLevelDef_
	messageBody_
}

func (Enum) TemplateName() string {
	return "enum"
}

type EnumBody interface{ enumBody() }
type enumBody_ struct{}

func (enumBody_) enumBody() {}

type EnumField struct {
	Name    string
	Number  int
	Options []Option

	enumBody_
}

func (EnumField) TemplateName() string {
	return "enum-field"
}

type Message struct {
	Name ident.Ident
	Body []MessageBody

	topLevelDef_
	messageBody_
}

func (Message) TemplateName() string {
	return "message"
}

type MessageBody interface{ messageBody() }
type messageBody_ struct{}

func (messageBody_) messageBody() {}

type MessageField struct {
	Labels  []Label
	Type    Type
	Name    ident.Ident
	Number  int
	Options []Option

	messageBody_
}

func (MessageField) TemplateName() string {
	return "message-field"
}

type MessageOneof struct {
	Name    ident.Ident
	Options []Option
	Body    []MessageOneofBody

	messageBody_
}

func (MessageOneof) TemplateName() string {
	return "message-oneof"
}

type MessageOneofBody interface{ messageOneofBody() }
type messageOneofBody_ struct{}

func (messageOneofBody_) messageOneofBody() {}

type MessageOneofField struct {
	Type    Type
	Name    ident.Ident
	Number  int
	Options []Option

	messageOneofBody_
}

func (MessageOneofField) TemplateName() string {
	return "message-oneof-field"
}

type Service struct {
	Name ident.Ident
	Body []ServiceBody

	topLevelDef_
}

func (Service) TemplateName() string {
	return "service"
}

type ServiceBody interface{ serviceBody() }
type serviceBody_ struct{}

func (serviceBody_) serviceBody() {}

type Rpc struct {
	Name     ident.Ident
	Request  RpcType
	Response RpcType

	Options []Option

	serviceBody_
}

func (Rpc) TemplateName() string {
	return "rpc"
}

type RpcType struct {
	Type
	Stream bool
}

type OptionValue interface{ optionValue() }
type optionValue_ struct{}

func (optionValue_) optionValue() {}

type Option struct {
	Name  ident.Full
	Value OptionValue
}

type Comment struct {
	Value     string
	Multiline bool

	topLevelDef_
	enumBody_
	messageBody_
	messageOneofBody_
	serviceBody_
}

func (Comment) TemplateName() string {
	return "comment"
}

type TopLevelDef interface{ topLevelDef() }
type topLevelDef_ struct{}

func (topLevelDef_) topLevelDef() {}

type Data struct {
	Fields []DataField

	dataValue_
}

func (Data) TemplateName() string {
	return "data"
}

type DataList struct {
	Values []Data

	dataValue_
}

func (DataList) TemplateName() string {
	return "data-list"
}

type DataField struct {
	Name  string
	Value DataValue
}

type DataValue interface{ dataValue() }
type dataValue_ struct {
	optionValue_
}

func (dataValue_) dataValue() {}

type DataString struct {
	Value string

	dataValue_
}

func (DataString) TemplateName() string {
	return "data-string"
}

type UnsafeLiteral struct {
	Value string

	dataValue_
}

func (UnsafeLiteral) TemplateName() string {
	return "unsafe-literal"
}
