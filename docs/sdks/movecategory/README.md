# MoveCategory

### Available Operations

* [MoveCategoryList](#movecategorylist)
* [MoveCategoryRead](#movecategoryread)

## MoveCategoryList

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
    res, err := s.MoveCategory.MoveCategoryList(ctx, operations.MoveCategoryListRequest{
        Limit: PokeAPI.Int64(60225),
        Offset: PokeAPI.Int64(969810),
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
| `request`                                                                                | [operations.MoveCategoryListRequest](../../models/operations/movecategorylistrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.MoveCategoryListResponse](../../models/operations/movecategorylistresponse.md), error**


## MoveCategoryRead

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
    res, err := s.MoveCategory.MoveCategoryRead(ctx, operations.MoveCategoryReadRequest{
        ID: 666767,
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
| `request`                                                                                | [operations.MoveCategoryReadRequest](../../models/operations/movecategoryreadrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.MoveCategoryReadResponse](../../models/operations/movecategoryreadresponse.md), error**

