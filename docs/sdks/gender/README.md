# Gender
(*.Gender*)

### Available Operations

* [GenderList](#genderlist)
* [GenderRead](#genderread)

## GenderList

### Example Usage

```go
package main

import(
	"context"
	"log"
	pokeapi "PokeAPI/v2"
	"PokeAPI/v2/pkg/models/operations"
)

func main() {
    s := pokeapi.New()

    ctx := context.Background()
    res, err := s.Gender.GenderList(ctx, operations.GenderListRequest{})
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
| `request`                                                                    | [operations.GenderListRequest](../../models/operations/genderlistrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GenderListResponse](../../models/operations/genderlistresponse.md), error**


## GenderRead

### Example Usage

```go
package main

import(
	"context"
	"log"
	pokeapi "PokeAPI/v2"
	"PokeAPI/v2/pkg/models/operations"
)

func main() {
    s := pokeapi.New()

    ctx := context.Background()
    res, err := s.Gender.GenderRead(ctx, operations.GenderReadRequest{
        ID: 403245,
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
| `request`                                                                    | [operations.GenderReadRequest](../../models/operations/genderreadrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GenderReadResponse](../../models/operations/genderreadresponse.md), error**

