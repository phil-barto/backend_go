export GOPATH := $(HOME)/go

reload_rest_air:
	air --build.cmd "go build -o bin/rest_api cmd/rest/main.go" --build.bin "./bin/rest_api"
