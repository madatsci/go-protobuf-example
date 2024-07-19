.PHONY: gen
gen:
	protoc --go_out=. ./addressbook.proto