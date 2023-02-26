RED  =  "\e[31;1m"
GREEN = "\e[32;1m"
YELLOW = "\e[33;1m"
BLUE  = "\e[34;1m"
PURPLE = "\e[35;1m"
CYAN  = "\e[36;1m"

docker-build:
	docker build -f app/Dockerfile -t ${DOCKER_USERNAME}/ops-api:${VERSION} .
	docker build -f rpc/Dockerfile -t ${DOCKER_USERNAME}/ops-rpc:${VERSION} .
	# docker build -f Dockerfile-job -t ${DOCKER_USERNAME}/core-job:${VERSION} .
	@printf $(GREEN)"[SUCCESS] build docker successfully"

docker-publish:
	echo "${DOCKER_PASSWORD}" | docker login --username ${DOCKER_USERNAME} --password-stdin https://${REPO}
	docker push ${DOCKER_USERNAME}/ops-rpc:${VERSION}
	docker push ${DOCKER_USERNAME}/ops-api:${VERSION}
	# docker push ${DOCKER_USERNAME}/ops-job:${VERSION}
	@printf $(GREEN)"[SUCCESS] publish docker successfully"

gen-api:
	goctl api go -api app/api/ops.api -dir ./app/
	@printf $(GREEN)"[SUCCESS] generate api successfully"

gen-rpc:
	goctl rpc protoc rpc/ops-rpc.proto --go_out=./rpc/types --go-grpc_out=./rpc/types --zrpc_out=./rpc/
	@printf $(GREEN)"[SUCCESS] generate rpc successfully"

gen-model:
	goctl
	@printf $(GREEN)"[SUCCESS] generate model successfully"