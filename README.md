
[How To Build Go Executables for Multiple Platforms on Ubuntu 16.04](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04)

```bash
go run app/guess.go
go build app/guess.go
env GOOS=windows GOARCH=amd64 go build app/guess.go
```

modules
```bash
 go mod init mod
 go mod tidy
 ```
