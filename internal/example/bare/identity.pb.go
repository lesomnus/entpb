// Code generated by "protoc-gen-entpb". DO NOT EDIT

package bare

import (
	context "context"
	uuid "github.com/google/uuid"
	ent "github.com/lesomnus/entpb/internal/example/ent"
	identity "github.com/lesomnus/entpb/internal/example/ent/identity"
	predicate "github.com/lesomnus/entpb/internal/example/ent/predicate"
	user "github.com/lesomnus/entpb/internal/example/ent/user"
	pb "github.com/lesomnus/entpb/internal/example/pb"
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
func (s *IdentityServiceServer) Create(ctx context.Context, req *pb.CreateIdentityRequest) (*pb.Identity, error) {
	q := s.db.Identity.Create()
	if v := req.Name; v != nil {
		q.SetName(*v)
	}
	if v := req.Email; v != nil {
		q.SetEmail(*v)
	}
	if v := req.GetDateUpdated(); v != nil {
		w := v.AsTime()
		q.SetDateUpdated(w)
	}
	if id, err := GetUserId(ctx, s.db, req.GetOwner()); err != nil {
		return nil, err
	} else {
		q.SetOwnerID(id)
	}

	res, err := q.Save(ctx)
	if err != nil {
		return nil, ToStatus(err)
	}

	return ToProtoIdentity(res), nil
}
func (s *IdentityServiceServer) Delete(ctx context.Context, req *pb.GetIdentityRequest) (*emptypb.Empty, error) {
	q := s.db.Identity.Delete()
	if v, err := uuid.FromBytes(req.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "id: %s", err)
	} else {
		q.Where(identity.IDEQ(v))
	}

	_, err := q.Exec(ctx)
	if err != nil {
		return nil, ToStatus(err)
	}

	return &emptypb.Empty{}, nil
}
func (s *IdentityServiceServer) Get(ctx context.Context, req *pb.GetIdentityRequest) (*pb.Identity, error) {
	q := s.db.Identity.Query()
	if p, err := GetIdentitySpecifier(req); err != nil {
		return nil, err
	} else {
		q.Where(p)
	}

	q.WithOwner(func(q *ent.UserQuery) { q.Select(user.FieldID) })

	res, err := q.Only(ctx)
	if err != nil {
		return nil, ToStatus(err)
	}

	return ToProtoIdentity(res), nil
}
func (s *IdentityServiceServer) Update(ctx context.Context, req *pb.UpdateIdentityRequest) (*pb.Identity, error) {
	id, err := GetIdentityId(ctx, s.db, req.GetKey())
	if err != nil {
		return nil, err
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
		return nil, ToStatus(err)
	}

	return ToProtoIdentity(res), nil
}
func ToProtoIdentity(v *ent.Identity) *pb.Identity {
	m := &pb.Identity{}
	m.Id = v.ID[:]
	m.DateCreated = timestamppb.New(v.DateCreated)
	m.Name = v.Name
	m.Email = v.Email
	if v.DateUpdated != nil {
		m.DateUpdated = timestamppb.New(*v.DateUpdated)
	}
	if v := v.Edges.Owner; v != nil {
		m.Owner = ToProtoUser(v)
	}
	return m
}
func GetIdentityId(ctx context.Context, db *ent.Client, req *pb.GetIdentityRequest) (uuid.UUID, error) {
	var r uuid.UUID
	if v, err := uuid.FromBytes(req.GetId()); err != nil {
		return r, status.Errorf(codes.InvalidArgument, "id: %s", err)
	} else {
		return v, nil
	}
}
func GetIdentitySpecifier(req *pb.GetIdentityRequest) (predicate.Identity, error) {
	if v, err := uuid.FromBytes(req.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "id: %s", err)
	} else {
		return identity.IDEQ(v), nil
	}
}
