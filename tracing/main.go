package main

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"tracing/controller"
	"tracing/middleware"
	"tracing/pkg/app"
	"tracing/pkg/http"
	"tracing/pkg/store"
	"tracing/repository"
	"tracing/repository/ent"
	"tracing/repository/service/game"
)

func main() {

	ctx := context.Background()
	tp, err := initTracer(ctx)
	if err != nil {
		panic(err)
	}

	source := 43
	signKey := "1234567890"
	host := "http://apisgame.qiyi.domain"

	components := &app.Components{
		Tracer: &app.Tracer{
			Ctr: tp.Tracer("ctr"),
			DB:  tp.Tracer("db"),
		},
		Servers: &app.Servers{
			BI: &app.ServerInfo{
				Host:    host + "/bi/",
				Source:  source,
				SignKey: signKey,
			},
			Payment: &app.ServerInfo{
				Host:    host + "/pay",
				Source:  source,
				SignKey: signKey,
			},
		},
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true",
		"root",
		"123456",
		"127.0.0.1:3306",
		"test",
	)

	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db := ent.NewClient(ent.Driver(NewDriver(drv, components.Tracer.DB)))

	//db.Intercept(ent.InterceptFunc(func(next ent.Querier) ent.Querier {
	//	return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
	//
	//		if q, ok := query.(*ent.UserQuery); ok {
	//			q.Limit(1)
	//		}
	//
	//		value, err := next.Query(ctx, query)
	//		pretty.Println(value)
	//		return value, err
	//	})
	//}))

	components.DB = db
	//cache := store.NewMemory()
	cache := store.NewRedis()
	components.Cache = cache

	request := http.New(nil)
	servBI := game.NewBI(request, components)
	servPayment := game.NewPayment(request, components)
	repoUser := repository.NewUser(db, servBI, servPayment)

	ctrUser := controller.NewUser(repoUser, components)

	handler := controller.NewHandler(components.Tracer.Ctr)

	engine := gin.New()
	engine.Use(middleware.Cors())
	engine.GET("/user/:id", handler.Wrapper(ctrUser.One))
	engine.GET("/users", handler.Wrapper(ctrUser.Many))

	engine.Run()

}
