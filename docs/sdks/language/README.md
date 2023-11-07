# Language
(*.Language*)

### Available Operations

* [LanguageList](#languagelist)
* [LanguageRead](#languageread)

## LanguageList

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
    res, err := s.Language.LanguageList(ctx, operations.LanguageListRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.LanguageListRequest](../../models/operations/languagelistrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.LanguageListResponse](../../models/operations/languagelistresponse.md), error**


## LanguageRead

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
    res, err := s.Language.LanguageRead(ctx, operations.LanguageReadRequest{
        ID: 424722,
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.LanguageReadRequest](../../models/operations/languagereadrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.LanguageReadResponse](../../models/operations/languagereadresponse.md), error**

