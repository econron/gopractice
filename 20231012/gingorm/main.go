package main

import(
	"gorm.io/gorm"
	"gingorm/connector"
	"github.com/gin-gonic/gin"
)

type Airport struct {
	gorm.Model
	Name string
	Place string
}

type AirportPostRequest struct {
	Name string `json:"name" binding:"required"`
	Place string `json:"place" binding:"required"`
}

type AirportPutRequest struct {
	ID uint `json:"id" binding:"required"`
	Name string `json:"name"`
	Place string `json:"place"`
}

func setupRouter() *gin.Engine {
	db := connector.Connect()

	r := gin.Default()
	r.GET("/airport/:id", func(c *gin.Context){
		var airport Airport
		result := db.First(&airport, c.Param("id"))
		if result.RowsAffected == 0 {
			c.JSON(404, gin.H{"status": "not found"})
			return
		}
		c.JSON(200, gin.H{"status": "ok", "airport-name": airport.Name, "airport-place": airport.Place})
	})

	r.POST("/airport/create", func(c *gin.Context){
		var airportRequest AirportPostRequest
		if err := c.ShouldBind(&airportRequest); err != nil {
			c.JSON(400, gin.H{"status": "bad request"})
			return
		}
		airport := Airport{Name: airportRequest.Name, Place: airportRequest.Place}
		result := db.Create(&airport)
		if result.RowsAffected == 0 {
			c.JSON(500, gin.H{"status": "Something went wrong"})
			return
		}
		c.JSON(200, gin.H{"status": "ok", "airport-name": airport.Name, "airport-place": airport.Place})
	})

	r.PUT("/airport/update", func(c *gin.Context){
		var airportRequest AirportPutRequest
		if err := c.ShouldBind(&airportRequest); err != nil {
			c.JSON(400, gin.H{"status": "bad request"})
			return
		}
		var airport Airport
		result := db.First(&airport, airportRequest.ID)
		if result.RowsAffected == 0 {
			c.JSON(404, gin.H{"status": "not found"})
			return
		}
		if airportRequest.Name != "" {
			airport.Name = airportRequest.Name
		}
		if airportRequest.Place != "" {
			airport.Place = airportRequest.Place
		}
		if err := db.Save(&airport).Error; err != nil {
			c.JSON(500, gin.H{"status": "something went wrong"})
			return
		}
		c.JSON(200, gin.H{"status": "ok", "airport-name": airport.Name, "airport-place": airport.Place})
	})

	r.DELETE("/airport/delete/:id", func(c *gin.Context){
		var airport Airport
		result := db.Delete(&airport, c.Param("id"))
		if result.RowsAffected == 0 {
			c.JSON(404, gin.H{"status": "not found"})
			return
		}
		if result.Error != nil {
			c.JSON(500, gin.H{"status": "something went wrong"})
			return
		}
		c.JSON(200, gin.H{"status": "ok"})
	})
	return r
}

func main(){
	r := setupRouter()

	r.Run(":8080")
}

