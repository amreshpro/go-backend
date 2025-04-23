package expense

import "net/http"

func CreateExpenseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("create expense handler yoyo"))
}

func GetExpensesHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("get expense handler"))
}
