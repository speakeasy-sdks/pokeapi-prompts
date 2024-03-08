# MoveBattleStyle
(*MoveBattleStyle*)

### Available Operations

* [MoveBattleStyleList](#movebattlestylelist)
* [MoveBattleStyleRead](#movebattlestyleread)

## MoveBattleStyleList

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
    res, err := s.MoveBattleStyle.MoveBattleStyleList(ctx, operations.MoveBattleStyleListRequest{})
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
| `request`                                                                                          | [operations.MoveBattleStyleListRequest](../../pkg/models/operations/movebattlestylelistrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.MoveBattleStyleListResponse](../../pkg/models/operations/movebattlestylelistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## MoveBattleStyleRead

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
    res, err := s.MoveBattleStyle.MoveBattleStyleRead(ctx, operations.MoveBattleStyleReadRequest{
        ID: 152461,
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
| `request`                                                                                          | [operations.MoveBattleStyleReadRequest](../../pkg/models/operations/movebattlestylereadrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.MoveBattleStyleReadResponse](../../pkg/models/operations/movebattlestylereadresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
