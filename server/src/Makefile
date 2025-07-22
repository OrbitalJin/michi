dev:
	@go run ./cmd/main.go


prod:
	@mkdir -p build && go build -o ./build/pow ./cmd/main.go && chmod +x ./build/pow && ./build/pow
