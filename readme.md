Dockerize python script from golang script using dockers and sanic


# How to Use

1- Build python image
docker build -t or-python:latest -f Dockerfile.python .

2- Build golang image
docker build -t or-go:latest -f Dockerfile.go .

3- Run container
docker-compose up
(run in other termnal )--> go run test.go