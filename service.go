package entpb

import (
	"github.com/lesomnus/entpb/pbgen/ident"
)

type ServiceOption interface {
	serviceOpt(*Service)
}

type Service struct {
	Filepath string
	Ident    ident.Ident
	Rpcs     map[ident.Ident]*Rpc

	File    *ProtoFile
	Message *MessageAnnotation // Message defined by Ent schema which defines this service.
}

func (s *Service) messageOpt(t *MessageAnnotation) {
	t.Service = s
}

func WithService(filepath string, opts ...ServiceOption) MessageOption {
	s := &Service{
		Filepath: filepath,

		Rpcs: map[ident.Ident]*Rpc{},
	}
	for _, opt := range opts {
		opt.serviceOpt(s)
	}
	return s
}

type Rpc struct {
	Ident   ident.Ident
	Comment string

	Req    PbType
	Res    PbType
	Stream Stream

	EntReq *MessageAnnotation
	EntRes *MessageAnnotation
}

func (r *Rpc) serviceOpt(t *Service) {
	t.Rpcs[r.Ident] = r
}

type Stream int

const (
	StreamNone Stream = iota
	StreamClient
	StreamServer
	StreamBoth
)

func RpcEntCreate() *Rpc {
	return &Rpc{Ident: "Create"}
}

func RpcEntGet() *Rpc {
	return &Rpc{Ident: "Get"}
}

func RpcEntUpdate() *Rpc {
	return &Rpc{Ident: "Update"}
}

func RpcEntDelete() *Rpc {
	return &Rpc{Ident: "Delete"}
}
