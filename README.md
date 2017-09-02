# GitHub Public SSH Key Lister
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bhttps%3A%2F%2Fgithub.com%2FJKhawaja%2Frest-example.svg?type=shield)](https://app.fossa.io/projects/git%2Bhttps%3A%2F%2Fgithub.com%2FJKhawaja%2Frest-example?ref=badge_shield)


Swagger documentation is also available in the `swagger` directory.

## API Routes

- Host: `localhost:8080`
- **`/keys`**
    + `POST`: provides a list of public ssh keys for a given list of github usernames
        * Request Format Example (json request body): `["dave", "tom"]`

## Build

Assuming a working Go Environment .... Inside repo's top directory: `go get ./...` + `go build .`

Can vendor dependencies using the `dep` tool. 

Get `dep`: `$ go get -u github.com/golang/dep/cmd/dep`

Run: `dep ensure` inside this repository (this will build a vendor directory based off the dependency manifest files adhering to any constraints).


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bhttps%3A%2F%2Fgithub.com%2FJKhawaja%2Frest-example.svg?type=large)](https://app.fossa.io/projects/git%2Bhttps%3A%2F%2Fgithub.com%2FJKhawaja%2Frest-example?ref=badge_large)