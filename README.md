# Project Archipelago Go API

An API written in Go that serves up data for the Project Archipelago frontend. 

## Installation & Run

```bash
# Download this project
go get https://github.com/do-community/project-archipelago-go-backend
```

```bash
# Build and Run
cd project-archipelago-go-backend
go mod init 
go run .

API Endpoint : http://127.0.0.1:8080/

``` 

## Structure

```
├── controllers
|   └── profilecontroller.go
├── database
|   └── mongo.go
├── entities
|   └── profile.go
└── config.go
└── config.json
└── main.go
└── README.md
``` 

## API

#### api/profiles
* `GET` : Get all profiles

#### api/profiles/{#id}
* `GET` : Get a specific profile using an ID 

#### api/profiles/{#id}
* `POST` : Create a new profile

#### api/profiles/{#id}
* `PUT` : Update a profile 

#### api/profies/{#id}
* `DELETE` : Delete a profile

#### api/health
* `GET` : Get the healthcheck endpoint, `OK`

## To Do
- [ ] Add a healthcheck endpoint at `api/health` that returns `OK` 
- [ ] Store the MongoDB Connection string as an ENV var and use the `os` pkg to get the value 
- [ ] Refactor `GetProfiles` code
- [ ] Write `GetProfileById` code
- [ ] Write `GetProfileById` code 
- [ ] Write `CreateProfile` code 
- [ ] Write `UpdateProfile` code
- [ ] Write `DeleteProfile` code 
- [ ] Containerize with Docker
- [ ] Build a deployment process to push Docker image to a container registry 
- [ ] Add yaml manifests for a Kubernetes Deployment and ClusterIP service