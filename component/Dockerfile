# Docker Steps: 
# docker build -t us-central1-docker.pkg.dev/prj-common-l-shared-73205/gabriel-dev-reg/pluggable-components:1.0 . 
# docker push us-central1-docker.pkg.dev/prj-common-l-shared-73205/gabriel-dev-reg/pluggable-components:1.0
# docker images
# docker run -it guergabo/pluggable-component-example

# Start on alpine linux with golang installed
FROM golang

# Move to working directory /pluggable-component
WORKDIR /pluggable-component

# Copy and download code and dependencies
COPY . /pluggable-component
RUN apt-get update
RUN apt install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
RUN export PATH="$PATH:$(go env GOPATH)/bin"
RUN go mod tidy
RUN make plug

# Command to run when starting the container
CMD ["/pluggable-component/bin/plug/server"]