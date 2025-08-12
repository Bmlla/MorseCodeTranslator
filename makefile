.PHONY: all build clean test run help

build:
	@echo "Building..."
	@go build -o .

run-demo-to:
	@echo "Running with text='HELLO WORLD', dict='default/latin', mode='to'"
	@go run . --text "HELLO WORLD" --mode to

run-demo-from:
	@echo "Running with text='HELLO WORLD in Morse Code', dict='default/latin', mode='from'"
	@go run . --text ".... . .-.. .-.. ---  .-- --- .-. .-.. -.." --mode from

