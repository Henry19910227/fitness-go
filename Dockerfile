FROM golang
WORKDIR /fitness-go
COPY . /fitness-go
RUN go build main.go
EXPOSE 9090
ENTRYPOINT ./main -m release