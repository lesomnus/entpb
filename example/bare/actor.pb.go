// Code generated by "protoc-gen-entpb". DO NOT EDIT

package bare

import (
	context "context"
	uuid "github.com/google/uuid"
	runtime "github.com/lesomnus/entpb/cmd/protoc-gen-entpb/runtime"
	ent "github.com/lesomnus/entpb/example/ent"
	user "github.com/lesomnus/entpb/example/ent/user"
	pb "github.com/lesomnus/entpb/example/pb"
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
func (s *ActorServiceServer) Create(ctx context.Context, req *pb.Actor) (*pb.Actor, error) {
	q := s.db.User.Create()

	res, err := q.Save(ctx)
	if err != nil {
		return nil, runtime.EntErrorToStatus(err)
	}

	return ToProtoUser(res), nil
}
func (s *ActorServiceServer) Delete(ctx context.Context, req *pb.DeleteActorRequest) (*emptypb.Empty, error) {
	q := s.db.User.Delete()
	if v, err := uuid.FromBytes(req.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "id: %s", err)
	} else {
		q.Where(user.IDEQ(v))
	}

	_, err := q.Exec(ctx)
	if err != nil {
		return nil, runtime.EntErrorToStatus(err)
	}

	return &emptypb.Empty{}, nil
}
func (s *ActorServiceServer) Get(ctx context.Context, req *pb.GetActorRequest) (*pb.Actor, error) {
	q := s.db.User.Query()
	if v, err := uuid.FromBytes(req.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "id: %s", err)
	} else {
		q.Where(user.IDEQ(v))
	}
	q.WithParent(func(q *ent.UserQuery) { q.Select(user.FieldID) })
	q.WithIdentities(func(q *ent.IdentityQuery) { q.Select(user.FieldID) })
	q.WithChildren(func(q *ent.UserQuery) { q.Select(user.FieldID) })

	res, err := q.Only(ctx)
	if err != nil {
		return nil, runtime.EntErrorToStatus(err)
	}

	return ToProtoUser(res), nil
}
func (s *ActorServiceServer) Update(ctx context.Context, req *pb.UpdateActorRequest) (*pb.Actor, error) {
	id, err := uuid.FromBytes(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "id: %s", err.Error())
	}

	q := s.db.User.UpdateOneID(id)

	res, err := q.Save(ctx)
	if err != nil {
		return nil, runtime.EntErrorToStatus(err)
	}

	return ToProtoUser(res), nil
}
func ToProtoUser(v *ent.User) *pb.Actor {
	m := &pb.Actor{}
	m.Id = v.ID[:]
	m.DateCreated = timestamppb.New(v.DateCreated)
	if v := v.Edges.Parent; v != nil {
		m.Referer = &pb.Actor{Id: v.ID[:]}
	}
	for _, v := range v.Edges.Identities {
		m.Identities = append(m.Identities, &pb.Identity{Id: v.ID[:]})
	}
	for _, v := range v.Edges.Children {
		m.Children = append(m.Children, &pb.Actor{Id: v.ID[:]})
	}
	return m
}
