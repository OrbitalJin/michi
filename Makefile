serve-dev:
	@go run ./cmd/server.go

build-prod:
	@mkdir -p build && export GIN_MODE=release && go build -o ./build/michi ./cmd/server.go

serve-prod:
	@mkdir -p build && export GIN_MODE=release && go build -o ./build/michi ./cmd/server.go && chmod +x ./build/michi && ./build/michi
