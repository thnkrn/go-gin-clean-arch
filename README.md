# Clean Architecture using Golang with Gin framework

## Template Structure

- [Gin](github.com/gin-gonic/gin) is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.
- [JWT](github.com/golang-jwt/jwt) A go (or 'golang' for search engine friendliness) implementation of JSON Web Tokens.
- [GORM](https://gorm.io/index.html) with [PostgresSQL](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL)The fantastic ORM library for Golang aims to be developer friendly.
- [Wire](https://github.com/google/wire) is a code generation tool that automates connecting components using dependency injection.
- [Viper](https://github.com/spf13/viper) is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application, and can handle all types of configuration needs and formats.
- [swag](https://github.com/swaggo/swag) converts Go annotations to Swagger Documentation 2.0 with [gin-swagger](https://github.com/swaggo/gin-swagger) and [swaggo files](github.com/swaggo/files)

## Using `bulletin-resource-center` project

To use `bulletin-resource-center` project, follow these steps:

```bash
# Navigate into the project
cd ./bulletin-resource-center

# Install dependencies
make deps

# Generate wire_gen.go for dependency injection
# Please make sure you are export the env for GOPATH
make wire

# Run the project in Development Mode
make run
```

Additional commands:

```bash
➔ make help
build                          Compile the code, build Executable File
run                            Start application
test                           Run tests
test-coverage                  Run tests and generate coverage file
deps                           Install dependencies
deps-cleancache                Clear cache in Go module
wire                           Generate wire_gen.go
swag                           Generate swagger docs
help                           Display this help screen
```

## Available Endpoint

In the project directory, you can call:

### `GET /`

For getting status page

### `POST /login`

For generating a JWT

### `GET /api/users`

For getting all of users

### `GET /api/users/:id`

For getting user by ID

### `POST /api/users`

For creating new user

### `DELETE /api/users/:id`

For removing existing user

### `PUT /api/users/:id`

For updating existing user

## Commit rules

### Conventional Commits

- We are using Conventional Commits rule to add readable meaning to commit messages
- We are following the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) rule

#### Format for commit

[type][optional scope]: [optional REFERENCE-1234] [description]

- ex. build(husky): [BO-000] add husky and commitlint

- List of commit type
  [
  'build',
  'ci',
  'chore',
  'docs',
  'feat',
  'fix',
  'perf',
  'refactor',
  'revert',
  'style',
  'test'
  ]

## Contributing to `bulletin-resource-center`

To contribute to `bulletin-resource-center`, follow these steps:'

1. Clone this repository.
2. Create a feature branch: `git checkout -b <branch_name>`.
3. Make your changes and commit them: `git commit -m '<commit_message>'`
4. Push to the original branch: `git push origin <branch_name>`
5. Create the pull request against `master`.
