# Movie-service

The Movie Service is a RESTful API that provides information about movies. It allows users to browse, search and CRUD operations on movies. This service also provides movie details such as title, release year, runtime, and genre.

## API Endpoints

This application provides the following API endpoints:

| Endpoint                 | Method | Description                                                     |
| ------------------------| ------ | --------------------------------------------------------------- |
| `/v1/healthcheck`        | GET    | Shows the current state of application.                         |
| `/debug/vars`            | GET    | Displays application metrics.                                   |
| `/movies`                | GET    | Returns a list of all movies.                                   |
| `/movies/{id}`           | GET    | Returns a movie with a specific ID.                             |
| `/movies/search?q={query}` | GET  | Returns a list of movies that match the search query.           |
| `/movies`                | POST   | Adds a new movie to the database.                               |
| `/movies/{id}`           | PATCH  | Updates an existing movie with a specific ID.                   |
| `/movies/{id}`           | DELETE | Deletes an existing movie with a specific ID.                   |
