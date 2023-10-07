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
	"context"
	"log"
	pokeapi "PokeAPI"
	"PokeAPI/pkg/models/operations"
)

func main() {
    s := pokeapi.New()

    ctx := context.Background()
    res, err := s.MoveBattleStyle.MoveBattleStyleList(ctx, operations.MoveBattleStyleListRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.MoveBattleStyleListRequest](../../models/operations/movebattlestylelistrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.MoveBattleStyleListResponse](../../models/operations/movebattlestylelistresponse.md), error**


## MoveBattleStyleRead

### Example Usage

```go
package main

import(
	"context"
	"log"
	pokeapi "PokeAPI"
	"PokeAPI/pkg/models/operations"
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

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.MoveBattleStyleReadRequest](../../models/operations/movebattlestylereadrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.MoveBattleStyleReadResponse](../../models/operations/movebattlestylereadresponse.md), error**

