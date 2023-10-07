# EvolutionChain
(*EvolutionChain*)

### Available Operations

* [EvolutionChainList](#evolutionchainlist)
* [EvolutionChainRead](#evolutionchainread)

## EvolutionChainList

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
    res, err := s.EvolutionChain.EvolutionChainList(ctx, operations.EvolutionChainListRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.EvolutionChainListRequest](../../models/operations/evolutionchainlistrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.EvolutionChainListResponse](../../models/operations/evolutionchainlistresponse.md), error**


## EvolutionChainRead

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
    res, err := s.EvolutionChain.EvolutionChainRead(ctx, operations.EvolutionChainReadRequest{
        ID: 663834,
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.EvolutionChainReadRequest](../../models/operations/evolutionchainreadrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.EvolutionChainReadResponse](../../models/operations/evolutionchainreadresponse.md), error**

