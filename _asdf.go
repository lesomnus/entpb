package _pb

type UserServiceServer {
	db *_ent.Client
	UnimplementedUserServiceServer
}

func NewUserServiceServer(db *_ent.Client) *UserServiceServer {
	return &UserServiceServer{
		db: db
	}
}


func (s *UserServiceServer) Create(context.Context, *Account) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (s *UserServiceServer) Lock(context.Context, *Account) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lock not implemented")
}
