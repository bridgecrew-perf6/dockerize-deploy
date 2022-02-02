
# Build nameserver remap script (golang)

FROM golang:alpine3.10 as gobuild
WORKDIR /tm1303/src/

COPY go.mod ./
COPY go.sum ./
COPY go/*.go ./

RUN go mod download
RUN go build -o remap *.go

# Packaged terraform and utils

FROM hashicorp/terraform:1.1.3

WORKDIR /tm1303/src/
COPY --from=gobuild /tm1303/src/remap .
COPY tf/ .
COPY entrypoint.sh ./

RUN ["chmod", "+x", "remap"] 
RUN ["chmod", "+x", "entrypoint.sh"] 

ENTRYPOINT ["sh", "./entrypoint.sh"]