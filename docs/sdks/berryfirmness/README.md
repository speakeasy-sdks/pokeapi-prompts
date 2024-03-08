# BerryFirmness
(*BerryFirmness*)

### Available Operations

* [BerryFirmnessList](#berryfirmnesslist)
* [BerryFirmnessRead](#berryfirmnessread)

## BerryFirmnessList

### Example Usage

```go
package main

import(
	pokeapi "PokeAPI/v3"
	"context"
	"PokeAPI/v3/pkg/models/operations"
	"log"
)

func main() {
    s := pokeapi.New()

    ctx := context.Background()
    res, err := s.BerryFirmness.BerryFirmnessList(ctx, operations.BerryFirmnessListRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.BerryFirmnessListRequest](../../pkg/models/operations/berryfirmnesslistrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.BerryFirmnessListResponse](../../pkg/models/operations/berryfirmnesslistresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## BerryFirmnessRead

### Example Usage

```go
package main

import(
	pokeapi "PokeAPI/v3"
	"context"
	"PokeAPI/v3/pkg/models/operations"
	"log"
)

func main() {
    s := pokeapi.New()

    ctx := context.Background()
    res, err := s.BerryFirmness.BerryFirmnessRead(ctx, operations.BerryFirmnessReadRequest{
        ID: 711277,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.BerryFirmnessReadRequest](../../pkg/models/operations/berryfirmnessreadrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.BerryFirmnessReadResponse](../../pkg/models/operations/berryfirmnessreadresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
