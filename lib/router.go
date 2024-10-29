package lib

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type RouterHandler func(c *gin.Context, db *gorm.DB)

func GetVersions(c *gin.Context, db *gorm.DB) {
	var versions []Version
	result := db.Find(&versions)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, []Version{})
	} else {
		c.IndentedJSON(http.StatusOK, versions)
	}
}

func GetReleases(c *gin.Context, db *gorm.DB) {
	var releases []Release
	result := db.Find(&releases)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, []Release{})
	} else {
		c.IndentedJSON(http.StatusOK, releases)
	}
}
