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
