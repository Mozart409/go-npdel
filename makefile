dev: clear
	go run ./main.go

clear:
	clear

build:
	go build -o ./dist/go-npdel

clean:
	rm -rf ./dist

folders:
	cd test && npm install

check:
	goreleaser check