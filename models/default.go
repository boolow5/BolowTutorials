package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/mattn/go-sqlite3"
)

var o orm.Ormer

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "data.db")
	orm.RegisterDataBase("test", "sqlite3", "test.db")

	orm.RegisterModelWithPrefix("bolow_", new(User), new(Profile))

	o = orm.NewOrm()

	globalSessions, _ := session.NewManager("redis", `{"cookieName":"gosessionid","gclifetime":3600,"ProviderConfig":"127.0.0.1:6379"}`)
	go globalSessions.GC()

	RecreateTables(false)

}

func RecreateTables(force bool) {
	// Database alias.
	name := "default"

	// Drop table and re-create.

	// Print log.
	verbose := true

	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}
