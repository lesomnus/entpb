#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset
# set -o xtrace

shopt -s globstar

__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)" # Directory where this script exists.
__root="$(cd "$(dirname "${__dir}")" && pwd)"         # Root directory of project.



MODULE_NAME=github.com/lesomnus/entpb
PROTO_ROOT="${__root}/example/protos/syntax_proto3"
OUTPUT_DIR="${__root}/example"
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
	"${PROTO_ROOT}"/**/*.proto
