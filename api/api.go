package api

//coin balance params the api will take
type CoinBalanceParams struct {
	Username string
}

//successful 200 response type
type CoinBalanceParams struct {
	//success code 200 int
	Code int
	//account balance
	Balance int64
}

//Error Response
type Error struct {
	//error code
	Code int
	//error message
	Message string
}
