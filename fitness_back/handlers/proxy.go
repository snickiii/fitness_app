package handlers

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

func ProxyHandler(target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		targetURL, err := url.Parse(target)
		if err != nil {
			log.Printf("Error parsing target URL: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		proxyReq, err := http.NewRequest(r.Method, targetURL.String()+r.RequestURI, r.Body)
		if err != nil {
			log.Printf("Error creating proxy request: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		proxyReq.Header = r.Header

		client := &http.Client{}
		resp, err := client.Do(proxyReq)
		if err != nil {
			log.Printf("Error sending proxy request: %v", err)
			http.Error(w, "Failed to connect to the target service", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		for key, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}
		w.WriteHeader(resp.StatusCode)

		if _, err := io.Copy(w, resp.Body); err != nil {
			log.Printf("Error copying response body: %v", err)
			http.Error(w, "Error processing response from target service", http.StatusInternalServerError)
			return
		}
	}
}
