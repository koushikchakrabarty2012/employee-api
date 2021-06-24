# employee-api
Step1 :This is a simple rest api with a set of Employee Data and having different operations on it .
  -GET -GetAllEmployees
  -POST -PostSingleEmployee
  -DELETE -DeleteSingleEmployee
  -PUT -UpdateSingleEmployee
  -GETByID -GetSingleEmployee with id
  -Another operation is GetAllManagers which returns all managers from system which has certain employee set .

Step2:Swagger 2.0 OpenApi Specification added on top of it .
    1.Visit https://goswagger.io/generate/spec.html for generating specification from source
    2.import github.com/go-openapi/runtime/middleware in main.go and add handler for swagger which will start a file server for serving swagger.yaml 
    //handler for documentation
      ops := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
      sh := middleware.Redoc(ops, nil)
      getRouter.Handle("/docs", sh)
      getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
    3. Swagger documentation is generated from the code annotations inside the source using go-swagger.

    Go swagger installed with the following command:

    go get -u github.com/go-swagger/go-swagger/cmd/swagger

    Generate the documentation using the command:

    swagger generate spec -o ./swagger.yaml --scan-models

    After running the application:

    go run main.go

Swagger documentation can be viewed using the ReDoc UI in your browser at http://localhost:port/docs.


Step3:Dockerised the application as an image.
  commands:
            1.docker build -t employee-api:1 .
            2.docker run -p 9090:8081 employee-api
                  rule:docker run -p <local-machine-port>:<container port> imageid/name
                       docker run -d -p 9090:8080 employee-api => detached mode
            3. docker login
            4. docker tag employee-api:1 thisisjgec06/employee-api:1
                  rule: docker tag SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]
            5. After you properly tag the image, use the docker push command to push your image to the Docker Hub registry:
              $ docker push thisisjgec06/employee-api:1

