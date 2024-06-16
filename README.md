
# Book List App

A Book List App designed to help users efficiently manage their books.

#### Features
- Register a new user
- Users able to Login
- Users able to manage(CRUD) book list

## Prerequisites
- Docker Compose ([link](https://docs.docker.com/compose/install/))


## How to Run

- Clone this repo

```bash
git clone https://github.com/joei1994/datawow-book-list-test.git
```

- Checkout `origin/main` branch
```bash
git fetch
```
```bash
git checkout -b main origin/main
```
- Rename `.env.dist` to `.env`

- Start Application
```bash
docker-compose up
```
Now you should have 
- **Book List App** running on `localhost:7788`
- **MySql Server** running on `localhost:33060`
  - **host**: localhost
  - **username**: local_user
  - **password**: password
  - **port**: 33060

## API Document via Swagger
- Browse to `http://localhost:7788/swagger/index.html`.

- Make requests to register a user and login.

- After login, copy a **token** from the response, then click "Authorize" and in a popup, enter the value for "apiKey" in a form: "Bearer {token}". For example:

Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODk0NDA5NjYsIm9yaWdfaWF0IjoxNTg5NDM5OTY2LCJ1c2VyX2lkIjo1fQ.f8dSG3NxFLHwyA5-XIYALT5GtXm4eiH-motqtqAUBOI

Then, click "Authorize" and close the popup. Now, you are able to make requests which require authentication
