package main

import (
	"log"
	"net/http"
	"strconv"

	//"os"
	"time"

	"github.com/firemanm/go_crud/database"
	"github.com/firemanm/go_crud/handlers"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// define http req and duration metrics
var (
	reg = prometheus.NewRegistry()

	requestsTotal = promauto.With(reg).NewCounterVec(
		prometheus.CounterOpts{
			Name: "gocrudapp_http_requests_total",
			Help: "Total number of HTTP requests to go-crud-app.",
		},
		[]string{"path", "method", "statuscode"},
	)
	requestDuration = promauto.With(reg).NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "gocrudapp_http_request_duration_seconds",
			Help: "Duration of HTTP requests to go-crud-app.",
		},
		[]string{"path", "method", "statuscode"},
	)
)

// Custom ResponseWriter to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// Middleware for metrics collection
func metricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// Create custom ResponseWriter to capture status code
		wrappedWriter := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK, // default status code
		}
		// call next handler
		next.ServeHTTP(wrappedWriter, r)

		// now get the route after matching
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()
		if path == "" {
			path = r.URL.Path
		}
		// calculate duration
		duration := time.Since(start).Seconds()

		statusCode := strconv.Itoa(wrappedWriter.statusCode)

		// write metrics
		requestDuration.WithLabelValues(path, r.Method, statusCode).Observe(duration)
		requestsTotal.WithLabelValues(path, r.Method, statusCode).Inc()
	})
}

func main() {

	// Load .env for local development
	godotenv.Load()

	// DB init
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// table create
	err = database.CreateTable(db)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	// handlers init
	userHandler := handlers.NewUserHandler(db)

	// mux routes init
	router := mux.NewRouter()

	// implement middleware to all routes
	router.Use(metricsMiddleware)

	router.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")
	log.Printf("Routes registered through /users...")

	// register custom health check handler
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")
	log.Printf("Health route registered...")

	//register metrics handler
	router.Path("/metrics").Handler(promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	log.Printf("/metrics route registered...")

	// run server

	port := "8080"
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(server.ListenAndServe())
}
