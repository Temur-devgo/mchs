package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/subosito/gotenv"
	"mchs-api/components/db"
	"mchs-api/components/helpers"
	_ "mchs-api/routers"
	"os"
	"path"
)

func init() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	currentDirectory := path.Dir(ex)

	gotenv.Load(currentDirectory + "/.env")
	_ = beego.LoadAppConfig("ini", currentDirectory+"/conf/app.conf")

	DbDriver := "postgres"
	DbUser := os.Getenv("POSTGRES_USER")
	DbPass := os.Getenv("POSTGRES_PASSWORD")
	DbName := os.Getenv("POSTGRES_DB")
	DbHost := os.Getenv("POSTGRES_HOST")
	DbPort := os.Getenv("POSTGRES_PORT")

	db.InitDB(DbDriver, "host="+DbHost+" port="+DbPort+" user="+DbUser+" password="+DbPass+" dbname="+DbName +" sslmode=disable")

	helpers.ContentPath = os.Getenv("CONTENT_URL")
	helpers.ContentDir = os.Getenv("CONTENT_DIR")
	helpers.MainUrl = os.Getenv("MAIN_URL")

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = currentDirectory + "/swagger"
	}
}

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}))
	beego.Run()
	//test
}
