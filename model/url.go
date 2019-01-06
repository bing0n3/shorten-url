package model

import "time"

// URL model
type URL struct {
	ID       int        `gorm:"primary_key"`
	Short    string     `gorm:"type:varchar(256)"`
	Type     string     `gorm:"type:varchar(256)"`
	InsertAt *time.Time `sql:"DEFAULT:current_timestamp"`
	UpdateAt *time.Time `sql:"DEFAULT:current_timestamp"`
}

func (*URL) CreateShort(url string) {

}
