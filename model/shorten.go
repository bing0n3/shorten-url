package model

import (
	"errors"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

const (
	SHORTEXIST    = 1
	FAIL          = 2
	WRONGSHORT    = 3
	PARSESCUSSCE  = 4
	SHORTNOTEXIST = 5
)

// Shorten model
type Shorten struct {
	UID         int64  `gorm:"primary_key"`
	ID          int64  `gorm:"type:bigint"`
	Short       string `gorm:"type:varchar(256)"`
	OriginalURL string `gorm:"type:varchar(256)"`
	Custom      bool
	InsertAt    *time.Time `sql:"DEFAULT:current_timestamp"`
	UpdateAt    *time.Time `sql:"DEFAULT:current_timestamp"`
}

func (shorten *Shorten) CreateShort(id int64) {
	encodedURL := Encode(id)
	shorten.Short = encodedURL
}

// get counter number when init the system
func GetCounter() (int64, error) {
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
func AddShortenInst(originalURL string, target string) (string, int) {
	var (
		shortURL string
		id       int64
		now      time.Time
		inst     Shorten
		custom   bool
	)

	if target != "" {
		custom = true
		if db.First(&Shorten{}, "short = ?", target).RecordNotFound() {
			id, err := Decode(target)
			if err != nil {
				log.Println("Decode Failed")
				return "", WRONGSHORT
			}
			log.Printf("Decode %s to %d", target, id)
			now = time.Now()
			inst = Shorten{ID: id, OriginalURL: originalURL, Short: target, Custom: custom, UpdateAt: &now, InsertAt: &now}
			db.Create(&inst)
			return target, PARSESCUSSCE
		} else {
			log.Printf("Short Code %s Exist", target)
			return "", SHORTEXIST
		}
	} else {
		custom = false
		id = lastID.UpdateCounter()
		now = time.Now()
		inst = Shorten{ID: id, OriginalURL: originalURL, Custom: custom, UpdateAt: &now, InsertAt: &now}
		inst.CreateShort(id)
		log.Printf("Insert id: %d", id)
		if db.First(&Shorten{}, "id= ?", id).RecordNotFound() {
			db.Create(&inst)
			shortURL = inst.Short
		} else {
			AddShortenInst(originalURL, target)
		}
		return shortURL, PARSESCUSSCE
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
func FindShortenByShort(short string) (string, error) {
	var inst Shorten
	if db.Where("short = ?", short).First(&inst).RecordNotFound() {
		return "", errors.New("Cannot find in the Databse by the short Url")
	} else {
		return inst.OriginalURL, nil
	}
}
