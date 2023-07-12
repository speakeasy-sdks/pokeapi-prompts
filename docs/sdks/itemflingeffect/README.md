# ItemFlingEffect

### Available Operations

* [ItemFlingEffectList](#itemflingeffectlist)
* [ItemFlingEffectRead](#itemflingeffectread)

## ItemFlingEffectList

### Example Usage

```go
package main

import(
	"context"
	"log"
	"PokeAPI"
	"PokeAPI/pkg/models/operations"
)

func main() {
    s := pokeapi.New()

    ctx := context.Background()
    res, err := s.ItemFlingEffect.ItemFlingEffectList(ctx, operations.ItemFlingEffectListRequest{
        Limit: pokeapi.Int64(473600),
        Offset: pokeapi.Int64(264555),
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
| `request`                                                                                      | [operations.ItemFlingEffectListRequest](../../models/operations/itemflingeffectlistrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.ItemFlingEffectListResponse](../../models/operations/itemflingeffectlistresponse.md), error**


## ItemFlingEffectRead

### Example Usage

```go
package main

import(
	"context"
	"log"
	"PokeAPI"
	"PokeAPI/pkg/models/operations"
)

func main() {
    s := pokeapi.New()

    ctx := context.Background()
    res, err := s.ItemFlingEffect.ItemFlingEffectRead(ctx, operations.ItemFlingEffectReadRequest{
        ID: 186332,
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
| `request`                                                                                      | [operations.ItemFlingEffectReadRequest](../../models/operations/itemflingeffectreadrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.ItemFlingEffectReadResponse](../../models/operations/itemflingeffectreadresponse.md), error**

