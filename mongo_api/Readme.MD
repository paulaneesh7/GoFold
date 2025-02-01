# Mongo API

This is a simple RESTful API built with Go and MongoDB for managing a collection of movies. It allows users to create, read, update, and delete movie records.

## Table of Contents

- [Mongo API](#mongo-api)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Usage](#usage)
  - [API Endpoints](#api-endpoints)
    - [Create a Movie](#create-a-movie)
    - [Get All Movies](#get-all-movies)
    - [Update a Movie](#update-a-movie)
    - [Delete a Movie](#delete-a-movie)
    - [Delete All Movies](#delete-all-movies)
  - [Contributing](#contributing)

## Installation

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd mongo_api
   ```

2. **Initialize Go modules:**
   ```bash
   go mod init <module-name>
   ```

3. **Install dependencies:**
   ```bash
   go get github.com/gorilla/mux
   go get go.mongodb.org/mongo-driver/mongo
   ```

4. **Run the application:**
   ```bash
   go run main.go
   ```

## Usage

Once the application is running, you can interact with the API using tools like Postman or curl. The API will be available at `http://localhost:8080`.

## API Endpoints

### Create a Movie

- **Endpoint:** `POST /movies`
- **Description:** Creates a new movie record.
- **Request Body:**
  ```json
  {
    "_id_": "Movie Id",
    "movie": "Movie Title",
    "watched": true
  }
  ```

### Get All Movies

- **Endpoint:** `GET /movies`
- **Description:** Retrieves a list of all movies.

### Update a Movie

- **Endpoint:** `PUT /movies/{id}`
- **Description:** Updates the movie as watched or not watched by Id.


### Delete a Movie

- **Endpoint:** `DELETE /movies/{id}`
- **Description:** Deletes a movie record by ID.

### Delete All Movies

- **Endpoint:** `DELETE /movies`
- **Description:** Deletes all movie records.

## Contributing

Contributions are welcome! Please fork the repository and create a pull request.