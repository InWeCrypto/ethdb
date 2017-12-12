package main

import (
	"flag"
	"fmt"

	"github.com/dynamicgo/config"
	"github.com/dynamicgo/slf4go"
	"github.com/go-xorm/xorm"
	ethdb "github.com/inwecrypto/ethdb"
	_ "github.com/lib/pq"
)

var logger = slf4go.Get("ethdb")
var configpath = flag.String("conf", "./ethdb.json", "geth indexer config file path")

func main() {

	flag.Parse()

	conf, err := config.NewFromFile(*configpath)

	if err != nil {
		logger.ErrorF("load eth indexer config err , %s", err)
		return
	}

	username := conf.GetString("ethdb.username", "xxx")
	password := conf.GetString("ethdb.password", "xxx")
	port := conf.GetString("ethdb.port", "6543")
	host := conf.GetString("ethdb.host", "localhost")
	scheme := conf.GetString("ethdb.schema", "postgres")

	engine, err := xorm.NewEngine("postgres", fmt.Sprintf("user=%v password=%v host=%v dbname=%v port=%v sslmode=disable", username, password, host, scheme, port))

	if err != nil {
		logger.ErrorF("create postgres orm engine err , %s", err)
		return
	}

	if err := engine.Sync2(new(ethdb.TableTx), new(ethdb.TableOrder), new(ethdb.TableWallet)); err != nil {
		logger.ErrorF("sync table schema error , %s", err)
		return
	}

}
