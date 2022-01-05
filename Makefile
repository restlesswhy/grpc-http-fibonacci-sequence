protobuf:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		internal/fib/proto/fib.proto

run:
	echo "Starting docker environment"
	docker-compose -f docker-compose.yml up --build


export TEST_CONTAINER_NAME=test_db

test.integration:
	docker run --rm -d -p 6379:6379 --name $$TEST_CONTAINER_NAME  redis:6-alpine

	go test -v ./tests/
	docker stop $$TEST_CONTAINER_NAME
