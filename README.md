# graphql [![GoDoc](https://godoc.org/github.com/cusspvz/graphql?status.png)](http://godoc.org/github.com/cusspvz/graphql) [![Build Status](https://travis-ci.org/cusspvz/graphql.svg?branch=master)](https://travis-ci.org/cusspvz/graphql) [![Go Report Card](https://goreportcard.com/badge/github.com/cusspvz/graphql)](https://goreportcard.com/report/github.com/cusspvz/graphql)

Low-level GraphQL client for Go, forked from https://github.com/machinebox/graphql/

## Features

* Simple, familiar API
* Respects `context.Context` timeouts and cancellation
* Build and execute any kind of GraphQL request
* Use strong Go types for response data
* Use variables and upload files
* Simple error handling
* Supports Subscriptions out-of-the-box

## Next Features
* Subscription connection handling
* 

## Installation
Make sure you have a working Go environment. To install graphql, simply run:

```
$ go get github.com/cusspvz/graphql
```

## Usage

```go
import "context"

const reqDoc := `
    query ($key: String!) {
        items (id:$key) {
            field1
            field2
            field3
        }
    }
`

type reqStructure struct {
    items struct {
        field1 string
        field2 string
        field3 string
    }
}


func main() {

    // create a client (safe to share across requests)
    client := graphql.NewClient("https://cusspvz.io/graphql")

    // make a request
    req := graphql.NewRequest(reqDoc)

    // set any variables
    req.Var("key", "value")

    // set header fields
    req.Header.Set("Cache-Control", "no-cache")

    // define a Context for the request
    ctx := context.Background()

    // run it and capture the response
    var respData reqStructure
    if err := client.Run(ctx, req, &respData); err != nil {
        log.Fatal(err)
    }

    //subscribe 

    subClient, err := client.SubscriptionClient(ctx, req.Header)
    if err != nil {
        log.Fatal(err)
    }


    sub, err := subClient.Subscribe(`
        subscription {
    `)
}
```

### File support via multipart form data

By default, the package will send a JSON body. To enable the sending of files, you can opt to
use multipart form data instead using the `UseMultipartForm` option when you create your `Client`:

```
client := graphql.NewClient("https://cusspvz.io/graphql", graphql.UseMultipartForm())
```

For more information, [read the godoc package documentation](http://godoc.org/github.com/cusspvz/graphql) or the [blog post](https://blog.cusspvz.io/a-graphql-client-library-for-go-5bffd0455878).

## Thanks

Thanks to everyone who contributed to https://github.com/machinebox/graphql.
