// Code generated by "protoc-gen-entpb". DO NOT EDIT

package bare

import (
	context "context"
	uuid "github.com/google/uuid"
	ent "github.com/lesomnus/entpb/internal/example/ent"
	account "github.com/lesomnus/entpb/internal/example/ent/account"
	predicate "github.com/lesomnus/entpb/internal/example/ent/predicate"
	pb "github.com/lesomnus/entpb/internal/example/pb"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type AccountServiceServer struct {
	db *ent.Client
	pb.UnimplementedAccountServiceServer
}

func NewAccountServiceServer(db *ent.Client) *AccountServiceServer {
	return &AccountServiceServer{db: db}
}
func (s *AccountServiceServer) Create(ctx context.Context, req *pb.CreateAccountRequest) (*pb.Account, error) {
	q := s.db.Account.Create()
	if v := req.Alias; v != nil {
		q.SetAlias(*v)
	}
	q.SetRole(toEntRole(req.GetRole()))
	if id, err := GetUserId(ctx, s.db, req.GetOwner()); err != nil {
		return nil, err
	} else {
		q.SetOwnerID(id)
	}

	res, err := q.Save(ctx)
	if err != nil {
		return nil, ToStatus(err)
	}

	return ToProtoAccount(res), nil
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
	default:
		return nil, status.Errorf(codes.InvalidArgument, "key not provided")
	}

	_, err := q.Exec(ctx)
	if err != nil {
		return nil, ToStatus(err)
	}

	return &emptypb.Empty{}, nil
}
func (s *AccountServiceServer) Get(ctx context.Context, req *pb.GetAccountRequest) (*pb.Account, error) {
	q := s.db.Account.Query()
	if p, err := GetAccountSpecifier(req); err != nil {
		return nil, err
	} else {
		q.Where(p)
	}

	q.WithOwner(func(q *ent.UserQuery) { q.Select(account.FieldID) })

	res, err := q.Only(ctx)
	if err != nil {
		return nil, ToStatus(err)
	}

	return ToProtoAccount(res), nil
}
func (s *AccountServiceServer) Update(ctx context.Context, req *pb.UpdateAccountRequest) (*pb.Account, error) {
	id, err := GetAccountId(ctx, s.db, req.GetKey())
	if err != nil {
		return nil, err
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
		return nil, ToStatus(err)
	}

	return ToProtoAccount(res), nil
}
func ToProtoAccount(v *ent.Account) *pb.Account {
	m := &pb.Account{}
	m.Id = v.ID[:]
	m.DateCreated = timestamppb.New(v.DateCreated)
	m.Alias = v.Alias
	m.Role = toPbGroupRole(v.Role)
	if v := v.Edges.Owner; v != nil {
		m.Owner = ToProtoUser(v)
	}
	return m
}
func GetAccountId(ctx context.Context, db *ent.Client, req *pb.GetAccountRequest) (uuid.UUID, error) {
	var r uuid.UUID
	k := req.GetKey()
	if t, ok := k.(*pb.GetAccountRequest_Id); ok {
		if v, err := uuid.FromBytes(t.Id); err != nil {
			return r, status.Errorf(codes.InvalidArgument, "id: %s", err)
		} else {
			return v, nil
		}
	}

	q := db.Account.Query()
	switch t := k.(type) {
	case *pb.GetAccountRequest_Alias:
		q.Where(account.AliasEQ(t.Alias))
	case nil:
		return r, status.Errorf(codes.InvalidArgument, "key not provided")
	default:
		return r, status.Errorf(codes.Unimplemented, "unknown type of key")
	}
	if v, err := q.OnlyID(ctx); err != nil {
		return r, ToStatus(err)
	} else {
		return v, nil
	}
}
func GetAccountSpecifier(req *pb.GetAccountRequest) (predicate.Account, error) {
	switch t := req.GetKey().(type) {
	case *pb.GetAccountRequest_Id:
		if v, err := uuid.FromBytes(t.Id); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "id: %s", err)
		} else {
			return account.IDEQ(v), nil
		}
	case *pb.GetAccountRequest_Alias:
		return account.AliasEQ(t.Alias), nil
	case nil:
		return nil, status.Errorf(codes.InvalidArgument, "key not provided")
	default:
		return nil, status.Errorf(codes.Unimplemented, "unknown type of key")
	}
}
