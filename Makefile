

build:
	@cd src; go build -o ../bin/todo-cli

run: build
	@./bin/todo-cli

