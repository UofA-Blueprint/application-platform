# application-platform
Platform for hiring process at Blueprint

# Prerequisites
Before starting the project, install:
## Backend
Install Go:  Follow instruction here

Windows | Mac | Linux

https://go.dev/doc/install

https://go.dev/dl/
 

### Navigate to backend folder
```Bash
    cd backend
```
### Install dependencies
```Bash
    go mod download
```


```Bash
    go get github.com/gin-gonic/gin
    go get github.com/lib/pq
    go mod tidy
```
Swagger doc
```Bash
swag init -g cmd/api/v1/main.go -o docs
```

### Run the project
```Bash
    go run cmd/api/v1/main.go
```
Swagger backend can visit via http://localhost:8080/swagger/index.html
and the request make to backend will be through http://localhost:8080

## Frontend
Install dependencies
```Bash
    npm install
```

Run the frontend:
```Bash
    cd frontend
    npm run dev
```
## Database
supabase folder (setup db local dev host):
RUN: 

Supa base set up command:
```Bash
    cd supabase
    npm init -y
    npm install supabase --save-dev
    npx supabase init
```

then use this command tostart local supabase
** Have to have docker open when run **
```Bash
    npx supabase start
```

access Access Supabase Studio (GUI)
with http://localhost:54323

Local Database Connection String
postgres://postgres:postgres@localhost:54322/postgres

### Migrate db
Ex:
```Bash
    npx supabase migration new create_users_table
```
Then apply changes to the new created migration file

Then, make sure it is runnning
```Bash 
    npx supabase start
```
Apply migration

```Bash
    npx supabase db reset
```


## Backend structure
- handlers (handle http requests and responses, HTTP status code)
- middlewwares (authentication, authorization)
- models (Define data structures used across the application like database entities Users, Application)
- routes (Define all API endpoints and connect them to handlers --> map API to handlers)
- services ( application logic like validate email format, password strength)
- utils ( helper function like Data formatting)
- repositories ( Execute SQL queries (SELECT, INSERT, UPDATE, DELETE))