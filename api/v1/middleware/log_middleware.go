package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/TanmoyTSSaha/GoBase/pkg/services"
)

func LoggingMiddleware(logService *services.LogService) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// LOGGING LOGIC
			// Start time
			start := time.Now()

			// Logging request details
			log.Printf("STARTED %s METHOD AT %s", r.Method, r.URL.Path)

			// Creating a response writer to capture the code
			rw := &responseWriter{w, http.StatusOK}

			// CALL THE NEXT HANDLER
			next.ServeHTTP(rw, r)



			duration := time.Since(start)
			message := fmt.Sprintf("COMPLETED %s METHOD AT %s WITH STATUS CODE %d IN %v", r.Method, r.URL.Path, rw.statusCode, duration)
			// INITIATING LOG STORING PROCESS || FOR DEVELOPMENT PURPOSE I HAVE STOPPED THE LOG STORING PROCESS
			// log_details := services.LogStruct{URL: r.URL.Path, Timestamp: start, Method: r.Method, StatusCode: rw.statusCode, ResponseDuration: duration, Message: message}
			// if err := logService.StoreLogs(log_details); err != nil {
			// 	log.Printf("FAILED TO STORE LOG: %v", err)
			// }

			// LOGGING RESPONSE DETAILS AND DURATION
			log.Printf(message)
		})
	}
}

// IT IS A WRAPPER TO CAPTURE THE STATUS CODE
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader IS OVERRIDDEN TO CAPTURE THE STATUS CODE
func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader((statusCode))
}
