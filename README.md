## Go-Grpc-Server-Stram 
* The Grpc server streams the data to the client peer.

## Dependencies 
* Install Golang latest version
* check out to this url =  https://github.com/golang/dep
* after the dependency installation do the following command
* Go to the project path then execute `dep ensure -v` command to install the dependency.

## Run Server

* to the server folder run the command `go run main.go`

## Run Client

* to the client folder run the command `go run main.go`

## Output - client

    {
        client program
    server streaming roc
    Response from greet many times %v HellobilalNumber0
    Response from greet many times %v HellobilalNumber1
    Response from greet many times %v HellobilalNumber2
    Response from greet many times %v HellobilalNumber3
    Response from greet many times %v HellobilalNumber4
    Response from greet many times %v HellobilalNumber5
    Response from greet many times %v HellobilalNumber6
    Response from greet many times %v HellobilalNumber7
    Response from greet many times %v HellobilalNumber8
    Response from greet many times %v HellobilalNumber9
    End of the stream
    }