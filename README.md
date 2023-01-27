# portfolio-backend-GO

# How to start

## Docker

### Build

`docker build -t backend .`

### Run from scratch

`docker run -p 8080:8080 backend &`

### Run existing

`docker start <container name> &`

### My networks public IP

`24.60.191.54`

### Enter into bash of container

`docker exec -it <container name> /bin/bash`

### Chcek if Docker is running on pi

`systemctl status docker`