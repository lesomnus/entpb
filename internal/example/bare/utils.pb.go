// Code generated by "protoc-gen-entpb". DO NOT EDIT

package bare

import (
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	ent "github.com/lesomnus/entpb/internal/example/ent"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func ToStatus(err error) error {
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
