
build-dev:
	@export ENV=dev && mkdir -p build && export GIN_MODE=release && go build -o ./build/michi-dev ./cmd/michi.go


build-prod:
	@export ENV=prod && mkdir -p build && export GIN_MODE=release && go build -o ./build/michi ./cmd/michi.go

