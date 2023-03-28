package routes

import (
	admin "golang_gin_gorm_jwt/Acontrollers"
	student "golang_gin_gorm_jwt/Scontrollers"

	"github.com/gin-gonic/gin"
)

func AuthStudentsRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "./static")

	//router.Use(middleware.AuthMiddleWare())
	router.GET("/", student.IndexHandler)

	router.GET("/signup", student.SignUp)
	router.POST("/signup", student.PostSignUp)
	router.GET("/succesfull", student.HandleLogin)

	router.GET("/login", student.Login)
	router.POST("/login", student.PostLogin)
	router.GET("/home", student.Home)
	router.GET("/logout", student.Logout)

}

func AuthAdminRoutes(router *gin.Engine) {

	router.GET("/admin", admin.Home)
	router.GET("/delete/:Id/:Table", admin.Delete)
	router.POST("/adddepartment", admin.AddDepartment)

	router.POST("/addadmin", admin.AddAdmin)
	router.POST("/adminlogin", admin.AdminLogin)
	router.POST("/ajax", admin.SendAjax)
	router.POST("/updatestudent", admin.UpdateStudent)
	router.POST("/addstudent", admin.AddStudent)
	router.POST("/search", admin.Search)
	router.GET("/adminlogout", admin.Logout)

}
