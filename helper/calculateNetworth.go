package helper

import (
	"calculators"
	"constants"
	"github.com/Tnze/go-mc/nbt"
	"github.com/PrismarineJS/minecraft-data/data/pc/1.16/blocks"
	"github.com/PrismarineJS/minecraft-data/data/pc/1.16/items"
	"github.com/PrismarineJS/minecraft-data/data/pc/1.16/recipes"
)

type Item struct {
	Name         string
	LoreName     string
	ID           string
	Price        int
	Base         int
	Calculation  []CalculationData
	Count        int
	Soulbound    bool
	IsPet        bool
	ExtraAttributes ExtraAttributes
}

type CalculationData struct {
	ID    string
	Type  string
	Price int
	Count int
	Star  int
}

type ExtraAttributes struct {
	PetInfo string
	Exp     int
}

type Category struct {
	Total            int
	UnsoulboundTotal int
	Items            []Item
}

type NetworthResult struct {
	NoInventory        bool
	Networth           int
	UnsoulboundNetworth int
	Purse              int
	Bank               int
	PersonalBank       int
	Types              map[string]Category
}

func CalculateNetworth(items map[string][]Item, purseBalance, bankBalance, personalBankBalance int, prices map[string]int, onlyNetworth, returnItemData bool) NetworthResult {
	categories := make(map[string]Category)

	for category, categoryItems := range items {
		categories[category] = Category{Total: 0, UnsoulboundTotal: 0, Items: []Item{}}

		for _, item := range categoryItems {
			var result Item
			switch category {
			case "pets":
				result = calculators.CalculatePet(item, prices)
			case "sacks":
				result = calculators.CalculateSackItem(item, prices)
			case "essence":
				result = calculators.CalculateEssence(item, prices)
			default:
				result = calculators.CalculateItem(item, prices, returnItemData)
			}

			categories[category].Total += result.Price
			if !result.Soulbound {
				categories[category].UnsoulboundTotal += result.Price
			}
			if !onlyNetworth && result.Price > 0 {
				categories[category].Items = append(categories[category].Items, result)
			}
		}

		if !onlyNetworth && len(categories[category].Items) > 0 {
			categories[category].Items = mergeItems(categories[category].Items)
		}

		if onlyNetworth {
			categories[category].Items = nil
		}
	}

	total := calculateTotalNetworth(categories, bankBalance, purseBalance, personalBankBalance)
	unsoulboundTotal := calculateUnsoulboundNetworth(categories, bankBalance, purseBalance, personalBankBalance)

	return NetworthResult{
		NoInventory:        len(items["inventory"]) == 0,
		Networth:           total,
		UnsoulboundNetworth: unsoulboundTotal,
		Purse:              purseBalance,
		Bank:               bankBalance,
		PersonalBank:       personalBankBalance,
		Types:              categories,
	}
}

func mergeItems(items []Item) []Item {
	mergedItems := []Item{}
	for _, item := range items {
		if len(mergedItems) > 0 {
			last := &mergedItems[len(mergedItems)-1]
			if last.Name == item.Name && last.Price/last.Count == item.Price/item.Count && !item.IsPet && last.Soulbound == item.Soulbound {
				last.Price += item.Price
				last.Count += item.Count
				last.Base = last.Base
				last.Calculation = last.Calculation
			} else {
				mergedItems = append(mergedItems, item)
			}
		} else {
			mergedItems = append(mergedItems, item)
		}
	}
	return mergedItems
}

func calculateTotalNetworth(categories map[string]Category, bankBalance, purseBalance, personalBankBalance int) int {
	total := 0
	for _, category := range categories {
		total += category.Total
	}
	return total + bankBalance + purseBalance + personalBankBalance
}

func calculateUnsoulboundNetworth(categories map[string]Category, bankBalance, purseBalance, personalBankBalance int) int {
	unsoulboundTotal := 0
	for _, category := range categories {
		unsoulboundTotal += category.UnsoulboundTotal
	}
	return unsoulboundTotal + bankBalance + purseBalance + personalBankBalance
}

func CalculateItemNetworth(item Item, prices map[string]int, returnItemNetworth bool) Item {
	isPet := item.ExtraAttributes.PetInfo != "" || item.ExtraAttributes.Exp != 0
	if isPet {
		petInfo := item.ExtraAttributes.PetInfo
		if petInfo != "" {
			petInfo = item.ExtraAttributes.PetInfo
		} else {
			petInfo = item
		}
		level := constants.GetPetLevel(petInfo)
		petInfo.Level = level.Level
		petInfo.XpMax = level.XpMax
		return calculators.CalculatePet(petInfo, prices)
	}
	return calculators.CalculateItem(item, prices, returnItemNetworth)
}
