server-debug:
	go mod download
	go run cmd/main.go
client-debug:
	cd ./web; yarn install
	cd ./web; yarn start
manual-build-push:
	./scripts/build.sh
