# Deployment of mini-roles-manager

## Requirements:
* Golang 1.16
* Node.js 15.13
* NPM 7.7
* Docker 19.03
* docker-compose 1.25
* GNU Make 4.3

## Steps:
* Clone repo:
```bash
$ git clone https://github.com/ilya-mezentsev/mini-roles-manager.git && cd mini-roles-manager
```

* Build project:
```bash
$ make build
```

* Run:
```bash
$ make run
```

## Tests:
### Run all project tests:
```bash
$ make tests
```

### Run backend tests:
```bash
$ make backend-tests
```

### Run frontend tests:
```bash
$ make frontend-tests
```

### Run frontend tests (with coverage):
```bash
$ make frontend-tests-coverage
```
