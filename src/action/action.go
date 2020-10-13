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

func Search(param string)([]interface{}, error) {
	var err error
	var res []interface{}
	if strings.HasPrefix(param, "http://") {
		res = append(res, "No match.")
		return res, nil
	}
	if !strings.HasPrefix(param, "https://www.youtube.com/") {
		param = "https://www.youtube.com/" + param + "/videos"
	}
	req, err := http.NewRequest("GET", param, nil)
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
	var ytInitialData string
	doc.Find("body script").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if len(text) > 0 && strings.HasPrefix(strings.TrimSpace(text), "window[\"ytInitialData\"]") {
			ytInitialData = strings.TrimSpace(text)[26:]
		}
	})
	if len(ytInitialData) > 0 {
		tmp := strings.Split(ytInitialData,"\n")[0]
		var tmpMap map[string]interface{}
		if err := json.Unmarshal([]byte(tmp[0:len(tmp)-1]), &tmpMap); err != nil {
			return res, err
		}
		slice := []interface{}{
			"contents",
			"twoColumnBrowseResultsRenderer",
			"tabs",
			1,
			"tabRenderer",
			"content",
			"sectionListRenderer",
			"contents",
			0,
			"itemSectionRenderer",
			"contents",
			0,
			"gridRenderer",
			"items",
		}
		if tmpSlice, ok := GetSliceInterface(tmpMap, slice); ok {
			for _, v := range tmpSlice {
				if videoId, ok_ := GetStringFromInterface(v.(map[string]interface{}), []interface{}{"gridVideoRenderer","videoId"}); ok_ {
					if title, ok__ := GetStringFromInterface(v.(map[string]interface{}), []interface{}{"gridVideoRenderer", "title", "runs", 0, "text"}); ok__ {
						res = append(res, "<a href='https://www.youtube.com/watch?v=" + videoId + "' target='_blank'>" + title + "</a>")
					}
				}
			}
		}
	}
	return res, nil
}

func GetStringFromInterface(i map[string]interface{}, slice []interface{})(string, bool) {
	last := slice[len(slice)-1].(string)
	result, ok := GetMapInterface(i, slice[:len(slice)-1])
	return result[last].(string), ok
}

func GetMapInterface(i map[string]interface{}, slice []interface{})(map[string]interface{}, bool) {
	itfMap := i
	var ok bool
	var itfSlice []interface{}
	for idx, v := range slice {
		var next interface{}
		if idx + 1 == len(slice) {
			next = "0"
		} else {
			next = slice[idx + 1]
		}
		itfMap, itfSlice, ok = ConvertInterface(itfMap, itfSlice, v, next)
		if !ok {
			return itfMap, false
		}
	}
	return itfMap, true
}

func GetSliceInterface(i map[string]interface{}, slice []interface{})([]interface{}, bool) {
	itfMap := i
	var ok bool
	var itfSlice []interface{}
	for idx, v := range slice {
		var next interface{}
		if idx + 1 == len(slice) {
			next = 0
		} else {
			next = slice[idx + 1]
		}
		itfMap, itfSlice, ok = ConvertInterface(itfMap, itfSlice, v, next)
		if !ok {
			return itfSlice, false
		}
	}
	return itfSlice, true
}

func ConvertInterface(itfMap map[string]interface{}, itfSlice []interface{}, idx interface{}, next interface{})(map[string]interface{}, []interface{}, bool) {
	var nextMap map[string]interface{}
	var nextSlice []interface{}
	var current interface{}
	switch idx.(type){
	case int:
		current = itfSlice[idx.(int)]
	case string:
		current = itfMap[idx.(string)]
	default:
		return nextMap, nextSlice, false
	}
	if current == nil {
		return nextMap, nextSlice, false
	}
	switch next.(type){
	case int:
		nextSlice = current.([]interface{})
	case string:
		nextMap = current.(map[string]interface{})
	default:
		return nextMap, nextSlice, false
	}
	return nextMap, nextSlice, true
}
