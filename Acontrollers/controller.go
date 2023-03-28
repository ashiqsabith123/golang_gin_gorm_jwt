package acontrollers

import (
	"fmt"
	"log"
	"net/http"

	"golang_gin_gorm_jwt/bycrpt"
	postgres "golang_gin_gorm_jwt/connectDb"
	"golang_gin_gorm_jwt/helpers"
	"golang_gin_gorm_jwt/models"

	"github.com/gin-gonic/gin"
)

type Students struct {
	Id       string
	Fname    string
	Lname    string
	Email    string
	Phone    string
	Place    string
	Dob      string
	Username string
	Password string
	Dep_id   string
}

type Departments struct {
	Id       string
	Dep_name string
	Hod_name string
	Dep_id   string
}

type Admin struct {
	Username string
	Password string
}

func Home(c *gin.Context) {
	ok := AdminLogged(c)

	if !ok {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	db := postgres.ConnectDb()

	var details []Students
	db.Find(&details)

	var dep []Departments
	db.Find(&dep)

	c.HTML(http.StatusOK, "adminpanel.html", gin.H{
		"details":    details,
		"department": dep,
	})

}

func AdminLogin(c *gin.Context) {

	var admin []Admin
	var status bool

	entered_admin := models.Admin{
		Username: c.Request.FormValue("admin_username"),
		Password: c.Request.FormValue("admin_password"),
	}

	fmt.Println(entered_admin)

	db := postgres.ConnectDb()
	db.Find(&admin)

	for _, i := range admin {

		if i.Username == entered_admin.Username {
			validPassword := bycrpt.VerifyPassword(i.Password, entered_admin.Password)
			if validPassword {
				status = true
				break
			}

		}
	}

	if !status {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"Erro": "Invalid username or password",
		})
	}

	token := helpers.GenerateTokens(entered_admin.Username, "Admin")

	c.SetCookie("admin", token, 600000, "/", "localhost", true, true)

	c.Redirect(http.StatusSeeOther, "/admin")

}

func Delete(c *gin.Context) {

	var st Students
	var dt Departments

	id := c.Param("Id")
	table := c.Param("Table")

	db := postgres.ConnectDb()

	if table == "stu" {
		db.Where("Id=?", id).Delete(&st)
	} else if table == "dep" {
		db.Where("Id=?", id).Delete(&dt)
	}

	typ := "Student"
	stat := "Deleted"
	path := "/admin"

	c.HTML(http.StatusOK, "succesfull.html", gin.H{
		"value": typ,
		"path":  path,
		"stat":  stat,
	})
}

func AddDepartment(c *gin.Context) {
	department := models.Departments{
		Dep_name: c.PostForm("department"),
		Hod_name: c.PostForm("hod_name"),
		Dep_id:   c.PostForm("dep_id"),
	}

	db := postgres.ConnectDb()
	err := db.AutoMigrate(&models.Departments{})

	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	typ := "Department"
	db.Create(&department)
	path := "/admin"
	c.HTML(http.StatusOK, "succesfull.html", gin.H{
		"value": typ,
		"path":  path,
	})

}

func AddAdmin(c *gin.Context) {

	admin := models.Admin{
		Name:     c.PostForm("admin_name"),
		Username: c.PostForm("admin_username"),
		Password: bycrpt.HashPassword(c.PostForm("admin_password")),
	}

	db := postgres.ConnectDb()
	err := db.AutoMigrate(&models.Admin{})
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	typ := "Admin"
	stat := "Added"
	path := "/admin"
	db.Create(&admin)
	c.HTML(http.StatusOK, "succesfull.html", gin.H{
		"value": typ,
		"path":  path,
		"stat":  stat,
	})
}

func AddStudent(c *gin.Context) {
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
	stat := "Added"
	path := "/admin"
	db.Create(&student)
	c.HTML(http.StatusOK, "succesfull.html", gin.H{
		"value": created,
		"path":  path,
		"stat":  stat,
	})
}

func SendAjax(c *gin.Context) {
	var ID int
	var item Students

	if err := c.ShouldBind(&ID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := postgres.ConnectDb()
	db.First(&item, ID)

	c.JSON(http.StatusOK, item)

}

func Search(c *gin.Context) {
	var Data string
	var searchData []Students

	if err := c.ShouldBind(&Data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := postgres.ConnectDb()

	db.Where("fname LIKE ?", fmt.Sprintf("%%%s%%", Data)).Find(&searchData)
	c.JSON(http.StatusOK, searchData)
}

func UpdateStudent(c *gin.Context) {
	values := models.Students{
		Fname:    c.PostForm("edfname"),
		Lname:    c.PostForm("edlname"),
		Email:    c.PostForm("edemail"),
		Phone:    c.PostForm("edphone"),
		Place:    c.PostForm("edplace"),
		Dob:      c.PostForm("eddate"),
		Username: c.PostForm("edusername"),
		Password: bycrpt.HashPassword(c.PostForm("edpassword")),
		Dep_id:   c.PostForm("dep_id"),
	}

	id := c.PostForm("edid")
	fmt.Println(values)

	db := postgres.ConnectDb()
	db.Table("students").Where("id=?", id).Updates(values)

	created := "Student"
	stat := "Updated"
	path := "/admin"

	c.HTML(http.StatusOK, "succesfull.html", gin.H{
		"value": created,
		"path":  path,
		"stat":  stat,
	})

}

func AdminLogged(c *gin.Context) bool {
	var stat bool
	cookie, err := c.Cookie("admin")

	if err != nil {
		return false
	}

	stat = helpers.ValidateTokens(cookie)

	return stat
}

func Logout(c *gin.Context) {
	_, err := c.Request.Cookie("admin")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}
	c.SetCookie("user", "", -1, "/", "localhost", false, false)
	c.Redirect(http.StatusSeeOther, "/login")
}
