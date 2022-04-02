# Documentation

The project include a Makefile

1. Run the build command. Example: make build
2. Up the database with the db-start command. Example: make db-start
3. Load data examples with testdata command. Example: make testdata
4. Up the local server with the run command. Example: make run
5. When you finish with the requests run the db-stop command. Example: make db-stop
6. And run the clean command. Example: make clean

# Url's

The url base is http://localhost:8080/v1/

  * GET /healthcheck: a dummy ping to test the local server
  * GET /v1/products/:id: returns the information of a product
  * POST /v1/products: creates a new product
  * PUT /v1/products/:id: updates an existing product
  * DELETE /v1/products/:id: deletes a product
  
Try the URL http://localhost:8080/healthcheck in a browser, and you should see something like "OK v1.0.0" displayed.

# Tests

Run the test command (Example: make test) to watch the coverage code in the command line. The project include a postman directory, inside there is a json file,  import it with postman.

# Requirements

The project need a Go compiler, Postman client and Docker Desktop. 

# References
I used this project as base https://github.com/qiangxue/go-rest-api
