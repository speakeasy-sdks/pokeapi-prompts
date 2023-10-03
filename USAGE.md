<!-- Start SDK Example Usage -->


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
    res, err := s.Ability.AbilityList(ctx, operations.AbilityListRequest{
        Limit: pokeapi.Int64(578676),
        Offset: pokeapi.Int64(855699),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->