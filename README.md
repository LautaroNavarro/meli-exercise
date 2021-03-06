# Meli exercise
[![Build Status](https://travis-ci.org/LautaroNavarro/meli-exercise.svg?branch=master)](https://travis-ci.org/LautaroNavarro/meli-exercise)

Here you will find a set of services for mutant determination and statistics generation.

- [Meli exercise](#meli-exercise)
  - [Services](#services)
    - [Mutant determination service](#mutant-determination-service)
      - [Domain](#domain)
      - [Domain knowledge](#domain-knowledge)
      - [Technical description](#technical-description)
    - [Mutant statistics service](#mutant-statistics-service)
      - [Domain](#domain-1)
      - [Technical description](#technical-description-1)
  - [Public actions vs private actions](#public-actions-vs-private-actions)
  - [Communication between services](#communication-between-services)
  - [Architecture](#architecture)
  - [Diagrams](#diagrams)
    - [Determine dna action](#determine-dna-action)
    - [Get stats action](#get-stats-action)
  - [Getting started](#getting-started)
    - [Local project](#local-project)
    - [Run tests](#run-tests)

## Services

### Mutant determination service

#### Domain

Calculate if a human dna is mutant or not, that includes validate if the dna is human.

#### Domain knowledge

* Human DNA: We consider a dna as human dna, when it is a matrix NxN and its items are inside this list ["A", "T", "C", "G"].
* Mutant human DNA: We consider a human dna as a mutant human dna, when there are four or more equals and consecutives letters vertically, horizontally or diagonally.

![image](https://drive.google.com/uc?export=view&id=1ubC0WNumqg_AVkCTPHjNgfMK9TYvkJbg)


#### Technical description

This is a golang service built with Docker. Data is stored in MongoDB.

### Mutant statistics service

#### Domain

Generate and return mutant human statistics.

#### Technical description

This is a golang service built with Docker. Data is stored in Redis.

## Public actions vs private actions

There are two kind of actions inside each service:
* Public actions: This actions are exposed to the outside world.
* Private actions: This actions are not exposed to the outside world and are only accessible from other services.

## Communication between services

Services communicate with each other using the HTTP protocol, an alternative for this could be use a message broker.
The benefits of using HTTP instead of a message broker are:
* Not having to maintain the communication layer.
* Not having a single point of failure.


## Architecture

![image](https://drive.google.com/uc?export=view&id=1TbSaHj9n3L4mtniB4cfhP-lB8ozhcwW4)


## Diagrams

### Determine dna action
![image](https://drive.google.com/uc?export=view&id=1ylzQoK-HMhZyYQ6jj29hKOOuNWKLohIP)

### Get stats action
![image](https://drive.google.com/uc?export=view&id=1__I12PAhhzpIidqH9MtPldVgXL1ZRoeA)

## Getting started

To get this project working on your local machine, you will need the following dependencies:
* Docker
* Skaffold
* Minikube | Docker for Mac >= 18.06 | A k8s cluster to deploy the apps

### Local project

You will need to create a secret for Redis connection and one for mongo connection

    kubectl create secret generic mongopassword --from-literal MONGOPASSWORD=mongo_password
    kubectl create secret generic redispassword --from-literal REDISPASSWORD=redis_password

If you are using Minikube, enable the ingress addon

    minikube addons enable ingress

To start the project you just need to run

    skaffold dev

To run unit tests, you just need docker and run the following command

    make test

That's it now you will be able to access your services through your ingress service.

Notes:
* If you are using Docker, you IP will be local host
* If you are using Minikube run ```minikube ip``` to get your IP
* If you are using another k8s cluster, get the public IP for your k8s cluster

### Run tests
To run unit tests, you just need docker and run the following command

    make test
