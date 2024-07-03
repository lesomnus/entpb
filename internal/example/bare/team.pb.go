// Code generated by "protoc-gen-entpb". DO NOT EDIT

package bare

import (
	context "context"
	uuid "github.com/google/uuid"
	ent "github.com/lesomnus/entpb/internal/example/ent"
	predicate "github.com/lesomnus/entpb/internal/example/ent/predicate"
	silo "github.com/lesomnus/entpb/internal/example/ent/silo"
	team "github.com/lesomnus/entpb/internal/example/ent/team"
	pb "github.com/lesomnus/entpb/internal/example/pb"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type TeamServiceServer struct {
	db *ent.Client
	pb.UnimplementedTeamServiceServer
}

func NewTeamServiceServer(db *ent.Client) *TeamServiceServer {
	return &TeamServiceServer{db: db}
}
func (s *TeamServiceServer) Create(ctx context.Context, req *pb.CreateTeamRequest) (*pb.Team, error) {
	q := s.db.Team.Create()
	if v := req.Alias; v != nil {
		q.SetAlias(*v)
	}
	if v := req.Name; v != nil {
		q.SetName(*v)
	}
	if v := req.Description; v != nil {
		q.SetDescription(*v)
	}
	if id, err := GetSiloId(ctx, s.db, req.GetSilo()); err != nil {
		return nil, err
	} else {
		q.SetSiloID(id)
	}

	res, err := q.Save(ctx)
	if err != nil {
		return nil, ToStatus(err)
	}

	return ToProtoTeam(res), nil
}
func (s *TeamServiceServer) Delete(ctx context.Context, req *pb.GetTeamRequest) (*emptypb.Empty, error) {
	p, err := GetTeamSpecifier(req)
	if err != nil {
		return nil, err
	}
	if _, err := s.db.Team.Delete().Where(p).Exec(ctx); err != nil {
		return nil, ToStatus(err)
	}

	return &emptypb.Empty{}, nil
}
func (s *TeamServiceServer) Get(ctx context.Context, req *pb.GetTeamRequest) (*pb.Team, error) {
	q := s.db.Team.Query()
	if p, err := GetTeamSpecifier(req); err != nil {
		return nil, err
	} else {
		q.Where(p)
	}

	res, err := QueryTeamWithEdgeIds(q).Only(ctx)
	if err != nil {
		return nil, ToStatus(err)
	}

	return ToProtoTeam(res), nil
}
func QueryTeamWithEdgeIds(q *ent.TeamQuery) *ent.TeamQuery {
	q.WithSilo(func(q *ent.SiloQuery) { q.Select(silo.FieldID) })

	return q
}
func (s *TeamServiceServer) Update(ctx context.Context, req *pb.UpdateTeamRequest) (*pb.Team, error) {
	id, err := GetTeamId(ctx, s.db, req.GetKey())
	if err != nil {
		return nil, err
	}

	q := s.db.Team.UpdateOneID(id)
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

	return ToProtoTeam(res), nil
}
func ToProtoTeam(v *ent.Team) *pb.Team {
	m := &pb.Team{}
	m.Id = v.ID[:]
	m.DateCreated = timestamppb.New(v.DateCreated)
	m.Alias = v.Alias
	m.Name = v.Name
	m.Description = v.Description
	if v := v.Edges.Silo; v != nil {
		m.Silo = ToProtoSilo(v)
	}
	return m
}
func GetTeamId(ctx context.Context, db *ent.Client, req *pb.GetTeamRequest) (uuid.UUID, error) {
	var r uuid.UUID
	if v, err := uuid.FromBytes(req.GetId()); err != nil {
		return r, status.Errorf(codes.InvalidArgument, "id: %s", err)
	} else {
		return v, nil
	}
}
func GetTeamSpecifier(req *pb.GetTeamRequest) (predicate.Team, error) {
	if v, err := uuid.FromBytes(req.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "id: %s", err)
	} else {
		return team.IDEQ(v), nil
	}
}