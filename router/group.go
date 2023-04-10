package router

import (
	"digital-gin/controllers"
	"github.com/gin-gonic/gin"
)

// router group
const (
	ManageGroup  = "manage"
	AIGroup      = "ai"
	ArteryGroup  = "artery"
	BigDataGroup = "bigdata"
	ITGroup      = "it"
	SupportGroup = "support"
)

func Groups(r *gin.Engine) {
	var (
		manageGroup  = r.Group(ManageGroup)
		aiGroup      = r.Group(AIGroup)
		arteryGroup  = r.Group(ArteryGroup)
		bigDataGroup = r.Group(BigDataGroup)
		itGroup      = r.Group(ITGroup)
		supportGroup = r.Group(SupportGroup)
	)

	r.GET("/index", controllers.Index)
	manageGroup.GET("/index", controllers.Index)
	//manageGroup.POST("/login", manage.Login)
	//manageGroup.POST("/logout", manage.Logout)
	//manageGroup.POST("/update", manage.Update)

	aiGroup.GET("/index", controllers.Index)
	arteryGroup.GET("/index", controllers.Index)
	bigDataGroup.GET("/index", controllers.Index)
	itGroup.GET("/index", controllers.Index)
	supportGroup.GET("/index", controllers.Index)
}
