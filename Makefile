test: install_buildx
docker_driver_name=ggorockee

REPOSITORY=ggorockee

############################
# backend setting
BACKEND_WORKSPACE=./backend
BACKEND_DOCKERFILE=backend.dockerfile
BACKEND_APPNAME=backend
BACKEND_TAG=$(shell date '+%y%m%d_%H%M')
#############################



create_driver:
	@echo "docker create builder"
ifneq ($(shell docker buildx create --name ${docker_driver_name} --driver docker-container --use; echo $$?), 0)
	@echo "If the docker driver is already in use"
	docker buildx use ${docker_driver_name}
endif


backend_all: create_driver backend_build_and_push

backend_build_and_push: create_driver
	@echo "backend docker image build"
	cd ${BACKEND_WORKSPACE} && docker buildx build --platform linux/arm64,linux/amd64 -t ${REPOSITORY}/${BACKEND_APPNAME}:${BACKEND_TAG} -f ${BACKEND_DOCKERFILE} . --push


#backend_push:
#	@echo "docker image build"
#	docker push ${REPOSITORY}/${BACKEND_APPNAME}:${BACKEND_TAG}

up:
	docker compose down
	docker compose up -d


down:
	docker compose down