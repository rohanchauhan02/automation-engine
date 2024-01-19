# Automation Engine

## Prerequisite

* Install go v1.20 or above to use go modules

* Make vendored copy of depedencies
    ```shell
        go mod vendor
    ```

* Copy app.config.sample yml into local environment. For example:
    ```shell
        cp app.config.sample.yml app.config.local.yml
    ```

* Fill the yml values
* Make sure your redis and mysql database is ready

* Install `goose` for db migrations :
    ```shell
        go get -u github.com/pressly/goose/cmd/goose
    ```

* Migration Up
    ```
        cd datatbase/migrations
        goose postgres "user=root dbname=automation_engine sslmode=disable" up
    ```
* Migration Down
    ```cd datatbase/migrations
        goose postgres "user=root dbname=automation_engine sslmode=disable" down
    ```

## Usage
1. Manual run
    ```shell
        go run app/main.go
    ```

## TODO
- [v] Initiate project
- [ ] Initiate postgresSQL as data store
- [ ] Initiate redis as cache repository
- [ ] Integrate with prometheus, grafana, elastic
- [ ] Unit test and integration test

## Directory Structure
```tree
.
|-- main execution file
|-- domain contains modular code based per domain/ feature
|   |-- domain-name is domain or feature name
|   |   |-- delivery is folder consist of delivery process of the domain
|   |   |   |-- http represent that domain feature will deliver with http endpoint
|   |   |-- repository is a domain's repository acting to store the data
|   |   |   |-- repository.go represent that domain feature using postgres depedencies as data store
|   |   `-- usecase is a domain's usecase acting for logical process of the feature
|   |       `-- usecase.go is a file for usecase interface implementation
|    `--`-- domain-name.go is main file of domain feature represent of interface implementation from usecase and repository
|-- models is a directory consist of struct models used in codebase
|-- shared is a directory consist of shared package for the codebase
|-- app.config.yml is a file containing the env value for the codebase
`-- README.md is a codebase documentation
