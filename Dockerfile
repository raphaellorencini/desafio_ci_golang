FROM scratch

COPY . .

RUN go test

RUN go build main.go

ENTRYPOINT ["./main"]