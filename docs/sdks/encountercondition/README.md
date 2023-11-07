# EncounterCondition
(*.EncounterCondition*)

### Available Operations

* [EncounterConditionList](#encounterconditionlist)
* [EncounterConditionRead](#encounterconditionread)

## EncounterConditionList

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
    res, err := s.EncounterCondition.EncounterConditionList(ctx, operations.EncounterConditionListRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.EncounterConditionListRequest](../../models/operations/encounterconditionlistrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.EncounterConditionListResponse](../../models/operations/encounterconditionlistresponse.md), error**


## EncounterConditionRead

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
    res, err := s.EncounterCondition.EncounterConditionRead(ctx, operations.EncounterConditionReadRequest{
        ID: 628973,
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.EncounterConditionReadRequest](../../models/operations/encounterconditionreadrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.EncounterConditionReadResponse](../../models/operations/encounterconditionreadresponse.md), error**

