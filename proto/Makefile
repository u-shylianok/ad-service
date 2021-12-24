.PHONY: generate svc-ads

generate: svc-ads

svc-ads:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    svc/ads/ad.proto
