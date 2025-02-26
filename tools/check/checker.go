package check

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

func writeFile(path string, lines []string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Failed to write to file:", err)
			return
		}
	}
	writer.Flush()
}

func fetchWithProxy(wg *sync.WaitGroup, proxy string, URL string, results chan<- string) {
	defer wg.Done()

	proxyURL, err := url.Parse(proxy)
	if err != nil {
		fmt.Println("Invalid proxy URL:", proxy)
		return
	}

	transport := &http.Transport{
		Proxy:             http.ProxyURL(proxyURL),
		DisableKeepAlives: true,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 5 * time.Second,
		}).DialContext,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}

	response, err := client.Get(URL)
	if err != nil {
		fmt.Println("Failed to connect via proxy:", proxy, "Error:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		fmt.Println("Working proxy:", proxy)
		results <- proxy
	} else {
		fmt.Println("Bad response from proxy:", proxy, "Status:", response.Status)
	}
}

func CheckProxies() {
	inputPath := "tools/check/input.txt"
	outputPath := "tools/check/output.txt"
	URL := "https://moodle.innopolis.university"

	proxies := readFile(inputPath)
	if proxies == nil {
		return
	}

	var wg sync.WaitGroup
	results := make(chan string, len(proxies))

	for _, proxy := range proxies {
		wg.Add(1)
		go fetchWithProxy(&wg, proxy, URL, results)
	}

	wg.Wait()
	close(results)

	var workingProxies []string
	for proxy := range results {
		workingProxies = append(workingProxies, proxy)
	}

	writeFile(outputPath, workingProxies)
	fmt.Println("Working proxies saved to", outputPath)
}
