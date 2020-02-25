.PHONY: build

build:
	go build -o ./bin/ctl main.go

install:
	install -C ./bin/ctl /usr/bin/ctl

completion:
	install -C ./bash/ctl.sh /etc/bash_completion.d/ctl
	install -d /usr/lib/ctl/
	install -C ./bash/complete.sh /usr/lib/ctl/complete.sh

run:
	go run main.go
