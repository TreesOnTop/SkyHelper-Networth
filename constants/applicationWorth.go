package constants

import (
	"github.com/Tnze/go-mc/nbt"
	"github.com/PrismarineJS/minecraft-data/data/pc/1.16/blocks"
	"github.com/PrismarineJS/minecraft-data/data/pc/1.16/items"
	"github.com/PrismarineJS/minecraft-data/data/pc/1.16/recipes"
)

type ApplicationWorth struct {
  Enrichment             float64
  FarmingForDummies      float64
  GemstonePowerScroll    float64
  WoodSingularity        float64
  ArtOfWar               float64
  FumingPotatoBook       float64
  GemstoneSlots          float64
  Runes                  float64
  TunedTransmission      float64
  PocketSackInASack      float64
  Essence                float64
  GoldenBounty           float64
  Silex                  float64
  ArtOfPeace             float64
  DivanPowderCoating     float64
  JalapenoBook           float64
  ManaDisintegrator      float64
  Recomb                 float64
  ThunderInABottle       float64
  Enchants               float64
  ShensAuctionPrice      float64
  Dye                    float64
  GemstoneChambers       float64
  Attributes             float64
  DrillPart              float64
  Etherwarp              float64
  MasterStar             float64
  Gemstone               float64
  HotPotatoBook          float64
  NecronBladeScroll      float64
  Polarvoid              float64
  PrestigeItem           float64
  Reforge                float64
  WinningBid             float64
  PetCandy               float64
  SoulboundPetSkins      float64
  PetItem                float64
}

var applicationWorth = ApplicationWorth{
  Enrichment:             0.5,
  FarmingForDummies:      0.5,
  GemstonePowerScroll:    0.5,
  WoodSingularity:        0.5,
  ArtOfWar:               0.6,
  FumingPotatoBook:       0.6,
  GemstoneSlots:          0.6,
  Runes:                  0.6,
  TunedTransmission:      0.7,
  PocketSackInASack:      0.7,
  Essence:                0.75,
  GoldenBounty:           0.75,
  Silex:                  0.75,
  ArtOfPeace:             0.8,
  DivanPowderCoating:     0.8,
  JalapenoBook:           0.8,
  ManaDisintegrator:      0.8,
  Recomb:                 0.8,
  ThunderInABottle:       0.8,
  Enchants:               0.85,
  ShensAuctionPrice:      0.85,
  Dye:                    0.9,
  GemstoneChambers:       0.9,
  Attributes:             1,
  DrillPart:              1,
  Etherwarp:              1,
  MasterStar:             1,
  Gemstone:               1,
  HotPotatoBook:          1,
  NecronBladeScroll:      1,
  Polarvoid:              1,
  PrestigeItem:           1,
  Reforge:                1,
  WinningBid:             1,
  PetCandy:               0.65,
  SoulboundPetSkins:      0.8,
  PetItem:                1,
}

type EnchantsWorth struct {
  CounterStrike          float64
  BigBrain               float64
  UltimateInferno        float64
  Overload               float64
  UltimateSoulEater      float64
  UltimateFatalTempo     float64
}

var enchantsWorth = EnchantsWorth{
  CounterStrike:          0.2,
  BigBrain:               0.35,
  UltimateInferno:        0.35,
  Overload:               0.35,
  UltimateSoulEater:      0.35,
  UltimateFatalTempo:     0.65,
}
