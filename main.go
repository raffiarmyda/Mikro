package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"mikro/app/middlewares"
	"mikro/app/routes"
	_userUsecase "mikro/business/users"
	_userDelivery "mikro/deliveries/users"
	"mikro/repository/databases/postgres"
	"mikro/repository/databases/records"
	"mikro/repository/drivers/mongo"
	_configDb "mikro/repository/drivers/postgres"
	"time"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service run on DEBUG MODE")
	}
}

func DbMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&records.Products{},
		&records.Users{},
		&records.Transactions{},
	)
	if err != nil {
		panic(err)
	}
}

func main() {
	postgresConfig := _configDb.ConfigDb{
		DbHost:     viper.GetString(`databases.postgres.host`),
		DbUser:     viper.GetString(`databases.postgres.user`),
		DbPassword: viper.GetString(`databases.postgres.password`),
		DbName:     viper.GetString(`databases.postgres.dbname`),
		DbPort:     viper.GetString(`databases.postgres.port`),
		DbSslMode:  viper.GetString(`databases.postgres.sslmode`),
		DbTimezone: viper.GetString(`databases.postgres.timezone`),
	}

	mongoConfig := mongo.ConfigDb{
		Cluster:  viper.GetString(`databases.mongodb.cluster`),
		Username: viper.GetString(`databases.mongodb.username`),
		Password: viper.GetString(`databases.mongodb.password`),
	}

	configJWT := middlewares.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	connPostgres := postgresConfig.InitialDb(viper.GetBool(`debug`))

	//MONGO
	logCol := middlewares.InitCollection(struct {
		DbName     string
		Collection string
	}{
		DbName:     viper.GetString(`databases.mongodb.dbname`),
		Collection: viper.GetString(`databases.mongodb.collection.logger`),
	})

	initMongo := mongoConfig.InitDb()
	loggerMiddleware := middlewares.InitConfig(initMongo, logCol, viper.GetDuration(`context.timeout`))

	DbMigrate(connPostgres)
	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userRepository := postgres.NewPostgresUserRepository(connPostgres)
	userUsecase := _userUsecase.NewUserUsecase(userRepository, timeoutContext, &configJWT)
	userDelivery := _userDelivery.NewUserController(userUsecase)

	routesInit := routes.ControllerList{
		UserController:   *userDelivery,
		LoggerMiddleware: *loggerMiddleware,
		JWTMiddleware:    configJWT.Init(),
	}

	routesInit.RouteUsers(e)
	address := fmt.Sprintf("%v:%v",
		viper.GetString("server.address.host"),
		viper.GetString("server.address.port"),
	)
	err := e.Start(address)
	if err != nil {
		panic(err)
	}
}
