# MoveDamageClass
(*MoveDamageClass*)

### Available Operations

* [MoveDamageClassList](#movedamageclasslist)
* [MoveDamageClassRead](#movedamageclassread)

## MoveDamageClassList

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
    res, err := s.MoveDamageClass.MoveDamageClassList(ctx, operations.MoveDamageClassListRequest{})
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
| `request`                                                                                      | [operations.MoveDamageClassListRequest](../../models/operations/movedamageclasslistrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.MoveDamageClassListResponse](../../models/operations/movedamageclasslistresponse.md), error**


## MoveDamageClassRead

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
    res, err := s.MoveDamageClass.MoveDamageClassRead(ctx, operations.MoveDamageClassReadRequest{
        ID: 349044,
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
| `request`                                                                                      | [operations.MoveDamageClassReadRequest](../../models/operations/movedamageclassreadrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.MoveDamageClassReadResponse](../../models/operations/movedamageclassreadresponse.md), error**

