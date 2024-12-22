package main

import (
	_ "fitness_back/docs"
	"fitness_back/handlers"
	middlewares "fitness_back/midlewares"
	"fitness_back/utils"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

// @title Auth Service API
// @version 1.0
// @description This is a sample auth service API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	utils.JwtKey = []byte(os.Getenv("JWT_KEY"))

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	serverPort := os.Getenv("SERVER_PORT")

	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s", dbHost, dbPort, dbUser, dbName, dbSSLMode, dbPassword)

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	handlers.InitDB(db)

	r := mux.NewRouter()

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/register", handlers.Register).Methods("POST")
	authRouter.HandleFunc("/login", handlers.Login).Methods("POST")

	userRouter := r.PathPrefix("/user").Subrouter()
	userRouter.Use(middlewares.AuthMiddleware)

	userRouter.HandleFunc("/char-history", handlers.CharsHistoryHandler).Methods("GET")
	userRouter.HandleFunc("/add-chars", handlers.CreateChars).Methods("POST")

	userRouter.HandleFunc("/search-food", handlers.FoodDataHandler).Methods("GET")
	userRouter.HandleFunc("/ration-history", handlers.RationHistoryHandler).Methods("GET")
	userRouter.HandleFunc("/meal", handlers.CreateMeal).Methods("POST")
	userRouter.HandleFunc("/meal", handlers.DeleteMeal).Methods("DELETE")

	userRouter.HandleFunc("/profile", handlers.ProfileHandler).Methods("GET")
	userRouter.HandleFunc("/update-password", handlers.UpdatePassword).Methods("PUT")
	userRouter.HandleFunc("/update-email", handlers.UpdateEmail).Methods("PUT")
	userRouter.HandleFunc("/update-name", handlers.UpdateName).Methods("PUT")
	userRouter.HandleFunc("/update-username", handlers.UpdateUsername).Methods("PUT")

	setupWorkoutProxy(r)

	log.Printf("Server started at %s", serverPort)
	if err := http.ListenAndServe(":"+serverPort, r); err != nil {
		log.Fatalf("failed start server: %v", err)
	}

}

func setupWorkoutProxy(r *mux.Router) {
	targetURL := os.Getenv("TRAINING_URL")
	target, err := url.Parse(targetURL)
	if err != nil {
		log.Fatalf("Failed to parse target URL: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	workoutRouter := r.PathPrefix("/workout").Subrouter()
	workoutRouter.Use(middlewares.AuthMiddleware)

	workoutRouter.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/workout") {
			r.URL.Scheme = target.Scheme
			r.URL.Host = target.Host
			proxy.ServeHTTP(w, r)
			return
		}

		http.NotFound(w, r)
	})
}
