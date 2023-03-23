# Go REST API

This is a simple REST API boilerplate built with Go language using GORM and Fiber. It uses PostgreSQL as the database.

## Prerequisites

- Go programming language installed on your machine
- Docker and Docker Compose installed on your machine
- Nodemon (NPM package) installed on your machine

## Getting Started

1. Clone this repository onto your local machine.

   ```
   git clone https://github.com/dduafa/go-server.git
   ```

2. Navigate to the project directory.

   ```
   cd go-server
   ```

3. Create an `.env` file with the sample in `.env.example`

4. Build and run the Docker container for the PostgreSQL database.

   ```
   docker-compose up -d --remove-orphans
   ```

5. Start the server.

   ```
   nodemon
   ```

6. You should now be able to access the API via `http://localhost:${SERVER_PORT}`.

## Sample Scripts

- `"docker-compose up -d --remove-orphans"` - Start postgres container.
- `"docker-compose down"` - Stop postgres container.
- `"docker exec -it postgres bash"` - Access docker bash.
- `"psql -U postgres db_name"` - Access postgre database.

## Built With

- [Go](https://golang.org/) - Programming language
- [GORM](https://gorm.io/) - Object-relational mapping (ORM) library
- [Fiber](https://gofiber.io/) - Web framework
- [Nodemon](https://nodemon.io/) - Automatic restarting of application
- [Docker](https://www.docker.com/) - Containerization platform
- [Docker Compose](https://docs.docker.com/compose/) - Tool for defining and running multi-container Docker applications

## To Do
- [ ] Fix nodemon not reflecting new changes
- [ ] Fix registration process failing on duplicate