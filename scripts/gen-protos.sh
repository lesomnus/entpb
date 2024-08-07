#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace

shopt -s globstar

__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)" # Directory where this script exists.
__root="$(cd "$(dirname "${__dir}")" && pwd)"         # Root directory of project.



MODULE_NAME=github.com/lesomnus/entpb/internal/example
PROTO_ROOT="${__root}/internal/example/protos/syntax_proto3"
OUTPUT_DIR="${__root}/internal/example"
cd "${PROTO_ROOT}"

protoc \
	--proto_path="${PROTO_ROOT}" \
	\
	--go_out="${OUTPUT_DIR}" \
	--go_opt=module="${MODULE_NAME}" \
	\
	--go-grpc_out="${OUTPUT_DIR}" \
	--go-grpc_opt=module="${MODULE_NAME}" \
	\
	--entpb_out="${OUTPUT_DIR}" \
	--entpb_opt=module="${MODULE_NAME}" \
	--entpb_opt=schema_path="${__root}/internal/example/schema" \
	--entpb_opt=ent_package="${MODULE_NAME}/ent" \
	--entpb_opt=package="${MODULE_NAME}/bare" \
	\
	--grpc-gateway_out="${OUTPUT_DIR}/gw" \
	--grpc-gateway_opt=module="${MODULE_NAME}" \
    --grpc-gateway_opt="standalone=true" \
	\
	--openapiv2_out="${OUTPUT_DIR}/openapiv2" \
	--openapiv2_opt="allow_merge=true" \
	\
	"${PROTO_ROOT}"/**/*.proto
