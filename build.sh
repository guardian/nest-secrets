env GOOS=linux GOARCH=amd64 go build -o nest-secrets-linux-amd64 main.go
mv nest-secrets-linux-amd64 bin/

env GOOS=darwin GOARCH=amd64 go build -o nest-secrets-darwin-amd64 main.go
mv nest-secrets-darwin-amd64 bin/
