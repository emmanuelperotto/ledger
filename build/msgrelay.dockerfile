## Build
FROM golang:1.19-alpine

WORKDIR /app

#Copy module files
COPY go.mod ./
COPY go.sum ./

#Download go modules
RUN go mod download

#Generate binary
COPY cmd/msgrelay/*.go ./
RUN go build -o /msgrelay

## Deploy
#FROM gcr.io/distroless/base-debian10
#
#WORKDIR /
#
#COPY --from=build /msgrelay /msgrelay
#
#USER nonroot:nonroot
#
ENTRYPOINT ["/msgrelay"]