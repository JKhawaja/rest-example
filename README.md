# GitHub Public SSH Key Lister

Swagger documentation is also available in the `swagger` directory.

## API Routes

- Host: `localhost:8080`
- **`/keys`**
    + `POST`: provides a list of public ssh keys for a given list of github usernames
        * Request Format Example (json request body): `["dave", "tom"]`

## Build

Assuming a working Go Environment ....

All dependencies have been vendored and committed to the public repo. This means that it should build with a simple `go build .`