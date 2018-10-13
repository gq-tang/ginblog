

dev-requirements:
	go get -u github.com/jmoiron/sqlx
	go get -u github.com/sirupsen/logrus
	go get -u github.com/go-sql-driver/mysql
	go get -u github.com/satori/go.uuid
	go get -u github.com/smartystreets/goconvey/convey
	go get -u github.com/jteeuwen/go-bindata/...

generate:
	@echo "Generating binary files"
	@go generate blog/main.go 