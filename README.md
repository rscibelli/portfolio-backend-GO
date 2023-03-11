# portfolio-backend-GO

# How to start

## Docker

### Build

`docker build -t backend .`

### Run from scratch

`docker run -d -p 8080:8080 --name backend backend-build`

### Run existing

`docker start <container name> &`

### Enter into bash of container

`docker exec -it <container name> /bin/bash`

### Chcek if Docker is running on pi

`systemctl status docker`

### List all docker contaiers

`docker ps -a`

### List all docker builds

`docker image ls -a`

### Delete docker container

`docker rm [containerID]`

### Delete docker image

`docker rmi [iamgeID]`

### How to stop a docker container

`docker stop [containerID]`

## My Network Stuff

### My networks public IP

`24.60.191.54`

### My very own domain name (Maps to my public IP)

`55mont.ddns.net`