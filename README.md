## Weather REST API Microservice

Tech stack: Golang, GIN (Framework), RapidAPI

### Setup information

* Clone this repo
* Install all dependency by this command: go get
* In the config folder, create a file named with `development.yml` and copy everything from `test.yaml`. Put proper information in `development.yml` file
* Run the server by `go run main.go`

### API Preview


* http://localhost:8080/health/ (To check system health)
* http://localhost:8080/api/v1/weather/?query=&days=&dt=
    Query Params:
    * query (Required): Query parameter based on which data is sent back. It could be following:
        * Latitude and Longitude (Decimal degree) e.g: q=48.8567,2.3508
        * city name e.g.: q=Paris
        * US zip e.g.: q=10001
        * UK postcode e.g: q=SW1
        * Canada postal code e.g: q=G2J
        * metar: e.g: q=metar:EGLL
        * iata: e.g: q=iata:DXB
        * auto:ip IP lookup e.g: q=auto:ip
        * IP address (IPv4 and IPv6 supported) e.g: q=100.0.0.1
    * days (Optional): Number of days of forecast required. Default value is 1 
    * dt (Optional): If passing 'dt', it should be between today and next 10 day in `yyyy-MM-dd` format. Default value is current local date.

* For response demo check `response.json` file


### Run unit-tests

Use this command: `go test tests/weather_test.go`

#### Some basic installation process:

* Init module: `go mod init example/weather-rest-api`
* INstall go: `go get github.com/gin-gonic/gin`
* Install GoDotEnv: `go get github.com/joho/godotenv`