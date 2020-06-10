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
	# Apply k8s config files
	kubectl apply -f k8s/ingress-service
	kubectl apply -f k8s/mongo
	kubectl apply -f k8s/mutant-determination-service
	kubectl apply -f k8s/mutant-statistics-service
	kubectl apply -f k8s/redis
	# Force image update to commit images
	kubectl set image deployments/mutant-statistics-service mss=lautaronavarro/mutant-statistics-service:$(SHA)
	kubectl set image deployments/mutant-determination-service mds=lautaronavarro/mutant-determination-service:$(SHA)
