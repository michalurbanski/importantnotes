# Nice tutorial on Makefile https://gist.github.com/Integralist/9e27ff5582d37ed26aef, 
# followed to create this file.

# parameters
ifndef GOPATH
$(error Please set GOPATH variable to use this installation method)
endif

# Special syntax $$ specific for Makefile, it's not bash
GOBIN=$$GOPATH/bin

# https://stackoverflow.com/a/29576870 - explains why @ at the beginning
hello:
	@echo "Please check makefile contents for details on how to install the program."
	@echo $(GOBIN)

# Simple approach - install application in /bin and copy config to the same place.
# Read as 'install step depends on config.development.yaml file'
install: config.development.yaml
	go install
	cp config.development.yaml $(GOBIN)

#run:
#	go run main.go
#
#all: hello build
