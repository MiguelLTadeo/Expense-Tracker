package utils

// USAGE
// w.Header().Set("Content-Type", "application/json")

// err = json.NewEncoder(w).Encode(utils.TokenResponse{
// 	Token: token,
// })

type TokenResponse struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SucessResponse struct {
	Message string `json:"message"`
}
