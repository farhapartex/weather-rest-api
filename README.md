## Weather Information REST API

Tech stack: Golang, GIN (Framework), RapidAPI

### Setup information

* Clone this repo
* Install all dependency by this command: go get
* In the config folder, create a file named with `development.yml` and copy everything from `test.yaml`. Put proper information in `development.yml` file
* Run the server by `go run main.go`


#### Some basic installation process:

* Init module: `go mod init example/weather-rest-api`
* INstall go: `go get github.com/gin-gonic/gin`
* Install GoDotEnv: `go get github.com/joho/godotenv`