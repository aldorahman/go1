package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go1/config"
	"go1/controller"

	"go1/middleware"
	"go1/repository"
	"go1/service"
)

var (
	db              *gorm.DB                   = config.SetupDatabaseConnection()
	userRepository  repository.UserRepository  = repository.NewUserRepository(db)
	colorRepository repository.ColorRepository = repository.NewColorRepository(db)
	todoRepository  repository.TodoRepository  = repository.NewTodoRepository(db)

	jwtService   service.JWTService   = service.NewJWTService()
	authService  service.AuthService  = service.NewAuthService(userRepository)
	userService  service.UserService  = service.NewUserService(userRepository)
	colorService service.ColorService = service.NewColorService(colorRepository)
	todoService  service.TodoService  = service.NewTodoService(todoRepository)

	authController  controller.AuthController  = controller.NewAuthController(authService, jwtService)
	userController  controller.UserController  = controller.NewUserController(userService, jwtService)
	colorController controller.ColorController = controller.NewColorController(colorService, jwtService)
	todoController  controller.TodoController  = controller.NewTodoController(todoService, jwtService)
)

func Routes(r *gin.Engine) {
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	colorRoutes := r.Group("api/color", middleware.AuthorizeJWT(jwtService))
	{
		colorRoutes.GET("/", colorController.All)
		colorRoutes.POST("/", colorController.Insert)
		colorRoutes.GET("/:id", colorController.FindByID)
		colorRoutes.PUT("/:id", colorController.Update)
		colorRoutes.DELETE("/:id", colorController.Delete)
	}

	todoRoutes := r.Group("api/todo", middleware.AuthorizeJWT(jwtService))
	{
		todoRoutes.GET("/", todoController.All)
		todoRoutes.POST("/", todoController.Insert)
		todoRoutes.GET("/:id", todoController.FindByID)
		todoRoutes.PUT("/:id", todoController.Update)
		todoRoutes.DELETE("/:id", todoController.Delete)
	}
}
