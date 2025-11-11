tuitune: ./cmd/
	go build tuitune.go

all: tuitune completion


bash-completion-path = $$HOME/.local/share/bash-completion/completion

# does not work. no idea where to put that file for it to work
completion: completion-bash

completion-bash: tuitune completion-dir
	./tuitune completion bash > ${bash-completion-path}/tuitune

completion-dir:
	mkdir -p ${bash-completion-path}
