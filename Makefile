setup-env:
	docker-compose -f tools/development/docker-compose.yaml up -d

sync-deps:
	go mod tidy
	go mod vendor

serve-web:
	go run main.go web