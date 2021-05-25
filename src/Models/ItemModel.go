package Models

import (
    "github.com/FuzzyStatic/blizzard/v2/wowgd"
)

type Item struct {
  ID uint
  ItemID int
  Name string
  SubclassName string
  SubclassID int
  QualityType string
  QualityName string
  Level int
  RequiredLevel int
  ItemClassName string
  ItemClassID int
  InventoryTypeName string
  InventoryType string
  PurchasePrice int
  SellPrice int
  MaxCount int
  IsEquippable bool
  IsStackable bool
  AuctionID int
  Buyout int
  Quantity int
  // UnitPrice int
  TimeLeft wowgd.TimeLeft
  Table    string `gorm:"-"`
}

func (item *Item) TableName() string {
  switch item.Table {
  case "aegwynn":
      item.Table = "aegwynn"
  case "aerie peak":
      item.Table = "aerie_peak"
  case "agamaggan":
      item.Table = "agamaggan"
  case "aggramar":
      item.Table = "aggramar"
  case "akama":
      item.Table = "akama"
  case "alexstrasza":
      item.Table = "alexstrasza"
  case "alleria":
      item.Table = "alleria"
  case "altar of storms":
      item.Table = "altar_of_storms"
  case "alterac mountains":
      item.Table = "alterac_mountains"
  case "andorhal":
      item.Table = "andorhal"
  case "anetheron":
      item.Table = "anetheron"
  case "antonidas":
      item.Table = "antonidas"
  case "anub'arak":
      item.Table = "anub_arak"
  case "anvilmar":
      item.Table = "anvilmar"
  case "arathor":
      item.Table = "arathor"
  case "archimonde":
      item.Table = "archimonde"
  case "area 52":
      item.Table = "area_52"
  case "argent dawn":
      item.Table = "argent_dawn"
  case "arthas":
      item.Table = "arthas"
  case "arygos":
      item.Table = "arygos"
  case "auchindoun":
      item.Table = "auchindoun"
  case "azgalor":
      item.Table = "azgalor"
  case "azjol-nerub":
      item.Table = "azjol_nerub"
  case "azshara":
      item.Table = "azshara"
  case "azuremyst":
      item.Table = "azuremyst"
  case "baelgun":
      item.Table = "baelgun"
  case "balnazzar":
      item.Table = "balnazzar"
  case "black dragonflight":
      item.Table = "black_dragonflight"
  case "blackhand":
      item.Table = "blackhand"
  case "blackrock":
      item.Table = "blackrock"
  case "blackwater-raiders":
      item.Table = "blackwater_raiders"
  case "blackwing lair":
      item.Table = "blackwing_lair"
  case "blades-edge":
      item.Table = "blades_edge"
  case "bleeding hollow":
      item.Table = "bleeding_hollow"
  case "blood furnace":
      item.Table = "blood_furnace"
  case "bloodhoof":
      item.Table = "bloodhoof"
  case "bloodscalp":
      item.Table = "bloodscalp"
  case "bonechewer":
      item.Table = "bonechewer"
  case "borean tundra":
      item.Table = "borean_tundra"
  case "boulderfist":
      item.Table = "boulderfist"
  case "bronzebeard":
      item.Table = "bronzebeard"
  case "burning blade":
      item.Table = "burning_blade"
  case "burning legion":
      item.Table = "burning_legion"
  case "cairne":
      item.Table = "cairne"
  case "cenarion circle":
      item.Table = "cenarion_circle"
  case "cenarius":
      item.Table = "cenarius"
  case "cho'gall":
      item.Table = "cho_gall"
  case "chromaggus":
      item.Table = "chromaggus"
  case "coilfang":
      item.Table = "coilfang"
  case "crushridge":
      item.Table = "crushridge"
  case "daggerspine":
      item.Table = "daggerspine"
  case "dalaran":
      item.Table = "dalaran"
  case "dalvengyr":
      item.Table = "dalvengyr"
  case "dark iron":
      item.Table = "dark_iron"
  case "darkspear":
      item.Table = "darkspear"
  case "darrowmere":
      item.Table = "darrowmere"
  case "dawnbringer":
      item.Table = "dawnbringer"
  case "deathwing":
      item.Table = "deathwing"
  case "demon soul":
      item.Table = "demon_soul"
  case "dentarg":
      item.Table = "dentarg"
  case "destromath":
      item.Table = "destromath"
  case "dethecus":
      item.Table = "dethecus"
  case "detheroc":
      item.Table = "detheroc"
  case "doomhammer":
      item.Table = "doomhammer"
  case "draenor":
      item.Table = "draenor"
  case "dragonblight":
      item.Table = "dragonblight"
  case "dragonmaw":
      item.Table = "dragonmaw"
  case "drak'tharon":
      item.Table = "drak_tharon"
  case "drak'thul":
      item.Table = "drak_thul"
  case "draka":
      item.Table = "draka"
  case "drenden":
      item.Table = "drenden"
  case "dunemaul":
      item.Table = "dunemaul"
  case "durotan":
      item.Table = "durotan"
  case "duskwood":
      item.Table = "duskwood"
  case "earthen ring":
      item.Table = "earthen_ring"
  case "echo isles":
      item.Table = "echo_isles"
  case "eitrigg":
      item.Table = "eitrigg"
  case "eldre'thalas":
      item.Table = "eldre_thalas"
  case "elune":
      item.Table = "elune"
  case "emerald dream":
      item.Table = "emerald_dream"
  case "eonar":
      item.Table = "eonar"
  case "eredar":
      item.Table = "eredar"
  case "executus":
      item.Table = "executus"
  case "exodar":
      item.Table = "exodar"
  case "farstriders":
      item.Table = "farstriders"
  case "feathermoon":
      item.Table = "feathermoon"
  case "fenris":
      item.Table = "fenris"
  case "firetree":
      item.Table = "firetree"
  case "fizzcrank":
      item.Table = "fizzcrank"
  case "frostmane":
      item.Table = "frostmane"
  case "frostwolf":
      item.Table = "frostwolf"
  case "galakrond":
      item.Table = "galakrond"
  case "garithos":
      item.Table = "garithos"
  case "garona":
      item.Table = "garona"
  case "garrosh":
      item.Table = "garrosh"
  case "ghostlands":
      item.Table = "ghostlands"
  case "gilneas":
      item.Table = "gilneas"
  case "gnomeregan":
      item.Table = "gnomeregan"
  case "gorefiend":
      item.Table = "gorefiend"
  case "gorgonnash":
      item.Table = "gorgonnash"
  case "greymane":
      item.Table = "greymane"
  case "grizzly hills":
      item.Table = "grizzly_hills"
  case "gul'dan":
      item.Table = "gul_dan"
  case "gurubashi":
      item.Table = "gurubashi"
  case "hakkar":
      item.Table = "hakkar"
  case "haomarush":
      item.Table = "haomarush"
  case "hellscream":
      item.Table = "hellscream"
  case "hydraxis":
      item.Table = "hydraxis"
  case "hyjal":
      item.Table = "hyjal"
  case "icecrown":
      item.Table = "icecrown"
  case "illidan":
      item.Table = "illidan"
  case "jaedenar":
      item.Table = "jaedenar"
  case "kael'thas":
      item.Table = "kael_thas"
  case "kalecgos":
      item.Table = "kalecgos"
  case "kargath":
      item.Table = "kargath"
  case "kel'thuzad":
      item.Table = "kel_thuzad"
  case "khadgar":
      item.Table = "khadgar"
  case "khaz modan":
      item.Table = "khaz_modan"
  case "kil'jaeden":
      item.Table = "kil_jaeden"
  case "kilrogg":
      item.Table = "kilrogg"
  case "korgath":
      item.Table = "korgath"
  case "korialstrasz":
      item.Table = "korialstrasz"
  case "kul tiras":
      item.Table = "kul_tiras"
  case "laughing skull":
      item.Table = "laughing_skull"
  case "lethon":
      item.Table = "lethon"
  case "lightbringer":
      item.Table = "lightbringer"
  case "lightninmaelstrom":
      item.Table = "lightninmaelstrom"
  case "lightning's blade":
      item.Table = "lightnings_blade"
  case "lightninghoof":
      item.Table = "lightninghoof"
  case "llane":
      item.Table = "llane"
  case "lothar":
      item.Table = "lothar"
  case "madoran":
      item.Table = "madoran"
  case "maelstrom":
      item.Table = "maelstrom"
  case "magtheridon":
      item.Table = "magtheridon"
  case "maiev":
      item.Table = "maiev"
  case "mal'ganis":
      item.Table = "mal_ganis"
  case "malfurion":
      item.Table = "malfurion"
  case "malorne":
      item.Table = "malorne"
  case "malygos":
      item.Table = "malygos"
  case "mannoroth":
      item.Table = "mannoroth"
  case "medivh":
      item.Table = "medivh"
  case "misha":
      item.Table = "misha"
  case "mok'nathal":
      item.Table = "mok_nathal"
  case "moon guard":
      item.Table = "moon_guard"
  case "moonrunner":
      item.Table = "moonrunner"
  case "mug'thol":
      item.Table = "mug_thol"
  case "muradin":
      item.Table = "muradin"
  case "nathrezim":
      item.Table = "nathrezim"
  case "nazgrel":
      item.Table = "nazgrel"
  case "nazjatar":
      item.Table = "nazjatar"
  case "ner'zhul":
      item.Table = "ner_zhul"
  case "nesingwary":
      item.Table = "nesingwary"
  case "nordrassil":
      item.Table = "nordrassil"
  case "norgannon":
      item.Table = "norgannon"
  case "onyxia":
      item.Table = "onyxia"
  case "perenolde":
      item.Table = "perenolde"
  case "proudmoore":
      item.Table = "proudmoore"
  case "quel'thalas":
      item.Table = "quel_thalas"
  case "quel'dorei":
      item.Table = "quel_dorei"
  case "ragnaros":
      item.Table = "ragnaros"
  case "ravencrest":
      item.Table = "ravencrest"
  case "ravenholdt":
      item.Table = "ravenholdt"
  case "rexxar":
      item.Table = "rexxar"
  case "rivendare":
      item.Table = "rivendare"
  case "runetotem":
      item.Table =  "runetotem"
  case "sargeras":
      item.Table = "sargeras"
  case "saurfang":
      item.Table = "saurfang"
  case "scarlet crusade":
      item.Table = "scarlet_crusade"
  case "scolla":
      item.Table = "scolla"
  case "sen'jin":
      item.Table = "sen_jin"
  case "sentinels":
      item.Table =  "sentinels"
  case "shadow council":
      item.Table = "shadow_council"
  case "shadowmoon":
      item.Table =  "shadowmoon"
  case "shadowsong":
      item.Table = "shadowsong"
  case "shandris":
      item.Table = "shandris"
  case "shattered halls":
      item.Table = "shattered_halls"
  case "shattered hand":
      item.Table =  "shattered_hand"
  case "shu'halo":
      item.Table = "shu_halo"
  case "silver hand":
      item.Table = "silver_hand"
  case "silvemoon":
      item.Table = "silvemoon"
  case "sister of elune":
      item.Table = "sister_of_elune"
  case "skullcrusher":
      item.Table = "skullcrusher"
  case "skywall":
      item.Table = "skywall"
  case "smolderthorn":
      item.Table = "smolderthorn"
  case "spinebreaker":
      item.Table = "spinebreaker"
  case "spirestone":
      item.Table = "spirestone"
  case "staghelm":
      item.Table = "staghelm"
  case "steamwheedle cartel":
      item.Table = "steamwheedle_cartel"
  case "stonemaul":
      item.Table = "stonemaul"
  case "stormrage":
      item.Table = "stormrage"
  case "stormreaver":
      item.Table = "stormreaver"
  case "stormscale":
      item.Table = "stormscale"
  case "suramar":
      item.Table = "suramar"
  case "tanaris":
      item.Table = "tanaris"
  case "terokkar":
      item.Table =  "terokkar"
  case "thaurissan":
      item.Table =  "thaurissan"
  case "the forgotten coast":
      item.Table = "the_forgotten_coast"
  case "the scryers":
      item.Table = "the_scryers"
  case "the underbog":
      item.Table = "the_underbog"
  case "the venture co":
      item.Table = "the_venture_co"
  case "thorium brotherhood":
      item.Table = "thorium_brotherhood"
  case "thrall":
      item.Table = "thrall"
  case "thunderhorn":
      item.Table = "thunderhorn"
  case "thunderlord":
      item.Table = "thunderlord"
  case "tichondrius":
      item.Table = "tichondrius"
  case "tol barad":
      item.Table =  "tol_barad"
  case "tortheldrin":
      item.Table = "tortheldrin"
  case "trollbane":
      item.Table = "trollbane"
  case "turalyon":
      item.Table = "turalyon"
  case "twisting nether":
      item.Table = "twisting_nether"
  case "uldaman":
      item.Table = "uldaman"
  case "uldum":
      item.Table = "uldum"
  case "undermine":
      item.Table = "undermine"
  case "ursin":
      item.Table = "ursin"
  case "uther":
      item.Table = "uther"
  case "vashj":
      item.Table = "vashj"
  case "vek'nilash":
      item.Table = "vek_nilash"
  case "velen":
      item.Table =  "velen"
  case "warsong":
      item.Table =  "warsong"
  case "whisperwind":
      item.Table =  "whisperwind"
  case "wildhammer":
      item.Table =  "wildhammer"
  case "windrunner":
      item.Table =  "windrunner"
  case "winterhoof":
      item.Table = "winterhoof"
  case "wyrmrest accord":
      item.Table = "wyrmrest_accord"
  case "ysera":
      item.Table =  "ysera"
  case "ysondre":
      item.Table = "ysondre"
  case "zangarmarsh":
      item.Table =  "zangarmarsh"
  case "zul'jin":
      item.Table =  "zul_jin"
  case "zuluhead":
      item.Table = "zuluhead"
  default:
      item.Table = "item"
  }

  return item.Table
}
