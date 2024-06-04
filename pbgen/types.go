package pbgen

import (
	"github.com/lesomnus/entpb/pbgen/ident"
)

type Type = ident.Full

var (
	TypeDouble   = Type{"double"}
	TypeFloat    = Type{"float"}
	TypeInt32    = Type{"int32"}
	TypeInt64    = Type{"int64"}
	TypeUint32   = Type{"uint32"}
	TypeUint64   = Type{"uint64"}
	TypeSint32   = Type{"sint32"}
	TypeSint64   = Type{"sint64"}
	TypeFixed32  = Type{"fixed32"}
	TypeFixed64  = Type{"fixed64"}
	TypeSfixed32 = Type{"sfixed32"}
	TypeSfixed64 = Type{"sfixed64"}
	TypeBool     = Type{"bool"}
	TypeString   = Type{"string"}
	TypeBytes    = Type{"bytes"}
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

	serviceBody_
}

func (Rpc) TemplateName() string {
	return "rpc"
}

type RpcType struct {
	Type
	Stream bool
}

type Option struct {
	Name  ident.Full
	Value string
}

func (Option) TemplateName() string {
	return "option"
}

type Comment struct {
	Value     string
	Multiline bool

	topLevelDef_
	enumBody_
	messageBody_
	serviceBody_
}

func (Comment) TemplateName() string {
	return "comment"
}

type TopLevelDef interface{ topLevelDef() }
type topLevelDef_ struct{}

func (topLevelDef_) topLevelDef() {}
