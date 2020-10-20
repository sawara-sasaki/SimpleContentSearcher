package main

import (
	"os"
	"fmt"
	"errors"

	"github.com/webview/webview"
)

func main() {
	fmt.Println(getWebviewResponse(os.Args[1]))
}

func getWebviewResponse(url string)(result string) {
	defer func() {
		if err := recover(); err != nil {
			result = fmt.Sprint(err)
		}
	}()
	wv := webview.New(false)
	wv.SetSize(100, 100, webview.HintNone)
	wv.Bind("response", func(s string) {
		wv.Dispatch(func() {
			response(s)
		})
	})
	wv.Init(`
	var counter = 7
	window.onload = function() {
		document.body.style.opacity = 0.0;
		checkReady();
	};
	var checkReady = function() {
		if (counter > 0) {
			counter--;
			if (!!document.getElementsByTagName('img')[0]) {
				counter--;
			}
			setTimeout(checkReady, 500);
		} else {
			response(document.documentElement.innerHTML);
		}
	};`)
	wv.Navigate(url)
	wv.Run()
	return ""
}

func response(s string) {
	panic(errors.New(s))
}
