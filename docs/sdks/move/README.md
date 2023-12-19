# Move
(*Move*)

### Available Operations

* [MoveList](#movelist)
* [MoveRead](#moveread)

## MoveList

### Example Usage

```go
package main

import(
	pokeapi "PokeAPI/v2"
	"context"
	"PokeAPI/v2/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := pokeapi.New()

    ctx := context.Background()
    res, err := s.Move.MoveList(ctx, operations.MoveListRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.MoveListRequest](../../pkg/models/operations/movelistrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.MoveListResponse](../../pkg/models/operations/movelistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## MoveRead

### Example Usage

```go
package main

import(
	pokeapi "PokeAPI/v2"
	"context"
	"PokeAPI/v2/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := pokeapi.New()

    ctx := context.Background()
    res, err := s.Move.MoveRead(ctx, operations.MoveReadRequest{
        ID: 552244,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.MoveReadRequest](../../pkg/models/operations/movereadrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.MoveReadResponse](../../pkg/models/operations/movereadresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
