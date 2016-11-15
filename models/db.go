package models

import "gopkg.in/pg.v4"

var Db *pg.DB

func init() {
	Db = pg.Connect(&pg.Options{
		Addr:     "127.0.0.1:5432",
		User:     "postgres",
		Password: "postgres",
		Database: "beegodb",
		SSL:      false,
	})
}
