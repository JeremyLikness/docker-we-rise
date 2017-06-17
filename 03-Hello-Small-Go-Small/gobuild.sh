export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
go build -a -installsuffix cgo -o main .
docker build -t we-rise-svc .
rm main