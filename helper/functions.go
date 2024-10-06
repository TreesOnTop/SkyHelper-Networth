package helper

import (
	"bytes"
	"encoding/base64"
	"strings"
	"unicode"

	"github.com/Tnze/go-mc/nbt"
	"github.com/PrismarineJS/minecraft-data/data/pc/1.16/blocks"
	"github.com/PrismarineJS/minecraft-data/data/pc/1.16/items"
	"github.com/PrismarineJS/minecraft-data/data/pc/1.16/recipes"
)

func TitleCase(str string) string {
	words := strings.Fields(strings.ToLower(strings.ReplaceAll(str, "_", " ")))
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	return strings.Join(words, " ")
}

func DecodeData(data string) ([]interface{}, error) {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	var parsed map[string]interface{}
	if err := nbt.Unmarshal(bytes.NewReader(decoded), &parsed); err != nil {
		return nil, err
	}

	return parsed["i"].([]interface{}), nil
}
