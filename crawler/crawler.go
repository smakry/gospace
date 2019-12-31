package crawler

import (
	"context"
	"errors"
	// "github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"time"
	"log"
	"github.com/chromedp/chromedp"
	"fmt"
	"os"
	"bufio"
)

var res string

func DoCrawlerTask(url string) {
	var err error

	// create link
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create new chrome instance
	c, _ := chromedp.NewContext(ctxt, chromedp.WithLogf(log.Printf))
	err = chromedp.Run(c, visitWeb(url))
	if err != nil {
		log.Fatal(err)
	}

	x, _ := chromedp.NewContext(ctxt, chromedp.WithLogf(log.Printf))
	fmt.Println("chromedp.NewContext!")
	err = chromedp.Run(x, DoCrawler())
	fmt.Println("chromedp.Run!")
	if err != nil {
		log.Fatal(err)
	}
	WirteTXT(res)
}

// set cookie, get account login page
func visitWeb(url string) chromedp.Tasks {
	tmpfun := func(ctxt context.Context) error {
		// expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
		success, err := network.SetCookie("ASP.NET_SessionId", "this is value"). // set cookie
			// WithExpires(&expr).
			WithDomain("http://forum.nderp.99.com/BugManage/Internal/bugOperationList.htm?taskId=100600"). //view web main
			WithHTTPOnly(true).
			Do(ctxt)

		if err != nil {
			return err
		}
		if !success {
			return errors.New("could not set cookie")
		}

		return nil	
	}
	fmt.Println("visitWeb is begin!")
	return chromedp.Tasks{
		chromedp.ActionFunc(tmpfun),
		chromedp.Navigate(url), //page turns
	}
}

// turn page, get html
func DoCrawler() chromedp.Tasks {
	fmt.Println("DoCrawler is begin!")
	defer func(){
		fmt.Println("DoCrawler is success!")
	}()
	return chromedp.Tasks{
		chromedp.Sleep(1 * time.Second),
		//wait "id=from1" page is view,  ByQuery is user DOM select query
		chromedp.WaitVisible(`#form1`, chromedp.ByQuery),
		chromedp.Sleep(1 * time.Second),
		chromedp.Click(`.pagination li:nth-last-child(4) a`, chromedp.ByQuery), //click turn page
		chromedp.OuterHTML(`tbody`, &res, chromedp.ByQuery), // get tbody label html
	}
}

func WirteTXT(txt string) {
	f, err := os.OpenFile("result.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("os create error: ", err)
		return
	}
	defer f.Close()

	bw := bufio.NewWriter(f)
	bw.WriteString(txt + "\n")
	bw.Flush()
}












