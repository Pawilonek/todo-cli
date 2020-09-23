

build:
	@cd src; go build -o ../bin/nozbe-cli

run: build
	@./bin/nozbe-cli

