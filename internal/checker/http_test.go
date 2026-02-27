package checker

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckLinkOK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	result := CheckLink(server.URL, &http.Client{})
	if !result.OK {
		t.Errorf("expected OK, got status %d, error: %s", result.StatusCode, result.Error)
	}
}

func TestCheckLink404(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	result := CheckLink(server.URL, &http.Client{})
	if result.OK {
		t.Error("expected not OK for 404")
	}
	if result.StatusCode != 404 {
		t.Errorf("status = %d, want 404", result.StatusCode)
	}
}

func TestCheckLinkRedirect(t *testing.T) {
	final := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer final.Close()

	redir := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, final.URL, http.StatusMovedPermanently)
	}))
	defer redir.Close()

	result := CheckLink(redir.URL, &http.Client{})
	if !result.OK {
		t.Errorf("expected OK after following redirect, error: %s", result.Error)
	}
	if !result.Redirected {
		t.Error("expected Redirected = true")
	}
}

func TestCheckLinks(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/bad" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	urls := []string{server.URL + "/good", server.URL + "/bad", server.URL + "/also-good"}
	results := CheckLinks(urls, 2, nil)
	if len(results) != 3 {
		t.Fatalf("results = %d, want 3", len(results))
	}

	for _, r := range results {
		if r.URL == server.URL+"/bad" && r.OK {
			t.Error("expected /bad to not be OK")
		}
		if r.URL == server.URL+"/good" && !r.OK {
			t.Error("expected /good to be OK")
		}
	}
}
