dev: folders clear
	go run ./main.go

clear:
	clear

build: clear clean
	goreleaser build --snapshot

clean:
	rm -rf ./dist

folders:
	go run ./demo/folders.go

check:
	goreleaser check

vet:
	go vet

# git tag -a "v0.0.1" -m "this is release v0.0.1"
# git push origin v0.0.1