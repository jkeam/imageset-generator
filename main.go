package main

import (
	"flag"
	"log"
	"log/slog"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
  "gorm.io/driver/postgres"
  "github.com/jkeam/imageset-generator/lib"
)

var db *gorm.DB

func main() {
	var err error
	dsn := "host=localhost user=postgres password=adminpassword dbname=imageset port=5432 sslmode=disable TimeZone=America/New_York"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
		log.Fatal(err)
    panic("failed to connect database")
  }

  migrateFlag := flag.Bool("migrate", false, "migrate boolean")
  flag.Parse()
  if *migrateFlag {
  	lib.Migrate(db)
    slog.Info("Database migrated")
  	return
  }

	router := gin.Default()
	router.GET("/versions", getVersions)
	router.GET("/releases", getReleases)
	router.Run("localhost:8000")
}

func baseGet(c *gin.Context, fn lib.RouterHandler) {
	c.Header("Content-Type", "application/json")
	fn(c, db)
}

func getVersions(c *gin.Context) {
	slog.Info("getVersions")
	baseGet(c, lib.GetVersions)
}

func getReleases(c *gin.Context) {
	slog.Info("getReleases")
	baseGet(c, lib.GetReleases)
}
