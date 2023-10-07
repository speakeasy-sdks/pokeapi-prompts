# ContestEffect
(*ContestEffect*)

### Available Operations

* [ContestEffectList](#contesteffectlist)
* [ContestEffectRead](#contesteffectread)

## ContestEffectList

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
    res, err := s.ContestEffect.ContestEffectList(ctx, operations.ContestEffectListRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ContestEffectListRequest](../../models/operations/contesteffectlistrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.ContestEffectListResponse](../../models/operations/contesteffectlistresponse.md), error**


## ContestEffectRead

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
    res, err := s.ContestEffect.ContestEffectRead(ctx, operations.ContestEffectReadRequest{
        ID: 338469,
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ContestEffectReadRequest](../../models/operations/contesteffectreadrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.ContestEffectReadResponse](../../models/operations/contesteffectreadresponse.md), error**

