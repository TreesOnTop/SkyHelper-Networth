package calculators

import (
	"strings"
	"constants"
	"github.com/Tnze/go-mc/nbt"
	"github.com/PrismarineJS/minecraft-data/data/pc/1.16/blocks"
	"github.com/PrismarineJS/minecraft-data/data/pc/1.16/items"
	"github.com/PrismarineJS/minecraft-data/data/pc/1.16/recipes"
)

type Essence struct {
	Name       string
	ID         string
	Price      int
	Calculation []interface{}
	Count      int
	Soulbound  bool
}

func titleCase(str string) string {
	words := strings.Split(strings.ToLower(str), " ")
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	return strings.Join(words, " ")
}

func CalculateEssence(item Essence, prices map[string]int) *Essence {
	itemPrice, exists := prices[strings.ToLower(item.ID)]
	if !exists {
		return nil
	}

	return &Essence{
		Name:       titleCase(strings.Split(item.ID, "_")[1]) + " Essence",
		ID:         item.ID,
		Price:      itemPrice * item.Count * constants.ApplicationWorth.Essence,
		Calculation: []interface{}{},
		Count:      item.Count,
		Soulbound:  false,
	}
}
