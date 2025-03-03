# tempo-news-api

## Overview
**tempo-news-api** is a RESTful API built with Golang to manage articles. It provides endpoints to create, retrieve, update, and delete articles.

## Requirements
- **Go** version **1.23** or later
- **PostgreSQL**
- **Redis 7**

## Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/sholehbaktiabadi/tempo-news-api.git
   cd tempo-news-api
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Set up environment variables (e.g., `.env` file):
   ```plaintext
   rename .env.example to .env and put your config
   ```
4. Run the application:
   ```sh
   go run main.go
   ```

## API Documentation
### Create Article
```sh
curl --location 'localhost:8000/api/v1/article' \
--header 'Content-Type: application/json' \
--data '{
    "title": "pertamina korup",
    "content": "mending pertamini",
    "author": "bekti"
}'
```

### Get One Article
```sh
curl --location 'localhost:8000/api/v1/article/1'
```

### Get All Articles (with filters)
```sh
curl --location 'localhost:8000/api/v1/article?author=bekti&search=pertamina'
```

### Update Article
```sh
curl --location --request PATCH 'localhost:8000/api/v1/article/1' \
--header 'Content-Type: application/json' \
--data '{
    "title": "pertamina updated",
    "content": "mending pertamini updated",
    "author": "bekti2"
}'
```

### Delete Article
```sh
curl --location --request DELETE 'localhost:8000/api/v1/article/1'
```