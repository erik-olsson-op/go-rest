dockertag = erikolssonop/go-rest

# Generate swagger files - see /docs
swagger-init:
	swag init -g cmd/main.go

# Build the docker image
docker-build:
	docker build -t $(dockertag) .

# Set exposed port 8080 default and pass the .env file
docker-run:
	docker run --env-file .env -it --rm -p 8080:8080 $(dockertag)