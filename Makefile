
run:
	go run main.go

release:
	docker build -t vingarcia/echoserver:latest .
	docker push vingarcia/echoserver:latest
