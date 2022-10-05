run: 
	go run main.go

swagger-install:
	go install github.com/swaggo/swag/cmd/swag@latest

swagger:
	swag fmt
	swag init

localstack:
	docker run --rm -it -p 4566:4566 -p 4510-4559:4510-4559 localstack/localstack

dynamo-create-table:
	aws dynamodb --endpoint-url=http://localhost:8000 create-table \
		--table-name TodoItems \
		--attribute-definitions \
			AttributeName=Description,AttributeType=S \
		--key-schema \
			AttributeName=Description,KeyType=HASH \
		--provisioned-throughput \
			ReadCapacityUnits=10,WriteCapacityUnits=5

docker-build:
	docker build -t todo-api .

docker-compose:
	docker compose up
