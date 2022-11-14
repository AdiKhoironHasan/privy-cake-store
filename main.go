package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/AdiKhoironHasan/privy-cake-store/pkg/database"

	integ "github.com/AdiKhoironHasan/privy-cake-store/internal/integration"
	SqlRepo "github.com/AdiKhoironHasan/privy-cake-store/internal/repository/mysql"
	"github.com/AdiKhoironHasan/privy-cake-store/internal/services"
	handlers "github.com/AdiKhoironHasan/privy-cake-store/internal/transport/http"
	"github.com/AdiKhoironHasan/privy-cake-store/internal/transport/http/middleware"

	"github.com/apex/log"
	"github.com/labstack/echo"

	"github.com/spf13/viper"
)

func main() {

	errChan := make(chan error)

	e := echo.New()
	m := middleware.NewMidleware()

	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config-dev")

	err := viper.ReadInConfig()
	if err != nil {
		e.Logger.Fatal(err)
	}

	dbhost, dbUser, dbPassword, dbName, dbPort :=
		viper.GetString("db.mysql.host"),
		viper.GetString("db.mysql.user"),
		viper.GetString("db.mysql.password"),
		viper.GetString("db.mysql.dbname"),
		viper.GetString("db.mysql.port")

	MySqlDB, err := database.MySqlInitialize(dbhost, dbUser, dbPassword, dbName, dbPort)

	if err != nil {
		log.Fatal("Failed to Connect MySQL Database: " + err.Error())
	}

	defer func() {
		err := MySqlDB.Conn.Close()
		if err != nil {
			log.Fatal(err.Error())
		}

		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	e.Use(m.CORS)

	sqlrepo := SqlRepo.NewRepo(MySqlDB.Conn)
	integSrv := integ.NewService()
	srv := services.NewService(sqlrepo, integSrv)
	handlers.NewHttpHandler(e, srv)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		errChan <- e.Start(":" + viper.GetString("server.port"))
	}()

	e.Logger.Print("Starting ", viper.GetString("appName"))
	err = <-errChan
	log.Error(err.Error())

}
