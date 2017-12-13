# Simple Server Mocking
**mocksrv** is a simple server which provides predefined responses to HTTP requests. URL routes and methods are provided with details as to how the server should respond, effectively mocking out remote servers. For example, given the following config

```
    /api/users:
      get:
        body: |
            {
              "id": 123
              "name": "tyndyll"
            }
        content-type: application/json
        response-code: 200
```

When the server is started and a request made to the route `http://localhost:8000/api/users`, the body will be set to the body in the YAML

```
➜  ~ curl -vv -X GET -vv http://localhost:8000/api/users
*   Trying ::1...
* Connected to localhost (::1) port 8000 (#0)
> GET /api/users HTTP/1.1
> Host: localhost:8000
> User-Agent: curl/7.43.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Wed, 13 Dec 2017 21:45:33 GMT
< Content-Length: 36
<
{
  "id": 123
  "name": "tyndyll"
}
* Connection #0 to host localhost left intact
```

## Install
The server can be installed using the go install tool
```
go install github.com/tyndyll/mocksrv/...
```

Releases are available through the Github releases page for a number of platforms

## Usage
**mocksrv** is an HTTP service that can be started by executing the following command line
```
mocksrv [--echo] [--port <PORT>] <config.yaml>

$ mocksrv -h
Usage of mocksrv:
  -echo
    	echo request (default true)
  -port int
    	port the server should listen on (default 8000)
```

A YAML configuration file must be passed or the server will not start.


### Config.yaml Syntax
Configuration for the mocksrv is provided via a YAML file. It is used to define the route, method and the appopriate response to a request

| *Field Name* | *Description* |
| ------------ | ------------- |
| content-type | MIME type that the response should be set to |
| body | Text that should be returned |
| response-code | HTTP Status Code that should be returned |
| method | Method on which this config should be invoked upon. The method does not have to be capitalised |

```
{
    "cors": {
        "allow-origins": [],
        "allow-headers": [],
        "allow-methods": [],
    },
    "routes": [
        {
            "content-type": "text/html",
            "body": "<!DOCTYPE html><html><body><h1>Hello world!</body></html>",
            "response-code": 200,
            "method": "GET",
            "path": "/path"
        }
    ]
}
```
