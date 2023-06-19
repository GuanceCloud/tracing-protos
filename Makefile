#  Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
#  You may obtain a copy of the License at
#     # http://www.apache.org/licenses/LICENSE-2.0
#  Unless required by applicable law or agreed to in writing, software
#  distributed under the License is distributed on an "AS IS" BASIS,
#  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#  See the License for the specific language governing permissions and
#  limitations under the License.

protoc := $(shell which protoc)
ifeq ($(protoc),)
$(error install protoc first)
endif

# go path
go_src = ${GOPATH}/src

# generate proto for open-telemetry
otel_remote       := git@github.com:CodapeWild/opentelemetry-proto.git
otel_proto_tag    := v0.19.0-fixed
otel_dir          := opentelemetry-proto
otel_proto_dir    := ${otel_dir}/opentelemetry/proto
otel_proto_files  = $(wildcard ${otel_proto_dir}/*/*/*.proto ${otel_proto_dir}/*/*/*/*.proto)
otel_gen_dir      := ./opentelemetry-gen-go

# generate proto for pinpoint
pp_remote       := git@github.com:CodapeWild/pinpoint-grpc-idl.git
pp_proto_tag    := v2.3.1-fixed
pp_dir          := pinpoint-grpc-idl
pp_proto_dir    := ${pp_dir}/proto
pp_proto_files  = $(wildcard ${pp_proto_dir}/*/*.proto)
pp_gen_dir      := ./pinpoint-gen-go

# generate proto for skywalking
sky_remote        := git@github.com:CodapeWild/skywalking-data-collect-protocol.git
sky_proto_tag     := v9.4.0-fixed
sky_dir           := skywalking-data-collect-protocol
sky_proto_dir     := ${sky_dir}
sky_proto_files   = $(wildcard ${sky_proto_dir}/*/*.proto ${sky_proto_dir}/*/*/*.proto)
sky_gen_dir       := ./skywalking-gen-go

# protoc env configuration
# "protoc": {
#   "path": "/usr/local/bin/protoc",
#   "options": [
#     "--proto_path=${workspaceRoot}/proto",
#     "--go_out=${env.GOPATH}/src",
#     "--go_opt=paths=import",
#     "--go-grpc_out=${env.GOPATH}/src",
#     "--go-grpc_opt=paths=import"
#   ]
# }

# parameters
# ${1} origin branch
# ${2} expected branch
# ${3} dir
# ${4} proto dir
define generate_proto
	echo "expected branch:${2}"
	echo "${protoc} --proto_path=./${3} --go_opt=paths=import --go_out=${go_src} --go-grpc_out==${go_src} $(shell cd ${3};git checkout -q ${2};cd ..;find ./${3} -type f -iname "*.proto")"
	${protoc} --proto_path=./${3} --go_opt=paths=import --go_out=${go_src} --go-grpc_out=${go_src} $(shell cd ${3};git checkout -q ${2};cd ..;find ./${3} -type f -iname "*.proto")
	cd ${3};git checkout ${1}

endef

.DEFAULT: gen-all
gen-all: rm gen-opentelemetry gen-pinpoint gen-skywalking

.PHONY: gen-opentelemetry
gen-opentelemetry: ${otel_dir}
	@cd ${otel_dir};git checkout -q ${otel_proto_tag}
	${protoc} --proto_path=${otel_dir} --go_opt=paths=import --go_out=${go_src} --go-grpc_out=${go_src} ${otel_proto_files}
	@cd ${otel_dir};git checkout -q main

${otel_dir}:
	git clone -v ${otel_remote}

.PHONY: gen-pinpoint
gen-pinpoint: ${pp_dir}
	@cd ${pp_dir};git checkout -q ${pp_proto_tag}
	${protoc} --proto_path=${pp_proto_dir} --go_opt=paths=import --go_out=${go_src} --go-grpc_out=${go_src} ${pp_proto_files}
	@cd ${pp_dir};git checkout -q master

${pp_dir}:
	git clone -v ${pp_remote}

.PHONY: gen-skywalking
gen-skywalking: ${sky_dir}
	@cd ${sky_dir};git checkout -q ${sky_proto_tag}
	${protoc} --proto_path=${sky_dir} --go_opt=paths=import --go_out=${go_src} --go-grpc_out=${go_src} ${sky_proto_files}
	@cd ${sky_dir};git checkout -q master

${sky_dir}:
	git clone -v ${sky_remote}

clean:
	@rm -rf ${otel_gen_dir} ${pp_gen_dir} ${sky_gen_dir}

rm:
	@rm -rf ${otel_dir} ${pp_dir} ${sky_dir}