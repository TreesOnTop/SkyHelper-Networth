package helper

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/PrismarineJS/minecraft-data/data/pc/1.16/blocks"
	"github.com/Tnze/go-mc/nbt"
	"github.com/PrismarineJS/minecraft-data/data/pc/1.16/items"
	"github.com/PrismarineJS/minecraft-data/data/pc/1.16/recipes"
	"constants"
)

type Item struct {
	ID             string
	Amount         int
	Tag            *ItemTag
	ExtraAttributes *ExtraAttributes
}

type ItemTag struct {
	ExtraAttributes *ExtraAttributes
	Display         *Display
}

type ExtraAttributes struct {
	NewYearCakeBagData string
	NewYearCakeBagYears []int
}

type Display struct {
	Name string
	Lore []string
}

type ProfileData struct {
	Inventory *Inventory
	SharedInventory *SharedInventory
	SacksCounts map[string]int
	Currencies *Currencies
	Pets []*Pet
}

type Inventory struct {
	BagContents map[string]*BagContent
	BackpackContents map[string]*BackpackContent
	BackpackIcons map[string]*BackpackIcon
}

type SharedInventory struct {
	CandyInventoryContents *BagContent
	CarnivalMaskInventoryContents *BagContent
}

type BagContent struct {
	Data string
}

type BackpackContent struct {
	Data string
}

type BackpackIcon struct {
	Data string
}

type Currencies struct {
	Essence map[string]*Essence
}

type Essence struct {
	Current int
}

type Pet struct {
	Type string
	Level int
	XPMax int
	Exp int
	HeldItem string
	Skin string
	CandyUsed int
}

var singleContainers = constants.SingleContainers
var bagContainers = constants.BagContainers
var sharedContainers = constants.SharedContainers

func parseItems(profileData *ProfileData, museumData *MuseumData, v2Endpoint bool) (map[string][]*Item, error) {
	items := make(map[string][]*Item)

	for container, key := range singleContainers {
		items[container] = []*Item{}
		if v2Endpoint {
			var containerData *BagContent
			if contains(bagContainers, key) {
				containerData = profileData.Inventory.BagContents[key]
			} else if contains(sharedContainers, key) {
				containerData = profileData.SharedInventory.CandyInventoryContents
			} else {
				containerData = profileData.Inventory.BagContents[key]
			}
			if containerData != nil {
				decodedData, err := decodeData(containerData.Data)
				if err != nil {
					return nil, err
				}
				items[container] = decodedData
			}
		} else {
			if profileData.Inventory.BagContents[key] != nil {
				decodedData, err := decodeData(profileData.Inventory.BagContents[key].Data)
				if err != nil {
					return nil, err
				}
				items[container] = decodedData
			}
		}
	}

	items["storage"] = []*Item{}
	if profileData.Inventory.BackpackContents != nil && profileData.Inventory.BackpackIcons != nil {
		for _, backpackContent := range profileData.Inventory.BackpackContents {
			decodedData, err := decodeData(backpackContent.Data)
			if err != nil {
				return nil, err
			}
			items["storage"] = append(items["storage"], decodedData...)
		}

		for _, backpackIcon := range profileData.Inventory.BackpackIcons {
			decodedData, err := decodeData(backpackIcon.Data)
			if err != nil {
				return nil, err
			}
			items["storage"] = append(items["storage"], decodedData...)
		}
	}

	items["museum"] = []*Item{}
	if museumData != nil && museumData.Items != nil {
		for _, data := range museumData.Items {
			if data.Items.Data == "" || data.Borrowing {
				continue
			}
			decodedItem, err := decodeData(data.Items.Data)
			if err != nil {
				return nil, err
			}
			items["museum"] = append(items["museum"], decodedItem...)
		}

		for _, data := range museumData.Special {
			if data.Items.Data == "" {
				continue
			}
			decodedItem, err := decodeData(data.Items.Data)
			if err != nil {
				return nil, err
			}
			items["museum"] = append(items["museum"], decodedItem...)
		}
	}

	err := postParseItems(profileData, items, v2Endpoint)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func postParseItems(profileData *ProfileData, items map[string][]*Item, v2Endpoint bool) error {
	for _, categoryItems := range items {
		for _, item := range categoryItems {
			if item.Tag == nil || item.Tag.ExtraAttributes == nil || item.Tag.ExtraAttributes.NewYearCakeBagData == "" {
				continue
			}
			cakes, err := decodeData(item.Tag.ExtraAttributes.NewYearCakeBagData)
			if err != nil {
				return err
			}
			item.Tag.ExtraAttributes.NewYearCakeBagYears = []int{}
			for _, cake := range cakes {
				if cake.ID != "" && cake.Tag != nil && cake.Tag.ExtraAttributes != nil && cake.Tag.ExtraAttributes.NewYearCakeBagYears != nil {
					item.Tag.ExtraAttributes.NewYearCakeBagYears = append(item.Tag.ExtraAttributes.NewYearCakeBagYears, cake.Tag.ExtraAttributes.NewYearCakeBagYears...)
				}
			}
		}
	}

	items["sacks"] = []*Item{}
	if profileData.SacksCounts != nil {
		for id, amount := range profileData.SacksCounts {
			if amount > 0 {
				items["sacks"] = append(items["sacks"], &Item{ID: id, Amount: amount})
			}
		}
	}

	items["essence"] = []*Item{}
	if v2Endpoint {
		if profileData.Currencies != nil && profileData.Currencies.Essence != nil {
			for id, essence := range profileData.Currencies.Essence {
				items["essence"] = append(items["essence"], &Item{ID: "essence_" + id, Amount: essence.Current})
			}
		}
	} else {
		for id, amount := range profileData.SacksCounts {
			if strings.HasPrefix(id, "essence_") {
				items["essence"] = append(items["essence"], &Item{ID: id, Amount: amount})
			}
		}
	}

	items["pets"] = []*Pet{}
	if profileData.Pets != nil {
		for _, pet := range profileData.Pets {
			newPet := *pet
			level := constants.GetPetLevel(&newPet)
			newPet.Level = level.Level
			newPet.XPMax = level.XPMax
			items["pets"] = append(items["pets"], &newPet)
		}
	}

	return nil
}

func decodeData(data string) ([]*Item, error) {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	var items []*Item
	err = json.Unmarshal(decoded, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
