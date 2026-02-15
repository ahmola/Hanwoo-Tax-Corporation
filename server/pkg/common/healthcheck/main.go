// cmd/healthcheck/main.go
package healthcheck

import (
	"net/http"
	"os"
)

func main() {
	// Gin 서버의 헬스체크 주소 (컨테이너 내부이므로 localhost)
	resp, err := http.Get("http://localhost:9090/health")
	if err != nil || resp.StatusCode != http.StatusOK {
		os.Exit(1)
	}
	os.Exit(0)
}
