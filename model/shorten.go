package model

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

// Shorten model
type Shorten struct {
	ID          int    `gorm:"primary_key"`
	Short       string `gorm:"type:varchar(256)"`
	OriginalURL string `gorm:"type:varchar(256)"`
	Custom      bool
	InsertAt    *time.Time `sql:"DEFAULT:current_timestamp"`
	UpdateAt    *time.Time `sql:"DEFAULT:current_timestamp"`
}

func (shorten *Shorten) CreateShort(id int) {
	encodedURL := Encode(id)
	shorten.Short = encodedURL
}

// get counter number when init the system
func GetCounter() (int, error) {
	var inst Shorten
	err := db.Where("custom = ?", "false").Last(&inst).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return 9999, nil
	} else if err != nil && err != gorm.ErrRecordNotFound {
		return -1, err
	}
	return inst.ID, nil
}

// add url
func AddShortenInst(originalURL string, custom bool) {
	id := lastID.UpdateCounter()
	now := time.Now()

	inst := Shorten{ID: id, OriginalURL: originalURL, Custom: custom, UpdateAt: &now, InsertAt: &now}
	inst.CreateShort(id)
	log.Printf("Insert id: %d", id)
	if db.First(&Shorten{}, "id= ?", id).RecordNotFound() {
		db.Create(&inst)
	} else {
		AddShortenInst(originalURL, custom)
	}
}

// CheckShortenByOriginalURL check url exist in db or not func
// if exist return short url and true, if noe return empty string and false
func CheckShortenByOriginalURL(originalURL string) (string, bool) {
	var inst Shorten
	err := db.Where("original_url = ?", originalURL).First(&inst).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return "", false
	} else if err != nil {
		log.Printf(err.Error())
		return "", false
	}

	return inst.Short, true
}

// CheckShortenByShort checks shorten exist or not by attribute short
// return: boolean, exist return true, not return false
func CheckShortenByShort(short string) bool {
	var inst Shorten
	err := db.Where("short = ?", short).First(&inst).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		log.Printf("Don't find Record by short = \"%s\"\n", short)
		return false
	} else if err != nil {
		log.Printf(err.Error())
		return false
	}
	return true
}
