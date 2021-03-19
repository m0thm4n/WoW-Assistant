package Routes

import (
  "Podify/src/Controllers"
  "github.com/gin-gonic/gin"
  // "github.com/gin-contrib/multitemplate"
  "net/http"
  "html/template"
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

  html := template.Must(template.ParseFiles("templates/index.gohtml", "templates/nav.gohtml",
    "templates/auction.gohtml", "templates/footer.gohtml"))
  r.SetHTMLTemplate(html)

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

