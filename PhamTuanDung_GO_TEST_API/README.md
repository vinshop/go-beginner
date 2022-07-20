### SIMPLE GO RESTFUL API

This is simple restful api practice with golang with:
1. **Gin Framework**
2. **Gorm**
3. **Redis**

### Installation & Run
```bash
# Download Gin Framework
$ go get github.com/gin-gonic/gin

# Download GORM
$ go get github.com/jinzhu/gorm

# Run program
$ Go run main.o
```
### Setting DB & connect redis
```bash
.env
#mysql
dbdriver := "mysql"
dbhost := "127.0.0.1"
dbpassword := your mysql user password
dbuser := your account mysql
dbname := your database
dbport := "3306"
#redis
redis_host := "127.0.0.1"
redis_port := "6379"
redis_password := your password
# Connect database mysql
DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
# Connect redis
$ sudo systemctl start redis-server
```
### STRUCTURE

```
├── application
│   ├── user_app.go
│   └── user_app_test.go
├── domain
│   ├── entity
│   │   └── user.go
│   └── repository
│       └── user_repository.go
├── go.mod
├── go.sum
├── infrastructure
│   ├── auth
│   │   ├── auth.go
│   │   ├── redisdb.go
│   │   └── token.go
│   ├── persistence
│   │   ├── db.go
│   │   ├── setup_test.go
│   │   ├── user_repository.go
│   │   └── user_repository_test.go
│   └── security
│       └── password.go
├── interfaces
│   ├── handler_setup_test.go
│   ├── login_handler.go
│   ├── login_handler_test.go
│   ├── middleware
│   │   └── middleware.go
│   ├── user_handler.go
│   └── user_handler_test.go
├── main.go
└── utils
    └── mock
        └── mock.go
```

### HTTP METHODS 

| METHOD | DESCRIPTION             | ENDPOINTS                         |
|--------|-------------------------|-----------------------------------|
| GET    | GET ALL USERS           | `{{}}/api/users`                  |
| GET    | GET AN USER BY ID       | `{{}}/api/user/:user_id`          |
| POST   | CREATE NEW USER         | `{{}}/api/users`                  |
| POST   | LOGIN USER              | `{{}}/api/auth/login`             |
| POST   | REFRESH TOKEN           | `{{}}/api/auth/refresh`           |
| DEL    | DELETE AN USER BY ID    | `{{}}/api/user/:user_id`          |
| PUT    | UPDATE AN INFO USER     | `{{}}/api/user/:user_id`          |
| PUT    | UPDATE AN PASSWORD USER | `{{}}/api/user/password/:user_id` |

### TEST API EXAMPLE

#### 1. GET ALL USERS

JSON Response
```
[
    {
        "id": 13,
        "name": "dung01"
    },
    {
        "id": 14,
        "name": "dung14"
    },
    {
        "id": 15,
        "name": "dung15@gmail.com"
    },
    {
        "id": 16,
        "name": "dung16@gmail.com"
    },
    {
        "id": 17,
        "name": "dung17@gmail.com"
    },
    {
        "id": 18,
        "name": "dung18@gmail.com"
    }
]
```
![screenshot-1](Image_Readme/1.png)<br />

#### 2. GET AN USER BY ID
API
```
http://localhost:8080/api/user/14
```
JSON Response
```
{
    "id": 14,
    "name": "dung14"
}
```
![screenshot-1](Image_Readme/2.png)<br />
#### 3. CREATE NEW USER
API
```
http://localhost:8080/api/users
```
```
{
    "email":"dung26@gmail.com",
    "name":"dung26@gmail.com",
    "password":"123456"
}
```
JSON Response
```
{
    "id": 26,
    "name": "dung26@gmail.com"
}
```
![screenshot-1](Image_Readme/3.png)<br />
#### 4. LOGIN USER
API
```
http://localhost:8080/api/auth/login
```
```
{
    "email":"dung14@gmail.com",
    "password":"123457"
}
```
JSON Response
```
{
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6IjNiNTlhMGRmLWY0NWItNDZjZC1iMDM3LTJjNzY4MTc0NjUwNSIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTY1ODEzMDA5NCwidXNlcl9pZCI6MTR9.Heuc-4Y3ki2NOm4201oknNhSoxcPpt8QnqlIP3ymMJ4",
    "id": 14,
    "name": "dung14",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTg3MzM5OTQsInJlZnJlc2hfdXVpZCI6IjNiNTlhMGRmLWY0NWItNDZjZC1iMDM3LTJjNzY4MTc0NjUwNSsrMTQiLCJ1c2VyX2lkIjoxNH0.X55fxAjojeaNs8h4p2Tm3OETCeerQeaSXg69itIa7To"
}
```
![screenshot-1](Image_Readme/4.png)<br />
#### 5. REFRESH TOKEN
API
```
http://localhost:8080/api/auth/refresh-token
```
Input
```
{
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6ImM3OTA5OGExLWQ4ZWItNGZhNi1hYzg4LTM1YTMyMDcxMmRmYiIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTY1ODI5OTE2OCwidXNlcl9pZCI6MTExfQ.KBYQQna5YoGV_u4j0-cu4mDwewjCUyJmsf-4fhon0W8",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTg5MDMwNjgsInJlZnJlc2hfdXVpZCI6ImM3OTA5OGExLWQ4ZWItNGZhNi1hYzg4LTM1YTMyMDcxMmRmYisrMTExIiwidXNlcl9pZCI6MTExfQ.2JtvPhV1kbUn19-IWXjGYURnZbMBR7T0QQzHA-xHHlw"
}
```
Response new token
```
{
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6ImY0MWMwM2ExLTNlYTYtNGQ4MC1hMDlkLWZkODg2ODA3YmU2ZSIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTY1ODI5OTI2MywidXNlcl9pZCI6MTExfQ.pbAq5EX633QhGC0MyFkcUuUM5fYEqQMBeT1QMgA1NO4",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTg5MDMxNjMsInJlZnJlc2hfdXVpZCI6ImY0MWMwM2ExLTNlYTYtNGQ4MC1hMDlkLWZkODg2ODA3YmU2ZSsrMTExIiwidXNlcl9pZCI6MTExfQ.uCXL_gLd5nBFLSaOXGrWB4KjnZXdA1Yws0PEJSpA9mg"
}
```
![screenshot-1](Image_Readme/5.png)<br />
#### 6. DELETE AN USER BY ID
Delete user who have id = 13 <br />
API
```
http://localhost:8080/api/user/13
```
JSON Response
```
"user deleted"
```
Result
![screenshot-1](Image_Readme/6.1.png)<br />
No user who have id=13 when query get all user in database
![screenshot-1](Image_Readme/6.2.png)<br />
#### 7. UPDATE AN USER INFO 
API
```
http://localhost:8080/api/user/14
```
```
{
    "email":"dung14@gmail.com",
    "name":"dung15"
}
```
JSON Response
```
{
    "id": 14,
    "name": "dung15"
}
```
![screenshot-1](Image_Readme/7.png)<br />
#### 8. UPDATE AN USER PASSWORD 
API
```
http://localhost:8080/api/user/password/14
```
```
{
    "email":"dung14@gmail.com",
    "name":"dung14",
    "password":"123456"
}
```
JSON Response
```
{
    "id": 14,
    "name": "dung15"
}
```
![screenshot-1](Image_Readme/8.png)<br />

### UNIT TEST
Example unit test at [Here]()


