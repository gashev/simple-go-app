all:
	go build main.go

clean:
	rm -f main

run:
	./main

docker:
	docker build . -t simple-go-app