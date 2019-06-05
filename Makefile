.PHONY: bindata build build_batch_linux build_document build_linux datafeed fmt lint restore test artifact

PATH_API_MAIN = application/api/server/main.go

build:
	make bindata
	go build -o bin/api $(PATH_API_MAIN)

build_linux:
	make bindata
	GOOS=linux GOARCH=amd64 go build -o bin/api $(PATH_API_MAIN)
restore:
	go get -u github.com/jteeuwen/go-bindata/...
	go get -u github.com/FiloSottile/gvt
	gvt restore

fmt:
	gofmt -w {application,domain,infrastructure,util}
	goimports -w {application,domain,infrastructure,util}

bindata:
	go-bindata -pkg bindata -o bindata/bindata.go conf/...

