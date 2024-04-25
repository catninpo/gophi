# Gophi
Gophi is intended to be a basic example of how a REST API could be laid out 
utilising the latest HTTP Routing in Go 1.22.x. The layout draws from the approach
put forward by Ben Johnson in [this article](https://www.gobeyond.dev/standard-package-layout/).

## Getting Started
If you wish to have a play around with the project you can clone it down and either build 
or run the application from the `cmd/` directory. You can specify an optional port but by
default this application will run on port `:8888`.

This application expects one environmental variable to be set for the connection to a 
local postgres database. It will look for this connection string in the `DATABASE_URL` 
environment variable. You can set this with the following:

```bash
export DATABASE_URL="{POSTGRES_CONNECTION_STRING}"
```

If you do not have a local postgres database instance, you can update the `postgres/user.go` 
file with the following temporarily to get the project running:

```go 
func (u *UserService) UserByID(id int) (*gophi.User, error) {
    return &gophi.User{
        ID: 1,
        Name: "Gophi",
    }, nil
}
```

The project also implements a very basic `templ` implementation, you can find
installastion instructions at the [a-h/templ](https://github.com/a-h/templ). To aid
in building the project a Makefile has been provided so you can simply run `make` to 
get the project running.

## Contributing
If you see areas of improvement within the codebase, feel free to submit a pull request with 
any changes and their benefits. 
