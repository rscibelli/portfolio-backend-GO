sshpass -p pi ssh pi@24.60.191.54 << EOF
    docker stop backend
    docker rm backend
    docker rmi backend-build
    cd repositories/portfolio-backend-GO
    git pull
    docker build -t backend-build .
    docker run -d -p 8080:8080 --name backend backend-build
    docker image ls
    docker container ls
EOF