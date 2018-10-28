package handler

import "fmt"

type SuccessResponse struct {
	Id      string `json:"id"`
	LongUrl string `json:"longURL"`
	Status  string `json:"status"`
}

type BadResponse struct {
	Err    string `json:"err"`
	Status string `json:"status"`
}

func CustomSuccessResponse(short_url, long_url string) SuccessResponse {
	return SuccessResponse{
		Id:      short_url,
		LongUrl: long_url,
		Status:  "OK",
	}
}

func CustomBadResponse(err error) BadResponse {
	return BadResponse{
		Err:    fmt.Sprintf("%v", err),
		Status: "ERROR",
	}
}
