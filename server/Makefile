dev:
	@go run ./cmd/api/main.go


prod:
	@mkdir -p build && go build -o ./build/pow ./cmd/api/main.go && chmod +x ./build/pow && ./build/pow
