package cmd

import (
	"fmt"
	"os"

	deliveryGin "github.com/HenkCord/GOServicePlaces/delivery/gin"
	placesRepository "github.com/HenkCord/GOServicePlaces/repository/mongo/places"
	placesUsecase "github.com/HenkCord/GOServicePlaces/usecases/places"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/HenkCord/GOServicePlaces/entities"

	mgo "gopkg.in/mgo.v2"
)

// Connections db
type Connections struct {
	MongoAddr        string
	MongoDB          string
	PlacesCollection string
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run server",
	Long:  ``,
	Run:   runServer,
}

func init() {
	RootCmd.AddCommand(serverCmd)
}

func runServer(ccmd *cobra.Command, args []string) {

	serverAddress := viper.GetString("server.host")
	if serverAddress == "" {
		fmt.Println("Bad server address")
		os.Exit(1)
	}

	hi(viper.GetString("environment"), serverAddress, viper.ConfigFileUsed())

	dbs := connectToDbs()

	placesRepo := placesRepository.InitMongoPlacesRepository(dbs.MongoAddr, dbs.MongoDB, dbs.PlacesCollection)
	placesUse := placesUsecase.InitPlacesUsecase(placesRepo)

	e := gin.Default()
	e.Use(deliveryGin.HandlerError())
	deliveryGin.InitPlacesServer(e, placesUse)

	if err := e.Run(serverAddress); err != nil {
		fmt.Println("Server was unable to start", err)
		os.Exit(1)
	}

}

func connectToDbs() *Connections {
	mongoAddress := viper.GetString("db.mongo")
	if mongoAddress == "" {
		fmt.Println("Bad mongo address:", mongoAddress)
		os.Exit(1)
	}
	fmt.Println("Check connecting to MongoDB:", mongoAddress)
	mongoSession, err := mgo.Dial(mongoAddress)
	defer mongoSession.Close()
	if err != nil {
		fmt.Println("Bad connection to mongo:", mongoAddress, err)
		os.Exit(1)
	}

	if err := mongoSession.Ping(); err != nil {
		fmt.Println(entities.BadConnectToMongo)
	}
	fmt.Println(entities.SuccessConnectToMongo)

	return &Connections{
		MongoAddr:        mongoAddress,
		MongoDB:          "places",
		PlacesCollection: "places",
	}
}

func hi(env string, host string, cfg string) {
	fmt.Println("--------------------------------")
	fmt.Printf("Server started on http://%s\n", host)
	fmt.Println("Environment:", env)
	fmt.Println("Using config file:", cfg)
	fmt.Println("--------------------------------")
}
