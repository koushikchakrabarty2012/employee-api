# employee-api
Step1 :This is a simple rest api with a set of Employee Data and having different operations on it .
  -GET -GetAllEmployees
  -POST -PostSingleEmployee
  -DELETE -DeleteSingleEmployee
  -PUT -UpdateSingleEmployee
  -GETByID -GetSingleEmployee with id
  -Another operation is GetAllManagers which returns all managers from system which has certain employee set .

Step2:Swagger 2.0 OpenApi Specification added on top of it .

    Swagger documentation is generated from the code annotations inside the source using go-swagger.

    Go swagger installed with the following command:

    go get -u github.com/go-swagger/go-swagger/cmd/swagger

    Generate the documentation using the command:

    swagger generate spec -o ./swagger.yaml --scan-models

    After running the application:

    go run main.go

      


Step3:Dockerised the application as an image.
  commands:
            1.docker build -t employee-api:1 .
            2.docker run -p 9090:8081 employee-api
                  rule:docker run -p <local-machine-port>:<container port> imageid/name
                       docker run -d -p 9090:8080 healthplanner-api => detached mode
            3. docker login
            4. docker tag employee-api:1 thisisjgec06/employee-api:1
                  rule: docker tag SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]
            5. After you properly tag the image, use the docker push command to push your image to the Docker Hub registry:
              $ docker push thisisjgec06/employee-api:1

