package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/ruhancs/api-pattern/configs"
	_ "github.com/ruhancs/api-pattern/docs"
	"github.com/ruhancs/api-pattern/internal/entity"
	"github.com/ruhancs/api-pattern/internal/infra/database"
	"github.com/ruhancs/api-pattern/internal/infra/webserver/handler"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title           API Pattern
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Ruhan CS
// @contact.url    ruhancorreasoares@gmail.com
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /

// @securityDefinitions.apikey  ApiKeyAuth
// @in  header
// @name  Authorization

func main() {
	config := configs.LoadConfig(".")
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	//evitar que a requisicao caia e nao retorne
	r.Use(middleware.Recoverer)

	productRepository := database.NewProductRepository(db)
	productHandler := handler.NewProductHandler(productRepository)

	r.Route("/products", func(r chi.Router) {
		//verificar o token de authenticacao no contexto,procura na requisicao o token
		r.Use(jwtauth.Verifier(config.TokenAuth))
		//authenticar o token, verificar se o token é valido
		r.Use(jwtauth.Authenticator)

		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.FindAll)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	userRepository := database.NewUserRepository(db)
	usertHandler := handler.NewUserHandler(userRepository, config.TokenAuth, config.JWTExpiresIn)

	r.Post("/users", usertHandler.CreateUser)
	r.Post("/users/get_token", usertHandler.GetJwt)

	//swagger
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	//server := &http.Server{
	//	Addr: ":8000",
	//	Handler: routes.Routes(*productHandler,*usertHandler,config.TokenAuth),
	//}
	//err = server.ListenAndServe()
	//if err != nil {
	//	log.Panic(err)
	//}

	http.ListenAndServe(":8000", r)
}
