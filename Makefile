build:
	DOCKER_BUILDKIT=1 docker build -t go-firebase .

pull:
	sudo docker pull eecs4312basedcode/go-firebase

push:
	sudo docker tag go-firebase teamdn-htn/go-firebase
	docker push eecs4312basedcode/go-firebase

run:
	sudo docker run  --rm -d -p 8081:8081 -e PORT='8081' \
		--name go-firebase go-firebase

kill:
	sudo docker kill go-firebase