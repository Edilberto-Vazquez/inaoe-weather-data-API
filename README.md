# INAOE-weather-data-API
Weather data REST API for INAOE weather stations.
This page was development with Golang 1.17, Gin Web Framework and Gorm as ORM.
The deploy was made to heroku with a Database in AWS.

For more information about how to deploy to heroku please refer to the official page https://devcenter.heroku.com/articles/getting-started-with-go

To put the necessary modules inside the main package, run `go mod vendor`

TO-DO
-------------

+ Change the DB from AWS

Libraries used
-------------

+ "gin"
+ "cors"
+ "gorm"
