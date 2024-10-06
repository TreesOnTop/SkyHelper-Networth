package calculators

import (
	"strings"
	"unicode"
	"constants"
	"github.com/Tnze/go-mc/nbt"
	"github.com/PrismarineJS/minecraft-data/data/pc/1.16/blocks"
	"github.com/PrismarineJS/minecraft-data/data/pc/1.16/items"
	"github.com/PrismarineJS/minecraft-data/data/pc/1.16/recipes"
)

type Item struct {
	Tag struct {
		ExtraAttributes struct {
			ID                   string
			PetInfo              string
			Skin                 string
			PartyHatEmoji        string
			NewYearsCake         int
			Edition              int
			IsShiny              bool
			Price                int
			PickonimbusDurability int
			Attributes           map[string]int
			SackPSS              int
			WoodSingularityCount int
			JalapenoCount        int
			TunedTransmission    int
			ManaDisintegratorCount int
			ThunderCharge        int
			HotPotatoCount       int
			DyeItem              string
			ArtOfWarCount        int
					ArtOfPeaceApplied    int
			FarmingForDummiesCount int
			Polarvoid            int
			DivanPowderCoating   int
			TalismanEnrichment   string
			RarityUpgrades       int
			Gems                 struct {
				Formatted     bool
				UnlockedSlots []string
				Gems          []struct {
					Type     string
					Tier     int
					SlotType string
				}
			}
			PowerAbilityScroll string
			Modifier           string
			DungeonItemLevel   int
			UpgradeLevel       int
			AbilityScroll      []string
			DrillPartUpgradeModule string
			DrillPartFuelTank  string
			DrillPartEngine    string
			Ethermerge         int
			NewYearCakeBagYears []int
			DonatedMuseum      bool
		}
		Display struct {
			Name string
			Lore []string
		}
	}
	Count int
}

type CalculationData struct {
	ID    string
	Type  string
	Price int
	Count int
	Star  int
	Shards int
}

type Pet struct {
	Level int
	XpMax int
}

type SkyblockItem struct {
	Category       string
	GemstoneSlots  []struct {
		SlotType string
		Costs    []struct {
			Type   string
			Coins  int
			ItemID string
			Amount int
		}
	}
	UpgradeCosts []struct {
		EssenceType string
		ItemID      string
		Amount      int
	}
	Prestige struct {
		Costs []struct {
			EssenceType string
			ItemID      string
			Amount      int
		}
	}
}

var (
	ApplicationWorth = constants.ApplicationWorth
)

var (
	EnchantsWorth = constants.EnchantsWorth
)

var (
	BlockedEnchants = constants.BlockedEnchants
	IgnoredEnchants = constants.IgnoredEnchants
	StackingEnchants = constants.StackingEnchants
	IgnoreSilex = constants.IgnoreSilex
	MasterStars = constants.MasterStars
	ValidRunes = constants.ValidRunes
	AllowedRecombTypes = constants.AllowedRecombTypes
	AllowedRecombIds = constants.AllowedRecombIds
	AttributesBaseCosts = constants.AttributesBaseCosts
	Enrichments = constants.Enrichments
	PickonimbusDurability = constants.PickonimbusDurability
	SpecialEnchantmentMatches = constants.SpecialEnchantmentMatches
)

var (
	Reforges = map[string]string{
		"stiff": "hardened_wood", "salty": "salt_cube", "aote_stone": "aote_stone", "blazing": "blazen_sphere", "waxed": "blaze_wax", "rooted": "burrowing_spores",
		"candied": "candy_corn", "perfect": "diamond_atom", "fleet": "diamonite", "fabled": "dragon_claw", "spiked": "dragon_scale", "royal": "dwarven_treasure",
		"hyper": "endstone_geode", "coldfusion": "entropy_suppressor", "blooming": "flowering_bouquet", "fanged": "full_jaw_fanging_kit", "jaded": "jaderald",
		"jerry": "jerry_stone", "magnetic": "lapis_crystal", "earthy": "large_walnut", "fortified": "meteor_shard", "gilded": "midas_jewel", "cubic": "molten_cube",
		"necrotic": "necromancer_brooch", "fruitful": "onyx", "precise": "optical_lens", "mossy": "overgrown_grass", "pitchin": "pitchin_koi", "undead": "premium_flesh",
		"blood_soaked": "presumed_gallon_of_red_paint", "mithraic": "pure_mithril", "reinforced": "rare_diamond", "ridiculous": "red_nose", "loving": "red_scarf",
		"auspicious": "rock_gemstone", "treacherous": "rusty_anchor", "headstrong": "salmon_opal", "strengthened": "searing_stone", "glistening": "shiny_prism",
		"bustling": "skymart_brochure", "spiritual": "spirit_decoy", "suspicious": "suspicious_vial", "snowy": "terry_snowglobe", "dimensional": "titanium_tesseract",
		"ambered": "amber_material", "beady": "beady_eyes", "blessed": "blessed_fruit", "bulky": "bulky_stone", "buzzing": "clipped_wings", "submerged": "deep_sea_orb",
		"renowned": "dragon_horn", "festive": "frozen_bauble", "giant": "giant_tooth", "lustrous": "gleaming_crystal", "bountiful": "golden_ball", "chomp": "kuudra_mandible",
		"lucky": "lucky_dice", "stellar": "petrified_starfall", "scraped": "pocket_iceberg", "ancient": "precursor_gear", "refined": "refined_amber", "empowered": "sadan_brooch",
		"withered": "wither_blood", "glacial": "frigid_husk", "heated": "hot_stuff", "dirty": "dirt_bottle", "moil": "moil_log", "toil": "toil_log", "greater_spook": "boo_stone",
	}
)

func titleCase(str string) string {
	words := strings.Fields(strings.ToLower(strings.ReplaceAll(str, "_", " ")))
	for i, word := range words {
		words[i] = strings.ToUpper(string(word[0])) + word[1:]
	}
	return strings.Join(words, " ")
}

func starCost(prices map[string]int, upgrade struct {
	EssenceType string
	ItemID      string
	Amount      int
}, star int) *CalculationData {
	upgradePrice := prices[strings.ToLower(upgrade.EssenceType)]
	if upgradePrice == 0 {
		upgradePrice = prices[strings.ToLower(upgrade.ItemID)]
	}
	if upgradePrice == 0 {
		return nil
	}

	calculationData := &CalculationData{
		ID:    upgrade.EssenceType + "_ESSENCE",
		Type:  "star",
		Price: upgrade.Amount * upgradePrice,
		Count: upgrade.Amount,
		Star:  star,
	}
	if upgrade.EssenceType == "" {
		calculationData.ID = upgrade.ItemID
		calculationData.Type = "prestige"
		calculationData.Price = upgrade.Amount * upgradePrice
		calculationData.Count = upgrade.Amount
	}
	return calculationData
}

func starCosts(prices map[string]int, calculation *[]CalculationData, upgrades []struct {
	EssenceType string
	ItemID      string
	Amount      int
}, prestigeItem string) int {
	price := 0
	star := 0
	datas := []*CalculationData{}
	for _, upgrade := range upgrades {
		star++
		var data *CalculationData
		if len(upgrade.EssenceType) > 0 {
			data = starCost(prices, upgrade, star)
			datas = append(datas, data)
			if prestigeItem == "" && data != nil {
				price += data.Price
				*calculation = append(*calculation, *data)
			}
		} else {
			data = starCost(prices, upgrade, 0)
			datas = append(datas, data)
			if prestigeItem == "" && data != nil {
				price += data.Price
				*calculation = append(*calculation, *data)
			}
		}
	}

	if prestigeItem != "" && len(datas) > 0 && datas[0] != nil {
		prestige := datas[0].Type == "prestige"
		calculationData := &CalculationData{
			ID:    prestigeItem,
			Type:  "stars",
			Price: 0,
			Count: star,
		}
		if prestige {
			calculationData.Type = "prestige"
			calculationData.Count = 1
		}
		for _, data := range datas {
			calculationData.Price += data.Price
		}
		if prestige && prices[strings.ToLower(prestigeItem)] != 0 {
			calculationData.Price += prices[strings.ToLower(prestigeItem)]
		}
		price += calculationData.Price
		*calculation = append(*calculation, *calculationData)
	}
	return price
}

func calculateItem(item Item, prices map[string]int, returnItemData bool) *Item {
	if item.Tag.ExtraAttributes.ID == "PET" && item.Tag.ExtraAttributes.PetInfo != "" {
		var petInfo Pet
		// Unmarshal petInfo from item.Tag.ExtraAttributes.PetInfo
		// Calculate pet level and xpMax
		petInfo.Level = 1 // Replace with actual calculation
		petInfo.XpMax = 100 // Replace with actual calculation
		return calculatePet(petInfo, prices)
	}

	if item.Tag.ExtraAttributes.ID != "" {
		if item.Tag.Display.Name == "" {
			return nil
		}
		itemName := strings.ReplaceAll(item.Tag.Display.Name, "§[0-9a-fk-or]", "")
		itemId := strings.ToLower(item.Tag.ExtraAttributes.ID)
		ExtraAttributes := item.Tag.ExtraAttributes
		skyblockItem := getHypixelItemInformationFromId(strings.ToUpper(itemId))

		if ExtraAttributes.Skin != "" {
			if prices[itemId+"_skinned_"+strings.ToLower(ExtraAttributes.Skin)] != 0 {
				itemId += "_skinned_" + strings.ToLower(ExtraAttributes.Skin)
			}
		}
		if itemId == "party_hat_sloth" && ExtraAttributes.PartyHatEmoji != "" {
			if prices[itemId+"_"+strings.ToLower(ExtraAttributes.PartyHatEmoji)] != 0 {
				itemId += "_" + strings.ToLower(ExtraAttributes.PartyHatEmoji)
			}
		}

		if itemName == "Beastmaster Crest" || itemName == "Griffin Upgrade Stone" || itemName == "Wisp Upgrade Stone" {
			itemName = itemName + " (" + titleCase(skyblockItem.Tier) + ")"
		} else if strings.HasSuffix(itemName, " Exp Boost") {
			itemName = itemName + " (" + titleCase(skyblockItem.ID) + ")"
		}

		if ExtraAttributes.ID == "RUNE" || ExtraAttributes.ID == "UNIQUE_RUNE" {
			for runeType, runeTier := range ExtraAttributes.Runes {
				itemId = "rune_" + runeType + "_" + runeTier
				break
			}
		}
		if ExtraAttributes.ID == "NEW_YEAR_CAKE" {
			itemId = "new_year_cake_" + strconv.Itoa(ExtraAttributes.NewYearsCake)
		}
		if ExtraAttributes.ID == "PARTY_HAT_CRAB" || ExtraAttributes.ID == "PARTY_HAT_CRAB_ANIMATED" || ExtraAttributes.ID == "BALLOON_HAT_2024" {
			if ExtraAttributes.PartyHatColor != "" {
				itemId = strings.ToLower(ExtraAttributes.ID) + "_" + ExtraAttributes.PartyHatColor
			}
		}
		if ExtraAttributes.ID == "DCTR_SPACE_HELM" && ExtraAttributes.Edition != 0 {
			itemId = "dctr_space_helm_editioned"
		}
		if ExtraAttributes.ID == "CREATIVE_MIND" && ExtraAttributes.Edition == 0 {
			itemId = "creative_mind_uneditioned"
		}
		if ExtraAttributes.IsShiny && prices[itemId+"_shiny"] != 0 {
			itemId += "_shiny"
		}
		if strings.HasPrefix(ExtraAttributes.ID, "STARRED_") && prices[itemId] == 0 && prices[strings.Replace(itemId, "starred_", "", 1)] != 0 {
			itemId = strings.Replace(itemId, "starred_", "", 1)
		}

		itemData := prices[itemId]
		price := itemData * item.Count
		base := itemData * item.Count
		if ExtraAttributes.Skin != "" {
			newPrice := prices[strings.ToLower(ExtraAttributes.ID)]
			if newPrice != 0 && newPrice > price {
				price = newPrice * item.Count
				base = newPrice * item.Count
			}
		}
		if price == 0 && ExtraAttributes.Price != 0 {
			price = ExtraAttributes.Price * 85 / 100
			base = ExtraAttributes.Price * 85 / 100
		}
		calculation := []CalculationData{}

		if ExtraAttributes.ID == "PICKONIMBUS" && ExtraAttributes.PickonimbusDurability != 0 {
			reduction := float64(ExtraAttributes.PickonimbusDurability) / float64(PickonimbusDurability)
			price += int(float64(price) * (reduction - 1))
			base += int(float64(price) * (reduction - 1))
		}

		if itemId != "attribute_shard" && ExtraAttributes.Attributes != nil {
			sortedAttributes := make([]string, 0, len(ExtraAttributes.Attributes))
			for attribute := range ExtraAttributes.Attributes {
				sortedAttributes = append(sortedAttributes, attribute)
			}
			sort.Strings(sortedAttributes)
			formattedId := strings.ReplaceAll(itemId, "(hot_|fiery_|burning_|infernal_)", "")
			godRollId := formattedId
			for _, attribute := range sortedAttributes {
				godRollId += "_roll_" + strings.ToLower(attribute)
			}
			godRollPrice := prices[godRollId]
			if godRollPrice > price {
				price = godRollPrice
				base = godRollPrice
				calculation = append(calculation, CalculationData{
					ID:    godRollId[len(formattedId)+1:],
					Type:  "god_roll",
					Price: godRollPrice,
					Count: 1,
				})
			}
		}

		if itemData == 0 {
			prestige := Prestiges[strings.ToUpper(itemId)]
			if prestige != nil {
				for _, prestigeItem := range prestige {
					foundItem := getHypixelItemInformationFromId(prestigeItem)
					if price == 0 {
						price = 0
					}
					if foundItem.UpgradeCosts != nil {
						price += starCosts(prices, &calculation, foundItem.UpgradeCosts, prestigeItem)
					}
					if foundItem.Prestige.Costs != nil {
						price += starCosts(prices, &calculation, foundItem.Prestige.Costs, prestigeItem)
					}
				}
			}
		}

		if ExtraAttributes.Price != 0 && ExtraAttributes.Auction != 0 && ExtraAttributes.Bid != 0 {
			pricePaid := ExtraAttributes.Price * ApplicationWorth.ShensAuctionPrice
			if pricePaid > price {
				price = pricePaid
				calculation = append(calculation, CalculationData{
					ID:    itemId,
					Type:  "shens_auction",
					Price: pricePaid,
					Count: 1,
				})
			}
		}

		if itemId == "midas_sword" || itemId == "starred_midas_sword" || itemId == "midas_staff" || itemId == "starred_midas_staff" {
			maxBid := 50000000
			if strings.Contains(itemId, "midas_staff") {
				maxBid = 100000000
			}
			typeStr := itemId + "_50m"
			if strings.Contains(itemId, "midas_staff") {
				typeStr = itemId + "_100m"
			}

			if ExtraAttributes.WinningBid+ExtraAttributes.AdditionalCoins >= maxBid {
				calculationData := CalculationData{
					ID:    itemId,
					Type:  typeStr,
					Price: prices[typeStr],
					Count: 1,
				}
				price = calculationData.Price
				calculation = append(calculation, calculationData)
			} else {
				calculationData := CalculationData{
					ID:    itemId,
					Type:  "winning_bid",
					Price: ExtraAttributes.WinningBid * ApplicationWorth.WinningBid,
					Count: 1,
				}
				price = calculationData.Price
				calculation = append(calculation, calculationData)

				if ExtraAttributes.AdditionalCoins != 0 {
					calculationData := CalculationData{
						ID:    itemId,
						Type:  "additional_coins",
						Price: ExtraAttributes.AdditionalCoins * ApplicationWorth.WinningBid,
						Count: 1,
					}
					price += calculationData.Price
					calculation = append(calculation, calculationData)
				}
			}
		}

		if itemId == "enchanted_book" && ExtraAttributes.Enchantments != nil {
			if len(ExtraAttributes.Enchantments) == 1 {
				for name, value := range ExtraAttributes.Enchantments {
					calculationData := CalculationData{
						ID:    strings.ToUpper(name + "_" + strconv.Itoa(value)),
						Type:  "enchant",
						Price: prices["enchantment_"+strings.ToLower(name)+"_"+strconv.Itoa(value)],
						Count: 1,
					}
					price = calculationData.Price
					calculation = append(calculation, calculationData)
					itemName = SpecialEnchantmentMatches[name]
					if itemName == "" {
						itemName = titleCase(strings.ReplaceAll(name, "_", " "))
					}
					break
				}
			} else {
				enchantmentPrice := 0
				for name, value := range ExtraAttributes.Enchantments {
					calculationData := CalculationData{
						ID:    strings.ToUpper(name + "_" + strconv.Itoa(value)),
						Type:  "enchant",
						Price: prices["enchantment_"+strings.ToLower(name)+"_"+strconv.Itoa(value)] * ApplicationWorth.Enchants,
						Count: 1,
					}
					enchantmentPrice += calculationData.Price
					calculation = append(calculation, calculationData)
				}
				price = enchantmentPrice
			}
		} else if ExtraAttributes.Enchantments != nil {
			for name, value := range ExtraAttributes.Enchantments {
				name = strings.ToLower(name)
				if contains(BlockedEnchants[itemId], name) {
					continue
				}
				if IgnoredEnchants[name] == value {
					continue
				}

				if contains(StackingEnchants, name) {
					value = 1
				}

				if name == "efficiency" && value > 5 && !contains(IgnoreSilex, itemId) {
					efficiencyLevel := value - 5
					if itemId == "stonk_pickaxe" {
						efficiencyLevel = value - 6
					}
					if efficiencyLevel > 0 {
						calculationData := CalculationData{
							ID:    "SIL_EX",
							Type:  "silex",
							Price: prices["sil_ex"] * efficiencyLevel * ApplicationWorth.Silex,
							Count: efficiencyLevel,
						}
						price += calculationData.Price
						calculation = append(calculation, calculationData)
					}
				}

				if name == "scavenger" && value >= 6 {
					calculationData := CalculationData{
						ID:    "GOLDEN_BOUNTY",
						Type:  "golden_bounty",
						Price: prices["GOLDEN_BOUNTY"] * ApplicationWorth.GoldenBounty,
						Count: 1,
					}
					price += calculationData.Price
					calculation = append(calculation, calculationData)
				}

				calculationData := CalculationData{
					ID:    strings.ToUpper(name + "_" + strconv.Itoa(value)),
					Type:  "enchant",
					Price: prices["enchantment_"+name+"_"+strconv.Itoa(value)] * EnchantsWorth[name],
					Count: 1,
				}
				if calculationData.Price != 0 {
					price += calculationData.Price
					calculation = append(calculation, calculationData)
				}
			}
		}

		if ExtraAttributes.Attributes != nil {
			for attribute, tier := range ExtraAttributes.Attributes {
				if tier == 1 {
					continue
				}
				shards := (1 << (tier - 1)) - 1
				baseAttributePrice := prices["attribute_shard_"+attribute]
				if AttributesBaseCosts[itemId] != "" && prices[AttributesBaseCosts[itemId]] < baseAttributePrice {
					baseAttributePrice = prices[AttributesBaseCosts[itemId]]
				} else if strings.HasPrefix(itemId, "aurora") && prices["kuudra_helmet_"+attribute] < baseAttributePrice {
					baseAttributePrice = prices["kuudra_helmet_"+attribute]
				} else if strings.HasPrefix(itemId, "aurora") {
					kuudraPrices := []int{prices["kuudra_chestplate_"+attribute], prices["kuudra_leggings_"+attribute], prices["kuudra_boots_"+attribute]}
					kuudraPrice := 0
					for _, price := range kuudraPrices {
						kuudraPrice += price
					}
					kuudraPrice /= len(kuudraPrices)
					if kuudraPrice != 0 && (baseAttributePrice == 0 || kuudraPrice < baseAttributePrice) {
						baseAttributePrice = kuudraPrice
					}
				}
				if baseAttributePrice == 0 {
					continue
				}
				attributePrice := baseAttributePrice * shards * ApplicationWorth.Attributes
				price += attributePrice
				calculation = append(calculation, CalculationData{
					ID:    strings.ToUpper(attribute + "_" + strconv.Itoa(tier)),
					Type:  "attribute",
					Price: attributePrice,
					Count: 1,
					Shards: shards,
				})
			}
		}

		if ExtraAttributes.SackPSS != 0 {
			calculationData := CalculationData{
				ID:    "POCKET_SACK_IN_A_SACK",
				Type:  "pocket_sack_in_a_sack",
				Price: prices["pocket_sack_in_a_sack"] * ExtraAttributes.SackPSS * ApplicationWorth.PocketSackInASack,
				Count: ExtraAttributes.SackPSS,
			}
			price += calculationData.Price
			calculation = append(calculation, calculationData)
		}

		if ExtraAttributes.WoodSingularityCount != 0 {
			calculationData := CalculationData{
				ID:    "WOOD_SINGULARITY",
				Type:  "wood_singularity",
				Price: prices["wood_singularity"] * ExtraAttributes.WoodSingularityCount * ApplicationWorth.WoodSingularity,
				Count: ExtraAttributes.WoodSingularityCount,
			}
			price += calculationData.Price
			calculation = append(calculation, calculationData)
		}

		if ExtraAttributes.JalapenoCount != 0 {
			calculationData := CalculationData{
				ID:    "JALAPENO_BOOK",
				Type:  "jalapeno_book",
				Price: prices["jalapeno_book"] * ExtraAttributes.JalapenoCount * ApplicationWorth.JalapenoBook,
				Count: ExtraAttributes.JalapenoCount,
			}
			price += calculationData.Price
			calculation = append(calculation, calculationData)
		}

		if ExtraAttributes.TunedTransmission != 0 {
			calculationData := CalculationData{
				ID:    "TRANSMISSION_TUNER",
				Type:  "tuned_transmission",
				Price: prices["transmission_tuner"] * ExtraAttributes.TunedTransmission * ApplicationWorth.TunedTransmission,
				Count: ExtraAttributes.TunedTransmission,
			}
			price += calculationData.Price
			calculation = append(calculation, calculationData)
		}

		if ExtraAttributes.ManaDisintegratorCount != 0 {
			calculationData := CalculationData{
				ID:    "MANA_DISINTEGRATOR",
				Type:  "mana_disintegrator",
				Price: prices["mana_disintegrator"] * ExtraAttributes.ManaDisintegratorCount * ApplicationWorth.ManaDisintegrator,
				Count: ExtraAttributes.ManaDisintegratorCount,
			}
			price += calculationData.Price
			calculation = append(calculation, calculationData)
		}

		if ExtraAttributes.ThunderCharge != 0 && itemId == "pulse_ring" {
			thunderUpgrades := ExtraAttributes.ThunderCharge / 50000
			calculationData := CalculationData{
				ID:    "THUNDER_IN_A_BOTTLE",
				Type:  "thunder_charge",
				Price: prices["thunder_in_a_bottle"] * thunderUpgrades * ApplicationWorth.ThunderInABottle,
				Count: thunderUpgrades,
			}
			price += calculationData.Price
			calculation = append(calculation, calculationData)
		}

		if ExtraAttributes.Runes != nil && !strings.HasPrefix(itemId, "rune") {
			for runeType, runeTier := range ExtraAttributes.Runes {
				runeId := runeType + "_" + runeTier
				if contains(ValidRunes, runeId) {
					calculationData := CalculationData{
						ID:    "RUNE_" + strings.ToUpper(runeId),
						Type:  "rune",
						Price: prices["rune_"+strings.ToLower(runeId)] * ApplicationWorth.Runes,
						Count: 1,
					}
					price += calculationData.Price
					calculation = append(calculation, calculationData)
				}
			}
		}

		if ExtraAttributes.HotPotatoCount != 0 {
			hotPotatoCount := ExtraAttributes.HotPotatoCount
			if hotPotatoCount > 10 {
				calculationData := CalculationData{
					ID:    "FUMING_POTATO_BOOK",
					Type:  "fuming_potato_book",
					Price: prices["fuming_potato_book"] * (hotPotatoCount - 10) * ApplicationWorth.FumingPotatoBook,
					Count: hotPotatoCount - 10,
				}
				price += calculationData.Price
				calculation = append(calculation, calculationData)
			}
			calculationData := CalculationData{
				ID:    "HOT_POTATO_BOOK",
				Type:  "hot_potato_book",
				Price: prices["hot_potato_book"] * min(hotPotatoCount, 10) * ApplicationWorth.HotPotatoBook,
				Count: min(hotPotatoCount, 10),
			}
			price += calculationData.Price
			calculation = append(calculation, calculationData)
		}

		if ExtraAttributes.DyeItem != "" {
			calculationData := CalculationData{
				ID:    ExtraAttributes.DyeItem,
				Type:  "dye",
				Price: prices[strings.ToLower(ExtraAttributes.DyeItem)] * ApplicationWorth.Dye,
				Count: 1,
			}
			price += calculationData.Price
			calculation = append(calculation, calculationData)
		}

		if ExtraAttributes.ArtOfWarCount != 0 {
			artOfWarCount := ExtraAttributes.ArtOfWarCount
			calculationData := CalculationData{
				ID:    "THE_ART_OF_WAR",
				Type:  "the_art_of_war",
				Price: prices["the_art_of_war"] * artOfWarCount * ApplicationWorth.ArtOfWar,
				Count: artOfWarCount,
			}
			price += calculationData.Price
			calculation = append(calculation, calculationData)
		}

		if ExtraAttributes.ArtOfPeaceApplied != 0 {
			calculationData := CalculationData{
				ID:    "THE_ART_OF_PEACE",
				Type:  "the_art_of_peace",
				Price: prices["the_art_of_peace"] * ExtraAttributes.ArtOfPeaceApplied * ApplicationWorth.ArtOfPeace,
				Count: ExtraAttributes.ArtOfPeaceApplied,
			}
			price += calculationData.Price
			calculation = append(calculation, calculationData)
		}

		if ExtraAttributes.FarmingForDummiesCount != 0 {
			calculationData := CalculationData{
				ID:    "FARMING_FOR_DUMMIES",
				Type:  "farming_for_dummies",
				Price: prices["farming_for_dummies"] * ExtraAttributes.FarmingForDummiesCount * ApplicationWorth.FarmingForDummies,
				Count: ExtraAttributes.FarmingForDummiesCount,
			}
			price += calculationData.Price
			calculation = append(calculation, calculationData)
		}

		if ExtraAttributes.Polarvoid != 0 {
			calculationData := CalculationData{
				ID:    "POLARVOID_BOOK",
				Type:  "polarvoid_book",
				Price: prices["polarvoid_book"] * ExtraAttributes.Polarvoid * ApplicationWorth.Polarvoid,
				Count: ExtraAttributes.Polarvoid,
			}
			price += calculationData.Price
			calculation = append(calculation, calculationData)
		}

		if ExtraAttributes.DivanPowderCoating != 0 {
			calculationData := CalculationData{
				ID:    "DIVAN_POWDER_COATING",
				Type:  "divan_powder_coating",
				Price: prices["divan_powder_coating"] * ApplicationWorth.DivanPowderCoating,
				Count: ExtraAttributes.DivanPowderCoating,
			}
			price += calculationData.Price
			calculation = append(calculation, calculationData)
		}

		if ExtraAttributes.TalismanEnrichment != "" {
			enrichmentPrice := math.MaxInt64
			for _, enrichment := range Enrichments {
				if prices[strings.ToLower(enrichment)] < enrichmentPrice {
					enrichmentPrice = prices[strings.ToLower(enrichment)]
				}
			}
			if enrichmentPrice != math.MaxInt64 {
				calculationData := CalculationData{
					ID:    strings.ToUpper(ExtraAttributes.TalismanEnrichment),
					Type:  "talisman_enrichment",
					Price: enrichmentPrice * ApplicationWorth.Enrichment,
					Count: 1,
				}
				price += calculationData.Price
				calculation = append(calculation, calculationData)
			}
		}

		if ExtraAttributes.RarityUpgrades > 0 && ExtraAttributes.ItemTier == 0 {
			lastLoreLine := item.Tag.Display.Lore[len(item.Tag.Display.Lore)-1]
			if ExtraAttributes.Enchantments != nil || contains(AllowedRecombTypes, skyblockItem.Category) || contains(AllowedRecombIds, itemId) || strings.Contains(lastLoreLine, "ACCESSORY") || strings.Contains(lastLoreLine, "HATCESSORY") {
				recombApplicationWorth := ApplicationWorth.Recomb
				if itemId == "bone_boomerang" {
					recombApplicationWorth *= 0.5
				}
				calculationData := CalculationData{
					ID:    "RECOMBOBULATOR_3000",
					Type:  "recombobulator_3000",
					Price: prices["recombobulator_3000"] * recombApplicationWorth,
					Count: 1,
				}
				price += calculationData.Price
				calculation = append(calculation, calculationData)
			}
		}

		if ExtraAttributes.Gems != nil {
			unlockedSlots := []string{}
			gems := []struct {
				Type     string
				Tier     int
				SlotType string
			}{}
			if skyblockItem.GemstoneSlots != nil {
				if ExtraAttributes.Gems.Formatted {
					unlockedSlots = ExtraAttributes.Gems.UnlockedSlots
					gems = ExtraAttributes.Gems.Gems
				} else {
					ExtraAttributesGems := ExtraAttributes.Gems
					for _, slot := range skyblockItem.GemstoneSlots {
						if slot.Costs != nil && ExtraAttributesGems.UnlockedSlots != nil {
							for index, typeStr := range ExtraAttributesGems.UnlockedSlots {
								if strings.HasPrefix(typeStr, slot.SlotType) {
									unlockedSlots = append(unlockedSlots, slot.SlotType)
									ExtraAttributesGems.UnlockedSlots = append(ExtraAttributesGems.UnlockedSlots[:index], ExtraAttributesGems.UnlockedSlots[index+1:]...)
									break
								}
							}
						}
						if slot.Costs == nil {
							unlockedSlots = append(unlockedSlots, slot.SlotType)
						}
						key := ""
						for k := range ExtraAttributesGems {
							if strings.HasPrefix(k, slot.SlotType) && !strings.HasSuffix(k, "_gem") {
								key = k
								break
							}
						}
						if key != "" {
							typeStr := slot.SlotType
							if contains([]string{"COMBAT", "OFFENSIVE", "DEFENSIVE", "MINING", "UNIVERSAL", "CHISEL"}, slot.SlotType) {
								typeStr = ExtraAttributesGems[key+"_gem"]
							}
							gems = append(gems, struct {
								Type     string
								Tier     int
								SlotType string
							}{
								Type:     typeStr,
								Tier:     ExtraAttributesGems[key],
								SlotType: slot.SlotType,
							})
							delete(ExtraAttributesGems, key)
							if slot.Costs != nil && ExtraAttributesGems.UnlockedSlots == nil {
								unlockedSlots = append(unlockedSlots, slot.SlotType)
							}
						}
					}
				}

				isDivansArmor := contains([]string{"divan_helmet", "divan_chestplate", "divan_leggings", "divan_boots"}, itemId)
				if isDivansArmor || strings.HasPrefix(itemId, "aurora") {
					application := ApplicationWorth.GemstoneChambers
					if !isDivansArmor {
						application = ApplicationWorth.GemstoneSlots
					}
					gemstoneSlots := skyblockItem.GemstoneSlots
					for _, unlockedSlot := range unlockedSlots {
						slotIndex := -1
						for i, slot := range gemstoneSlots {
							if slot.SlotType == unlockedSlot {
								slotIndex = i
								break
							}
						}
						if slotIndex != -1 {
							total := 0
							for _, cost := range gemstoneSlots[slotIndex].Costs {
								if cost.Type == "COINS" {
									total += cost.Coins
								} else if cost.Type == "ITEM" {
									total += prices[strings.ToLower(cost.ItemID)] * cost.Amount
								}
							}
							calculationData := CalculationData{
								ID:    unlockedSlot,
								Type:  "gemstone_slot",
								Price: total * application,
								Count: 1,
							}
							price += calculationData.Price
							calculation = append(calculation, calculationData)
							gemstoneSlots = append(gemstoneSlots[:slotIndex], gemstoneSlots[slotIndex+1:]...)
						}
					}
				}

				for _, gemstone := range gems {
					calculationData := CalculationData{
						ID:    strings.ToUpper(gemstone.Tier + "_" + gemstone.Type + "_GEM"),
						Type:  "gemstone",
						Price: prices[strings.ToLower(gemstone.Tier+"_"+gemstone.Type+"_gem")] * ApplicationWorth.Gemstone,
						Count: 1,
					}
					price += calculationData.Price
					calculation = append(calculation, calculationData)
				}
			}
		}

		if ExtraAttributes.PowerAbilityScroll != "" {
			calculationData := CalculationData{
				ID:    ExtraAttributes.PowerAbilityScroll,
				Type:  "gemstone_power_scroll",
				Price: prices[ExtraAttributes.PowerAbilityScroll] * ApplicationWorth.GemstonePowerScroll,
				Count: 1,
			}
			price += calculationData.Price
			calculation = append(calculation, calculationData)
		}

		if ExtraAttributes.Modifier != "" && skyblockItem.Category != "ACCESSORY" {
			reforge := ExtraAttributes.Modifier
			if Reforges[reforge] != "" {
				calculationData := CalculationData{
					ID:    Reforges[reforge],
					Type:  "reforge",
					Price: prices[Reforges[reforge]] * ApplicationWorth.Reforge,
					Count: 1,
				}
				price += calculationData.Price
				calculation = append(calculation, calculationData)
			}
		}

		dungeonItemLevel := ExtraAttributes.DungeonItemLevel
		upgradeLevel := ExtraAttributes.UpgradeLevel
		if skyblockItem.UpgradeCosts != nil && (dungeonItemLevel > 5 || upgradeLevel > 5) {
			starsUsedDungeons := dungeonItemLevel - 5
			starsUsedUpgrade := upgradeLevel - 5
			starsUsed := max(starsUsedDungeons, starsUsedUpgrade)
			if len(skyblockItem.UpgradeCosts) <= 5 {
				for star := 0; star < starsUsed; star++ {
					calculationData := CalculationData{
						ID:    MasterStars[star],
						Type:  "master_star",
						Price: prices[MasterStars[star]] * ApplicationWorth.MasterStar,
						Count: 1,
					}
					price += calculationData.Price
					calculation = append(calculation, calculationData)
				}
			}
		}

		if skyblockItem.UpgradeCosts != nil && (dungeonItemLevel > 0 || upgradeLevel > 0) {
			level := max(dungeonItemLevel, upgradeLevel)
			price += starCosts(prices, &calculation, skyblockItem.UpgradeCosts[:level+1], "")
		}

		if ExtraAttributes.AbilityScroll != nil {
			for _, id := range ExtraAttributes.AbilityScroll {
				calculationData := CalculationData{
					ID:    id,
					Type:  "necron_scroll",
					Price: prices[strings.ToLower(id)] * ApplicationWorth.NecronBladeScroll,
					Count: 1,
				}
				price += calculationData.Price
				calculation = append(calculation, calculationData)
			}
		}

		drillPartTypes := []string{"drill_part_upgrade_module", "drill_part_fuel_tank", "drill_part_engine"}
		for _, typeStr := range drillPartTypes {
			if ExtraAttributes[typeStr] != "" {
				calculationData := CalculationData{
					ID:    strings.ToUpper(ExtraAttributes[typeStr]),
					Type:  "drill_part",
					Price: prices[ExtraAttributes[typeStr]] * ApplicationWorth.DrillPart,
					Count: 1,
				}
				price += calculationData.Price
				calculation = append(calculation, calculationData)
			}
		}

		if ExtraAttributes.Ethermerge > 0 {
			calculationData := CalculationData{
				ID:    "ETHERWARP_CONDUIT",
				Type:  "etherwarp_conduit",
				Price: prices["etherwarp_conduit"] * ApplicationWorth.Etherwarp,
				Count: 1,
			}
			price += calculationData.Price
			calculation = append(calculation, calculationData)
		}

		if ExtraAttributes.NewYearCakeBagYears != nil {
			cakesPrice := 0
			for _, year := range ExtraAttributes.NewYearCakeBagYears {
				cakesPrice += prices["new_year_cake_"+strconv.Itoa(year)]
			}
			calculationData := CalculationData{
				ID:    "NEW_YEAR_CAKES",
				Type:  "new_year_cakes",
				Price: cakesPrice,
				Count: 1,
			}
			price += calculationData.Price
			calculation = append(calculation, calculationData)
		}

		isSoulbound := ExtraAttributes.DonatedMuseum || contains(item.Tag.Display.Lore, "§8§l* §8Co-op Soulbound §8§l*") || contains(item.Tag.Display.Lore, "§8§l* §8Soulbound §8§l*")
		data := &Item{
			Tag: struct {
				ExtraAttributes struct {
					ID                   string
					PetInfo              string
					Skin                 string
					PartyHatEmoji        string
					NewYearsCake         int
					Edition              int
					IsShiny              bool
					Price                int
					PickonimbusDurability int
					Attributes           map[string]int
					SackPSS              int
					WoodSingularityCount int
					JalapenoCount        int
					TunedTransmission    int
					ManaDisintegratorCount int
					ThunderCharge        int
					HotPotatoCount       int
					DyeItem              string
					ArtOfWarCount        int
					ArtOfPeaceApplied    int
					FarmingForDummiesCount int
					Polarvoid            int
					DivanPowderCoating   int
					TalismanEnrichment   string
					RarityUpgrades       int
					Gems                 struct {
						Formatted     bool
						UnlockedSlots []string
						Gems          []struct {
							Type     string
							Tier     int
							SlotType string
						}
					}
					PowerAbilityScroll string
					Modifier           string
					DungeonItemLevel   int
					UpgradeLevel       int
					AbilityScroll      []string
					DrillPartUpgradeModule string
					DrillPartFuelTank  string
					DrillPartEngine    string
					Ethermerge         int
					NewYearCakeBagYears []int
					DonatedMuseum      bool
				}
				Display struct {
					Name string
					Lore []string
				}
			}{
				ExtraAttributes: struct {
					ID                   string
					PetInfo              string
					Skin                 string
					PartyHatEmoji        string
					NewYearsCake         int
					Edition              int
					IsShiny              bool
					Price                int
					PickonimbusDurability int
					Attributes           map[string]int
					SackPSS              int
					WoodSingularityCount int
					JalapenoCount        int
					TunedTransmission    int
					ManaDisintegratorCount int
					ThunderCharge        int
					HotPotatoCount       int
					DyeItem              string
					ArtOfWarCount        int
					ArtOfPeaceApplied    int
					FarmingForDummiesCount int
					Polarvoid            int
					DivanPowderCoating   int
					TalismanEnrichment   string
					RarityUpgrades       int
					Gems                 struct {
						Formatted     bool
						UnlockedSlots []string
						Gems          []struct {
							Type     string
							Tier     int
							SlotType string
						}
					}
					PowerAbilityScroll string
					Modifier           string
					DungeonItemLevel   int
					UpgradeLevel       int
					AbilityScroll      []string
					DrillPartUpgradeModule string
					DrillPartFuelTank  string
					DrillPartEngine    string
					Ethermerge         int
					NewYearCakeBagYears []int
					DonatedMuseum      bool
				}{
					ID:                   itemId,
					PetInfo:              ExtraAttributes.PetInfo,
					Skin:                 ExtraAttributes.Skin,
					PartyHatEmoji:        ExtraAttributes.PartyHatEmoji,
					NewYearsCake:         ExtraAttributes.NewYearsCake,
					Edition:              ExtraAttributes.Edition,
					IsShiny:              ExtraAttributes.IsShiny,
					Price:                price,
					PickonimbusDurability: ExtraAttributes.PickonimbusDurability,
					Attributes:           ExtraAttributes.Attributes,
					SackPSS:              ExtraAttributes.SackPSS,
					WoodSingularityCount: ExtraAttributes.WoodSingularityCount,
					JalapenoCount:        ExtraAttributes.JalapenoCount,
					TunedTransmission:    ExtraAttributes.TunedTransmission,
					ManaDisintegratorCount: ExtraAttributes.ManaDisintegratorCount,
					ThunderCharge:        ExtraAttributes.ThunderCharge,
					HotPotatoCount:       ExtraAttributes.HotPotatoCount,
					DyeItem:              ExtraAttributes.DyeItem,
					ArtOfWarCount:        ExtraAttributes.ArtOfWarCount,
					ArtOfPeaceApplied:    ExtraAttributes.ArtOfPeaceApplied,
					FarmingForDummiesCount: ExtraAttributes.FarmingForDummiesCount,
					Polarvoid:            ExtraAttributes.Polarvoid,
					DivanPowderCoating:   ExtraAttributes.DivanPowderCoating,
					TalismanEnrichment:   ExtraAttributes.TalismanEnrichment,
					RarityUpgrades:       ExtraAttributes.RarityUpgrades,
					Gems:                 ExtraAttributes.Gems,
					PowerAbilityScroll:   ExtraAttributes.PowerAbilityScroll,
					Modifier:             ExtraAttributes.Modifier,
					DungeonItemLevel:     ExtraAttributes.DungeonItemLevel,
					UpgradeLevel:         ExtraAttributes.UpgradeLevel,
					AbilityScroll:        ExtraAttributes.AbilityScroll,
					DrillPartUpgradeModule: ExtraAttributes.DrillPartUpgradeModule,
					DrillPartFuelTank:    ExtraAttributes.DrillPartFuelTank,
					DrillPartEngine:      ExtraAttributes.DrillPartEngine,
					Ethermerge:           ExtraAttributes.Ethermerge,
					NewYearCakeBagYears:  ExtraAttributes.NewYearCakeBagYears,
					DonatedMuseum:        ExtraAttributes.DonatedMuseum,
				},
				Display: struct {
					Name string
					Lore []string
				}{
					Name: itemName,
					Lore: item.Tag.Display.Lore,
				}
			},
			Count: item.Count,
		}
		if returnItemData {
			data.Tag.ExtraAttributes = ExtraAttributes
		}
		return data
	}
	return nil
}
