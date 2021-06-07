// Command click2 is a chromedp example demonstrating how to use a selector to
// click on an element.
package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/cdproto/network"

	"github.com/chromedp/chromedp"
)

func main() {
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		//chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var res string
	err := chromedp.Run(ctx,
		setheaders(
			ctx,
			"http://album.zhenai.com/u/1416899073",
			map[string]interface{}{
				"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36",
			},
			&res,
		))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Res:\n%s", res)
}

// setheaders returns a task list that sets the passed headers.
func setheaders(ctx context.Context, host string, headers map[string]interface{}, res *string) chromedp.Tasks {
	return chromedp.Tasks{
		network.Enable(),
		network.SetExtraHTTPHeaders(network.Headers(headers)),
		chromedp.Navigate(host),
		chromedp.WaitVisible(`#app`, chromedp.ByID),
		//chromedp.ActionFunc(func(ctx context.Context) error {
		//	cookies, err := network.GetAllCookies().Do(ctx)
		//	if err != nil {
		//		return err
		//	}
		//
		//	for i, cookie := range cookies {
		//		log.Printf("chrome cookie %d: %+v", i, cookie)
		//	}
		//
		//	return nil
		//}),

	}
}
