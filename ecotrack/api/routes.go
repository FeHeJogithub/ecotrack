package api

import (
	"encoding/json"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/habits", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			json.NewEncoder(w).Encode(habits)
		case http.MethodPost:
			var h Habit
			if err := json.NewDecoder(r.Body).Decode(&h); err != nil {
				http.Error(w, "Error en JSON", http.StatusBadRequest)
				return
			}
			h.ID = idCounter
			idCounter++
			habits = append(habits, h)
			json.NewEncoder(w).Encode(h)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})
}
