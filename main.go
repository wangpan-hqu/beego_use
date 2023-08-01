package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/casdoor/casdoor-go-sdk/auth"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

//go:embed token_jwt_key.pem
var JwtPublicKey string

func init() {
	casdoorEndpoint := beego.AppConfig.String("casdoorEndpoint")
	clientId := beego.AppConfig.String("clientId")
	clientSecret := beego.AppConfig.String("clientSecret")
	casdoorOrganization := beego.AppConfig.String("casdoorOrganization")
	casdoorApplication := beego.AppConfig.String("casdoorApplication")
	auth.InitConfig(casdoorEndpoint, clientId, clientSecret, JwtPublicKey, casdoorOrganization, casdoorApplication)
}

// Adapter represents the MySQL adapter for policy storage.
type Adapter struct {
	driverName     string
	dataSourceName string
	dbName         string
	Engine         *xorm.Engine
}

var adapter *Adapter

func initAdapter() {
	adapter = &Adapter{
		driverName:     beego.AppConfig.String("driverName"),
		dataSourceName: beego.AppConfig.String("dataSourceName"),
		dbName:         beego.AppConfig.String("casdoor")}

	dataSourceName := adapter.dataSourceName + adapter.dbName
	if adapter.driverName != "mysql" {
		dataSourceName = adapter.dataSourceName
	}
	engine, err := xorm.NewEngine(adapter.driverName, dataSourceName)
	if err != nil {
		panic(err)
	}

	adapter.Engine = engine

	_, err = engine.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s default charset utf8mb4 COLLATE utf8mb4_general_ci", adapter.dbName))
}

func main() {
	//initAdapter()
	// beego.Router()
}
