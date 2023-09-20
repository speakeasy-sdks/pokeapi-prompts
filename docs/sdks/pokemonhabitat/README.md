# PokemonHabitat

### Available Operations

* [PokemonHabitatList](#pokemonhabitatlist)
* [PokemonHabitatRead](#pokemonhabitatread)

## PokemonHabitatList

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
    res, err := s.PokemonHabitat.PokemonHabitatList(ctx, operations.PokemonHabitatListRequest{
        Limit: PokeAPI.Int64(623510),
        Offset: PokeAPI.Int64(158969),
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
| `request`                                                                                    | [operations.PokemonHabitatListRequest](../../models/operations/pokemonhabitatlistrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.PokemonHabitatListResponse](../../models/operations/pokemonhabitatlistresponse.md), error**


## PokemonHabitatRead

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
    res, err := s.PokemonHabitat.PokemonHabitatRead(ctx, operations.PokemonHabitatReadRequest{
        ID: 338007,
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
| `request`                                                                                    | [operations.PokemonHabitatReadRequest](../../models/operations/pokemonhabitatreadrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.PokemonHabitatReadResponse](../../models/operations/pokemonhabitatreadresponse.md), error**

