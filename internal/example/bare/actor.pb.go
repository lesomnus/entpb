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

type ActorServiceServer struct {
	db *ent.Client
	pb.UnimplementedActorServiceServer
}

func NewActorServiceServer(db *ent.Client) *ActorServiceServer {
	return &ActorServiceServer{db: db}
}
func (s *ActorServiceServer) Create(ctx context.Context, req *pb.CreateActorRequest) (*pb.Actor, error) {
	q := s.db.User.Create()
	if v := req.GetReferer(); v != nil {
		if id, err := GetUserId(ctx, s.db, v); err != nil {
			return nil, err
		} else {
			q.SetParentID(id)
		}
	}

	res, err := q.Save(ctx)
	if err != nil {
		return nil, ToStatus(err)
	}

	return ToProtoUser(res), nil
}
func (s *ActorServiceServer) Delete(ctx context.Context, req *pb.GetActorRequest) (*emptypb.Empty, error) {
	p, err := GetUserSpecifier(req)
	if err != nil {
		return nil, err
	}
	if _, err := s.db.User.Delete().Where(p).Exec(ctx); err != nil {
		return nil, ToStatus(err)
	}

	return &emptypb.Empty{}, nil
}
func (s *ActorServiceServer) Get(ctx context.Context, req *pb.GetActorRequest) (*pb.Actor, error) {
	q := s.db.User.Query()
	if p, err := GetUserSpecifier(req); err != nil {
		return nil, err
	} else {
		q.Where(p)
	}

	res, err := QueryUserWithEdgeIds(q).Only(ctx)
	if err != nil {
		return nil, ToStatus(err)
	}

	return ToProtoUser(res), nil
}
func QueryUserWithEdgeIds(q *ent.UserQuery) *ent.UserQuery {
	q.WithParent(func(q *ent.UserQuery) { q.Select(user.FieldID) })
	q.WithIdentities(func(q *ent.IdentityQuery) { q.Select(identity.FieldID) })
	q.WithChildren(func(q *ent.UserQuery) { q.Select(user.FieldID) })

	return q
}
func (s *ActorServiceServer) Update(ctx context.Context, req *pb.UpdateActorRequest) (*pb.Actor, error) {
	id, err := GetUserId(ctx, s.db, req.GetKey())
	if err != nil {
		return nil, err
	}

	q := s.db.User.UpdateOneID(id)
	if v := req.Referer; v != nil {
		if id, err := GetUserId(ctx, s.db, req.Referer); err != nil {
			return nil, err
		} else {
			q.SetParentID(id)
		}
	}

	res, err := q.Save(ctx)
	if err != nil {
		return nil, ToStatus(err)
	}

	return ToProtoUser(res), nil
}
func ToProtoUser(v *ent.User) *pb.Actor {
	m := &pb.Actor{}
	m.Id = v.ID[:]
	m.DateCreated = timestamppb.New(v.DateCreated)
	if v := v.Edges.Parent; v != nil {
		m.Referer = ToProtoUser(v)
	}
	for _, v := range v.Edges.Identities {
		m.Identities = append(m.Identities, ToProtoIdentity(v))
	}
	for _, v := range v.Edges.Children {
		m.Children = append(m.Children, ToProtoUser(v))
	}
	return m
}
func GetUserId(ctx context.Context, db *ent.Client, req *pb.GetActorRequest) (uuid.UUID, error) {
	var r uuid.UUID
	if v, err := uuid.FromBytes(req.GetId()); err != nil {
		return r, status.Errorf(codes.InvalidArgument, "id: %s", err)
	} else {
		return v, nil
	}
}
func GetUserSpecifier(req *pb.GetActorRequest) (predicate.User, error) {
	if v, err := uuid.FromBytes(req.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "id: %s", err)
	} else {
		return user.IDEQ(v), nil
	}
}
