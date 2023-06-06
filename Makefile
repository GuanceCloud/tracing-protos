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

# generate proto for open-telemetry
otel_proto_tag    := v0.19.0
otel_dir          := opentelemetry-proto
otel_proto_dir    := ${otel_dir}/opentelemetry/proto
otel_proto_files  = $(wildcard ${otel_proto_dir}/*/*/*.proto ${otel_proto_dir}/*/*/*/*.proto)
otel_gen_dir      := ${otel_dir}/gen

# generate proto for pinpoint
pp_proto_tag    := v2.3.1-fixed
pp_dir          := pinpoint-grpc-idl
pp_proto_dir    := ${pp_dir}/proto
pp_proto_files  = $(wildcard ${pp_proto_dir}/*/*.proto)
pp_gen_dir      := ${pp_dir}/gen

# generate proto for skywalking
sky_proto_tag   := v8.3.0-fixed v9.4.0-fixed
sky_dir         := skywalking-data-collect-protocol
sky_gen_dir     := ${sky_dir}/gen

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
# ${5} gen dir
define generate_proto
	rm -rf ./${5}/${2}
	mkdir -p ./${5}/${2}
	${protoc} --proto_path=./${3} --go_out=./${5}/${2} --go-grpc_out=./${5}/${2} $(shell cd ${3};git checkout ${2};cd ..;find ./${3} -type f -iname "*.proto")

endef

.DEFAULT: gen-all
gen-all: gen-opentelemetry gen-pinpoint gen-skywalking

.PHONY: gen-opentelemetry
gen-opentelemetry: ${otel_dir}
	@cd ${otel_dir};git checkout tags/${otel_proto_tag}
	@rm -rf ./${otel_gen_dir}
	@mkdir ./${otel_gen_dir}
	@${protoc} --proto_path=${otel_dir} --go_out=${otel_gen_dir} --go-grpc_out=${otel_gen_dir} ${otel_proto_files}
	@cd ${otel_dir};git checkout main

${otel_dir}:
	git clone -v git@github.com:open-telemetry/opentelemetry-proto.git

.PHONY: gen-pinpoint
gen-pinpoint: ${pp_dir}
	@cd ${pp_dir};git checkout ${pp_proto_tag}
	@rm -rf ./${pp_gen_dir}
	@mkdir ./${pp_gen_dir}
	@${protoc} --proto_path=${pp_proto_dir} --go_out=${pp_gen_dir} --go-grpc_out=${pp_gen_dir} ${pp_proto_files}
	@cd ${pp_dir};git checkout master

${pp_dir}:
	git clone -v git@github.com:pinpoint-apm/pinpoint-grpc-idl.git

.PHONY: gen-skywalking
gen-skywalking: ${sky_dir}
	@$(foreach tag, ${sky_proto_tag}, $(call generate_proto, master,${tag},${sky_dir},${sky_dir},${sky_gen_dir}))

${sky_dir}:
	git clone -v git@github.com:apache/skywalking-data-collect-protocol.git

clean:
	@rm -rf ${otel_gen_dir} ${pp_gen_dir} ${sky_gen_dir}