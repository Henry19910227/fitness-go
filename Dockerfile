FROM golang:1.16.6-alpine
ARG mode
ENV build_mode=${mode}
WORKDIR /fitness-go
COPY . /fitness-go
RUN go build main.go
EXPOSE 9090
ENTRYPOINT ./main -m ${build_mode}