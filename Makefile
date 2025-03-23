clean-all: clean-data clean-config

clean-data:
	rm -rf ~/.taskman

clean-config:
	rm -rf ~/.config/taskman

run-init:
	go run main.go init
