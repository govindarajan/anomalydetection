package metrics

import (
	"fmt"
)

// Request sends metrics about all the requests
// resource,application_id,method,status_code
func Request(requestID string, resource string, method string, statusCode int) {
	fmt.Println("METRIC", request, requestID, resource, method, statusCode)
}
