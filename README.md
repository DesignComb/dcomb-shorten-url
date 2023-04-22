# URL Shortener Project

This is a URL shortening tool built using Vue3 Nuxt for the frontend, Go for the backend, and PostgreSQL for the database. 

The project is containerized using Docker.

## Getting Started

These instructions will help you set up the project on your local machine for development and testing purposes.

### Prerequisites

Make sure you have the following installed:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Setup

1. Clone the repository:

```
git clone <repo_url>
```

2. Change to the project directory:

```
cd <project_directory>
```

3. Navigate to the `frontend` folder, rename `Dockerfile.example` to `Dockerfile`, and fill in the necessary information:

```
cd frontend
mv Dockerfile.example Dockerfile
```

4. Run the Docker Compose file to set up the services:

```
docker-compose up -d
```

This will create and run the following services:

- `nginx`: Reverse proxy to handle incoming requests.
- `backend`: Go-based backend API service.
- `db`: PostgreSQL database.
- `frontend`: Vue3 Nuxt-based frontend service.

## Services

### Nginx

- Image: `nginx:latest`
- Ports:
    - `80:80`
    - `443:443`
- Volumes:
    - `./nginx.conf:/etc/nginx/nginx.conf`
    - `./ssl:/etc/nginx/ssl/`
- Depends on: `frontend`, `backend`

### Backend

- Build context: `./backend/Dockerfile`
- Working directory: `/usr/local/go/src/main`
- Volumes: `./backend:/usr/local/go/src/main`
- Ports: `8000:5000`
- Depends on: `db`
- Entrypoint: `go run main.go`

### Database

- Image: `postgres:12.4-alpine`
- Ports: `5432:5432`
- Environment:
    - `POSTGRES_USER=user`
    - `POSTGRES_PASSWORD=mysecretpassword`
    - `PGDATA=/var/lib/postgresql/data/pgdata`

### Frontend

- Image: `node:18-alpine`
- Build context: `./frontend/Dockerfile`
- Ports: `8001:80`
- Environment:
    - `HOST=0.0.0.0`
    - `PORT=80`
- Depends on: `backend`

## Usage

After setting up the services, visit `http://localhost` to access the URL shortening tool.

## Contributing

Feel free to contribute by submitting pull requests or reporting issues.

## License
This project is licensed under the MIT License - see the LICENSE file for details.