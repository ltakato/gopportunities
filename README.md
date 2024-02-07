# prerequisities
install make
```
sudo apt install make
```

install swag
```
go install github.com/swaggo/swag/cmd/swag@latest
```

enable configs on shell start
```
export GOPATH=$(asdf where golang)/packages
export GOROOT=$(asdf where golang)/go
export PATH="${PATH}:$(go env GOPATH)/bin"
export PATH="${PATH}:$(go env GOROOT)/bin"
```