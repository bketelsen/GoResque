include $(GOROOT)/src/Make.inc

TARG=worker
GOFILES=worker.go

include $(GOROOT)/src/Make.cmd

format:
	gofmt -spaces=true -tabindent=false -tabwidth=4 -w worker.go

