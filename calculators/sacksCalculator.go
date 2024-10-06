package calculators

import (
	"strings"
	"unicode"

	"../constants"
	"../helper"
)

type SackItem struct {
	Name       string
	ID         string
	Price      int
	Calculation []interface{}
	Count      int
	Soulbound  bool
}

func calculateSackItem(item SackItem, prices map[string]int) *SackItem {
	itemPrice := prices[strings.ToLower(item.ID)]
	if strings.HasPrefix(item.ID, "RUNE_") && !contains(constants.ValidRunes, item.ID) {
		return nil
	}
	name := item.Name
	if name == "" {
		name = helper.TitleCase(item.ID)
	}
	if itemPrice > 0 {
		return &SackItem{
			Name:       stripColorCodes(name),
			ID:         item.ID,
			Price:      itemPrice * item.Count,
			Calculation: []interface{}{},
			Count:      item.Count,
			Soulbound:  false,
		}
	} else {
		return nil
	}
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func stripColorCodes(str string) string {
	var result strings.Builder
	for _, r := range str {
		if r == '§' {
			continue
		}
		if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r) {
			result.WriteRune(r)
		}
	}
	return result.String()
}
