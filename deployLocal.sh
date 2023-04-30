docker stop backend
docker rm backend
docker rmi backend-build
docker build -t backend-build .
docker run -d -p 8080:8080 --name backend -v files:/app/files backend-build
docker image ls
docker container ls