# Example Go RESTful API
This project was setup to illustrate share my standard of practice for building a fully containerized RESTful API structure built in Go. It includes things like Echo, sqlc, migrations, seeds, transformers, as well as a handful of helpful make commands which are documented below. It also features a fully containerized local development environment which you can learn more about below.

If you have questions, feel free to reachout to me on [Twitter](https://twitter.com/ATLChris) or [Mastodon](https://mas.to/@ATLChris).

## Dependencies
- [Go Echo](https://echo.labstack.com/) - Echo is a highly performant, extensible, minimalist Go web framework.
- [CompileDeamon](https://github.com/githubnemo/CompileDaemon) - CompileDaemon enables a live reloading development environment. It is a tool that watches project files and invokes `go build` if file changes are detected.
- [sqlc](https://sqlc.dev/) - sqlc generates fully type-safe idiomatic Go code from SQL.

## Local Development Setup (via VS Code)
1. Install and launch [Docker](https://www.docker.com/products/docker-desktop).
2. Install and launch [VS Code](https://code.visualstudio.com/).
3. Install the VS Code [Remote Development](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack) extension.
4. Open the project in VS Code.
5. Run "Remote-Containers: Open Workspace in Container" from th VS Code Command Palette.
6. Open a new terminal inside of VS Code and run the following command to launch the project in development mode: `make development`
7. Visit [http://localhost:8181/](http://localhost:8181/) via your browser.

## Local Development Setup (Other)
1. Install and launch [Docker](https://www.docker.com/products/docker-desktop).
2. Open a new terminal and browse to the "docker" directory inside the main project directory.
3. Run the following command to launch the projects Docker container: `docker compose up -d` or `docker-compose up -d`
4. Run the following command to open a ZSH prompt in the projects Docker container: `docker exec -it example-go-restful-api /bin/zsh`
5. Inside of the projects Docker container ZSH prompt, run the following command to launch the project in development mode: `make development`
6. Visit [http://localhost:8181/](http://localhost:8181/) via your browser.
7. [Optional] Run the following command to stop the projects Docker container: `docker kill example-go-restful-api`

## Troubleshooting Local Development Setup

### Reset Local Development Environment
Should a problem arise with your local development setup, close VS Code. Then via a terminal on your host machine, browse to the "docker" directory inside the main project directory. Run the below code to remove all containers and all volumes (including your database, so you will need to re-migrate and re-seed) associated with this development environment. Run `docker compose down -v` or `docker-compose down -v` to reset the development environment.

### Update Local Development Environment Containers
If the images in this project are upgraded, you have 2 options to rebuild the containers. If you are using VS Code, you can open the Command Palette and run the "Dev Containers: Rebuild Without Cache and Reopen in Container" command. If you are using the other methind, you can browse to the "docker" directory inside the main project directory and execute `docker compose build --no-cache` or `docker-compose build --no-cache`.

## Available Endpoints
- List users ([http://localhost:8080/v1/users](http://localhost:8080/v1/users)) - This endpoint will list all users in the PostgreSQL database, which in this example is two.
- Get user ([http://localhost:8080/v1/users/3d7fdbb6-8b8d-490d-93fc-59c779fbc5c8](http://localhost:8080/v1/users/3d7fdbb6-8b8d-490d-93fc-59c779fbc5c8)) - This endpoint will get a user by their UUID. This UUID will return my user record from the PostgreSQL database.

## Available Commands
- `make generate` - Generate the compiled sqlc code.
- `make migrate` - Migrate the local containerized PostgreSQL database. Don't forget to setup up your ".env" file.
- `make seed` - Seed the local containerized PostgreSQL database with two example users.
- `make rollback` - Rollback a single migration.
- `make drop` - Drop all tables in the local containerized PostgreSQL database.
- `make development` - Start the Echo server with live reloading via CompileDeamon.
- `make reset` - Reset the local containerised PostgreSQL database. Runs drop, generate, migrate, and seed in a single command.
- `make clean` - Delete all compiled Go executables in the projects directory.
