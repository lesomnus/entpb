// Code generated by "protoc-gen-entpb". DO NOT EDIT

package bare

import (
	context "context"
	uuid "github.com/google/uuid"
	runtime "github.com/lesomnus/entpb/cmd/protoc-gen-entpb/runtime"
	ent "github.com/lesomnus/entpb/example/ent"
	account "github.com/lesomnus/entpb/example/ent/account"
	pb "github.com/lesomnus/entpb/example/pb"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type AccountServiceServer struct {
	db *ent.Client
	pb.UnimplementedAccountServiceServer
}

func (s *AccountServiceServer) Create(ctx context.Context, req *pb.Account) (*pb.Account, error) {
	q := s.db.Account.Create()
	q.SetAlias(req.Alias)
	q.SetRole(toEntRole(req.Role))
	if v, err := uuid.FromBytes(req.Owner.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "owner: %s", err)
	} else {
		q.SetOwnerID(v)
	}

	res, err := q.Save(ctx)
	if err != nil {
		return nil, runtime.EntErrorToStatus(err)
	}

	return toProtoAccount(res), nil
}
func (s *AccountServiceServer) Delete(ctx context.Context, req *pb.DeleteAccountRequest) (*emptypb.Empty, error) {
	q := s.db.Account.Delete()
	switch t := req.GetKey().(type) {
	case *pb.DeleteAccountRequest_Id:
		if v, err := uuid.FromBytes(t.Id); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "id: %s", err)
		} else {
			q.Where(account.IDEQ(v))
		}
	case *pb.DeleteAccountRequest_Alias:
		q.Where(account.AliasEQ(t.Alias))
	}

	_, err := q.Exec(ctx)
	if err != nil {
		return nil, runtime.EntErrorToStatus(err)
	}

	return &emptypb.Empty{}, nil
}
func (s *AccountServiceServer) Get(ctx context.Context, req *pb.GetAccountRequest) (*pb.Account, error) {
	q := s.db.Account.Query()
	switch t := req.GetKey().(type) {
	case *pb.GetAccountRequest_Id:
		if v, err := uuid.FromBytes(t.Id); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "id: %s", err)
		} else {
			q.Where(account.IDEQ(v))
		}
	case *pb.GetAccountRequest_Alias:
		q.Where(account.AliasEQ(t.Alias))
	}

	q.WithOwner(func(q *ent.UserQuery) { q.Select(account.FieldID) })

	res, err := q.Only(ctx)
	if err != nil {
		return nil, runtime.EntErrorToStatus(err)
	}

	return toProtoAccount(res), nil
}
func (s *AccountServiceServer) Update(ctx context.Context, req *pb.UpdateAccountRequest) (*pb.Account, error) {
	id, err := uuid.FromBytes(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "id: %s", err.Error())
	}

	q := s.db.Account.UpdateOneID(id)
	if v := req.Alias; v != nil {
		q.SetAlias(*v)
	}
	if v := req.Role; v != nil {
		q.SetRole(toEntRole(*v))
	}

	res, err := q.Save(ctx)
	if err != nil {
		return nil, runtime.EntErrorToStatus(err)
	}

	return toProtoAccount(res), nil
}
func toProtoAccount(v *ent.Account) *pb.Account {
	m := &pb.Account{}
	m.Id = v.ID[:]
	m.DateCreated = timestamppb.New(v.DateCreated)
	m.Alias = v.Alias
	m.Role = toPbGroupRole(v.Role)
	if v := v.Edges.Owner; v != nil {
		m.Owner = &pb.Actor{Id: v.ID[:]}
	}
	return m
}