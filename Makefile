.PHONY: watch
watch:
	tmux has-session -t myapp 2>/dev/null || tmux new-session -d -s myapp 'air'

.PHONY: attach
	tmux attach -t myapp

stop:
	tmux kill-session -t myapp

.PHONY: build
build:
	go build -o build ./cmd/myapp

.PHONY: serve
serve: build
	myapp serve

