// Code generated by "protoc-gen-entpb". DO NOT EDIT

package bare

import (
	context "context"
	uuid "github.com/google/uuid"
	runtime "github.com/lesomnus/entpb/cmd/protoc-gen-entpb/runtime"
	ent "github.com/lesomnus/entpb/example/ent"
	pb "github.com/lesomnus/entpb/example/pb"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type ActorServiceServer struct {
	db *ent.Client
	pb.UnimplementedActorServiceServer
}

func (s *ActorServiceServer) Create(ctx context.Context, req *pb.Actor) (*pb.Actor, error) {
	q := s.db.User.Create()
	if v, err := uuid.FromBytes(req.Referer.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "referer: %s", err)
	} else {
		q.SetParentID(v)
	}
	for _, v := range req.Identities {
		if w, err := uuid.FromBytes(v.GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "identities: %s", err)
		} else {
			q.AddIdentityIDs(w)
		}
	}
	for _, v := range req.Children {
		if w, err := uuid.FromBytes(v.GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "children: %s", err)
		} else {
			q.AddChildIDs(w)
		}
	}
	res, err := q.Save(ctx)
	if err != nil {
		return nil, runtime.EntErrorToStatus(err)
	}

	return toProtoUser(res), nil
}
func toProtoUser(v *ent.User) *pb.Actor {
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
