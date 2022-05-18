
### API docs
POST /login (form-data: username, password)
POST /logout

To be added

## Setup
1. Create a PostgreSQL database named `masimelrowoo`
1. Make a .env file according to .env.example
1. Run this code to install dependencies
```
go get
```

## Run
```
go run main.go
```

## Deploy to Heroku
Add heroku on git
```
heroku git:remote -a Masimelrowoo
```
Push changes made to heroku
```
git push heroku master
```
