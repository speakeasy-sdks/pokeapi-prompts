# MoveLearnMethod
(*MoveLearnMethod*)

### Available Operations

* [MoveLearnMethodList](#movelearnmethodlist)
* [MoveLearnMethodRead](#movelearnmethodread)

## MoveLearnMethodList

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
    res, err := s.MoveLearnMethod.MoveLearnMethodList(ctx, operations.MoveLearnMethodListRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.MoveLearnMethodListRequest](../../pkg/models/operations/movelearnmethodlistrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.MoveLearnMethodListResponse](../../pkg/models/operations/movelearnmethodlistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## MoveLearnMethodRead

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
    res, err := s.MoveLearnMethod.MoveLearnMethodRead(ctx, operations.MoveLearnMethodReadRequest{
        ID: 268867,
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.MoveLearnMethodReadRequest](../../pkg/models/operations/movelearnmethodreadrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.MoveLearnMethodReadResponse](../../pkg/models/operations/movelearnmethodreadresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
