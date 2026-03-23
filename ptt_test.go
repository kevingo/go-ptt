package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchPages(t *testing.T) {
	// 創建一個假的 HTTP 回應
	htmlContent := `
		<html>
			<body>
				<div class="action-bar">
					<a class="btn" href="/example/page/3.html">Page 3</a>
				</div>
			</body>
		</html>
	`
	resp := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(htmlContent))
	}))
	defer resp.Close()

	// 創建一個用於測試的通道
	ch := make(chan int)

	// 呼叫 fetchPages 函數進行測試
	go fetchPages(resp.URL, ch)

	// 從通道中讀取結果
	pages := <-ch

	// 驗證結果是否正確
	expectedPages := 3
	if pages != expectedPages {
		t.Errorf("Expected pages: %d, but got: %d", expectedPages, pages)
	}
}
