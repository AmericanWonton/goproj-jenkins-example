run:
	go run *.go

gomod-exp:
	export GO111MODULE=on
gobuild:
	GOOS=linux GOARCH=amd64 go build -o gojenkinspipelinetest
gotest:
	go test -v
dockerbuild:
	docker build -t gojenkinspipelinetest .
dockerbuildandpush:
	docker build -t gojenkinspipelinetest .
	docker tag gojenkinspipelinetest americanwonton/gojenkinspipelinetest
	docker push americanwonton/gojenkinspipelinetest
dockerrun:
	docker run -it -p 80:80 gojenkinspipelinetest
dockerrundetached:
	docker run -d -p 80:80 gojenkinspipelinetest
dockerrunitvolume:
	docker run -it -p 80:80 -v photo-images:/static/images gojenkinspipelinetest
dockerrundetvolume:
	docker run -d -p 80:80 -v photo-images:/static/images gojenkinspipelinetest
dockertagimage:
	docker tag gojenkinspipelinetest americanwonton/gojenkinspipelinetest
dockerimagepush:
	docker push americanwonton/gojenkinspipelinetest
dockerallpush:
	docker tag gojenkinspipelinetest americanwonton/gojenkinspipelinetest
	docker push americanwonton/gojenkinspipelinetest
dockerseeshell:
	docker run -it gojenkinspipelinetest sh