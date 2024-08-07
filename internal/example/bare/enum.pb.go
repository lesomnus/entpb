// Code generated by "protoc-gen-entpb". DO NOT EDIT

package bare

import (
	pb "github.com/lesomnus/entpb/internal/example/pb"
	role "github.com/lesomnus/entpb/internal/example/role"
)

func toPbRole(v role.Role) pb.Role {
	switch v {
	case "UNSPECIFIED":
		return 0
	case "MEMBER":
		return 10
	case "ADMIN":
		return 20
	case "OWNER":
		return 30
	default:
		return 0
	}
}

func toEntRole(v pb.Role) role.Role {
	switch v {
	case 0:
		return "UNSPECIFIED"
	case 10:
		return "MEMBER"
	case 20:
		return "ADMIN"
	case 30:
		return "OWNER"
	default:
		return ""
	}
}
