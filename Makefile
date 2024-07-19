.PHONY: gen
gen:
	protoc --go_out=. --go_opt=module=github.com/madatsci/go-protobuf-example ./addressbook.proto