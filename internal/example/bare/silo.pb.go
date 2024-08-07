// Code generated by "protoc-gen-entpb". DO NOT EDIT

package bare

import (
	context "context"
	uuid "github.com/google/uuid"
	ent "github.com/lesomnus/entpb/internal/example/ent"
	predicate "github.com/lesomnus/entpb/internal/example/ent/predicate"
	silo "github.com/lesomnus/entpb/internal/example/ent/silo"
	pb "github.com/lesomnus/entpb/internal/example/pb"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	strings "strings"
)

type SiloServiceServer struct {
	db *ent.Client
	pb.UnimplementedSiloServiceServer
}

func NewSiloServiceServer(db *ent.Client) *SiloServiceServer {
	return &SiloServiceServer{db: db}
}
func (s *SiloServiceServer) Create(ctx context.Context, req *pb.CreateSiloRequest) (*pb.Silo, error) {
	q := s.db.Silo.Create()
	if v := req.Alias; v != nil {
		q.SetAlias(*v)
	}
	if v := req.Name; v != nil {
		q.SetName(*v)
	}
	if v := req.Description; v != nil {
		q.SetDescription(*v)
	}

	res, err := q.Save(ctx)
	if err != nil {
		return nil, ToStatus(err)
	}

	return ToProtoSilo(res), nil
}
func (s *SiloServiceServer) Delete(ctx context.Context, req *pb.GetSiloRequest) (*emptypb.Empty, error) {
	p, err := GetSiloSpecifier(req)
	if err != nil {
		return nil, err
	}
	if _, err := s.db.Silo.Delete().Where(p).Exec(ctx); err != nil {
		return nil, ToStatus(err)
	}

	return &emptypb.Empty{}, nil
}
func (s *SiloServiceServer) Get(ctx context.Context, req *pb.GetSiloRequest) (*pb.Silo, error) {
	q := s.db.Silo.Query()
	if p, err := GetSiloSpecifier(req); err != nil {
		return nil, err
	} else {
		q.Where(p)
	}

	res, err := QuerySiloWithEdgeIds(q).Only(ctx)
	if err != nil {
		return nil, ToStatus(err)
	}

	return ToProtoSilo(res), nil
}
func QuerySiloWithEdgeIds(q *ent.SiloQuery) *ent.SiloQuery {

	return q
}
func (s *SiloServiceServer) Update(ctx context.Context, req *pb.UpdateSiloRequest) (*pb.Silo, error) {
	id, err := GetSiloId(ctx, s.db, req.GetKey())
	if err != nil {
		return nil, err
	}

	q := s.db.Silo.UpdateOneID(id)
	if v := req.Alias; v != nil {
		q.SetAlias(*v)
	}
	if v := req.Name; v != nil {
		q.SetName(*v)
	}
	if v := req.Description; v != nil {
		q.SetDescription(*v)
	}

	res, err := q.Save(ctx)
	if err != nil {
		return nil, ToStatus(err)
	}

	return ToProtoSilo(res), nil
}
func ToProtoSilo(v *ent.Silo) *pb.Silo {
	m := &pb.Silo{}
	m.Id = v.ID[:]
	m.DateCreated = timestamppb.New(v.DateCreated)
	m.Alias = v.Alias
	m.Name = v.Name
	m.Description = v.Description
	return m
}
func GetSiloId(ctx context.Context, db *ent.Client, req *pb.GetSiloRequest) (uuid.UUID, error) {
	var r uuid.UUID
	k := req.GetKey()
	if t, ok := k.(*pb.GetSiloRequest_Id); ok {
		if v, err := uuid.FromBytes(t.Id); err != nil {
			return r, status.Errorf(codes.InvalidArgument, "id: %s", err)
		} else {
			return v, nil
		}
	}

	p, err := GetSiloSpecifier(req)
	if err != nil {
		return r, err
	}

	v, err := db.Silo.Query().Where(p).OnlyID(ctx)
	if err != nil {
		return r, ToStatus(err)
	}

	return v, nil
}

func GetSiloSpecifier(req *pb.GetSiloRequest) (predicate.Silo, error) {
	switch t := req.GetKey().(type) {
	case *pb.GetSiloRequest_Id:
		if v, err := uuid.FromBytes(t.Id); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "id: %s", err)
		} else {
			return silo.IDEQ(v), nil
		}
	case *pb.GetSiloRequest_Alias:
		return silo.AliasEQ(t.Alias), nil
	case *pb.GetSiloRequest_Query:
		if req, err := ResolveGetSiloQuery(req); err != nil {
			return nil, err
		} else {
			return GetSiloSpecifier(req)
		}
	case nil:
		return nil, status.Errorf(codes.InvalidArgument, "key not provided")
	default:
		return nil, status.Errorf(codes.Unimplemented, "unknown type of key")
	}
}

func ResolveGetSiloQuery(req *pb.GetSiloRequest) (*pb.GetSiloRequest, error) {
	t, ok := req.Key.(*pb.GetSiloRequest_Query)
	if !ok {
		return req, nil
	}

	q := t.Query

	if v, ok := strings.CutPrefix(q, "@"); ok {
		return pb.SiloByAlias(v), nil
	}
	v, err := uuid.Parse(q)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid query string: %s", err)
	}
	return pb.SiloById(v), nil
}
