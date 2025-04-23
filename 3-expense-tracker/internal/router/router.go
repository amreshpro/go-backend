package router

import (
	"net/http"
	"github.com/amreshpro/go-backend/3-expense-tracker/internal/auth"
	"github.com/amreshpro/go-backend/3-expense-tracker/internal/expense"
	"github.com/gorilla/mux"
)

func AppRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", healthCheck).Methods("GET")

	// auth
	router.HandleFunc("/login/", auth.LoginHandler).Methods("POST")
	router.HandleFunc("/singup", auth.SignupHandler).Methods("POST")

	// protected routes
	api := router.PathPrefix("/api").Subrouter()
	// api.Use(auth.JwtMiddleware)

	api.HandleFunc("/expenses", expense.CreateExpenseHandler).Methods("POST")
	api.HandleFunc("/expenses", expense.GetExpensesHandler).Methods("GET")

	return router
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Health is ok"))
}
