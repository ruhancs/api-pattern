package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"github.com/ruhancs/api-pattern/internal/infra/webserver/handler"
)

func Routes(productHandler handler.ProductHandler, userHandler handler.UserHandler,tokenAuth *jwtauth.JWTAuth) http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET","POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-TOKEN"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	}))

	//para checar o servico
	mux.Use(middleware.Heartbeat("/ping"))
	mux.Use(middleware.Logger)
	//evitar que a requisicao caia e nao retorne
	mux.Use(middleware.Recoverer)

	mux.Route("/products",func(r chi.Router) {
		//verificar o token de authenticacao no contexto,procura na requisicao o token
		r.Use(jwtauth.Verifier(tokenAuth))
		//authenticar o token, verificar se o token é valido
		r.Use(jwtauth.Authenticator)
		
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.FindAll)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	mux.Post("/user",userHandler.CreateUser)
	mux.Post("/user/get_token", userHandler.GetJwt)

	return mux
}