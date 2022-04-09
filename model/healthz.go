package model

// A HealthzResponse expresses health check message.
type HealthzResponse struct {
	Message string `json:"message"`
}

// func ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	task := HealthzResponse{"OK"}
// 	encoder := json.NewEncoder(w)
// 	err := encoder.Encode(task)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// }
