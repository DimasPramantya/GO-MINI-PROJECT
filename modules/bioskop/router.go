package bioskop

import (
	"github.com/gin-gonic/gin"

	"formative-14/database/connection"
)

func Initiator(router *gin.Engine) {
	api := router.Group("/api/bioskop")
	{
		api.POST("", CreateBioskopRouter)
		api.GET("", GetAllBioskopRouter)
		api.GET("/:id", GetBioskopByIdRouter)
		api.DELETE("/:id", HardDeleteBioskopRouter)
		api.PUT("/:id", UpdateBioskopRouter)
	}
}

func CreateBioskopRouter(c *gin.Context) {
	var (
		bioskopRepo    = NewRepository(connection.DBConnections)
		bioskopService = NewService(bioskopRepo)
	)

	result, err := bioskopService.CreateBioskop(c)
	if err != nil {
		return
	}

	c.JSON(201, gin.H{
		"message": "Bioskop created successfully",
		"data":    result,
	})
}

func GetAllBioskopRouter(c *gin.Context) {
	var (
		bioskopRepo    = NewRepository(connection.DBConnections)
		bioskopService = NewService(bioskopRepo)
	)

	result, err := bioskopService.GetAllBioskop(c)
	if err != nil {
		return
	}

	c.JSON(200, gin.H{
		"message": "Bioskops retrieved successfully",
		"data":    result,
	})
}

func GetBioskopByIdRouter(c *gin.Context) {
	var (
		bioskopRepo    = NewRepository(connection.DBConnections)
		bioskopService = NewService(bioskopRepo)
	)

	result, err := bioskopService.GetBioskopById(c)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"message": "Bioskop retrieved successfully",
		"data":    result,
	})
}

func HardDeleteBioskopRouter(c *gin.Context) {
	var (
		bioskopRepo    = NewRepository(connection.DBConnections)
		bioskopService = NewService(bioskopRepo)
	)

	err := bioskopService.HardDeleteBioskop(c)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"message": "Bioskop deleted successfully",
	})
}

func UpdateBioskopRouter(c *gin.Context) {
	var (
		bioskopRepo    = NewRepository(connection.DBConnections)
		bioskopService = NewService(bioskopRepo)
	)

	result, err := bioskopService.UpdateBioskop(c)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"message": "Bioskop updated successfully",
		"data":    result,
	})
}