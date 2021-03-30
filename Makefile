all: sleep

sleep:
	go build -buildmode c-shared -o sleep.so sleep.go
