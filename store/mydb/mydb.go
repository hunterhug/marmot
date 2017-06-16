package mydb

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"os"
	"time"
)

const (
	PG    = "postgres"
	MYSQL = "mysql"
)

type MyDbConfig struct {
	DriverName string
	DbConfig
	MaxIdleConns int
	MaxOpenConns int
	DebugToFile  bool
	Debug        bool
}

type MyDb struct {
	Config MyDbConfig
	Client *xorm.Engine
}

type DbConfig struct {
	Name    string
	Host    string
	User    string
	Pass    string
	Port    string
	Sslmode string // sslmode=verify-full require
}

func NewMysqlUrl(c DbConfig) string {
	if c.Port == "" {
		c.Port = "3306"
	}
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", c.User, c.Pass, c.Host, c.Port, c.Name)
	return dns
}

func NewPqUrl(c DbConfig) string {
	if c.Port == "" {
		c.Port = "5432"
	}
	//if c.Sslmode == "" {
	//	c.Sslmode = "verify-full"
	//}
	dns := fmt.Sprintf("dbname=%s host=%s user=%s password=%s port=%s sslmode=%s", c.Name, c.Host, c.User, c.Pass, c.Port, c.Sslmode)
	return dns
}

func NewDb(config MyDbConfig) (*MyDb, error) {
	db := new(MyDb)
	db.Config = config
	dns := ""
	if config.DriverName == MYSQL {
		dns = NewMysqlUrl(config.DbConfig)
	}
	if config.DriverName == PG {
		dns = NewPqUrl(config.DbConfig)
	}
	if config.DriverName == MYSQL || config.DriverName == PG {

		engine, err := xorm.NewEngine(config.DriverName, dns)
		if err != nil {
			return db, err
		}

		if config.Debug {
			if config.DebugToFile {
				f, err := os.Create("/tmp/" + config.DriverName + ".log")
				if err != nil {
					panic(err)
				}
				engine.SetLogger(xorm.NewSimpleLogger(f))
			}
			engine.ShowSQL(true) //会在控制台打印出生成的SQL语句
		}

		engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai") //标准时区,或者"Asia/Shanghai"

		engine.SetMaxIdleConns(config.MaxIdleConns) //  Mysql连接池
		engine.SetMaxOpenConns(config.MaxOpenConns)

		if err := engine.Ping(); err != nil {
			return db, err
		}
		db.Client = engine
		return db, nil
	} else {
		return db, errors.New("Not support this drive:" + config.DriverName)
	}
}

func (db *MyDb) Ping() error {
	if db.Client == nil {
		newdb, err := NewDb(db.Config)
		db = newdb
		return err
	} else {
		return db.Client.Ping()
	}
}

func (db *MyDb) IsTableExist(beanOrTableName interface{}) (bool, error) {
	return db.Client.IsTableExist(beanOrTableName)

}

func (db *MyDb) DropTables(beans ...interface{}) error {
	err := db.Client.DropTables(beans...)
	return err
}

func (db *MyDb) CreateTables(beanOrTableName interface{}) error {
	err := db.Client.CreateTables(beanOrTableName)
	return err
}

func (db *MyDb) Insert(beans ...interface{}) (int64, error) {
	return db.Client.Insert(beans...)
}

func (db *MyDb) InsertOne(beans interface{}) (int64, error) {
	return db.Client.InsertOne(beans)
}

func (db *MyDb) Update(bean interface{}, condiBean ...interface{}) (int64, error) {
	return db.Client.Update(bean, condiBean...)

}

func (db *MyDb) Delete(bean interface{}) (int64, error) {
	return db.Client.Delete(bean)

}

//sql := "select * from userinfo"
//results, err := engine.Query(sql)

func (db *MyDb) Query(sql string, paramStr ...interface{}) (resultsSlice []map[string][]byte, err error) {
	return db.Client.Query(sql, paramStr...)

}

//sql = "update `userinfo` set username=? where id=?"
//res, err := engine.Exec(sql, "xiaolun", 1)
func (db *MyDb) Exec(sql string, args ...interface{}) (sql.Result, error) {
	return db.Client.Exec(sql, args...)

}
