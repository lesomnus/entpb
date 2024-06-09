// Code generated by "protoc-gen-entpb". DO NOT EDIT

package bare

import (
	context "context"
	uuid "github.com/google/uuid"
	runtime "github.com/lesomnus/entpb/cmd/protoc-gen-entpb/runtime"
	ent "github.com/lesomnus/entpb/example/ent"
	identity "github.com/lesomnus/entpb/example/ent/identity"
	pb "github.com/lesomnus/entpb/example/pb"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type IdentityServiceServer struct {
	db *ent.Client
	pb.UnimplementedIdentityServiceServer
}

func NewIdentityServiceServer(db *ent.Client) *IdentityServiceServer {
	return &IdentityServiceServer{db: db}
}
func (s *IdentityServiceServer) Create(ctx context.Context, req *pb.Identity) (*pb.Identity, error) {
	q := s.db.Identity.Create()
	q.SetName(req.Name)
	if v := req.Email; v != nil {
		q.SetEmail(*v)
	}
	if v := req.DateUpdated; v != nil {
		w := v.AsTime()
		q.SetDateUpdated(w)
	}

	res, err := q.Save(ctx)
	if err != nil {
		return nil, runtime.EntErrorToStatus(err)
	}

	return ToProtoIdentity(res), nil
}
func (s *IdentityServiceServer) Delete(ctx context.Context, req *pb.DeleteIdentityRequest) (*emptypb.Empty, error) {
	q := s.db.Identity.Delete()
	if v, err := uuid.FromBytes(req.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "id: %s", err)
	} else {
		q.Where(identity.IDEQ(v))
	}

	_, err := q.Exec(ctx)
	if err != nil {
		return nil, runtime.EntErrorToStatus(err)
	}

	return &emptypb.Empty{}, nil
}
func (s *IdentityServiceServer) Get(ctx context.Context, req *pb.GetIdentityRequest) (*pb.Identity, error) {
	q := s.db.Identity.Query()
	if v, err := uuid.FromBytes(req.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "id: %s", err)
	} else {
		q.Where(identity.IDEQ(v))
	}
	q.WithOwner(func(q *ent.UserQuery) { q.Select(identity.FieldID) })

	res, err := q.Only(ctx)
	if err != nil {
		return nil, runtime.EntErrorToStatus(err)
	}

	return ToProtoIdentity(res), nil
}
func (s *IdentityServiceServer) Update(ctx context.Context, req *pb.UpdateIdentityRequest) (*pb.Identity, error) {
	id, err := uuid.FromBytes(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "id: %s", err.Error())
	}

	q := s.db.Identity.UpdateOneID(id)
	if v := req.Name; v != nil {
		q.SetName(*v)
	}
	if v := req.Email; v != nil {
		q.SetEmail(*v)
	}
	if v := req.DateUpdated; v != nil {
		w := v.AsTime()
		q.SetDateUpdated(w)
	}

	res, err := q.Save(ctx)
	if err != nil {
		return nil, runtime.EntErrorToStatus(err)
	}

	return ToProtoIdentity(res), nil
}
func ToProtoIdentity(v *ent.Identity) *pb.Identity {
	m := &pb.Identity{}
	m.Id = v.ID[:]
	m.DateCreated = timestamppb.New(v.DateCreated)
	m.Name = v.Name
	m.Email = v.Email
	m.DateUpdated = timestamppb.New(v.DateUpdated)
	if v := v.Edges.Owner; v != nil {
		m.Owner = &pb.Actor{Id: v.ID[:]}
	}
	return m
}
