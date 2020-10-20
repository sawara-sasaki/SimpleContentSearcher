package action

import (
	"errors"
	"github.com/webview/webview"
)

var w webview.WebView

func runWebview(url string) {
	w := webview.New(false)
	w.SetSize(100, 100, webview.HintNone)
	w.Navigate(url)
	w.Bind("response", func(s string) {
		w.Dispatch(func() {
			response(s)
		})
	})
	w.Init(`
	window.onload = function() {
		document.body.style.opacity = 0.0;
		checkReady();
	};
	var checkReady = function() {
		if (!document.getElementsByTagName('img')[0]) {
			setTimeout(checkReady, 500);
		} else {
			var eBody = document.getElementsByTagName('body')[0];
			if (!!eBody) {
				response(eBody.innerHTML);
			}
		}
	};`)
	w.Run()
}

func destroy() {
	w.Destroy()
}

func terminate() {
	w.Terminate()
}

func response(s string) {
	panic(errors.New(s))
}
