package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func ReverseProxy(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		parsedURL, err := url.Parse(target)
		if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
			log.Printf("invalid TARGET_SERVICE %q: %v", target, err)
			c.JSON(http.StatusBadGateway, gin.H{
				"error": "invalid upstream target configuration",
			})
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(parsedURL)

		// Catch errors from the proxy itself (e.g. upstream unreachable),
		// so a downed backend returns 502 instead of crashing the gateway.
		proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
			log.Printf("proxy error forwarding to %s: %v", target, err)
			w.WriteHeader(http.StatusBadGateway)
			w.Write([]byte(`{"error": "upstream service unavailable"}`))
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
