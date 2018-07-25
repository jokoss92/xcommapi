package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"xcommapi/model"

	_ "github.com/go-sql-driver/mysql"
	gorp "gopkg.in/gorp.v1"
)

func InitDb() *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("mysql", "root:admin12345@tcp(127.0.0.1:3306)/db_company")
	CheckErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(model.Department{}, "tbl_m_department").SetKeys(true, "DepartmentID")
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "xcommapi:", log.Lmicroseconds))
	return dbmap
}

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
		fmt.Println("Koneksi ke DB Company gagal.")
	} else {
		fmt.Println("Koneksi ke DB Company sukses.")
	}
}
