// Code generated by "protoc-gen-entpb". DO NOT EDIT

package bare

import (
	context "context"
	uuid "github.com/google/uuid"
	ent "github.com/lesomnus/entpb/internal/example/ent"
	predicate "github.com/lesomnus/entpb/internal/example/ent/predicate"
	token "github.com/lesomnus/entpb/internal/example/ent/token"
	user "github.com/lesomnus/entpb/internal/example/ent/user"
	pb "github.com/lesomnus/entpb/internal/example/pb"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	strings "strings"
)

type TokenServiceServer struct {
	db *ent.Client
	pb.UnimplementedTokenServiceServer
}

func NewTokenServiceServer(db *ent.Client) *TokenServiceServer {
	return &TokenServiceServer{db: db}
}
func (s *TokenServiceServer) Create(ctx context.Context, req *pb.CreateTokenRequest) (*pb.Token, error) {
	q := s.db.Token.Create()
	q.SetValue(req.GetValue())
	q.SetType(req.GetType())
	if v := req.Name; v != nil {
		q.SetName(*v)
	}
	if v := req.UseCountLimit; v != nil {
		q.SetUseCountLimit(*v)
	}
	if v := req.GetDateExpired(); v != nil {
		w := v.AsTime()
		q.SetDateExpired(w)
	}
	if id, err := GetUserId(ctx, s.db, req.GetOwner()); err != nil {
		return nil, err
	} else {
		q.SetOwnerID(id)
	}
	if v := req.GetParent(); v != nil {
		if id, err := GetTokenId(ctx, s.db, v); err != nil {
			return nil, err
		} else {
			q.SetParentID(id)
		}
	}

	res, err := q.Save(ctx)
	if err != nil {
		return nil, ToStatus(err)
	}

	return ToProtoToken(res), nil
}
func (s *TokenServiceServer) Delete(ctx context.Context, req *pb.GetTokenRequest) (*emptypb.Empty, error) {
	p, err := GetTokenSpecifier(req)
	if err != nil {
		return nil, err
	}
	if _, err := s.db.Token.Delete().Where(p).Exec(ctx); err != nil {
		return nil, ToStatus(err)
	}

	return &emptypb.Empty{}, nil
}
func (s *TokenServiceServer) Get(ctx context.Context, req *pb.GetTokenRequest) (*pb.Token, error) {
	q := s.db.Token.Query()
	if p, err := GetTokenSpecifier(req); err != nil {
		return nil, err
	} else {
		q.Where(p)
	}

	res, err := QueryTokenWithEdgeIds(q).Only(ctx)
	if err != nil {
		return nil, ToStatus(err)
	}

	return ToProtoToken(res), nil
}
func QueryTokenWithEdgeIds(q *ent.TokenQuery) *ent.TokenQuery {
	q.WithOwner(func(q *ent.UserQuery) { q.Select(user.FieldID) })
	q.WithParent(func(q *ent.TokenQuery) { q.Select(token.FieldID) })
	q.WithChildren(func(q *ent.TokenQuery) { q.Select(token.FieldID) })

	return q
}
func (s *TokenServiceServer) Update(ctx context.Context, req *pb.UpdateTokenRequest) (*pb.Token, error) {
	id, err := GetTokenId(ctx, s.db, req.GetKey())
	if err != nil {
		return nil, err
	}

	q := s.db.Token.UpdateOneID(id)
	if v := req.Name; v != nil {
		q.SetName(*v)
	}
	if v := req.UseCountLimit; v != nil {
		q.SetUseCountLimit(*v)
	}
	if v := req.DateExpired; v != nil {
		w := v.AsTime()
		q.SetDateExpired(w)
	}

	res, err := q.Save(ctx)
	if err != nil {
		return nil, ToStatus(err)
	}

	return ToProtoToken(res), nil
}
func ToProtoToken(v *ent.Token) *pb.Token {
	m := &pb.Token{}
	m.Id = v.ID[:]
	m.DateCreated = timestamppb.New(v.DateCreated)
	m.Value = v.Value
	m.Type = v.Type
	m.Name = v.Name
	m.UseCountLimit = v.UseCountLimit
	m.DateExpired = timestamppb.New(v.DateExpired)
	if v := v.Edges.Owner; v != nil {
		m.Owner = ToProtoUser(v)
	}
	if v := v.Edges.Parent; v != nil {
		m.Parent = ToProtoToken(v)
	}
	for _, v := range v.Edges.Children {
		m.Children = append(m.Children, ToProtoToken(v))
	}
	return m
}
func GetTokenId(ctx context.Context, db *ent.Client, req *pb.GetTokenRequest) (uuid.UUID, error) {
	var r uuid.UUID
	k := req.GetKey()
	if t, ok := k.(*pb.GetTokenRequest_Id); ok {
		if v, err := uuid.FromBytes(t.Id); err != nil {
			return r, status.Errorf(codes.InvalidArgument, "id: %s", err)
		} else {
			return v, nil
		}
	}

	p, err := GetTokenSpecifier(req)
	if err != nil {
		return r, err
	}

	v, err := db.Token.Query().Where(p).OnlyID(ctx)
	if err != nil {
		return r, ToStatus(err)
	}

	return v, nil
}

func GetTokenSpecifier(req *pb.GetTokenRequest) (predicate.Token, error) {
	switch t := req.GetKey().(type) {
	case *pb.GetTokenRequest_Id:
		if v, err := uuid.FromBytes(t.Id); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "id: %s", err)
		} else {
			return token.IDEQ(v), nil
		}
	case *pb.GetTokenRequest_Value:
		return token.ValueEQ(t.Value), nil
	case *pb.GetTokenRequest_Query:
		if req, err := ResolveGetTokenQuery(req); err != nil {
			return nil, err
		} else {
			return GetTokenSpecifier(req)
		}
	case nil:
		return nil, status.Errorf(codes.InvalidArgument, "key not provided")
	default:
		return nil, status.Errorf(codes.Unimplemented, "unknown type of key")
	}
}

func ResolveGetTokenQuery(req *pb.GetTokenRequest) (*pb.GetTokenRequest, error) {
	t, ok := req.Key.(*pb.GetTokenRequest_Query)
	if !ok {
		return req, nil
	}

	q := t.Query

	if v, ok := strings.CutPrefix(q, ""); ok {
		return pb.TokenByValue(v), nil
	}
	v, err := uuid.Parse(q)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid query string: %s", err)
	}
	return pb.TokenById(v), nil
}
