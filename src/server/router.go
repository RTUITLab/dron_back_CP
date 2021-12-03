package server

import (
	"github.com/0B1t322/CP-Rosseti-Back/controllers/auth"
	"github.com/0B1t322/CP-Rosseti-Back/controllers/module"
	"github.com/0B1t322/CP-Rosseti-Back/controllers/role"
	"github.com/0B1t322/CP-Rosseti-Back/controllers/user"
	_ "github.com/0B1t322/CP-Rosseti-Back/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type Controllers struct {
	User   *user.UserController
	Role   *role.RoleController
	Auth   *auth.AuthController
	Module *module.ModuleController
}

func NewRouter(c *Controllers) *gin.Engine {
	router := gin.New()

	baseRouter := router.Group("/api/dron")

	baseRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := baseRouter.Group("/v1")
	{
		user := v1.Group("/user").Use(c.Auth.TokenAuthMiddleware())
		{
			user.POST("", c.User.CreateUser)
			user.DELETE("/:id", c.User.DeleteUser)
			user.PUT("/:id", c.User.UpdateUser)
			user.GET("", c.User.GetUsers)
		}

		auth := v1.Group("/auth")
		{
			auth.POST("/login", c.Auth.Login)
			auth.POST("/refresh", c.Auth.Refresh)
		}

		module := v1.Group("/module").Use(c.Auth.TokenAuthMiddleware())
		{
			module.POST("", c.Module.CreateModule)
			module.GET("", c.Module.HTTPGetModules)
			module.GET("/:id", c.Module.HTTPGetModule)
			module.PUT("/:id/dependecy", c.Module.HTTPAddModuleDependecy)
			module.DELETE("/:id/dependecy", c.Module.HTTPDeleteModuleDependecy)
			module.POST("/:id/submodule", c.Module.AddSubModule)
			module.POST("/submodule/:id/test", c.Module.AddSubModuleTest)
			module.POST("/:id/test", c.Module.AddModuleTest)
			module.POST("/submodule/:id/test/theor", c.Module.AddTheorTest)
			module.POST("/submodule/:id/test/pract", c.Module.AddPractTest)
			module.PUT("/submodule/:id/test/pract", c.Module.UpdateConfigToPractTest)
			module.PUT("/:id/test/pract", c.Module.UpdateModuleConfigToPractTest)
			module.PUT("/submodule/:id/test/theor", c.Module.UpdateTheorTest)
			module.PUT("/:id/test/theor", c.Module.UpdateModuleTheorTest)

			module.DELETE("/submodule/:id/test/theor", c.Module.DeleteTheorTest)
			module.DELETE("/submodule/:id/test/pract", c.Module.DeletePractTest)
			module.DELETE("/:id/test/theor", c.Module.DeleteModuleTheorTest)
			module.DELETE("/:id/test/pract", c.Module.DeleteModulePractTest)
			module.GET("/submodule/:id", c.Module.GetSubModule)

			module.POST("/:id/test/theor/try", c.Module.TryTheoreticalTest)
			module.POST("/test/theor/try/:id", c.Module.TryAnswerQuestion)
		}
	}

	return router
}
