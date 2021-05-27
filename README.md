**Table of Contents**

- [NearestPrime Project](#nearestprime-project)
  - [About the project](#about-the-project)
    - [Design](#design)
    - [API docs](#api-docs)
    - [Image](#images)
  - [Running](#running)
    - [Layout](#layout)
  - [Additional Notes](#notes)
  - [Reference](#reference)

# NearestPrime Project

## About the project

The project is simply a web application which returns the nearest prime number closest to a number
provided by a user.

For instances, when a user supplies 100 to the service, it should return a value of 97.

This is purely written in golang

### Design

My initial approach in tackling this problem was to iterate from n-1 to 2 (where n is the number supplied 
by the user) and then check if the number was a prime. If it is a prime, I return and break out of the loop. 

The approach worked fine well but was really inefficient for large numbers. I then decide to rather pre-comupte 
the primes and store them in a file on disk but realized the size of the file grew large. Reading this at runtime 
will put stress on the application which migh take a long while to start.

Finally, I decided that I will generate the primes stored in an array in increasing order which is configurable. 
Then I can use a binary search to find the nearest prime number less the user supplied value.

### API docs

Swagger was the used for API documenation.

### Images

When application starts up, the swagger ui is initialized and you can access this as shown below

![](/img/img1.png?raw=true "Default Page")

To check to see if the service is up, the status check endpoint is used

![](/img/img2.png?raw=true "Server status")

To find the nearest prime, the nearest prime endpoint is supplied with a number

![](/img/img3.png?raw=true "Finding the neareset prime")


## Running

Simply use the start script or docker script in the base directory

```bash
m3d@MyMac:~/Projects/Go/nearestprimes (master)  
$ ./start.sh 
Building project
Running project
2021/05/27 14:20:17 {ServerAddress:localhost:8080 LargestPrime:9999999}
2021/05/27 14:20:17 Initializing primes
2021/05/27 14:20:17 Done prime initialization
2021/05/27 14:20:17 Starting application @ localhost:8080
```

### Layout

Below is the program layout

```tree
├── Dockerfile
├── README.md
├── app.env
├── cmd
│   └── appserver
│       └── main.go
├── docker.sh
├── docs
│   └── nearestprimes
│       ├── docs.go
│       ├── swagger.json
│       └── swagger.yaml
├── go.mod
├── go.sum
├── img
│   ├── img1.png
│   ├── img2.png
│   └── img3.png
├── pkg
│   └── http
│       └── rest
│           ├── controller
│           │   └── apicontroller.go
│           ├── model
│           │   └── apimodel.go
│           ├── service
│           │   └── apiservice.go
│           └── util
│               └── config.go
├── start.sh
└── tests
    ├── integration_test.go
    └── unit_test.go
```

A brief description of the layout:

* `Dockerfile` contains docker configuration.
* `cmd` contains main packages, each subdirecoty of `cmd` is a main package.
* `docs` for api project documentation.
* `img` for the swagger image snapshots.
* `pkg` places most of project business logic and locate `http/rest` package.
* `test` holds all tests.
* `.gitignore` these are files or directories git will ignore.
* `app.env` is the configuration file for project containing the binding address and the largest prime number to generate.
* `go.mod` golang modules used.
* `go.sum` module checksum.
* `README.md` is a detailed description of the project.
* `start.sh` is used to build and run the application.
* `docker.sh` is used to build and run the docker version of the application.


## Notes
* Configurations for start up can over-written at runtime through the use of environment variables
```bash
m3d@MyMac:~/Projects/Go/nearestprimes (master)  
$ SERVER_ADDRESS=localhost:9000 LARGEST_PRIME=1000000 GIN_MODE=release  ./nearestprime 
2021/05/27 14:15:29 {ServerAddress:localhost:9000 LargestPrime:1000000}
2021/05/27 14:15:29 Initializing primes
2021/05/27 14:15:29 Done prime initialization
2021/05/27 14:15:29 Starting application @ localhost:9000
```
* They could be instances where the user supplied input might not exist in the generated prime list
  * In such scenarios the application returns the last value in the list
* For optimal perform I suggest that these primes be computed and cache in a server(redis) for each lookup

## References

Library used in the generation of the primes

https://pkg.go.dev/github.com/kavehmz/prime


I relied on the sieve of Eratosthenes to generate these primes because it is better at generate large prime as 
compared to sieve of Sundaram. Below is a like to a white paper where comparism are made

https://iopscience.iop.org/article/10.1088/1742-6596/978/1/012123/pdf
