package runtime

import (
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/lesomnus/entpb/example/ent"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func EntErrorToStatus(err error) error {
	switch {
	case sqlgraph.IsUniqueConstraintError(err):
		return status.Errorf(codes.AlreadyExists, "already exists: %s", err)

	case ent.IsConstraintError(err):
		return status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)

	case ent.IsNotFound(err):
		return status.Errorf(codes.NotFound, "not found: %s", err)

	default:
		return status.Errorf(codes.Internal, "internal error: %s", err)
	}
}
