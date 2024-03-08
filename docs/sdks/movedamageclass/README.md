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
	pokeapi "PokeAPI/v3"
	"context"
	"PokeAPI/v3/pkg/models/operations"
	"log"
)

func main() {
    s := pokeapi.New()

    ctx := context.Background()
    res, err := s.MoveDamageClass.MoveDamageClassList(ctx, operations.MoveDamageClassListRequest{})
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
| `request`                                                                                          | [operations.MoveDamageClassListRequest](../../pkg/models/operations/movedamageclasslistrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.MoveDamageClassListResponse](../../pkg/models/operations/movedamageclasslistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## MoveDamageClassRead

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
    res, err := s.MoveDamageClass.MoveDamageClassRead(ctx, operations.MoveDamageClassReadRequest{
        ID: 349044,
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
| `request`                                                                                          | [operations.MoveDamageClassReadRequest](../../pkg/models/operations/movedamageclassreadrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.MoveDamageClassReadResponse](../../pkg/models/operations/movedamageclassreadresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
