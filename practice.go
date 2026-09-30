package main

import (
	"net/http"
)

type corporate struct {
	Company string
	Age     int
}

var cr = corporate{
	Company: "AT&T",
	Age:     80,
}

func login(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "test.html")
	//	w.Header().Set("Content-Type", "application/json")
	//	w.WriteHeader(http.StatusOK)
	//	json.NewEncoder(w).Encode(cr)
	//
	// w.Write([]byte(`{"vaghtemoon kheili kame"}`))
}

func customers(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "customers.html")

}
func dashboard(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dashboard.html")

}
func orders(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "orders.html")

}
func reports(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "reports.html")

}
func sales(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "sales.html")

}
func settings(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "settings.html")

}

// func login(w http.ResponseWriter, r *http.Request) {
// 	test := corporate{}
// 	json.NewDecoder(r.Body).Decode(&test)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusAccepted)
// 	json.NewEncoder(w).Encode(test)

// }
func main() {

	mux := http.NewServeMux()

	//mux.HandleFunc("/info", handler)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/customers", customers)
	mux.HandleFunc("/dashboard", dashboard)
	mux.HandleFunc("/orders", orders)
	mux.HandleFunc("/reports", reports)
	mux.HandleFunc("/sales", sales)
	mux.HandleFunc("/settings", settings)

	mux.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("./static")),
		),
	)

	server := &http.Server{
		Addr:    ":61573",
		Handler: mux,
	}
	server.ListenAndServe()

}
