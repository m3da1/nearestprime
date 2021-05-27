**Table of Contents**

- [NearestPrime Project](#nearestprime-project)
  - [About the project](#about-the-project)
    - [Design](#design)
    - [API docs](#api-docs)
    - [Image](#images)
  - [Getting started](#getting-started)
    - [Layout](#layout)
  - [Additional Notes](#notes)
  - [Reference](#reference)

# NearestPrime Project

## About the project

The project is simply a web application which returns the nearest prime number closest to a number
provided by a user.

For instances, when a user supplies 100 to the services, it should return a value of 97.

This is purely written in golang

### Design

My initial approach in tackling this problem was to iterate from n-1 to 2 (where n is the number supplied 
by the user) and then check if the number was a prime. If it is a prime, I return and break out of the loop. 

The approach worked fine well but was really inefficient for large numbers. I then decide to rather pre-comupte 
the primes and store them in a file on disk but realized the size of the file was large. Reading this and runtime 
will put stress and the application making it take a long while to start up.

Finally, I decided that I will generate the primes stored in an array in increasing order which is configurable. 
Then I can use a binary search to find the nearest prime number less the user supplied value.

I relied on the sieve of Eratosthenes to generate these primes because it is better at generate large prime as 
compared to sieve of Sundaram

### API docs

Swagger was the used for API documenation.

### Images

When application starts up, the swagger ui is initialized and you can access this as shown below

![](/img/img1.png?raw=true "Default Page")

To check to see if the service is up, the status check endpoint is used

![](/img/img2.png?raw=true "Server status")

To find the nearest prime, the nearest prime endpoint is supplied with a number

![](/img/img3.png?raw=true "Finding the neareset prime")


## Getting started

Below we describe the conventions or tools specific to golang project.

### Layout

```tree
├── README.md
├── app.env
├── cmd
│   └── appserver
│       └── main.go
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

* `.github` has two template files for creating PR and issue. Please see the files for more details.
* `.gitignore` varies per project, but all projects need to ignore `bin` directory.
* `.golangci.yml` is the golangci-lint config file.
* `Makefile` is used to build the project. **You need to tweak the variables based on your project**.
* `CHANGELOG.md` contains auto-generated changelog information.
* `OWNERS` contains owners of the project.
* `README.md` is a detailed description of the project.
* `bin` is to hold build outputs.
* `cmd` contains main packages, each subdirecoty of `cmd` is a main package.
* `build` contains scripts, yaml files, dockerfiles, etc, to build and package the project.
* `docs` for project documentations.
* `hack` contains scripts used to manage this repository, e.g. codegen, installation, verification, etc.
* `pkg` places most of project business logic and locate `api` package.
* `release` [chart](https://github.com/caicloud/charts) for production deployment.
* `test` holds all tests (except unit tests), e.g. integration, e2e tests.
* `third_party` for all third party libraries and tools, e.g. swagger ui, protocol buf, etc.
* `vendor` contains all vendored code.

## Notes

* Makefile **MUST NOT** change well-defined command semantics, see Makefile for details.
* Every project **MUST** use `dep` for vendor management and **MUST** checkin `vendor` direcotry.
* `cmd` and `build` **MUST** have the same set of subdirectories for main targets
  * For example, `cmd/admin,cmd/controller` and `build/admin,build/controller`.
  * Dockerfile **MUST** be put under `build` directory even if you have only one Dockerfile.

## References

library used

https://pkg.go.dev/github.com/kavehmz/prime

comparism of sieves

https://iopscience.iop.org/article/10.1088/1742-6596/978/1/012123/pdf