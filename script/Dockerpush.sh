echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker push goboilerplates/micro-stream
docker push goboilerplates/micro-stream:0.0.0