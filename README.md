# Catinder Backend

## Overview

This is the backend for the Catinder application. Catinder is a fun and interactive application designed for cat lovers. This backend is implemented in Go and provides RESTful APIs for the Catinder frontend.

## Project Structure

The project is structured as follows:

- `cmd/`: Contains the application's entry point (`main.go`). This initiates the server.
- `internal/`: Contains the application logic.
  - `handler/`: Contains the HTTP handlers for each endpoint.
  - `middleware/`: Contains middleware functions for the HTTP handlers.
  - `model/`: Contains the data models used in the application.
  - `repository/`: Contains the database operations.
- `pkg/`: Contains packages that can be used by external applications.
- `scripts/`: Contains scripts for various purposes (e.g., database migration).
- `db.go`: Contains the database connection logic.
- `init.sql`: Contains SQL commands to initialize the database.
- `Dockerfile`: Contains the Docker instructions to build the application image.
- `docker-compose.yml`: Contains the Docker Compose configuration to run the application and its dependencies (e.g., database) in containers.

## Setup

1. Clone the repository: `git clone https://github.com/username/catinder-backend.git`
2. Navigate to the project directory: `cd catinder-backend`
3. Install the dependencies: `go mod download`
4. Set up your environment variables in a `.env` file. You can use the `.env.example` as a base. Make sure to replace the placeholder values with your actual values.
5. Run the `init.sql` script to set up your database: `mysql -u yourusername -p yourdatabase < init.sql`
6. Build and run the project with Docker: `docker-compose up --build`

## Development

This project uses [Air](https://github.com/cosmtrek/air) for hot reloading. You can run the project in development mode with `air`.

## Contributing

Contributions are welcome. If you find a bug or have a feature request, please open an issue. If you want to contribute to the code, please fork the repository, make your changes, and submit a pull request.

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.
