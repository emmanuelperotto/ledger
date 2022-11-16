.PHONY: kube-apply
kube-apply:
	kubectl config use-context docker-desktop
	kubectl apply -f deployment/k8s/$(service)

.PHONY: docker-build
docker-build:
	@chmod +x bin/$(service)
	@docker build -t emmanuelperotto/$(service) -f ./build/$(service).dockerfile .

.PHONY: build
build:
	@go build  -o bin/msgrelay github.com/emmanuelperotto/ledger/cmd/msgrelay

.PHONY: up
up:
	./bin/$(service)