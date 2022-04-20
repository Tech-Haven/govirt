package responses


type TokenResponse struct {
	AccessToken		string	`json:"access_token"`
	Exp						string	`json:"exp"`
	TokenType			string	`json:"token_type"`
}

type ErrorResponse struct {
	ErrorCode			string	`json:"error_code"`
	Error					string	`json:"error"`
}