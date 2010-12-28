include $(GOROOT)/src/Make.inc

TARG=goresque
GOFILES=goresque.go

include $(GOROOT)/src/Make.cmd

format:
	gofmt -spaces=true -tabindent=false -tabwidth=4 -w goresque.go

