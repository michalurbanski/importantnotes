# Nice tutorial on Makefile https://gist.github.com/Integralist/9e27ff5582d37ed26aef, 
# followed to create this file.

# parameters
GOBIN=$$GOPATH/bin

hello:
	@echo "Please check makefile contents for details on how to install the program."
	@echo $(GOBIN)

# Simple approach - install application in /bin and copy config to the same place.
install: config.development.yaml
	go install
	cp config.development.yaml $(GOBIN)

#run:
#	go run main.go
#
#all: hello build
