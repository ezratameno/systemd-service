build:
	go build -o ./.dist/sleep-service ./main.go

run: build
	./.dist/sleep-service

# kill will send a kill signal to the
kill:
	killall --signal SIGTRAP sleep-service