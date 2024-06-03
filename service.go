package entpb

import (
	"github.com/lesomnus/entpb/pbgen/ident"
)

type service struct {
	Filepath string
	Name     ident.Ident
	Rpcs     map[ident.Ident]*Rpc

	// Marks key named RPC should be implemented.
	BuiltIn map[ident.Ident]rpc_built_in

	message *messageAnnotation
}

type Service struct {
	Filepath string      // If empty, filepath to message is set.
	Name     ident.Ident // If empty, "{message name}Service" is set; e.g. User -> UserService.
}

func (s *Service) messageOpt(a *messageAnnotation) {
	if a.Service == nil {
		a.Service = &service{
			Rpcs:    map[ident.Ident]*Rpc{},
			BuiltIn: map[ident.Ident]rpc_built_in{},
		}
	}

	a.Service.Filepath = s.Filepath
}

type Rpc struct {
	Name    ident.Ident
	Comment string

	Req    PbType
	Res    PbType
	Stream Stream

	// TODO: can it be separated into plugin?
	builtin rpc_built_in
}

func (r Rpc) NameAs(v ident.Ident) *Rpc {
	r.Name = v
	return &r
}

func (r *Rpc) messageOpt(a *messageAnnotation) {
	if a.Service == nil {
		a.Service = &service{
			Rpcs:    map[ident.Ident]*Rpc{},
			BuiltIn: map[ident.Ident]rpc_built_in{},
		}
	}

	a.Service.Rpcs[r.Name] = r
	if r.builtin != rpcBuiltInUnspecified {
		a.Service.BuiltIn[r.Name] = r.builtin
	}
}

type Stream int

const (
	StreamNone Stream = iota
	StreamClient
	StreamServer
	StreamBoth
)

func RpcEntCreate() *Rpc {
	return &Rpc{
		Name: "Create",
		Req:  PbThis,
		Res:  PbThis,

		builtin: rpcBuiltInCreate,
	}
}

type rpc_built_in int

const (
	rpcBuiltInUnspecified rpc_built_in = iota
	rpcBuiltInCreate
)
