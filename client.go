package splitTools

import (
	"github.com/cibersfaxa/splitTools/splitTools/api/add"
	"github.com/cibersfaxa/splitTools/splitTools/api/fetch"
	"github.com/imroc/req/v3"
)

type Client struct {
	AddAPI   *add.Add
	FetchAPI *fetch.Fetch
}

func NewClient() *Client {
	client := req.C().SetBaseURL("http://localhost:8080")
	client.EnableDumpAll() // Optional: Enable this to debug requests and responses
	return &Client{
		AddAPI:   &add.Add{Client: client},
		FetchAPI: &fetch.Fetch{Client: client},
	}
}
