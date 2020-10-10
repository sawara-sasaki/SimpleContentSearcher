package action

import (
	"fmt"
	"encoding/json"
)

type ActionRequest struct {
	Action     string   `json:"action"`
	Parameters []string `json:"parameters"`
}

type ActionResponse struct {
	Status int           `json:"status"`
	Data   []interface{} `json:"data"`
}

func Handle(request []byte)(ActionResponse, error) {
	var req ActionRequest
	var res ActionResponse
	var err error
	json.Unmarshal(request, &req)
	switch req.Action {
	case "search":
		if len(req.Parameters) == 1 {
			res.Data = Search(req.Parameters[0])
		} else {
			err = fmt.Errorf("err %s", "Bad Parameters")
		}
	default:
		err = fmt.Errorf("err %s", "Bad Action")
	}
	return res, err
}

func Search(word string) []interface{} {
	var res []interface{}
	switch word {
	case "match":
		res = append(res, "match!")
	default:
		res = append(res, "No match.")
	}
	return res
}
