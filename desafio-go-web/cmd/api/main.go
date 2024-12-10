package main

import (
	"app/internal/controller"
	"app/internal/loader"
	"app/internal/repository"
	"app/internal/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	// env
	godotenv.Load()

	// application
	// - config
	cfg := &ConfigAppDefault{
		ServerAddr: os.Getenv("SERVER_ADDR"),
		DbFile:     os.Getenv("DB_FILE"),
	}

	app := NewApplicationDefault(cfg)

	// - setup
	err := app.SetUp()
	if err != nil {
		fmt.Println(err)
		return
	}

	// - run
	err = app.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// ConfigAppDefault represents the configuration of the default application
type ConfigAppDefault struct {
	// serverAddr represents the address of the server
	ServerAddr string
	// dbFile represents the path to the database file
	DbFile string
}

// NewApplicationDefault creates a new default application
func NewApplicationDefault(cfg *ConfigAppDefault) *ApplicationDefault {
	// default values
	defaultRouter := chi.NewRouter()
	defaultConfig := &ConfigAppDefault{
		ServerAddr: ":8080",
		DbFile:     "",
	}
	if cfg != nil {
		if cfg.ServerAddr != "" {
			defaultConfig.ServerAddr = cfg.ServerAddr
		}
		if cfg.DbFile != "" {
			defaultConfig.DbFile = cfg.DbFile
		}
	}			

	return &ApplicationDefault{
		rt:         defaultRouter,
		serverAddr: defaultConfig.ServerAddr,
		dbFile:     defaultConfig.DbFile,
	}
}


// ApplicationDefault represents the default application
type ApplicationDefault struct {
	// router represents the router of the application
	rt *chi.Mux
	// serverAddr represents the address of the server
	serverAddr string
	// dbFile represents the path to the database file
	dbFile string
}


// SetUp sets up the application
func (a *ApplicationDefault) SetUp() (err error) {
	// dependencies
	db, lastId := loader.NewLoaderTicketCSV(a.dbFile)
	tickets,err := db.Load()
	if err != nil {
		log.Println("failed to load db")
		return 
	}
	rp := repository.NewRepositoryTicketMap(a.dbFile, lastId, tickets)
	// service ...x
	service := service.NewServiceTicketDefault(rp)
	// handler ...
	ctrl := controller.NewControllerTicketDefault(service)

	// routes
	(*a).rt.Get("/health", ctrl.GetHealth)
	(*a).rt.Get("/ticket/getByCountry/{dest}", ctrl.GetByCountry)
	(*a).rt.Get("/ticket/getAverage/{dest}", ctrl.GetAverageByCountry)
	
	return
}

// Run runs the application
func (a *ApplicationDefault) Run() (err error) {
	err = http.ListenAndServe(a.serverAddr, a.rt)
	return
}
