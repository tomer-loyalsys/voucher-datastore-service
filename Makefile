all:
	make -B	\
	dep \
	proto \
	grpc_mock \
	go_test \
	go_build \

dep:
	dep ensure

proto:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/loyalsys/voucher-datastore-service grpc/proto/voucherds.proto

grpc_mock:
	mockery -dir ./grpc/proto -output ./grpc/proto/mocks -name GrpcClient

go_build:
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-X main.version=$(version)" -installsuffix cgo -o voucher-datastore-service .

go_test:
	go test ./test/unit_test -v

.PHONY:
	all
	dep
	proto
	grpc_mock
	go_build
	go_test