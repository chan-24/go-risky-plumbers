# go-risky-plumbers
Risky Plumbers is simple a web application in golang to GET and POST risks

### Prerequisites
Need Go (version 1.22.8 or later)

### Steps
1. Clone the repository:
    `git clone https://github.com/chan-24/go-risky-plumbers.git`
2. Navigate to the project directory:
    `cd go-risky-plumbers`
3. Install dependencies:
    `go mod tidy`

## Usage
To start the application and start using
    `go run main.go`
The application runs on port 8080 by default, if its in use and you need to set it to something else before running application,
    `export PORT=8081`

## Features
1. Get all risks in system - Returns list of risks present in memory
    `curl -X GET 'http://localhost:8080/v1/risks'`
    
2. Post a new risk to the system - Adds new risk to memory
    `curl -X POST 'http://localhost:8080/v1/risks' -d '{"state": "open", "title":"risk1", "description": "Very risky risk!"}'`

3. Get a risk by ID - Fetches the risk if present
     `curl -X GET 'http://localhost:8080/v1/risks/ba40ff3e-327a-46b4-926b-b22970cc6100'`