<!-- Start SDK Example Usage -->


```go
package main

import (
	pokeapi "PokeAPI"
	"PokeAPI/pkg/models/operations"
	"context"
	"log"
)

func main() {
	s := pokeapi.New()

	ctx := context.Background()
	res, err := s.Ability.AbilityList(ctx, operations.AbilityListRequest{})
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusOK {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->