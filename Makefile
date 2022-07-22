build-image:
	docker build -t "keda-playground" .
build:
	mkdir -p ./dist/
	go build -o ./dist/ .
# lol
test:
	go test .
clean:
	rm -rf ./dist
	docker rmi keda-playground