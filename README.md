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


### Config YAML Syntax
The configuration for mocksrv is provided via a YAML file, defining _routes_, _directories_ and _proxys_.

| *Section Name* | *Description* |
| -------------- | ------------- |
| routes | Requests to these routes should return the response as dictated by the user |
| directories | Any request to these routes will serve static content from the filesystem. |
| proxys | A request to a proxy route will be forwarded on to a remote URL |

The base underlying configuration to these sections is similar - a list of endpoints that can then be configured based up its type. The most basic example of this can look like

```
routes:
  /path/to/somewhere:
    # route configuration. See below
directories:
  /path/to/somewhere/else:
    # directory configuration. See below
proxy:
  /path/to/nowhere:
    # proxy configuration. See below

```

#### routes
Routes con

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

## TODO:

* Add Tests
* Finish this README
* Enhance Routes
   * CORS Support
   * Provide Auth Validation (string matching)
* Add a certificate option which would spin up the service with SSL
* Provide more configuration examples
* Finalise Logging Format
* Provide a SIGHUP handler on the service to reload config changes
* Provide a web interface, including YAML generator?