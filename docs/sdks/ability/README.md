# Ability

### Available Operations

* [AbilityList](#abilitylist)
* [AbilityRead](#abilityread)

## AbilityList

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
    s := PokeAPI.New()

    ctx := context.Background()
    res, err := s.Ability.AbilityList(ctx, operations.AbilityListRequest{
        Limit: PokeAPI.Int64(715190),
        Offset: PokeAPI.Int64(844266),
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.AbilityListRequest](../../models/operations/abilitylistrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.AbilityListResponse](../../models/operations/abilitylistresponse.md), error**


## AbilityRead

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
    s := PokeAPI.New()

    ctx := context.Background()
    res, err := s.Ability.AbilityRead(ctx, operations.AbilityReadRequest{
        ID: 602763,
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.AbilityReadRequest](../../models/operations/abilityreadrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.AbilityReadResponse](../../models/operations/abilityreadresponse.md), error**

