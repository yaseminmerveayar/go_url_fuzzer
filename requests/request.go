package requests

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/common-nighthawk/go-figure"
	"github.com/gookit/color"
	"github.com/valyala/fasthttp"
	"github.com/yaseminmerveayar/fuzzer/config"
)

var (
	wg *sync.WaitGroup = &sync.WaitGroup{}
)

func Execute() {
	myFigure := figure.NewColorFigure("-- FUZZER --", "", "cyan", true)
	myFigure.Print()

	fmt.Printf("\nFuzzing URL : %s \n", config.AppFlag.URL)
	fmt.Printf("Method  : %s \n", config.AppFlag.RequestType)
	fmt.Printf("Status Filter  : %d \n\n", config.AppFlag.StatusHide)
	readFile, err := os.Open(config.AppFlag.Wordlist)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		wg.Add(1)
		go doRequest(fileScanner.Text())
	}
	wg.Wait()
}

func doRequest(fuzzWord string) error {
	rawURL := strings.Replace(config.AppFlag.URL, "FUZZ", fuzzWord, -1)

	defer wg.Done()

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod(config.AppFlag.RequestType)
	req.Header.SetContentType("application/json")
	req.SetRequestURI(rawURL)

	err := fasthttp.Do(req, resp)

	if err != nil {
		return err
	}

	if resp.StatusCode() != config.AppFlag.StatusHide {
		if config.AppFlag.StatusShow == 0 {
			wg.Add(1)
			go outputResult(resp.StatusCode(), rawURL)
		} else {
			if resp.StatusCode() == config.AppFlag.StatusShow {
				wg.Add(1)
				go outputResult(resp.StatusCode(), rawURL)
			}
		}
	}

	return nil
}
func outputResult(responseCode int, URL string) {
	defer wg.Done()

	red := color.FgRed.Render
	green := color.FgGreen.Render
	cyan := color.FgCyan.Render

	fmt.Printf("[ Status: ")
	if responseCode >= 200 && responseCode < 300 {
		fmt.Printf("%s", green(responseCode))
	} else {
		fmt.Printf("%s", red(responseCode))
	}
	fmt.Printf(" \t Request: %-50s ]\n", cyan(URL))

}
