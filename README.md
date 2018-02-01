# Gopher-Requests
Custom HTTP client for Golang

## Installation
```sh
go get github.com/TomOrth/Gopher-Requests
```

To use it in your project:
``` go
import {
    ...
    greq "github.com/TomOrth/Gopher-Requests"
    ...
}
```

## Usage
For this example, lets assume we have an api at the url "http://www.api.dog.com", an endpoint "/dog", and struct called `Dog` to store the results into.  With all this, we do:

```go
client := greq.Client{"http://www.api.dog.com"}
```

This makes a client to use to make requests.  Now we need to send it 2 headers, User-Agent and Content-Type. Once we do, we then store the results inside the instance of our `Dog` struct called `dog`. You're able to chain the Header creations along with the execution of the request into `dog` like so:

```go
err := client.Get("/dog").Header("User-Agent", "<Name>").Header("Content-Type", "JSON").Exec(&dog)
```

Please see `example` for another usage example

## Licensing
This software is available under the MIT License.
