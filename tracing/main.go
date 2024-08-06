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
	"tracing/pkg/store"
	"tracing/repository"
	"tracing/repository/ent"
)

func main() {

	ctx := context.Background()
	tp, err := initTracer(ctx)
	if err != nil {
		panic(err)
	}

	components := &app.Components{
		Tracer: &app.Tracer{
			Ctr: tp.Tracer("ctr"),
			DB:  tp.Tracer("db"),
		},
	}

	host := "127.0.0.1:3306"
	name := "test"
	username := "root"
	password := "123456"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true",
		username,
		password,
		host,
		name,
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
	cache := store.New()
	components.Cache = cache

	repoUser := repository.NewUser(db)
	components.RepoUser = repoUser

	ctrUser := controller.NewUser(components)

	handler := controller.NewHandler(components.Tracer.Ctr)

	engine := gin.New()
	engine.Use(middleware.Cors())
	engine.GET("/user/:id", handler.Wrapper(ctrUser.One))
	engine.GET("/users", handler.Wrapper(ctrUser.Many))

	engine.Run()

}
