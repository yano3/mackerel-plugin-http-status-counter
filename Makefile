deps:
	go get ./...

build: deps
	gox -osarch="linux/amd64" -output="pkg/{{.OS}}_{{.Arch}}/{{.Dir}}"
