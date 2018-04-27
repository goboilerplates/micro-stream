CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main.out .
docker build -t goboilerplates/micro-stream .
docker tag  goboilerplates/micro-stream goboilerplates/micro-stream:0.0.0