# Postman documentation generator

Postman documentation generator is a tool for generating models from the structures in source files and adding them to postman collection schema.

## Download and install

```go get github.com/madappgang/postman-doc-generator```

## Usage

For example, if we have a struct type called `User`,

```go

type User struct {
    // Unique identifier of the user
    ID           string    `json:"id"`
    // Email of the user
    Email        string    `json:"email"`
    HashPassword string    `json:"-"`
    // Status of the user account
    IsActive     bool      `json:"is_active"`
    // Time of user creation
    CreatedAt    time.Time `json:"created_at"`
}

```

run the command below in the same directory

postman-doc-generator -struct=User

it will create the model using Markdown syntax

```text

### User

Parameter | Type | Description
--------- | ---- | -----------
id | string | Unique identifier of the user
email | string | Email of the user
is_active | bool | Status of the user account
created_at | Time | Time of user creation

```

and put it into the file postman_collection.json.

Typically this process would be run using go generate, like this:

```go

//go:generate postman-doc-generator -struct=User

```

With no arguments, it processes the package in the current directory. Otherwise, the flag -source accepts the name of directory with Go source files that belong to a single Go package.

The -struct flag accepts a comma-separated list of structs for generating multiple structs at once.

The -output flag accepts a path to the postman schema file. By default, it is postman_collection.json.
