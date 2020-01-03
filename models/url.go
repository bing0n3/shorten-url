package models

import (
	"database/sql"
	db "github.com/bing0n3/shorten-url/stores/mysql"
	"github.com/bing0n3/shorten-url/utils"
	"github.com/pkg/errors"
	"time"
)

type URL struct {
	ID 			int 		`db:"id"`
	Alias       int64
	OriginalUrl string       `db:"original_url"`
	CreatedAt   time.Time    `db:"created_at"`
	ExpireAt    sql.NullTime `db:"expire_at"`
	UpdatedAt   time.Time    `db:"updated_at"`
}

var ErrNoEntryFound = errors.New("no entry found with this ID")

func GetURLByAlias(alias int64) (*URL, error) {
	if alias < 0 {
		return nil, errors.New("Alias is negative")
	}
	var url URL

	err := db.GetMySQLClient().Get(&url, "SELECT * FROM url WHERE alias=?", alias)
	if err != nil {
		return nil, errors.Wrap(err, "Read Failed from URL table")
	}
	return &url, nil
}

func InsertAlias(alias int64, original_url string, expire string) error{
	var exp_at sql.NullTime
	now := time.Now().UTC()

	if expire == "" {
		exp_at = sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		}
	} else {
		dur, err_dur := time.ParseDuration(expire)
		if err_dur != nil {
			utils.Warning.Println(errors.Wrap(err_dur, "Parse Duration Failed"))
		}
		exp_at = sql.NullTime{
			Time:  now.Add(dur),
			Valid: true,
		}
	}

	url := &URL{0, alias, original_url, now, exp_at, now}
	tx := db.GetMySQLClient().MustBegin()
	_, err := tx.NamedExec("INSERT INTO url (alias, original_url,created_at,expire_at, updated_at) VALUES (:alias, :original_url,:created_at, :expire_at,:updated_at)", url)
	if err != nil {
		return errors.Wrapf(err,"Failed Inserting url: %s, alias: %d", original_url, alias)
	}
	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err,"Failed Commit Transition for inserting url")
	}
	return nil
}

func GetMaxAlias() (int64, error) {
	var res sql.NullInt64

	err := db.GetMySQLClient().QueryRow("SELECT MAX(alias) FROM url").Scan(&res)

	if err != nil {
		return 0, errors.Wrap(err, "Can't get max id from dataset")
	}

	if res.Valid {
		return res.Int64, nil
	}
	return 0, errors.New("No entity in the dataset")
}