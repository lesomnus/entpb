package entpb

import (
	ent "entgo.io/ent/schema/field"
)

type PbType struct {
	Name   string // Type name
	Import string // Import path for this type.
}

var (
	PbUuid      = PbType{Name: "bytes"}
	PbEmpty     = PbType{Name: "google.protobuf.Empty", Import: "google/protobuf/empty.proto"}
	PbTimestamp = PbType{Name: "google.protobuf.Timestamp", Import: "google/protobuf/timestamp.proto"}
)

var pb_types = [...]PbType{
	ent.TypeBool:    {Name: "bool"},
	ent.TypeInt8:    {Name: "sint32"},
	ent.TypeInt16:   {Name: "sint32"},
	ent.TypeInt32:   {Name: "sint32"},
	ent.TypeInt:     {Name: "sint64"},
	ent.TypeInt64:   {Name: "sint64"},
	ent.TypeUint8:   {Name: "uint32"},
	ent.TypeUint16:  {Name: "uint32"},
	ent.TypeUint32:  {Name: "uint32"},
	ent.TypeUint:    {Name: "uint64"},
	ent.TypeUint64:  {Name: "uint64"},
	ent.TypeFloat32: {Name: "float"},
	ent.TypeFloat64: {Name: "double"},
	ent.TypeBytes:   {Name: "bytes"},
	ent.TypeString:  {Name: "string"},

	ent.TypeUUID: PbUuid,
	ent.TypeTime: PbTimestamp,
}
