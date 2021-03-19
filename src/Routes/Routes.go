package Routes

import (
  "WoW-Assistant/src/Controllers"
  "context"
  "fmt"
  "github.com/FuzzyStatic/blizzard/v2"
  "github.com/FuzzyStatic/blizzard/v2/wowgd"
  "github.com/gin-gonic/gin"
  "log"

  "html/template"
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
  "Blackwater Raiders": "blackwater raiders",
  "Blackwing Lair": "blackwing lair",
  "Blade's Edge": "blade's edge",
  "Bleeding Hollow": "bleeding hollow",
  "Blood Furnace": "blood furnace",
  "bloodhoof">Bloodhoof</option>
  "bloodscalp">Bloodscalp</option>
  "bonechewer">Bonechewer</option>
  "borean tundra">Borean Tundra</option>
  "boulderfist">Boulderfist</option>
  "bronzebeard">Bronzebeard</option>
  "burning blade">Burning Blade</option>
  "burning legion">Burning Legion</option>
  "cairne">Cairne</option>
  "cenarion circle">Cenarion Circle</option>
  "cenarius">Cenarius</option>
  "cho'gall">Cho'gall</option>
  "chromaggus">Chromaggus</option>
  "coilfang">Coilfang</option>
  "crushridge">Crushridge</option>
  "daggerspine">Daggerspine</option>
  "dalaran">Dalaran</option>
  "dalvengyr">Dalvengyr</option>
  "dark iron">Dark Iron</option>
  "darkspear">Darkspear</option>
  "darrowmere">Darrowmere</option>
  "dawnbringer">Dawnbringer</option>
  "deathwing">Deathwing</option>
  "demon soul">Demon Soul</option>
  "dentarg">Dentarg</option>
  "destromath">Destromath</option>
  "dethecus">Dethecus</option>
  "detheroc">Detheroc</option>
  "doomhammer">Doomhammer</option>
  "draenor">Draenor</option>
  "dragonblight">Dragonblight</option>
  "dragonmaw">Dragonmaw</option>
  "drak'tharon">Drak'Tharon</option>
  "drak'thul">Drak'thul</option>
  "draka">Draka</option>
  "drenden">Drenden</option>
  "dunemaul">Dunemaul</option>
  "durotan">Durotan</option>
  "duskwood">Duskwood</option>
  "earthen ring">Dethecus</option>
  "earthen ring">Earthen Ring</option>
  "echo isles">Echo Isles</option>
  "eitrigg">Eitrigg</option>
  "eldre'thalas">Eldre'Thalas</option>
  "elune">Elune</option>
  "emerald dream">Emerald Dream</option>
  "eonar">Eonar</option>
  "eredar">Eredar</option>
  "executus">Executus</option>
  "exodar">Exodar</option>
  "farstriders">Farstriders</option>
  "feathermoon">Feathermoon</option>
  "fenris">Fenris</option>
  "firetree">Firetree</option>
  "fizzcrank">Fizzcrank</option>
  "frostmane">Frostmane</option>
  "frostwolf">Frostwolf</option>
  "galakrond">Galakrond</option>
  "garithos">Garithos</option>
  "garona">Garona</option>
  "garrosh">Garrosh</option>
  "ghostlands">Ghostlands</option>
  "gilneas">Gilneas</option>
  "gnomeregan">Gnomeregan</option>
  "frostwolf">Frostwolf</option>
  "gorefiend">Gorefiend</option>
  "gorgonnash">Gorgonnash</option>
  "greymane">Greymane</option>
  "grizzly hills">Grizzly Hills</option>
  "gul'dan">Gul'dan</option>
  "gurubashi">Gurubashi</option>
  "hakkar">Hakkar</option>
  "haomarush">Haomarush</option>
  "hellscream">Hellscream</option>
  "hydraxis">Hydraxis</option>
  "hyjal">Hyjal</option>
  "icecrown">Icecrown</option>
  "illidan">Illidan</option>
  "jaedenar">Jaedenar</option>
  "kael'thas">Kael'thas</option>
  "kalecgos">Kalecgos</option>
  "kargath">Kargath</option>
  "kel'thuzad">Kel'Thuzad</option>
  "khadgar">Khadgar</option>
  "khaz modan">Khaz Modan</option>
  "kil'jaeden">Kil'jaeden</option>
  "Kilrogg">kilrogg</option>
  "korgath">Korgath</option>
  "korialstrasz">Korialstrasz</option>
  "kul tiras">Kul Tiras</option>
  "laughing skull">Laughing Skull</option>
  "lethon">Lethon</option>
  "lightbringer">Lightbringer</option>
  "lightning's blade">Lightning's Blade</option>
  "lightninghoof">Lightninghoof</option>
  "llane">Llane</option>
  "lothar">Lothar</option>
  "madoran">Madoran</option>
  "maelstrom">Maelstrom</option>
  "magtheridon">Magtheridon</option>
  "maiev">Maiev</option>
  "mal'ganis">Mal'Ganis</option>
  "malfurion">Malfurion</option>
  "malorne">Malorne</option>
  "malygos">Malygos</option>
  "mannoroth">Mannoroth</option>
  "medivh">Medivh</option>
  "misha">Misha</option>
  "mok'nathal">Mok'Nathal</option>
  "moon guard">Moon Guard</option>
  "moonrunner">Moonrunner</option>
  "mug'thol">Mug'thol</option>
  "muradin">Muradin</option>
  "nathrezim">Nathrezim</option>
  "nazgrel">Nazgrel</option>
  "nazjatar">Nazjatar</option>
  "ner'zhul">Ner'zhul</option>
  "nesingwary">Nesingwary</option>
  "nordrassil">Nordrassil</option>
  "norgannon">Norgannon</option>
  "onyxia">Onyxia</option>
  "perenolde">Perenolde</option>
  "proudmoore">Proudmoore</option>
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
  "Stormscale": "stormscale"
  "Suramar": "suramar",
  "Tanaris": "tanaris",
  "Terokkar": "terokkar",
  "Thaurissan": "thaurissan",
  "The Forgotten Coast": "the forgotten coast",
  "The Scryers": "the scryers"
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
  "Zangarmarsh": "zangarmarsh"
  "Zul'jin": "zul'jin",
  "Zuluhead": "zuluhead"
}

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
  // render := multitemplate.NewRenderer()
  // render.AddFromFiles("home", "templates/index.gohtml", "templates/nav.gohtml")

  r := gin.Default()

  // r.HTMLRender = render

  //server := gin.New()
  //
  //server.LoadHTMLGlob("templates/*.gohtml")

  r.SetFuncMap(template.FuncMap{
    "auction": WowAuctions,
  })

  r.LoadHTMLGlob("templates/*.gohtml")

  r.Static("/img", "templates/img/")
  r.Static("/css", "templates/css/")
  r.Static("/js", "templates/js/")

  r.GET("/", func(c *gin.Context) {
	  c.HTML(http.StatusOK, "index.gohtml", gin.H{
	    "title": "Wow-Assistant",
    })
  })

  r.GET("/auction", func(c *gin.Context) {
    c.HTML(http.StatusOK, "auction.gohtml", gin.H{
      "title": "Auction House",
    })
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

func getClient() *blizzard.Client {
  ctx := context.Background()

  blizz := blizzard.NewClient("d2a18a7c4f364725822fc2ef3740ecc5", "Bv5t7l2AifIWsgESnosKhUbYnlRVFLeG", blizzard.US, blizzard.EnUS)

  err := blizz.AccessTokenRequest(ctx)
  if err != nil {
    fmt.Println(err)
  }

  return blizz
}

// WowAuctions gets all auctions from a realm
func WowAuctions(realmToGet string) *wowgd.AuctionHouse {
  ctx := context.Background()

  client := getClient()

  realm, _, err := client.WoWRealm(ctx, realmToGet)
  if err != nil {
    log.Fatal(err)
  }

  wowAuctions, _, err := client.WoWAuctions(ctx, realm.ID)
  if err != nil {
    log.Fatal(err)
  }

  return wowAuctions
}
