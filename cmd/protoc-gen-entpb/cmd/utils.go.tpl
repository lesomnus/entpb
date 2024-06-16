func ToStatus(err error) error {
	{{ $status := import "google.golang.org/grpc/status" -}}
	{{ $codes := import "google.golang.org/grpc/codes" -}}
	switch {
	case {{ import "entgo.io/ent/dialect/sql/sqlgraph" | ident "IsUniqueConstraintError" }}(err):
		return {{ $status | ident "Errorf" }}({{ $codes | ident "AlreadyExists" }}, "already exists: %s", err)

	case {{ $.Ent.Ident "IsConstraintError" | use }}(err):
		return {{ $status | ident "Errorf" }}({{ $codes | ident "InvalidArgument" }}, "invalid argument: %s", err)

	case {{ $.Ent.Ident "IsNotFound" | use }}(err):
		return {{ $status | ident "Errorf" }}({{ $codes | ident "NotFound" }}, "not found: %s", err)

	default:
		return {{ $status | ident "Errorf" }}({{ $codes | ident "Internal" }}, "internal error: %s", err)
	}
}
