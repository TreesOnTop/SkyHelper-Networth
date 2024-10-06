package calculators

import (
	"strings"
	"unicode"

	"../constants"
	"../helper"
	"github.com/Tnze/go-mc/nbt"
	"github.com/PrismarineJS/minecraft-data/data/pc/1.16/blocks"
	"github.com/PrismarineJS/minecraft-data/data/pc/1.16/items"
	"github.com/PrismarineJS/minecraft-data/data/pc/1.16/recipes"
)

type Pet struct {
	Name       string
	ID         string
	Level      int
	Exp        float64
	XpMax      float64
	Tier       string
	Type       string
	HeldItem   string
	Skin       string
	CandyUsed  int
	Price      float64
	Base       float64
	Calculation []CalculationData
	Soulbound  bool
}

type CalculationData struct {
	ID    string
	Type  string
	Price float64
	Count int
}

func titleCase(str string) string {
	words := strings.Fields(strings.ToLower(str))
	for i, word := range words {
		words[i] = strings.ToUpper(string(word[0])) + word[1:]
	}
	return strings.Join(words, " ")
}

func getPetLevelPrices(pet Pet, prices map[string]float64) (float64, float64, float64, string) {
	tier := pet.Tier
	if pet.HeldItem == "PET_ITEM_TIER_BOOST" {
		tierIndex := indexOf(tier, constants.Tiers)
		if tierIndex > 0 {
			tier = constants.Tiers[tierIndex-1]
		}
	}
	skin := strings.ToLower(pet.Skin)
	tierName := strings.ToLower(tier + "_" + pet.Type)
	basePrices := struct {
		ID     string
		Lvl1   float64
		Lvl100 float64
		Lvl200 float64
	}{
		ID:     tierName + skinSuffix(skin),
		Lvl1:   prices["lvl_1_"+tierName],
		Lvl100: prices["lvl_100_"+tierName],
		Lvl200: prices["lvl_200_"+tierName],
	}

	if skin != "" {
		return max(basePrices.Lvl1, prices["lvl_1_"+tierName+skinSuffix(skin)]),
			max(basePrices.Lvl100, prices["lvl_100_"+tierName+skinSuffix(skin)]),
			max(basePrices.Lvl200, prices["lvl_200_"+tierName+skinSuffix(skin)]),
			basePrices.ID
	}
	return basePrices.Lvl1, basePrices.Lvl100, basePrices.Lvl200, basePrices.ID
}

func calculatePet(pet Pet, prices map[string]float64) Pet {
	lvl1, lvl100, lvl200, id := getPetLevelPrices(pet, prices)
	pet.Name = "[Lvl " + strconv.Itoa(pet.Level) + "] " + titleCase(pet.Tier+" "+pet.Type) + petSkinSuffix(pet.Skin)
	if lvl1 == 0 || lvl100 == 0 {
		return Pet{}
	}

	price := lvl200
	if price == 0 {
		price = lvl100
	}
	calculation := []CalculationData{}

	if pet.Level < 100 && pet.XpMax > 0 {
		baseFormula := (lvl100 - lvl1) / pet.XpMax
		if baseFormula > 0 {
			price = baseFormula*pet.Exp + lvl1
		}
	}

	if pet.Level > 100 && pet.Level < 200 {
		level := pet.Level % 100
		if level != 1 {
			baseFormula := (lvl200 - lvl100) / 100
			if baseFormula > 0 {
				price = baseFormula*float64(level) + lvl100
			}
		}
	}

	base := price
	soulbound := contains(constants.SoulboundPets, pet.Type)

	if pet.Skin != "" && soulbound {
		calculationData := CalculationData{
			ID:    pet.Skin,
			Type:  "soulbound_pet_skin",
			Price: prices["pet_skin_"+strings.ToLower(pet.Skin)] * constants.ApplicationWorth.SoulboundPetSkins,
			Count: 1,
		}
		price += calculationData.Price
		calculation = append(calculation, calculationData)
	}

	if pet.HeldItem != "" {
		calculationData := CalculationData{
			ID:    pet.HeldItem,
			Type:  "pet_item",
			Price: prices[strings.ToLower(pet.HeldItem)] * constants.ApplicationWorth.PetItem,
			Count: 1,
		}
		price += calculationData.Price
		calculation = append(calculation, calculationData)
	}

	maxPetCandyXp := float64(pet.CandyUsed) * 1000000
	xpLessPetCandy := pet.Exp - maxPetCandyXp
	if pet.CandyUsed > 0 && !contains(constants.BlockedCandyReducePets, pet.Type) && xpLessPetCandy < pet.XpMax {
		reducedValue := price * constants.ApplicationWorth.PetCandy
		if !math.IsNaN(price) {
			if pet.Level == 100 {
				price = math.Max(reducedValue, price-5000000)
			} else {
				price = math.Max(reducedValue, price-2500000)
			}
		}
	}

	pet.ID = id
	pet.Price = price
	pet.Base = base
	pet.Calculation = calculation
	pet.Soulbound = soulbound

	return pet
}

func skinSuffix(skin string) string {
	if skin != "" {
		return "_skinned_" + skin
	}
	return ""
}

func petSkinSuffix(skin string) string {
	if skin != "" {
		return " ✦"
	}
	return ""
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
