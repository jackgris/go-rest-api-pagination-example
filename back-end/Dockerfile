# syntax=docker/dockerfile:1
FROM golang:1.19 AS build-stage

# Set destination for COPY
WORKDIR /api

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
# Download Go modules
COPY ./ ./
RUN go mod download

WORKDIR /api/app/todo-api

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /todo-api

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /todo-api /todo-api

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 3000

USER nonroot:nonroot

# Run
ENTRYPOINT ["/todo-api"]
