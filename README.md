# Enron-challenge
This application allows users to search through a database of emails using ZincSearch, a lightweight alternative to Elasticsearch. The backend of the application is developed using Go, while the frontend is built with Vue.js.

## Technologies

- **Backend**: Go
- **Frontend**: Vue.js
- **Search Engine**: ZincSearch

## Requirements
- Docker
- Go 1.21

## Installation
Clone the repository, run ```docker-compose build``` and ```docker-compose up -d``` in the root folder where the docker-compose.yml file is located.

- The backend is running on port 8080
- The frontend is running on port 5173
- ZincSearch is running on port 4080

The database will be empty, to populate it download and unzip the [data](http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz) in enron-challenge/indexer/data. Then run ```go run main.go``` inside the indexer folder. This program will upload the data to ZincSearch. This process is a resource-intensive task and can take up a few minutes depending on your machine specifications.

Try the [application](http://localhost:5173) (After running the docker compose).
