package calculators

import (
	"strings"
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
		Price:      itemPrice * item.Count,
		Calculation: []interface{}{},
		Count:      item.Count,
		Soulbound:  false,
	}
}
