package api

import (
	"../auth"
	"../scheduler"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type ScheduleTaskRequest struct {
	Task scheduler.Task `json:"task"`
}
type ScheduleTaskResponse struct {
	Success bool `json:"success"`
}

func scheduleTask(w http.ResponseWriter, r *http.Request) {

	var request ScheduleTaskRequest
	var response ScheduleTaskResponse

	// rudimentary for DDOS prevention
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, DATA_LIMIT))
	if err != nil {
		log.Println("Error: Data exceeds limit", err)
		returnResponse(response, http.StatusBadRequest, w)
		return
	}
	log.Println("body", string(body))

	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Println("ERROR:", err)
		returnResponse(response, http.StatusBadRequest, w)
		return
	}

	//var valid bool
	if _, _, _ = auth.ValidateToken(r.Header.Get("x-api-key")); false {
		log.Println("ERROR:", err)
		returnResponse(response, http.StatusUnauthorized, w)
		return
	}

	log.Println("%+v", request)
	if err := scheduler.HandleTask(request.Task); err != nil {
		log.Println("ERROR:", err)
		returnResponse(response, http.StatusInternalServerError, w)
		return
	} else {
		log.Println("ERROR:", err)
		response.Success = true
		returnResponse(response, http.StatusOK, w)
		return
	}
}
