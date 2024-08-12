package main

import (
	"base-socket-service/conf"
	"base-socket-service/socket"
	"flag"
	"fmt"
)

var ENV string

func init() {
	flag.StringVar(&ENV, "env", "testnet", "EnvironmentEnum")
}

func InitEnv() {
	flag.Parse()
	if ENV == "loc" {
		conf.SystemEnvironmentEnum = conf.ExampleEnvironmentEnum
	} else if ENV == "mainnet" {
		conf.SystemEnvironmentEnum = conf.MainnetEnvironmentEnum
	} else if ENV == "testnet" {
		conf.SystemEnvironmentEnum = conf.TestnetEnvironmentEnum

	}
	fmt.Println(fmt.Sprintf("%s%v", "Env : ", ENV))
}

func InitAll() {
	conf.InitConfig()
	//major.InitMongo()
	//major.InitSqlConfig()
	//logName := "asset-base-service"
	//major.InitLogger(logName)
}

func main() {
	InitEnv()
	InitAll()
	socket.Run()
}
