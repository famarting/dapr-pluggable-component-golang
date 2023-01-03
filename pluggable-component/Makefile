PROJECT = plug
PACKAGE = github.com/guergabo/pluggable-component
BIN_DIR = bin
PROTO_DIR = protos/dapr/proto/components/v1
SERVER_DIR = server
CLIENT_DIR = client

${PROJECT}:
	@${CHECK_DIR_CMD}
	protoc -I${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. --go-grpc_opt=module=${PACKAGE} --go-grpc_out=. ${PROTO_DIR}/*.proto
	go build -o ${BIN_DIR}/${PROJECT}/${SERVER_BIN} ./${SERVER_DIR}
	go build -o ${BIN_DIR}/${PROJECT}/${CLIENT_BIN} ./${CLIENT_DIR}