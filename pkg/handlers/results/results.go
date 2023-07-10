package results

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/imad-almansi/backend-test-golang/pkg/model"
)

func Positive(rw http.ResponseWriter, result []model.Fact) {
	jsonResult, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		Negative(rw, fmt.Errorf("failed to marshal response, error was: %s", err), 500)
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	_, err = rw.Write(jsonResult)
	if err != nil {
		Negative(rw, fmt.Errorf("failed to write response, error was: %s", err), 500)
		return
	}
}

func Negative(rw http.ResponseWriter, err error, statusCode int) {
	log.Printf("error occured: %s", err)

	http.Error(rw, http.StatusText(statusCode), statusCode)
}
