.PHONY: build generate dev-requirements
VERSION=$(shell git describe --abbrev=0 --tags)
GO_EXTRA_BUILD_ARGS=-a -installsuffix cgo

build: generate
	mkdir -p build
	CGO_ENABLED=0 go build $(GO_EXTRA_BUILD_ARGS) -ldflags "-s -w -X main.version=$(VERSION)" -o build/blog blog/main.go 

dev-requirements:
	go get -u github.com/jmoiron/sqlx
	go get -u github.com/sirupsen/logrus
	go get -u github.com/go-sql-driver/mysql
	go get -u github.com/satori/go.uuid
	go get -u github.com/smartystreets/goconvey/convey
	go get -u github.com/jteeuwen/go-bindata/...
	go get -u github.com/gin-contrib/sessions
	go get -u github.com/gin-gonic/gin 
	go get -u github.com/spf13/viper
	go get -u github.com/spf13/cobra
	go get -u github.com/rubenv/sql-migrate
	go get -u github.com/elazarl/go-bindata-assetfs

generate: clean 
	@echo "Generating binary files"
	@go generate blog/main.go 

clean:
	@rm -f static/static_gen.go	
	@rm -f migrations/migrations_gen.go
	@rm -f views/views_gen.go 
	- @rm -rf build/