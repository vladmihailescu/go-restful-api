# Go RESTful API Boilerplate

This boilerplate serves as an example for the Distributed Systems laboratory and it provides the following features:
* RESTful endpoints for users: register, login, get-profile and get-all
* Unit test example
* JWT-based auth
* Error handling 
* Structured logging
* CRUD operations on a database table

Uses the following Go packages:
* Web framework: [gin](https://github.com/gin-gonic/gin)
* Database access: [gorm](https://gorm.io/index.html)
* Database management system: [sqlite](https://www.sqlite.org/index.html) 
* Logging: [logrus](https://github.com/sirupsen/logrus)
* JWT: [jwt](https://github.com/golang-jwt/jwt)

The project is designed to aid the development of small to medium size Go projects. [Standard Go Project Layout](https://github.com/golang-standards/project-layout) is recommended otherwise.