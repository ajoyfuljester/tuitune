completion-path = $$HOME/.local/share/bash-completion

tuitune: ./cmd/
	go build tuitune.go


completion: completion-bash

completion-bash: tuitune completion-dir
	./tuitune completion bash > ${completion-path}/tuitune

completion-dir:
	mkdir -p ${completion-path}
