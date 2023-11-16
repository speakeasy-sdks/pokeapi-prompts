# EvolutionTrigger
(*EvolutionTrigger*)

### Available Operations

* [EvolutionTriggerList](#evolutiontriggerlist)
* [EvolutionTriggerRead](#evolutiontriggerread)

## EvolutionTriggerList

### Example Usage

```go
package main

import(
	pokeapi "PokeAPI/v2"
	"context"
	"PokeAPI/v2/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := pokeapi.New()

    ctx := context.Background()
    res, err := s.EvolutionTrigger.EvolutionTriggerList(ctx, operations.EvolutionTriggerListRequest{})
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
| `request`                                                                                            | [operations.EvolutionTriggerListRequest](../../pkg/models/operations/evolutiontriggerlistrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.EvolutionTriggerListResponse](../../pkg/models/operations/evolutiontriggerlistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## EvolutionTriggerRead

### Example Usage

```go
package main

import(
	pokeapi "PokeAPI/v2"
	"context"
	"PokeAPI/v2/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := pokeapi.New()

    ctx := context.Background()
    res, err := s.EvolutionTrigger.EvolutionTriggerRead(ctx, operations.EvolutionTriggerReadRequest{
        ID: 848636,
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
| `request`                                                                                            | [operations.EvolutionTriggerReadRequest](../../pkg/models/operations/evolutiontriggerreadrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.EvolutionTriggerReadResponse](../../pkg/models/operations/evolutiontriggerreadresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
