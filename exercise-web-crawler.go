package main

import (
	"errors"
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type SafeFetcher struct {
	fetched map[string]error
	sync.Mutex
}

var safe_fetcher = SafeFetcher{
	fetched: make(map[string]error),
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		fmt.Printf("Fetch %v done. Depth 0.\n", url)
		return
	}

	safe_fetcher.Lock()
	if _, ok := safe_fetcher.fetched[url]; ok {
		safe_fetcher.Unlock()
		fmt.Printf("Fetch %v done. Already fetched.\n", url)
		return
	}
	safe_fetcher.fetched[url] = errors.New("Loading URL ...")
	safe_fetcher.Unlock()

	body, urls, err := fetcher.Fetch(url)

	safe_fetcher.Lock()
	safe_fetcher.fetched[url] = err
	safe_fetcher.Unlock()

	if err != nil {
		fmt.Printf("Error on %v. Error is %v\n", url, err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	q := make(chan bool)

	for i, u := range urls {
		fmt.Printf("Crawling %v/%v of %v: %v\n", i, len(urls), url, u)
		go func(url string) {
			Crawl(url, depth-1, fetcher)
			q <- true
		}(u)
	}

	for i, u := range urls {
		fmt.Printf("[%v] %v/%v Waiting %v\n", url, i, len(urls), u)
		<-q
	}
	fmt.Printf("Done with %v\n", url)
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
	for url, err := range safe_fetcher.fetched {
		if err != nil {
			fmt.Printf("%v failed: %v\n", url, err)
		} else {
			fmt.Printf("%v was fetched\n", url)
		}
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

// $ go run exercise-web-crawler.go
// found: http://golang.org/ "The Go Programming Language"
// Crawling 0/2 of http://golang.org/: http://golang.org/pkg/
// Crawling 1/2 of http://golang.org/: http://golang.org/cmd/
// [http://golang.org/] 0/2 Waiting http://golang.org/pkg/
// found: http://golang.org/pkg/ "Packages"
// Crawling 0/4 of http://golang.org/pkg/: http://golang.org/
// Crawling 1/4 of http://golang.org/pkg/: http://golang.org/cmd/
// Crawling 2/4 of http://golang.org/pkg/: http://golang.org/pkg/fmt/
// Crawling 3/4 of http://golang.org/pkg/: http://golang.org/pkg/os/
// [http://golang.org/pkg/] 0/4 Waiting http://golang.org/
// found: http://golang.org/pkg/os/ "Package os"
// Crawling 0/2 of http://golang.org/pkg/os/: http://golang.org/
// Crawling 1/2 of http://golang.org/pkg/os/: http://golang.org/pkg/
// [http://golang.org/pkg/os/] 0/2 Waiting http://golang.org/
// Fetch http://golang.org/pkg/ done. Already fetched.
// [http://golang.org/pkg/os/] 1/2 Waiting http://golang.org/pkg/
// Error on http://golang.org/cmd/. Error is not found: http://golang.org/cmd/
// [http://golang.org/] 1/2 Waiting http://golang.org/cmd/
// Fetch http://golang.org/ done. Already fetched.
// [http://golang.org/pkg/] 1/4 Waiting http://golang.org/cmd/
// Fetch http://golang.org/cmd/ done. Already fetched.
// [http://golang.org/pkg/] 2/4 Waiting http://golang.org/pkg/fmt/
// found: http://golang.org/pkg/fmt/ "Package fmt"
// Crawling 0/2 of http://golang.org/pkg/fmt/: http://golang.org/
// Fetch http://golang.org/ done. Already fetched.
// Done with http://golang.org/pkg/os/
// [http://golang.org/pkg/] 3/4 Waiting http://golang.org/pkg/os/
// Crawling 1/2 of http://golang.org/pkg/fmt/: http://golang.org/pkg/
// [http://golang.org/pkg/fmt/] 0/2 Waiting http://golang.org/
// Fetch http://golang.org/pkg/ done. Already fetched.
// [http://golang.org/pkg/fmt/] 1/2 Waiting http://golang.org/pkg/
// Fetch http://golang.org/ done. Already fetched.
// Done with http://golang.org/pkg/fmt/
// Done with http://golang.org/pkg/
// Done with http://golang.org/
// http://golang.org/pkg/fmt/ was fetched
// http://golang.org/ was fetched
// http://golang.org/cmd/ failed: not found: http://golang.org/cmd/
// http://golang.org/pkg/ was fetched
// http://golang.org/pkg/os/ was fetched
