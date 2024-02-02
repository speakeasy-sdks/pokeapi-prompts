<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	pokeapi "PokeAPI/v3"
	"PokeAPI/v3/pkg/models/operations"
	"context"
	"log"
	"net/http"
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
<!-- End SDK Example Usage [usage] -->