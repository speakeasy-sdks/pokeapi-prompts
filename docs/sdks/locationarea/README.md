# LocationArea
(*LocationArea*)

### Available Operations

* [LocationAreaList](#locationarealist)
* [LocationAreaRead](#locationarearead)

## LocationAreaList

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
    res, err := s.LocationArea.LocationAreaList(ctx, operations.LocationAreaListRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.LocationAreaListRequest](../../models/operations/locationarealistrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.LocationAreaListResponse](../../models/operations/locationarealistresponse.md), error**


## LocationAreaRead

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
    res, err := s.LocationArea.LocationAreaRead(ctx, operations.LocationAreaReadRequest{
        ID: 689350,
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.LocationAreaReadRequest](../../models/operations/locationareareadrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.LocationAreaReadResponse](../../models/operations/locationareareadresponse.md), error**

