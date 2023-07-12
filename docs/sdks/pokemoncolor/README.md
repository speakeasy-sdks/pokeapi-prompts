# PokemonColor

### Available Operations

* [PokemonColorList](#pokemoncolorlist)
* [PokemonColorRead](#pokemoncolorread)

## PokemonColorList

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
    s := pokeapi.New()

    ctx := context.Background()
    res, err := s.PokemonColor.PokemonColorList(ctx, operations.PokemonColorListRequest{
        Limit: pokeapi.Int64(581850),
        Offset: pokeapi.Int64(253291),
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
| `request`                                                                                | [operations.PokemonColorListRequest](../../models/operations/pokemoncolorlistrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.PokemonColorListResponse](../../models/operations/pokemoncolorlistresponse.md), error**


## PokemonColorRead

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
    s := pokeapi.New()

    ctx := context.Background()
    res, err := s.PokemonColor.PokemonColorRead(ctx, operations.PokemonColorReadRequest{
        ID: 414369,
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
| `request`                                                                                | [operations.PokemonColorReadRequest](../../models/operations/pokemoncolorreadrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.PokemonColorReadResponse](../../models/operations/pokemoncolorreadresponse.md), error**
