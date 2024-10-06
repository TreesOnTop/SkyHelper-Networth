package constants

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"github.com/Tnze/go-mc/nbt"
	"github.com/PrismarineJS/minecraft-data/data/pc/1.16/blocks"
	"github.com/PrismarineJS/minecraft-data/data/pc/1.16/items"
	"github.com/PrismarineJS/minecraft-data/data/pc/1.16/recipes"
)

var itemsMap map[string]Item

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func init() {
	itemsMap = make(map[string]Item)
	data, err := ioutil.ReadFile("constants/items.json")
	if err != nil {
		log.Fatalf("Failed to read items.json: %v", err)
	}
	var items []Item
	if err := json.Unmarshal(data, &items); err != nil {
		log.Fatalf("Failed to unmarshal items.json: %v", err)
	}
	for _, item := range items {
		itemsMap[item.ID] = item
	}
}

func GetHypixelItemInformationFromId(id string) *Item {
	item, exists := itemsMap[id]
	if !exists {
		return nil
	}
	return &item
}
