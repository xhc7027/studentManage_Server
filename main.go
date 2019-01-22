package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"time"
	"xuhaochen/models"
	_ "xuhaochen/routers"
)

func init() {
	//初始化数据模型
	var StartTime = time.Now().Unix()
	models.Init(StartTime)
}

//type httpServer struct {
//}
//
//func (s *httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	// Stop here if its Preflighted OPTIONS request
//	if origin := r.Header.Get("Origin"); origin != "" {
//		w.Header().Set("Access-Control-Allow-Origin", "*")
//		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT")
//		w.Header().Set("Access-Control-Allow-Headers", "x-requested-with,content-type")
//		w.Header().Set("Access-Control-Allow-Credentials", "true")
//		//w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
//		w.Header().Set("content-type", "application/json")             //返回数据格式是json
//	}
//
//	if r.Method == "OPTIONS" {
//		return
//	}
//
//	w.Write([]byte("hello"))
//}

func main() {
	//addr := flag.String("http-address", "", "")
	//flag.Parse()
	//
	//var h httpServer
	//
	//httpListener, err := net.Listen("tcp", *addr)
	//server := http.Server{
	//	Handler: &h,
	//}
	//server.Serve(httpListener)
	//
	//fmt.Println("finish ", err)
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
