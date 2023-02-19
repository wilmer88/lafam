package controllers

import (
	"errors"
	"log"

	// "gorm-test/database"
	// "gorm-test/models"

	"net/http"
	"strconv"
	
	"github.com/gin-gonic/gin"
	"github.com/wilmer88/lafam/database"
	"github.com/wilmer88/lafam/models"

	_ "github.com/heroku/x/hmetrics/onload"
	"gorm.io/gorm"
)

type Familia struct {
	Db *gorm.DB
}

func New() *Familia {
	db := database.InitDb()
if err := db.AutoMigrate(&models.Fammember{}); err != nil {
	log.Fatalf("error acured during database migrate: %s", err.Error())
}
	return &Familia{Db: db}
}

//create user
func (repository *Familia) CreateUser(c *gin.Context) {
	var member models.Fammember
	if err := c.ShouldBindJSON(&member); err != nil {
		log.Fatalf("error when ShuldBindJson: %s", err.Error())
	}
	err := models.CreateUser(repository.Db, &member)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, member)
	// return member nil,

}

//get users
func (tabla *Familia) GetUsers(c *gin.Context) {
	var member []models.Fammember
	err := models.GetUsers(tabla.Db, &member)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, member)
	// c.HTML(http.StatusOK, "index.tmpl.html", gin.H{"n": i, "prime": p, "likes": likes[n] })
}

//get user by id
func (repository *Familia) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var member models.Fammember
	err := models.GetUser(repository.Db, &member, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, member)
}

// update user
func (repository *Familia) UpdateUser(c *gin.Context) {
	var member models.Fammember
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.GetUser(repository.Db, &member, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}


	if err := c.BindJSON(&member); err != nil {
		 log.Fatalf("c.BindJSON was not able to happen: %s", err.Error() )
	}
	err = models.UpdateUser(repository.Db, &member)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, member)
}

// delete user
func (repository *Familia) DeleteUser(c *gin.Context) {
	var member models.Fammember
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteUser(repository.Db, &member, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}