# BerryFlavor
(*BerryFlavor*)

### Available Operations

* [BerryFlavorList](#berryflavorlist)
* [BerryFlavorRead](#berryflavorread)

## BerryFlavorList

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
    res, err := s.BerryFlavor.BerryFlavorList(ctx, operations.BerryFlavorListRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.BerryFlavorListRequest](../../models/operations/berryflavorlistrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.BerryFlavorListResponse](../../models/operations/berryflavorlistresponse.md), error**


## BerryFlavorRead

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
    res, err := s.BerryFlavor.BerryFlavorRead(ctx, operations.BerryFlavorReadRequest{
        ID: 678235,
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.BerryFlavorReadRequest](../../models/operations/berryflavorreadrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.BerryFlavorReadResponse](../../models/operations/berryflavorreadresponse.md), error**

