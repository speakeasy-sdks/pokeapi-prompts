# MoveAilment
(*MoveAilment*)

### Available Operations

* [MoveAilmentList](#moveailmentlist)
* [MoveAilmentRead](#moveailmentread)

## MoveAilmentList

### Example Usage

```go
package main

import(
	pokeapi "PokeAPI/v3"
	"context"
	"PokeAPI/v3/pkg/models/operations"
	"log"
)

func main() {
    s := pokeapi.New()

    ctx := context.Background()
    res, err := s.MoveAilment.MoveAilmentList(ctx, operations.MoveAilmentListRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.MoveAilmentListRequest](../../pkg/models/operations/moveailmentlistrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.MoveAilmentListResponse](../../pkg/models/operations/moveailmentlistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## MoveAilmentRead

### Example Usage

```go
package main

import(
	pokeapi "PokeAPI/v3"
	"context"
	"PokeAPI/v3/pkg/models/operations"
	"log"
)

func main() {
    s := pokeapi.New()

    ctx := context.Background()
    res, err := s.MoveAilment.MoveAilmentRead(ctx, operations.MoveAilmentReadRequest{
        ID: 654813,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.MoveAilmentReadRequest](../../pkg/models/operations/moveailmentreadrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.MoveAilmentReadResponse](../../pkg/models/operations/moveailmentreadresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
