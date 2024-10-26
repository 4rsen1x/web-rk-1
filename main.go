package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Output struct {
	Result int `json:"result"`
}

func ChessTableHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	nParam := r.URL.Query().Get("N")
	if nParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing parameter N"))
		return
	}

	n, err := strconv.Atoi(nParam)
	if err != nil || n <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid parameter N"))
		return
	}

	// totalCells := n * n
	// closedCells := 2
	// availableCells := totalCells - closedCells

	// result := availableCells / 2

	var result int
	totalCells := n * n

	
	if n == 1 {
		result = 0;
	} else if n%2 == 0 {
		result = ((totalCells-2) - 2) / 2
	} else {
		result = ((totalCells - 2) - 1) / 2
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Output{Result: result})
}

func main() {
	http.HandleFunc("/chess_table", ChessTableHandler)

	fmt.Println("starting server on 127.0.0.1:8081...")
	err := http.ListenAndServe("127.0.0.1:8081", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
