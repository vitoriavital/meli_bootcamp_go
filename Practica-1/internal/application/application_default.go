package application

import (
	"app/internal/handler"
	"app/internal/repository"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"database/sql"
 	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"os"
)
// NewApplicationDefault creates a new default application.
func NewApplicationDefault(addr string) (a *ApplicationDefault) {
	// default config
	defaultRouter := chi.NewRouter()
	defaultAddr := ":8080"
	if addr != "" {
		defaultAddr = addr
	}
	dataSource := os.Getenv("DB")
	storageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	if err = storageDB.Ping(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	log.Println("Database connected")

	a = &ApplicationDefault{
		rt: defaultRouter,
		addr: defaultAddr,
		db: storageDB,
	}
	return
}
		
// ApplicationDefault is the default application.
type ApplicationDefault struct {
	// rt is the router.
	rt *chi.Mux
	// addr is the address to listen.
	addr string
	// filePathStore is the file path to store.
	// filePathStore string
	db *sql.DB
}

// TearDown tears down the application.
func (a *ApplicationDefault) TearDown() (err error) {
	return
}

// SetUp sets up the application.
func (a *ApplicationDefault) SetUp() (err error) {
	// dependencies
	// - store
	// st := store.NewStoreProductJSON(a.filePathStore)
	// - repository
	rp := repository.NewRepositoryProductStore(a.db)
	// - handler
	hd := handler.NewHandlerProduct(rp)

	// router
	// - middlewares
	a.rt.Use(middleware.Logger)
	a.rt.Use(middleware.Recoverer)
	// - endpoints
	a.rt.Route("/products", func(r chi.Router) {
		// GET /products/{id}
		r.Get("/{id}", hd.GetById())
		// POST /products
		r.Post("/", hd.Create())
		// PUT /products/{id}
		r.Put("/{id}", hd.UpdateOrCreate())
		// PATCH /products/{id}
		r.Patch("/{id}", hd.Update())
		// DELETE /products/{id}
		r.Delete("/{id}", hd.Delete())
	})

	return
}

// Run runs the application.
func (a *ApplicationDefault) Run() (err error) {
	err = http.ListenAndServe(a.addr, a.rt)
	return
}