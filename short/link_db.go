package short

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

var schema = `
CREATE TABLE links (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  url varchar(255) NOT NULL,
  short varchar(255) NOT NULL,
  type varchar(10) NOT NULL,
  insert_at timestamp NULL DEFAULT NULL,
  update_at timestamp NULL DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY id (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
`

type Link struct {
	ID        int64     `db:"id"`
	Url       string    `db:"url"`
	Short     string    `db:"short"`
	custom    string    `db:"type"`
	INSERT_AT time.Time `db:"insert_at"`
	UPDATE_AT time.Time `db:"update_at"`
}

type LinkDB struct {
	DriverName     string
	DataSourceName string
	MaxOpenConns   int
	MaxIdleConns   int
	ADB            *sqlx.DB
}

type LinkTX struct {
	TX *sqlx.Tx
}

func InitMySQLPool(host, database, user, password, charset string, maxOpenConns, maxIdleConns int) *LinkDB {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&autocommit=true", user, password, host, database, charset)
	db := &LinkDB{
		DriverName:     "mysql",
		DataSourceName: dataSourceName,
		MaxOpenConns:   maxOpenConns,
		MaxIdleConns:   maxIdleConns,
	}
	if err := db.Open(); err != nil {
		log.Panicln("Init mysql pool failed.", err.Error())
	}
	return db
}

func (db *LinkDB) Open() error {
	var err error
	db.ADB, err = sqlx.Open(db.DriverName, db.DataSourceName)
	if err != nil {
		return err
	}
	db.ADB.SetMaxIdleConns(db.MaxIdleConns)
	db.ADB.SetMaxOpenConns(db.MaxOpenConns)
	return nil
}

func (db *LinkDB) CreateTable() {
	db.ADB.MustExec(schema)
}

func (db *LinkDB) GetByUrl(link *Link, url string) error {
	querySql := "SELECT * FROM links WHERE url=$1"
	return db.ADB.Get(link, querySql, url)
}

func (db *LinkDB) GetByShort(link *Link, short string) error {
	querySql := "SELECT * FROM links WHERE short=$1"
	return db.ADB.Get(link, querySql, short)
}

func (db *LinkDB) Begin() (*LinkTX, error) {
	var oneLinkTX = &LinkTX{}
	var err error
	if pingErr := db.ADB.Ping(); pingErr != nil {
		oneLinkTX.TX, err = db.ADB.Beginx()
	}
	return oneLinkTX, err
}

func (tx *LinkTX) Commit() error {
	return tx.TX.Commit()
}

func (tx *LinkTX) InsertLink(link Link) (int64, error) {
	insertSql := "INSERT INTO links (url, short, type,insert_at,update_at) VALUES (:url, :short, :type, :insert_at, :update_at)"
	result, err := tx.TX.NamedExec(insertSql, link)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
