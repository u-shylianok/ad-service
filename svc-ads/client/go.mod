module github.com/u-shylianok/ad-service/svc-ads/client

go 1.17

require github.com/u-shylianok/ad-service/svc-auth/client v0.0.1

replace github.com/u-shylianok/ad-service/svc-auth/client v0.0.1 => ../../svc-auth/client

require (
	github.com/golang/protobuf v1.5.0
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.27.1
)

require (
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
)
