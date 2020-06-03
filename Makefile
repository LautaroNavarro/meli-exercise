SHA=$(shell git rev-parse HEAD)

test:
	# Build and test mds
	docker build -t mutant-determination-service-test -f mutant_determination_service/Dockerfile.travis mutant_determination_service/.
	docker run mutant-determination-service-test
	# Build and test mss
	docker build -t mutant-statistics-service-test -f mutant_statistics_service/Dockerfile.travis mutant_statistics_service/.
	docker run mutant-statistics-service-test

deploy:
	# Build and push mds 
	docker build -t lautaronavarro/mutant-determination-service:$(SHA) -f mutant_determination_service/Dockerfile mutant_determination_service/.
	docker push lautaronavarro/mutant-determination-service:$(SHA)
	docker image tag lautaronavarro/mutant-determination-service:$(SHA) lautaronavarro/mutant-determination-service
	docker push lautaronavarro/mutant-determination-service
	# Build and push mss
	docker build -t lautaronavarro/mutant-statistics-service:$(SHA) -f mutant_statistics_service/Dockerfile mutant_statistics_service/.
	docker push lautaronavarro/mutant-statistics-service:$(SHA)
	docker image tag lautaronavarro/mutant-statistics-service:$(SHA) lautaronavarro/mutant-statistics-service
	docker push lautaronavarro/mutant-statistics-service
