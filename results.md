# Go Programming Challenge Results

This is a documentation of the solution to solve the challenge.

## Required tools

To be able to build and run the solution, you need these tools:
* golang 1.20
* docker
* docker-compose

## Build and Run

The commands to control the application lifecycle is done via make to simplify it, you can use the following:
* `make build`: This command will build the image locally and write the binary to bin/facts
* `make clean`: This command removes the containers form preious runs, and also cleans the bin directory
* `make image`: This command builds the docker image from the current resources, and tag it `facts:latest` by default.
* * The tag can be modified by passing `IMAGE_NAME`, and `VERSION` to the command, which will create the tag `IMAGE_NAME:$VERSION`
* `make run`: This command brings the docker-compose cluster up, by default the application image `facts:latest` is used. for further details about the docker compose config check [docker-compose.yml](docker-compose.yml)
* `make test`: This command runs the tests locally

## Infrastructure

There are 2 main components for the application to work(`mongodb`,`facts app`), and an optional component(`mongo-express`), which are deployed in the end as docker containers using docker compose to achive portability and consistency.

I used `docker-compose` as it's the simplist container orcastration tool for me, and it works well with small projects, however for bigger project `Kubernetes` might be a better tool.

## Secrets and Variables

All the secrets and variables are currently injected at deployment time as environment variables, by default by `.env` file with default values, or they can be overriden by exporting them as shell environment. I have choosen this method as it's the easiest way to inject and access the variables, and also it's flixability with different environments, as it can work with `docker-compose`, `Kubernetes`, `AWS Lambda Function`, `Local machine`, ...etc, also for production secrets, depends on the deployment pipline you can use different methods to store them, like `Kubernates Secrets`, `Encrypted Files`, `GitHub or Gitlab Secrets`, `Secret managers in Cloud Platforms`, ...etc, and then inject them on deployment or request them on-demand.

## Development choices

The database server is initialised with a new database and collection to hold the data, and the provided sample data is injected directly at deployment time via a small `js` init script.

I used mongo 5 for the project even though there is a stable version 6, however, I wanted to use the `mongo-express` tool, which is currently not working with version 6 because of using some legace operations that were removed in mongo version 6

For the application connection with the database I used the official mongo-driver fo golang. The connection initiation is done once per application run via the `init()` method (`singilton`), so you can use the same connection for all further commands, which is saving a lot of processing time and resources instead of creating a new connection each time.

All the filters that are done on the results are running on the database server side, so that the application is lighter and more scalable.

## Challenges and limitations

As this is my first time working with MongoDB, it was the most challenging part to get it up and running with the initial data. I expect there is a way to have a validation scheme for the collection so that invalid data is ignored, triggers a warning or triggers an error, however unfortunatly I couldn't find an easy way of doing it. Therefore I corrected 3 data entries in the sample data where the numbers where written as strings, which caused some failures with decoding data later.

I am aware that security wise, the api is poor, as it has no authentiaction and has no TLS protocol. The reasons are:
* I think have a propper user athentication needs a complete auth system with a propper auth mechanisim, and implementing such solution is increasing the complexity and the time needed.
* For TLS it's relatively easy to enable it and implement it, however I would still self-sign the certificates and would make it harder on your side to run the application, unless I upload all the certificates with the code, which is making it useless.

I omitted the integration tests for the solution as I tested it manually using `Postman`, and as I am not familiar with the MongoDB libraries I am not familier with how to test it (mock it), which will take a extra time for me to implement. However, I beleave that the application implementation is testable.