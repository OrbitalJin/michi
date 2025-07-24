serve-dev:
	@go run ./cmd/server.go


serve-prod:
	@mkdir -p build && export GIN_MODE=release && go build -o ./build/qmux ./cmd/server.go && chmod +x ./build/qmux && ./build/qmux
