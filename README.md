# Interest points API
Interest point API is an API service that provides ability to work with points of interest. Point of interest is an object that potentially might be interesting for user (e.g. museum, theatre, graffity etc).

## Deploy
To deploy application you need to use docker and docker-compose. Docker-compose.yml file contains 2 services: mysql (for database API service works with) and API server, written on golang. You can configure some variables that uses to deploying application in .env file. By default, mysql database deploys on port 3306, API server - on port 3000.

*Note:* if you'd like to change any parameters for database, you, probably, need to change database_url parameter in configs/apiserver.toml file.

**Note:** by default in configs/apiserver.toml file connection to database establishes through host.docker.internal:3306 address. If you'd like to deploy application in machine with any OS but MacOS, you need to change this address to the one, that represents localhost for docker.

After environment configuration, call **docker-compose up** in terminal from the main project folder.

After successful deploying, there will be 6 points in database (3 in Moscow and 3 in Perm) and default user **admin** with password **admin**.

## Developing process
To work with points of interest, there were implemented RESTfull API service with layer architecture. To work with database package *database/sql* was used. I didn't use ORM like gorm, because we have small number of objects (only point, point type, city and user) and I find it no sense to use ORM with a small number of objects.

In the end of doing this test task, I remebered about gRPC technology that I used in one project while studying at the university. Maybe it would be better (in terms of speed and server load) and easier to use gRPC, but I didn't have enough time to rewrite service, using gRPC.

The service supports the next methods (firstly I wanted to create Swagger documentation, but didn't have enough time for it):
**GET /api/points** - returns list of points by city. City is mandatory parameter that must contains city ID.
**GET /api/points/{id}** - return point by its ID.
**POST /api/points **- creates new point of interest. Request body must correspond to struct in file models/domain/point-request-body.go
**PUT /api/points/{id}** - updates existent point. Request body must correspond to struct in file models/requests/point-request-body.go.
**GET /api/nearest-point** - returns list of points that placed in mentioned radius. Lon, lat and radius are mandatory parameters.
**GET /api/health** - health endpoint.
**POST /api/authenticate** - authenticates user. Request body must correspond to struct in file models/domain/auth-request.go.

There were some mandatory requirments and some additional. Let's take a look on each of them.

### Mandatory requirements
#### Project must be published on GitHub od GitLab
I guess, you're reading this readme on GitHub right now ;).

#### Stack: Go, MySQL, Redis
I used Go and MySQL, but didn't use Redis, because didn't use cache in that project and decided to store data in relational database rather than in-memory storage.

#### README.md file with project description and instruction of building and launching application (preferably on English)
I guess, you're reading this readme on English right now ;).

#### All exceptions must be handled correctly, all error format in API responses must be in consistent format
All exception messages returns in consistent JSON format and has an appropriate http status code. Check function JSONError in utils/json-error.go file. 

### Additional requirements
#### Http caching repeating requests to get data
Not implemented.

#### SQL requests caching
Not implemented.

#### Implement authorization mechanism through JWT
Implemented /api/authenticate endpoint to authenticate. Implemented only receiving new token, refreshing not implemented because of lack of time. 

Token refreshing must be implemented, especially if API service will be used by SPA application, due to not force user to authenticate every time his token will be expired.

Token lifetime is 1 hour, this value is hardcoded. Not the best way to do that, it's better to store token lifetime in config file. It was done just to simplify - I guess, 1 hour wiil be enough to test API service :).

#### Restrict number of requests to API for client or IP-address
Not implemented.

#### Add HTTP headers to server responses for safe API methods calling from web applications
CORS headers are adding during the router configuration, allowed origins, methods and headers mentioned in config/apiserver.toml file.

#### Implement health-checking endpoint
Heatlh-checking endpoint implemented by address /api/health and doesn't required authentication. In case in our application checks just connection to database.

## What should have been done better?
- Domain models should be validated. Haven't done that due to lack of time, also creating domain models happens only after validation request models, so, in case of our service there should be no errors related to this.
- For model validations should be used some package.
- Probably should have used gRPC.
- Refresh token mechanism must be implemented.
- Swagger documentation for API methods is better, than describing them in readme.