package Routes

import (
  "WoW-Assistant/src/Controllers"
  "WoW-Assistant/src/Wow"
  "html/template"

  // "WoW-Assistant/src/Wow"
  "fmt"
  "github.com/gin-gonic/gin"
  "log"
  "sync"
  // "github.com/gin-contrib/multitemplate"
  "net/http"
)

var realms = map[string]interface{}{
  "Aegwynn": "aegwynn",
  "Aerie Peak": "aerie peak",
  "Agamaggan": "agamaggan",
  "Aggramar": "aggramar",
  "Akama": "akama",
  "Alexstrasza": "alexstrasza",
  "Alleria": "alleria",
  "Altar of Storms": "altar of storms",
  "Alterac Mountains": "alterac mountains",
  "Andorhal": "andorhal",
  "Anetheron": "anetheron",
  "Antonidas": "antonidas",
  "Anub'arak": "anub'arak",
  "Anvilmar": "anvilmar",
  "Arathor": "arathor",
  "Archimonde": "archimonde",
  "Area 52": "area 52",
  "Argent Dawn": "argent dawn",
  "Arthas": "arthas",
  "Arygos": "arygos",
  "Auchindoun": "auchindoun",
  "Azgalor": "azgalor",
  "Azjol-Nerub": "azjol-nerub",
  "Azshara": "azshara",
  "Azuremyst": "azuremyst",
  "Baelgun": "baelgun",
  "Balnazzar": "balnazzar",
  "Black Dragonflight": "black dragonflight",
  "Blackhand": "blackhand",
  "Blackrock": "blackrock",
  "Blackwater Raiders": "blackwater-raiders",
  "Blackwing Lair": "blackwing lair",
  "Blade's Edge": "blades-edge",
  "Bleeding Hollow": "bleeding hollow",
  "Blood Furnace": "blood furnace",
  "Bloodhoof": "bloodhoof",
  "Bloodscalp": "bloodscalp",
  "Bonechewer": "bonechewer",
  "Borean Tundra": "borean tundra",
  "Boulderfist": "boulderfist",
  "Bronzebeard": "bronzebeard",
  "Burning Blade": "burning blade",
  "Burning Legion": "burning legion",
  "Cairne": "cairne",
  "Cenarion Circle": "cenarion circle",
  "Cenarius": "cenarius",
  "Cho'gall": "cho'gall",
  "Chromaggus": "chromaggus",
  "Coilfang": "coilfang",
  "Crushridge": "crushridge",
  "Daggerspine": "daggerspine",
  "Dalaran": "dalaran",
  "Dalvengyr": "dalvengyr",
  "Dark Iron": "dark iron",
  "Darkspear": "darkspear",
  "Darrowmere": "darrowmere",
  "Dawnbringer": "dawnbringer",
  "Deathwing": "deathwing",
  "Demon Soul": "demon soul",
  "Dentarg": "dentarg",
  "Destromath": "destromath",
  "Dethecus": "dethecus",
  "Detheroc": "detheroc",
  "Doomhammer": "doomhammer",
  "Draenor": "draenor",
  "Dragonblight": "dragonblight",
  "Dragonmaw": "dragonmaw",
  "Drak'Tharon": "drak'tharon",
  "Drak'thul": "drak'thul",
  "Draka": "draka",
  "Drenden": "drenden",
  "Dunemaul": "dunemaul",
  "Durotan": "durotan",
  "Duskwood": "duskwood",
  "Earthen Ring": "earthen ring",
  "Echo Isles": "echo isles",
  "Eitrigg": "eitrigg",
  "Eldre'Thalas": "eldre'thalas",
  "Elune": "elune",
  "Emerald Dream": "emerald dream",
  "Eonar": "eonar",
  "Eredar": "eredar",
  "Executus": "executus",
  "Exodar": "exodar",
  "Farstriders": "farstriders",
  "Feathermoon": "feathermoon",
  "Fenris": "fenris",
  "Firetree": "firetree",
  "Fizzcrank": "fizzcrank",
  "Frostmane": "frostmane",
  "Frostwolf": "frostwolf",
  "Galakron": "galakrond",
  "Garithos": "garithos",
  "Garona": "garona",
  "Garrosh": "garrosh",
  "Ghostlands": "ghostlands",
  "Gilneas": "gilneas",
  "Gnomeregan": "gnomeregan",
  "Gorefiend": "gorefiend",
  "Gorgonnash": "gorgonnash",
  "Greymane": "greymane",
  "Grizzly Hills": "grizzly hills",
  "Gul'dan": "gul'dan",
  "Gurubashi": "gurubashi",
  "Hakkar": "hakkar",
  "Haomarush": "haomarush",
  "Hellscream": "hellscream",
  "Hydraxis": "hydraxis",
  "Hyjal": "hyjal",
  "Icecrown": "icecrown",
  "Illidan": "illidan",
  "Jaedenar": "jaedenar",
  "Kael'thas": "kael'thas",
  "Kalecgos": "kalecgos",
  "Kargath": "kargath",
  "Kel'Thuzad": "kel'thuzad",
  "Khadgar": "khadgar",
  "Khaz Modan": "khaz modan",
  "Kil'jaeden": "kil'jaeden",
  "Kilrogg": "kilrogg",
  "Korgath": "korgath",
  "Korialstrasz": "korialstrasz",
  "Kul Tiras": "kul tiras",
  "Laughing Skull": "laughing skull",
  "Lethon": "lethon",
  "Lightbringer": "lightbringer",
  "Lightning's Blade": "lightning's blade",
  "Lightninghoof": "lightninghoof",
  "Llane": "llane",
  "Lothar": "lothar",
  "Madoran": "madoran",
  "Maelstrom": "maelstrom",
  "Magtheridon": "magtheridon",
  "Maiev": "maiev",
  "Mal'Ganis": "mal'ganis",
  "Malfurion": "malfurion",
  "Malorne": "malorne",
  "Malygos": "malygos",
  "Mannoroth": "mannoroth",
  "Medivh": "medivh",
  "Misha": "misha",
  "Mok'Nathal": "mok'nathal",
  "Moon Guard": "moon guard",
  "Moonrunner": "moonrunner",
  "Mug'thol": "mug'thol",
  "Muradin": "muradin",
  "Nathrezim": "nathrezim",
  "Nazgrel": "nazgrel",
  "Nazjatar": "nazjatar",
  "Ner'zhul": "ner'zhul",
  "Nesingwary": "nesingwary",
  "Nordrassil": "nordrassil",
  "Norgannon": "norgannon",
  "Onyxia": "onyxia",
  "Perenolde": "perenolde",
  "Proudmoore": "proudmoore",
  "Quel'Thalas": "quel'thalas",
  "Quel'dorei": "quel'dorei",
  "Ragnaros": "ragnaros",
  "Ravencrest": "ravencrest",
  "Ravenholdt": "ravenholdt",
  "Rexxar": "rexxar",
  "Rivendare": "rivendare",
  "Runetotem": "runetotem",
  "Sargeras": "sargeras",
  "Saurfang": "saurfang",
  "Scarlet Crusade": "scarlet crusade",
  "Scilla": "scolla",
  "Sen'jin": "sen'jin",
  "Sentinels": "sentinels",
  "Shadow Council": "shadow council",
  "Shadowmoon": "shadowmoon",
  "Shadowsong": "shadowsong",
  "Shandris": "shandris",
  "Shattered Halls": "shattered halls",
  "Shattered Hand": "shattered hand",
  "Shu'halo": "shu'halo",
  "Silver Hand": "silver hand",
  "Silvermoon": "silvemoon",
  "Sister of Elune": "sister of elune",
  "Skullcrusher": "skullcrusher",
  "Skywall": "skywall",
  "Smolderthorn": "smolderthorn",
  "Spinebreaker": "spinebreaker",
  "Spirestone": "spirestone",
  "Staghelm": "staghelm",
  "Steamwheedle Cartel": "steamwheedle cartel",
  "Stonemaul": "stonemaul",
  "Stormrage": "stormrage",
  "Stormreaver": "stormreaver",
  "Stormscale": "stormscale",
  "Suramar": "suramar",
  "Tanaris": "tanaris",
  "Terokkar": "terokkar",
  "Thaurissan": "thaurissan",
  "The Forgotten Coast": "the forgotten coast",
  "The Scryers": "the scryers",
  "The Underbog": "the underbog",
  "The Venture Co": "the venture co",
  "Thorium Brotherhood": "thorium brotherhood",
  "Thrall": "thrall",
  "Thunderhorn": "thunderhorn",
  "Thunderlord": "thunderlord",
  "Tichondrius": "tichondrius",
  "Tol Barad": "tol barad",
  "Tortheldrin": "tortheldrin",
  "Trollbane": "trollbane",
  "Turalyon": "turalyon",
  "Twisting Nether": "twisting nether",
  "Uldaman": "uldaman",
  "Uldum": "uldum",
  "Undermine": "undermine",
  "Ursin": "ursin",
  "Uther": "uther",
  "Vashj": "vashj",
  "Vek'nilash": "vek'nilash",
  "Velen": "velen",
  "Warsong": "warsong",
  "Whisperwind": "whisperwind",
  "Wildhammer": "wildhammer",
  "Wildrunner": "windrunner",
  "Winterhoof": "winterhoof",
  "Wyrmrest Accord": "wyrmrest accord",
  "Ysera": "ysera",
  "Ysondre": "ysondre",
  "Zangarmarsh": "zangarmarsh",
  "Zul'jin": "zul'jin",
  "Zuluhead": "zuluhead",
}

// {
//    id: 110241996,
//    item: {id: 172092, context: ""},
//    buyout: 47,
//    quantity: 21000,
//    time left: VERY_LONG
// }

var items []string
var id int

type Realm struct {
  RealmName string `form:"realm-name"`
}

var realmStruct Realm

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
  // render := multitemplate.NewRenderer()
  // render.AddFromFiles("home", "templates/index.gohtml", "templates/nav.gohtml")

  r := gin.Default()

  // r.HTMLRender = render

  //server := gin.New()
  //
  //server.LoadHTMLGlob("templates/*.gohtml")

  r.LoadHTMLGlob("templates/*.gohtml")

  r.SetFuncMap(template.FuncMap{
    "getItems": getItems,
  })

  r.Static("/img", "templates/img/")
  r.Static("/css", "templates/css/")
  r.Static("/js", "templates/js/")

  r.GET("/", func(c *gin.Context) {
	  c.HTML(http.StatusOK, "index.gohtml", gin.H{
	    "title": "Wow-Assistant",
    })
  })

  r.GET("/realm", func(c *gin.Context) {
    c.HTML(http.StatusOK, "auctionlist.gohtml", gin.H{
      "title": "Realm List",
      "realms": realms,
    })
  })

  r.POST("/realm", func(c *gin.Context) {
    err := c.Request.ParseForm()
    if err != nil {
      log.Fatal(err)
    }

    realmStruct.RealmName = c.Request.FormValue("realm")

    fmt.Println(realmStruct.RealmName)

    // realmStruct.RealmName = realm

    c.Redirect(301, "http://localhost:8080/auction")
  })

  r.GET("/auction", func(c *gin.Context) {
    auctionHouseList := Wow.WowAuctions(realmStruct.RealmName)

    fmt.Println(items)

    fmt.Println(auctionHouseList.Auctions[0])

    c.HTML(http.StatusOK, "auction.gohtml", gin.H{
      "title": "Auction House",
      "auctions": auctionHouseList.Auctions,
      "items": items,
      "getItems": getItems,
    })

    //c.JSON(http.StatusOK, gin.H{
    //  "code": http.StatusOK,
    //  "id": string(ids),
    //  })
  })

	apiRoutes := r.Group("/api")
	{
		apiRoutes.GET("users", Controllers.GetUsers)
    apiRoutes.POST("user", Controllers.CreateUser)
    apiRoutes.GET("user/:id", Controllers.GetUserByID)
    apiRoutes.PUT("user/:id", Controllers.UpdateUser)
    apiRoutes.DELETE("user/:id", Controllers.DeleteUser)
	}

	// viewRoutes := server.Group("/view")
	// {
	  // viewRoutes.GET("/", )
	  // viewRoutes.GET("/podcasts")
  // }

	return r
}

func getItems(id int) []string {
  var wg sync.WaitGroup

  defer wg.Done()

  item := Wow.GetItem(id)

  fmt.Println(item)

  items = append(items, item.Name)

  return items
}
