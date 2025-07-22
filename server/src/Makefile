dev:
	@go run ./main.go


prod:
	@mkdir -p build && go build -o ./build/pow ./main.go && chmod +x ./build/pow && ./build/pow
