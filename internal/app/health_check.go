package app

import (
	"net/http"
	"os"

	"github.com/kholiqcode/go-todolist/utils"
)

func (s *httpServerImpl) runHealthCheck() {

	s.route.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		host, err := os.Hostname()
		utils.LogAndPanicIfError(err, "failed to get hostname")

		resp := map[string]interface{}{
			"startAt": s.startAt,
			"host":    host,
		}

		utils.GenerateJsonResponse(w, resp, http.StatusOK, "OK")
	})
}
