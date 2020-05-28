SHA=$(shell git rev-parse HEAD)

test:
	docker build -t mutant-determination-service-test -f mutant_determination_service/Dockerfile.travis mutant_determination_service/.
	docker run mutant-determination-service-test

deploy:
	docker build -t lautaronavarro/mutant-determination-service:$(SHA) -f mutant_determination_service/Dockerfile mutant_determination_service/.
	docker push lautaronavarro/mutant-determination-service:$(SHA)
