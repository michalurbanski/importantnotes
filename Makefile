# Nice tutorial on Makefile https://gist.github.com/Integralist/9e27ff5582d37ed26aef
hello:
	@echo "Please check makefile contents for details on how to install the program."

build:
	go install

run:
	go run main.go

all: hello build
