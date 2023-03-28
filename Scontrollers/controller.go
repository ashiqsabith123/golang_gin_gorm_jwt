package controllers

import (
	"fmt"
	acontrollers "golang_gin_gorm_jwt/Acontrollers"
	"golang_gin_gorm_jwt/bycrpt"
	postgres "golang_gin_gorm_jwt/connectDb"
	"golang_gin_gorm_jwt/helpers"
	"golang_gin_gorm_jwt/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Sigins struct {
	Username string
	Password string
}

func Login(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	ok := UserLogged(c)

	if !ok {
		c.HTML(http.StatusOK, "index.html", nil)
		return
	}
	c.Redirect(303, "/home")

}

func SignUp(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.HTML(http.StatusOK, "signup.html", nil)
}

func IndexHandler(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	ok := UserLogged(c)

	stat := acontrollers.AdminLogged(c)

	if ok {
		c.Redirect(http.StatusSeeOther, "/home")
	} else if stat {
		c.Redirect(http.StatusSeeOther, "/admin")
	} else {
		c.Redirect(http.StatusSeeOther, "/login")
	}

}

func PostLogin(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	var students []Sigins
	var status bool

	sign := models.Signin{
		Username: c.Request.FormValue("username"),
		Password: c.Request.FormValue("password"),
	}

	db := postgres.ConnectDb()
	db.Find(&students)

	fmt.Println(sign.Password)

	for _, i := range students {

		if i.Username == sign.Username {
			validPassword := bycrpt.VerifyPassword(i.Password, sign.Password)
			if validPassword {
				status = true
				break
			}

		}
	}

	if !status {

		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"Error": "Invalid username or password",
		})
		return
	}

	token := helpers.GenerateTokens(sign.Username, "User")

	c.SetCookie("user", token, 600000, "/", "localhost", true, true)

	c.Redirect(http.StatusSeeOther, "/home")

}

func PostSignUp(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	student := models.Students{
		Fname:    c.Request.FormValue("fname"),
		Lname:    c.Request.FormValue("lname"),
		Email:    c.Request.FormValue("email"),
		Phone:    c.Request.FormValue("phone"),
		Place:    c.Request.FormValue("place"),
		Dob:      c.Request.FormValue("date"),
		Username: c.Request.FormValue("username"),
		Password: bycrpt.HashPassword(c.Request.FormValue("password")),
		Dep_id:   c.Request.FormValue("dep_id"),
	}

	//fmt.Println(student)
	db := postgres.ConnectDb()
	err := db.AutoMigrate(&models.Students{})

	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	created := "Student"
	path := "/login"
	db.Create(&student)
	c.HTML(http.StatusOK, "succesfull.html", gin.H{
		"value": created,
		"path":  path,
	})
}

func HandleLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "message.html", gin.H{
		"message": "Sign successful!",
		"script":  "<script>alert('Login successful!');</script>",
	})

}

func Home(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	ok := UserLogged(c)

	if !ok {
		c.Redirect(303, "/login")
	}
	c.HTML(http.StatusOK, "home.html", nil)
}

func UserLogged(c *gin.Context) bool {
	var stat bool
	cookie, err := c.Cookie("user")

	if err != nil {
		return false
	}

	stat = helpers.ValidateTokens(cookie)

	return stat
}

func Logout(c *gin.Context) {

	_, err := c.Request.Cookie("user")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}
	c.SetCookie("user", "", -1, "/", "localhost", false, false)
	c.Redirect(http.StatusSeeOther, "/login")
}
