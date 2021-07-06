package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// tạo phiên bản chrome
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// tạo thời gian chờ
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// điều hướng đến một trang, đợi một phần tử, nhấp vào
	var example string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://golang.org/pkg/time/`),
		// đợi phần tử chân trang hiển thị (tức là trang được tải)
		chromedp.WaitVisible(`body > footer`),
		// tìm và nhấp vào liên kết "Expand All"
		chromedp.Click(`#pkg-examples > div`, chromedp.NodeVisible),
		// lấy giá trị của textarea
		chromedp.Value(`#example_After .play .input textarea`, &example),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Go's time.After example:\n%s", example)
}
