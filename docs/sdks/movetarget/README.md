# MoveTarget

### Available Operations

* [MoveTargetList](#movetargetlist)
* [MoveTargetRead](#movetargetread)

## MoveTargetList

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
    res, err := s.MoveTarget.MoveTargetList(ctx, operations.MoveTargetListRequest{
        Limit: PokeAPI.Int64(750686),
        Offset: PokeAPI.Int64(315428),
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.MoveTargetListRequest](../../models/operations/movetargetlistrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.MoveTargetListResponse](../../models/operations/movetargetlistresponse.md), error**


## MoveTargetRead

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
    res, err := s.MoveTarget.MoveTargetRead(ctx, operations.MoveTargetReadRequest{
        ID: 607831,
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.MoveTargetReadRequest](../../models/operations/movetargetreadrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.MoveTargetReadResponse](../../models/operations/movetargetreadresponse.md), error**

