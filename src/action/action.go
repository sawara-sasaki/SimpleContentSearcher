package action

import (
	"fmt"
	"strings"
	"net/http"
	"encoding/json"

	"github.com/PuerkitoBio/goquery"
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
			res.Data, err = Search(req.Parameters[0])
		} else {
			err = fmt.Errorf("err %s", "Bad Parameters")
		}
	default:
		err = fmt.Errorf("err %s", "Bad Action")
	}
	return res, err
}

func Search(url string)([]interface{}, error) {
	var err error
	var res []interface{}
	if !strings.HasPrefix(url, "http") {
		res = append(res, "No match.")
		return res, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return res, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return res, err
	}
	doc.Find("body a").Each(func(i int, s *goquery.Selection) {
		url_, _ := s.Attr("href")
		if len(url_) > 0 {
			text := strings.Trim(strings.TrimSpace(s.Text()), "\n")
			if len(text) < 1 {
				text = "_"
			}
			res = append(res, "<a href='" + url_ + "' target='_blank'>" + text + "</a>")
		}
	})
	doc.Find("body img").Each(func(i int, s *goquery.Selection) {
		url_, _ := s.Attr("src")
		if len(url_) > 0 {
			alt, _ := s.Attr("alt")
			if len(alt) < 1 {
				alt = "_"
			}
			res = append(res, "<a href='" + url_ + "' target='_blank'>" + alt + "</a>")
		}
	})
	return res, nil
}
