package web

type errorResponse struct {
	Code        string `json:"ERR_CODE"`
	Description string `json:"Description"`
}

type createResponse struct {
	Alias       string `json:"alias"`
	Original    string `json:"original"`
	Shortened   string `json:"shortened"`
	ElapsedTime string `json:"elapsed_time"`
}
