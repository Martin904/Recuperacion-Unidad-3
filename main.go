package main

import (
	"fmt"
						//PAQUETES QUE SE UTILIZARON//
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type propiedades struct {
	id       int   `json:"id"`
	precio string `json:"precio"`
	descripcion   string `json:"descripcion"`
	propietario    string `json:"propetario"`
	vendedor		string `json:"vendedor"`
	creado			string`json:"creado"`
}
func main() {

	
	db, err = gorm.Open("mysql", "root:martino@/bienes_raices?charset=utf8&parseTime=True&loc=Local")
				//"mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")//
	

	if err != nil {
	tcp	fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&propiedades{})

	r := gin.Default()
	r.GET("/propiedades/", MostrarPropiedades)
	r.GET("/propiedades/:id", MostrarUnaPropiedad)
	r.POST("/propiedades", AgregarPropiedades)
	r.PUT("/propiedades/:id", ActualizarPropiedad)
	r.DELETE("/propiedades/:id", BorrarPropiedad)

	r.Run(":8080")
}

func BorrarPropiedad(c *gin.Context) {
	id := c.Params.ByName("id")
	var propiedades Propiedades
	d := db.Where("id = ?", id).Delete(&propiedades)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "Borrado"})
}

func ActualizarPropiedad(c *gin.Context) {

	var propiedades Propiedades
	id := c.Params.ByName("id")

	if err := db.Where("id = ?", id).First(&propiedades).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&propiedades)

	db.Save(&propiedades)
	c.JSON(200, propiedades)

}

func AgregarPropiedades(c *gin.Context) {

	var propiedades Propiedades
	c.BindJSON(&propiedades)

	db.Create(&propiedades)
	c.JSON(200, propiedades)
}

func MostrarUnaPropiedad(c *gin.Context) {
	id := c.Params.ByName("id")
	var propiedades Propiedades
	if err := db.Where("id = ?", id).First(&propiedades).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, propiedades)
	}
}
func MostrarPropiedades(c *gin.Context) {
	var propiedades []Propiedades
	if err := db.Find(&propiedades).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, propiedades)
	}

}
