protobuf:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		internal/fib/proto/fib.proto

run:
	echo "Starting docker environment"
	docker-compose -f docker-compose.yml up --build