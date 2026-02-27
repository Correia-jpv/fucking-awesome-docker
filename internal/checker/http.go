package checker

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/veggiemonk/awesome-docker/internal/cache"
)

const (
	defaultTimeout     = 30 * time.Second
	defaultConcurrency = 10
	userAgent          = "awesome-docker-checker/1.0"
)

// LinkResult holds the result of checking a single URL.
type LinkResult struct {
	URL         string
	OK          bool
	StatusCode  int
	Redirected  bool
	RedirectURL string
	Error       string
}

// CheckLink checks a single URL. Uses HEAD first, falls back to GET.
func CheckLink(url string, client *http.Client) LinkResult {
	result := LinkResult{URL: url}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	// Try HEAD first
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, url, nil)
	if err != nil {
		result.Error = err.Error()
		return result
	}
	req.Header.Set("User-Agent", userAgent)

	// Track redirects
	var finalURL string
	origCheckRedirect := client.CheckRedirect
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		finalURL = req.URL.String()
		if len(via) >= 10 {
			return http.ErrUseLastResponse
		}
		return nil
	}
	defer func() { client.CheckRedirect = origCheckRedirect }()

	resp, err := client.Do(req)
	if err != nil {
		// Fallback to GET
		req, err2 := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err2 != nil {
			result.Error = err.Error()
			return result
		}
		req.Header.Set("User-Agent", userAgent)
		resp, err = client.Do(req)
		if err != nil {
			result.Error = err.Error()
			return result
		}
	}
	defer resp.Body.Close()

	result.StatusCode = resp.StatusCode
	result.OK = resp.StatusCode >= 200 && resp.StatusCode < 400

	if finalURL != "" && finalURL != url {
		result.Redirected = true
		result.RedirectURL = finalURL
	}

	return result
}

// CheckLinks checks multiple URLs concurrently.
func CheckLinks(urls []string, concurrency int, exclude *cache.ExcludeList) []LinkResult {
	if concurrency <= 0 {
		concurrency = defaultConcurrency
	}

	results := make([]LinkResult, len(urls))
	sem := make(chan struct{}, concurrency)
	var wg sync.WaitGroup

	for i, url := range urls {
		if exclude != nil && exclude.IsExcluded(url) {
			results[i] = LinkResult{URL: url, OK: true}
			continue
		}

		wg.Add(1)
		go func(idx int, u string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			client := &http.Client{Timeout: defaultTimeout}
			results[idx] = CheckLink(u, client)
		}(i, url)
	}

	wg.Wait()
	return results
}
