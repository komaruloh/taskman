clean-all: clean-data clean-config

clean-data:
	rm -rf ~/.local/share/taskman

clean-config:
	rm -rf ~/.config/taskman

run:
	go run main.go 

build:
	go build -o taskman main.go

test:
	go test -cover ./...
