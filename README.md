# Postman documents generator

Postman documents generator is a tool to generate models from the structures of the source files and add them to postman collection schema.

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

running this command in the same directory

postman-doc-generator -struct=User

creates the model used Markdown syntax

```text
    ### User

    Parameter | Type | Description
    --------- | ---- | -----------
    id | string | Unique identifier of the user
    email | string | Email of the user
    is_active | bool | Status of the user account
    created_at | Time | Time of user creation
```

and puts it into the file postman_collection.json.

Typically this process would be run using go generate, like this:

```go
    //go:generate postman-doc-generator -struct=User
```

With no arguments, it processes the package in the current directory. Otherwise, the flag -source accepts name a single directory holding a Go package or a set of Go source files that represent a single Go package.

The -struct flag accepts a comma-separated list of structs so a single run can generate methods for multiple structs.
