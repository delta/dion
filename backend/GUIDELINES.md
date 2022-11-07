# Dion

## Guidelines

### General Guidelines

1. Test everything possible
2. Follow the below package specific guidelines

### controllers package

1. This will be handling all the logical part.
2. Should call functions from `repository` package to do db related stuff
3. Should never try to actually connect to db and do things with it inside this
   package
4. The functions defined here shouldn't be the ones writing the response.
   Neither should they be accepting `*gin.Context` as parameter. Doing this will
   make it easier for us to test everything.

### repository package

1. All the db related things are to be done only within this package. Nothing
   should leak to other packages.

### models package

1. All the Request,Response, DB model types are to be defined within this
   package.

### server package

1. Add handler functions for the routes. All the processing for the handler
   function is to be delegated to controller.
2. The return value from the controller should be sent as response
   appropriately, with some transformation if necessary.
3. Annotate all the handlers properly. This is so as to generate api
   documentation for the server.
4. Register the route in routes array which is present in routes.go
