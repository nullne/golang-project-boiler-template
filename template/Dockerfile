# This dockefile is for k8s
FROM golang:1.15-buster

# Add source code to src and build the app
COPY . src/{{ServiceName}}
RUN cd src/{{ServiceName}} && make build

FROM debian:buster-20191014-slim

# Set the APP_NAME and copy the binary under /bin
WORKDIR /app/console.cloud.sea.com
ENV APP_NAME console.cloud.sea.com
COPY --from=0 /go/src/{{ServiceName}}/build/bin/* /bin/

ENTRYPOINT /bin/{{ServiceName}}

# Document that the service listens on port 8080.
EXPOSE 8080
