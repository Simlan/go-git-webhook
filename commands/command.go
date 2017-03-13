package commands

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"go-git-webhook/models"
	_ "github.com/go-sql-driver/mysql"
	_ "go-git-webhook/routers"
	_ "go-git-webhook/modules/filters"
	"github.com/astaxie/beego/logs"
)

//注册数据库
func RegisterDataBase()  {
	host := beego.AppConfig.String("db_host");
	database := beego.AppConfig.String("db_database");
	username := beego.AppConfig.String("db_username");
	password := beego.AppConfig.String("db_password");
	port := beego.AppConfig.String("db_port");

	dataSource := username + ":" + password + "@tcp(" + host + ":" + port +")/" + database + "?charset=utf8&parseTime=true";

	orm.RegisterDataBase("default", "mysql", dataSource)
}

//注册Model
func RegisterModel()  {
	orm.RegisterModel(new(models.Member))
	orm.RegisterModel(new(models.Server))
	orm.RegisterModel(new(models.WebHook))
	orm.RegisterModel(new(models.Scheduler))
	orm.RegisterModel(new(models.Relation))
}

func RegisterLogger()  {
	logs.SetLogger("console")
	logs.SetLogger("file",`{"filename":"logs/log.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Async()
}
//注册orm命令行工具
func RunCommand()  {
	orm.RunCommand()
	Install()
}

//启动Web监听
func Run()  {
	beego.Run()
}
