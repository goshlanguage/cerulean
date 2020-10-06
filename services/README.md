# Creating a New Service

1. Create a directory/package that will contain your service (e.g. `/services/subscriptions`)
1. Make a `README.md` file (e.g. `/services/subscriptions/README.md`) that gives an overview of your service along with a sample JSON response
1. Create a `models.go` file in your new service package (e.g. `/services/subscriptions/models.go`) and add in a constructor function that returns a struct that represents your service's JSON response populated with sane defaults
1. Create a `handlers.go` file and implement whatever handler functions you need (return JSON using the models you created constructors for)
    - We use [Echo](https://echo.labstack.com/) for our HTTP routing so you'll need to make your handler functions match a structure similar to the one below:
    ```
    func hello(c echo.Context) error {
  	return c.String(http.StatusOK, "Hello, World!")
    }
    ```
1. Make a `service.go` file in your new directory (e.g. `/services/subscriptions/service.go`), implement `GetHandlers()` as defined in `/services/interface.go`, and add your new handler functions (defined in `handlers.go` that you created in a step above) to the map that's returned
1. To finish things off, be a good person and create a `handlers_test.go` file and write a couple unit tests
