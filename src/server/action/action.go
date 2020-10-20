package action

import (
	"fmt"
	"strings"

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

func Handle(req ActionRequest, webviewResponse string)(ActionResponse, error) {
	var res ActionResponse
	var err error
	switch req.Action {
	case "search":
		if len(req.Parameters) == 1 {
			res.Data, err = Search(webviewResponse)
		} else {
			err = fmt.Errorf("err %s", "Bad Parameters")
		}
	default:
		err = fmt.Errorf("err %s", "Bad Action")
	}
	return res, err
}

func Search(webviewResponse string)([]interface{}, error) {
	var res []interface{}
	if len(webviewResponse) == 0 {
		res = append(res, "No match.")
		return res, nil
	}
	stringReader := strings.NewReader(webviewResponse)
	doc, err := goquery.NewDocumentFromReader(stringReader)
	if err == nil {
		doc.Find("div[data-testid='primaryColumn']").Each(func(i int, doc_ *goquery.Selection) {
			doc_.Find("div[data-testid='image'] img").Each(func(i_ int, s *goquery.Selection) {
				url_, _ := s.Attr("src")
				if len(url_) > 0 {
					alt, _ := s.Attr("alt")
					if len(alt) < 1 {
						alt = "image"
					}
					res = append(res, "<a href='" + url_ + "' target='_blank'>" + alt + "</a>")
				}
			})
			doc_.Find("div[dir='ltr'] > span").Each(func(i_ int, s *goquery.Selection) {
				if len(s.Text()) > 0 {
					// TODO href
					res = append(res, "<a href='#'>" + s.Text() + "</a>")
				}
			})
		})
	}
	return res, nil
}
