package lib

import (
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	resetDatabase(db)
	createFakeData(db)

	// fetch
	// var version Version
	// db.First(&version, "name = ?", "4.17")
	// log.Println("version: ", version)

	// // delete it
	// db.Delete(&version)

	// // get count
	// var count int64
	// db.Model(&Version{}).Where("name = ?", "4.17").Count(&count)
	// log.Println("count: ", count)
}

func resetDatabase(db *gorm.DB) {
	// migrate
	db.AutoMigrate(&Version{})
	db.AutoMigrate(&Release{})
	db.AutoMigrate(&Image{})
	db.AutoMigrate(&Catalog{})
	db.AutoMigrate(&Operator{})

	// delete all
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Version{})
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Release{})
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Image{})
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Catalog{})
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Operator{})
}

func createFakeData(db *gorm.DB) {
	db.Create(&Version{Name: "4.17"})

	// Versions
	var version Version
	db.First(&version, "name = ?", "4.17")

	// Releases
	db.Create(&Release{Name: "4.17.0", VersionID: version.ID, PullSpec: "quay.io/openshift-release-dev/ocp-release@sha256:fbad931c725b2e5b937b295b58345334322bdabb0b67da1c800a53686d7397da"})
	db.Create(&Release{Name: "4.17.1", VersionID: version.ID, PullSpec: "quay.io/openshift-release-dev/ocp-release@sha256:e16ac60ac6971e5b6f89c1d818f5ae711c0d63ad6a6a26ffe795c738e8cc4dde"})

	// Catalogs
	db.Create(&Catalog{Name: "registry.redhat.io/redhat/redhat-operator-index:v4.17", VersionID: version.ID})
	db.Create(&Catalog{Name: "registry.redhat.io/redhat/certified-operator-index:v4.17", VersionID: version.ID})
	db.Create(&Catalog{Name: "registry.redhat.io/redhat/community-operator-index:v4.17", VersionID: version.ID})
	db.Create(&Catalog{Name: "registry.redhat.io/redhat/redhat-marketplace-index:v4.17", VersionID: version.ID})
}
