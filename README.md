# AMSYS

AMSYS is a web-based application designed to manage version control and changelogs for projects. It provides a user-friendly interface to track changes, manage users, and display project information dynamically.

## Features

- **Version Control**: View and manage versions of projects.
- **Dynamic Content**: Load and paginate data dynamically using HTMX.
- **User Management**: Display and manage user information.
- **Changelog Display**: Visualize changelogs.
- **API Integration**: RESTful API endpoints for project and version management.
- **Responsive Design**: Built with Bootstrap for a modern and responsive UI.

## Technologies Used

- **Backend**: Go (Golang)
- **Frontend**: HTML, Bootstrap, HTMX
- **Database**: PostgreSQL
- **Web Server**: Nginx
- **Containerization**: Docker and Docker Compose

## Prerequisites

- Docker and Docker Compose installed
- PostgreSQL database running
- Go (Golang) installed for local development

## Setup Instructions

1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/amsys.git
   cd amsys
   ```

2. Configure the environment variables:
   - Edit the `.env` file to set up database credentials and other configurations.

3. Build and run the application using Docker Compose:
   ```bash
   docker-compose up --build
   ```

4. Access the application:
   - Open your browser and navigate to `http://localhost:3001`.

## API Endpoints

- **GET /version**: Fetch project version information.
- **POST /version**: Create a new version for a project.

## File Structure

- **cmd/**: Entry point for the application.
- **internal/api/**: API handlers, routes, and services.
- **internal/web/**: Web templates and static files.
- **configs/**: Configuration files.
- **.env**: Environment variables.

## Development

1. Install dependencies:
   ```bash
   go mod tidy
   ```

2. Run the application locally:
   ```bash
   go run ./cmd/main.go
   ```

3. Use [Air](https://github.com/cosmtrek/air) for live reloading during development:
   ```bash
   air
   ```

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
