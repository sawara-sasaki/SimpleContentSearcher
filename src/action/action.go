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
		param = "https://www.youtube.com/results?search_query=" + param
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
		text := strings.TrimSpace(strings.Replace(s.Text(), "\n", "", -1))
		if len(text) > 0 && strings.HasPrefix(text, "// scraper_data_begin") {
			text = text[21:]
		}
		if len(text) > 0 && strings.HasPrefix(text, "window[\"ytInitialData\"]") {
			ytInitialData = text[26:]
		} else if len(text) > 0 && strings.HasPrefix(text, "var ytInitialData") {
			ytInitialData = text[20:]
		}
	})

	if len(ytInitialData) > 0 {
		tmp := strings.Split(ytInitialData,"};")[0]
		var tmpMap map[string]interface{}
		if err := json.Unmarshal([]byte(tmp + "}"), &tmpMap); err != nil {
			return res, err
		}
		for idx := 0; idx < 2; idx++ {
			baseSlice, idSlice, titleSlice := GetInterfaceSlice(param, idx)
			if tmpSlice, ok := GetSliceInterface(tmpMap, baseSlice); ok {
				for _, v := range tmpSlice {
					if videoId, ok_ := GetStringFromInterface(v.(map[string]interface{}), idSlice); ok_ {
						if title, ok__ := GetStringFromInterface(v.(map[string]interface{}), titleSlice); ok__ {
							res = append(res, "<a href='/player.html?v=" + videoId + "'>" + title + "</a>")
						}
					}
				}
			}
			if !strings.HasPrefix(param, "https://www.youtube.com/results?search_query") || len(res) > 0 {
				break
			}
		}
	}
	return res, nil
}

func GetInterfaceSlice(param string, pattern int)([]interface{}, []interface{}, []interface{}) {
	if strings.HasPrefix(param, "https://www.youtube.com/watch?v") {
		return []interface{}{
			"contents",
			"twoColumnWatchNextResults",
			"secondaryResults",
			"secondaryResults",
			"results",
		}, []interface{}{
			"compactVideoRenderer",
			"videoId",
		}, []interface{}{
			"compactVideoRenderer",
			"title",
			"simpleText",
		}
	} else if strings.HasPrefix(param, "https://www.youtube.com/results?search_query") {
		if pattern == 0 {
			return []interface{}{
				"contents",
				"twoColumnSearchResultsRenderer",
				"primaryContents",
				"sectionListRenderer",
				"contents",
				0,
				"itemSectionRenderer",
				"contents",
				-2,
				"shelfRenderer",
				"content",
				"verticalListRenderer",
				"items",
			}, []interface{}{
				"videoRenderer",
				"videoId",
			}, []interface{}{
				"videoRenderer",
				"title",
				"runs",
				0,
				"text",
			}
		} else {
			return []interface{}{
				"contents",
				"twoColumnSearchResultsRenderer",
				"primaryContents",
				"sectionListRenderer",
				"contents",
			}, []interface{}{
				"itemSectionRenderer",
				"contents",
				-1,
				"videoRenderer",
				"videoId",
			}, []interface{}{
				"itemSectionRenderer",
				"contents",
				-1,
				"videoRenderer",
				"title",
				"runs",
				0,
				"text",
			}
		}
	} else {
		return []interface{}{
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
		}, []interface{}{
			"gridVideoRenderer",
			"videoId",
		}, []interface{}{
			"gridVideoRenderer",
			"title",
			"runs",
			0,
			"text",
		}
	}
}

func GetStringFromInterface(i map[string]interface{}, slice []interface{})(string, bool) {
	last := slice[len(slice)-1].(string)
	result, ok := GetMapInterface(i, slice[:len(slice)-1])
	if !ok {
		return "", false
	} else if _, ok_ := result[last]; !ok_ {
		return "", false
	}
	return result[last].(string), true
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
		if idx.(int) < 0 {
			for idx_, v_ := range itfSlice {
				for idx__, _ := range v_.(map[string]interface{}) {
					if next.(string) == idx__ {
						current = itfSlice[idx_]
						break
					}
				}
				if current != nil {
					break
				}
			}
		} else if len(itfSlice) > idx.(int) {
			current = itfSlice[idx.(int)]
		} else {
			return nextMap, nextSlice, false
		}
	case string:
		if c, ok := itfMap[idx.(string)]; ok {
			current = c
		} else {
			return nextMap, nextSlice, false
		}
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
