# Plane-Golang-Issues-API
The current project is done as part of a Task from plane.so, which aims to recreate plane's issues api built using django in golang such that it targets to reduce the latency to under 500ms

# Running the Project
## Completing the Prerequisites
### Start Postgres DB with Docker
```
chmod +x startpg.sh
./startpg.sh
```
### Performing Migrations for the db
```
brew install golang-migrate
migrate -path ./db/migration -database postgres://root:plane@localhost:5432/plane?sslmode=disable
```
## Booting up the api server
### Install the dependencies 
```
go get
```
### Starting the project
```
go run . start
```
### for help
```
go run . --help
```
