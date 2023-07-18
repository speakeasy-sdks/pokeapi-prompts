# PokeAPI

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/pokeapi-prompts
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->


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
    res, err := s.Ability.AbilityList(ctx, operations.AbilityListRequest{
        Limit: pokeapi.Int64(548814),
        Offset: pokeapi.Int64(592845),
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

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Ability](docs/sdks/ability/README.md)

* [AbilityList](docs/sdks/ability/README.md#abilitylist)
* [AbilityRead](docs/sdks/ability/README.md#abilityread)

### [Berry](docs/sdks/berry/README.md)

* [BerryList](docs/sdks/berry/README.md#berrylist)
* [BerryRead](docs/sdks/berry/README.md#berryread)

### [BerryFirmness](docs/sdks/berryfirmness/README.md)

* [BerryFirmnessList](docs/sdks/berryfirmness/README.md#berryfirmnesslist)
* [BerryFirmnessRead](docs/sdks/berryfirmness/README.md#berryfirmnessread)

### [BerryFlavor](docs/sdks/berryflavor/README.md)

* [BerryFlavorList](docs/sdks/berryflavor/README.md#berryflavorlist)
* [BerryFlavorRead](docs/sdks/berryflavor/README.md#berryflavorread)

### [Characteristic](docs/sdks/characteristic/README.md)

* [CharacteristicList](docs/sdks/characteristic/README.md#characteristiclist)
* [CharacteristicRead](docs/sdks/characteristic/README.md#characteristicread)

### [ContestEffect](docs/sdks/contesteffect/README.md)

* [ContestEffectList](docs/sdks/contesteffect/README.md#contesteffectlist)
* [ContestEffectRead](docs/sdks/contesteffect/README.md#contesteffectread)

### [ContestType](docs/sdks/contesttype/README.md)

* [ContestTypeList](docs/sdks/contesttype/README.md#contesttypelist)
* [ContestTypeRead](docs/sdks/contesttype/README.md#contesttyperead)

### [EggGroup](docs/sdks/egggroup/README.md)

* [EggGroupList](docs/sdks/egggroup/README.md#egggrouplist)
* [EggGroupRead](docs/sdks/egggroup/README.md#egggroupread)

### [EncounterCondition](docs/sdks/encountercondition/README.md)

* [EncounterConditionList](docs/sdks/encountercondition/README.md#encounterconditionlist)
* [EncounterConditionRead](docs/sdks/encountercondition/README.md#encounterconditionread)

### [EncounterConditionValue](docs/sdks/encounterconditionvalue/README.md)

* [EncounterConditionValueList](docs/sdks/encounterconditionvalue/README.md#encounterconditionvaluelist)
* [EncounterConditionValueRead](docs/sdks/encounterconditionvalue/README.md#encounterconditionvalueread)

### [EncounterMethod](docs/sdks/encountermethod/README.md)

* [EncounterMethodList](docs/sdks/encountermethod/README.md#encountermethodlist)
* [EncounterMethodRead](docs/sdks/encountermethod/README.md#encountermethodread)

### [EvolutionChain](docs/sdks/evolutionchain/README.md)

* [EvolutionChainList](docs/sdks/evolutionchain/README.md#evolutionchainlist)
* [EvolutionChainRead](docs/sdks/evolutionchain/README.md#evolutionchainread)

### [EvolutionTrigger](docs/sdks/evolutiontrigger/README.md)

* [EvolutionTriggerList](docs/sdks/evolutiontrigger/README.md#evolutiontriggerlist)
* [EvolutionTriggerRead](docs/sdks/evolutiontrigger/README.md#evolutiontriggerread)

### [Gender](docs/sdks/gender/README.md)

* [GenderList](docs/sdks/gender/README.md#genderlist)
* [GenderRead](docs/sdks/gender/README.md#genderread)

### [Generation](docs/sdks/generation/README.md)

* [GenerationList](docs/sdks/generation/README.md#generationlist)
* [GenerationRead](docs/sdks/generation/README.md#generationread)

### [GrowthRate](docs/sdks/growthrate/README.md)

* [GrowthRateList](docs/sdks/growthrate/README.md#growthratelist)
* [GrowthRateRead](docs/sdks/growthrate/README.md#growthrateread)

### [Item](docs/sdks/item/README.md)

* [ItemList](docs/sdks/item/README.md#itemlist)
* [ItemRead](docs/sdks/item/README.md#itemread)

### [ItemAttribute](docs/sdks/itemattribute/README.md)

* [ItemAttributeList](docs/sdks/itemattribute/README.md#itemattributelist)
* [ItemAttributeRead](docs/sdks/itemattribute/README.md#itemattributeread)

### [ItemCategory](docs/sdks/itemcategory/README.md)

* [ItemCategoryList](docs/sdks/itemcategory/README.md#itemcategorylist)
* [ItemCategoryRead](docs/sdks/itemcategory/README.md#itemcategoryread)

### [ItemFlingEffect](docs/sdks/itemflingeffect/README.md)

* [ItemFlingEffectList](docs/sdks/itemflingeffect/README.md#itemflingeffectlist)
* [ItemFlingEffectRead](docs/sdks/itemflingeffect/README.md#itemflingeffectread)

### [ItemPocket](docs/sdks/itempocket/README.md)

* [ItemPocketList](docs/sdks/itempocket/README.md#itempocketlist)
* [ItemPocketRead](docs/sdks/itempocket/README.md#itempocketread)

### [Language](docs/sdks/language/README.md)

* [LanguageList](docs/sdks/language/README.md#languagelist)
* [LanguageRead](docs/sdks/language/README.md#languageread)

### [Location](docs/sdks/location/README.md)

* [LocationList](docs/sdks/location/README.md#locationlist)
* [LocationRead](docs/sdks/location/README.md#locationread)

### [LocationArea](docs/sdks/locationarea/README.md)

* [LocationAreaList](docs/sdks/locationarea/README.md#locationarealist)
* [LocationAreaRead](docs/sdks/locationarea/README.md#locationarearead)

### [Machine](docs/sdks/machine/README.md)

* [MachineList](docs/sdks/machine/README.md#machinelist)
* [MachineRead](docs/sdks/machine/README.md#machineread)

### [Move](docs/sdks/move/README.md)

* [MoveList](docs/sdks/move/README.md#movelist)
* [MoveRead](docs/sdks/move/README.md#moveread)

### [MoveAilment](docs/sdks/moveailment/README.md)

* [MoveAilmentList](docs/sdks/moveailment/README.md#moveailmentlist)
* [MoveAilmentRead](docs/sdks/moveailment/README.md#moveailmentread)

### [MoveBattleStyle](docs/sdks/movebattlestyle/README.md)

* [MoveBattleStyleList](docs/sdks/movebattlestyle/README.md#movebattlestylelist)
* [MoveBattleStyleRead](docs/sdks/movebattlestyle/README.md#movebattlestyleread)

### [MoveCategory](docs/sdks/movecategory/README.md)

* [MoveCategoryList](docs/sdks/movecategory/README.md#movecategorylist)
* [MoveCategoryRead](docs/sdks/movecategory/README.md#movecategoryread)

### [MoveDamageClass](docs/sdks/movedamageclass/README.md)

* [MoveDamageClassList](docs/sdks/movedamageclass/README.md#movedamageclasslist)
* [MoveDamageClassRead](docs/sdks/movedamageclass/README.md#movedamageclassread)

### [MoveLearnMethod](docs/sdks/movelearnmethod/README.md)

* [MoveLearnMethodList](docs/sdks/movelearnmethod/README.md#movelearnmethodlist)
* [MoveLearnMethodRead](docs/sdks/movelearnmethod/README.md#movelearnmethodread)

### [MoveTarget](docs/sdks/movetarget/README.md)

* [MoveTargetList](docs/sdks/movetarget/README.md#movetargetlist)
* [MoveTargetRead](docs/sdks/movetarget/README.md#movetargetread)

### [Nature](docs/sdks/nature/README.md)

* [NatureList](docs/sdks/nature/README.md#naturelist)
* [NatureRead](docs/sdks/nature/README.md#natureread)

### [PalParkArea](docs/sdks/palparkarea/README.md)

* [PalParkAreaList](docs/sdks/palparkarea/README.md#palparkarealist)
* [PalParkAreaRead](docs/sdks/palparkarea/README.md#palparkarearead)

### [PokeathlonStat](docs/sdks/pokeathlonstat/README.md)

* [PokeathlonStatList](docs/sdks/pokeathlonstat/README.md#pokeathlonstatlist)
* [PokeathlonStatRead](docs/sdks/pokeathlonstat/README.md#pokeathlonstatread)

### [Pokedex](docs/sdks/pokedex/README.md)

* [PokedexList](docs/sdks/pokedex/README.md#pokedexlist)
* [PokedexRead](docs/sdks/pokedex/README.md#pokedexread)

### [Pokemon](docs/sdks/pokemon/README.md)

* [PokemonList](docs/sdks/pokemon/README.md#pokemonlist)
* [PokemonRead](docs/sdks/pokemon/README.md#pokemonread)

### [PokemonColor](docs/sdks/pokemoncolor/README.md)

* [PokemonColorList](docs/sdks/pokemoncolor/README.md#pokemoncolorlist)
* [PokemonColorRead](docs/sdks/pokemoncolor/README.md#pokemoncolorread)

### [PokemonForm](docs/sdks/pokemonform/README.md)

* [PokemonFormList](docs/sdks/pokemonform/README.md#pokemonformlist)
* [PokemonFormRead](docs/sdks/pokemonform/README.md#pokemonformread)

### [PokemonHabitat](docs/sdks/pokemonhabitat/README.md)

* [PokemonHabitatList](docs/sdks/pokemonhabitat/README.md#pokemonhabitatlist)
* [PokemonHabitatRead](docs/sdks/pokemonhabitat/README.md#pokemonhabitatread)

### [PokemonShape](docs/sdks/pokemonshape/README.md)

* [PokemonShapeList](docs/sdks/pokemonshape/README.md#pokemonshapelist)
* [PokemonShapeRead](docs/sdks/pokemonshape/README.md#pokemonshaperead)

### [PokemonSpecies](docs/sdks/pokemonspecies/README.md)

* [PokemonSpeciesList](docs/sdks/pokemonspecies/README.md#pokemonspecieslist)
* [PokemonSpeciesRead](docs/sdks/pokemonspecies/README.md#pokemonspeciesread)

### [Region](docs/sdks/region/README.md)

* [RegionList](docs/sdks/region/README.md#regionlist)
* [RegionRead](docs/sdks/region/README.md#regionread)

### [Stat](docs/sdks/stat/README.md)

* [StatList](docs/sdks/stat/README.md#statlist)
* [StatRead](docs/sdks/stat/README.md#statread)

### [SuperContestEffect](docs/sdks/supercontesteffect/README.md)

* [SuperContestEffectList](docs/sdks/supercontesteffect/README.md#supercontesteffectlist)
* [SuperContestEffectRead](docs/sdks/supercontesteffect/README.md#supercontesteffectread)

### [Type](docs/sdks/type/README.md)

* [TypeList](docs/sdks/type/README.md#typelist)
* [TypeRead](docs/sdks/type/README.md#typeread)

### [Version](docs/sdks/version/README.md)

* [VersionList](docs/sdks/version/README.md#versionlist)
* [VersionRead](docs/sdks/version/README.md#versionread)

### [VersionGroup](docs/sdks/versiongroup/README.md)

* [VersionGroupList](docs/sdks/versiongroup/README.md#versiongrouplist)
* [VersionGroupRead](docs/sdks/versiongroup/README.md#versiongroupread)
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
