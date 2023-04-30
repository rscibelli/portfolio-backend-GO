export $( grep -vE "^(#.*|\s*)$" .env )

sshpass -p ${ssh_pass} ssh pi@10.0.0.251 << EOF
    docker stop backend
    docker rm backend
    docker rmi backend-build
    cd repositories/portfolio-backend-GO
    git pull
    docker build -t backend-build .
    docker run -d -p 8080:8080 --name backend -v files:/app/files backend-build
    docker image ls
    docker container ls
EOF