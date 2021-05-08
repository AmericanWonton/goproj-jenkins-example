run:
	go run *.go

gomod-exp:
	export GO111MODULE=on

gobuild:
	GOOS=linux GOARCH=amd64 go build -o gojenkinspipeline
dockerbuild:
	docker build -t gojenkinspipeline .
dockerbuildandpush:
	docker build -t gojenkinspipeline .
	docker tag gojenkinspipeline americanwonton/gojenkinspipeline
	docker push americanwonton/gojenkinspipeline
dockerrun:
	docker run -it -p 80:80 gojenkinspipeline
dockerrundetached:
	docker run -d -p 80:80 gojenkinspipeline
dockerrunitvolume:
	docker run -it -p 80:80 -v photo-images:/static/images gojenkinspipeline
dockerrundetvolume:
	docker run -d -p 80:80 -v photo-images:/static/images gojenkinspipeline
dockertagimage:
	docker tag gojenkinspipeline americanwonton/gojenkinspipeline
dockerimagepush:
	docker push americanwonton/gojenkinspipeline
dockerallpush:
	docker tag gojenkinspipeline americanwonton/gojenkinspipeline
	docker push americanwonton/gojenkinspipeline
dockerseeshell:
	docker run -it gojenkinspipeline sh