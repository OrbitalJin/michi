serve-dev:
	@go run ./cmd/server.go


serv-prod:
	@mkdir -p build && go build -o ./build/qmux ./cmd/server.go && chmod +x ./build/qmux && ./build/qmux
