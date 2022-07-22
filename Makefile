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
port-forward:
	kubectl port-forward -n argocd deploy/argocd-server 8080:8080 2>&1 > /dev/null &