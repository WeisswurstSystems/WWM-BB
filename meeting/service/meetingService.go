package service

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/WeisswurstSystems/WWM-BB/meeting/store"
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/WeisswurstSystems/WWM-BB/product"
	"github.com/WeisswurstSystems/WWM-BB/security"
	"github.com/WeisswurstSystems/WWM-BB/util"
	"strings"
	"strconv"
	"github.com/gorilla/mux"
)

func ReadAll(w http.ResponseWriter, req *http.Request) {
	results, err := store.FindAllReduced()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var js []byte
	js, err = json.Marshal(results)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func ReadSingle(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	result, err := store.FindOne(vars["meetingId"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var js []byte
	js, err = json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func Create(w http.ResponseWriter, req *http.Request) {
	if req.Body == nil {
		http.Error(w, "No body received", http.StatusBadRequest)
		return
	}

	var meeting meeting.Meeting

	err := json.NewDecoder(req.Body).Decode(&meeting)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	meetingErrors := validateMeeting(meeting, req)
	if len(meetingErrors) != 0 {
		http.Error(w, strings.Join(meetingErrors[:],"\n"), http.StatusBadRequest)
		return
	}

	meeting.ID = util.GetUID(12)
	meeting, err = store.Create(meeting)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(meeting)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(js)
}

func validateMeeting(meeting meeting.Meeting, req *http.Request) ([]string) {
	var errors []string
	 hasCreator, hasDate, hasProducts := true, true, true

	if meeting.Place == "" {
		errors = append(errors, "Missing field: place")
	}

	if meeting.Creator == "" {
		errors = append(errors, "Missing field: creator")
		hasCreator = false
	}

	if meeting.Date.IsZero() {
		errors = append(errors, "Missing field: date")
		hasDate = false
	}

	if hasCreator && meeting.Creator != security.GetCurrentUser(req) {
		errors = append(errors, "Field creator must equal current user")
	}

	if hasDate && meeting.Date.Before(time.Now()) {
		errors = append(errors, "Invalid Date: Must be in the future")
	}

	if len(meeting.Products) == 0 {
		errors = append(errors, "Missing data: at least one product is needed")
		hasProducts = false;
	}

	if hasProducts {
		productErrors := validateProducts(meeting.Products)
		if len(productErrors) != 0 {
			errors = append(errors, productErrors...)
		}
	}

	has, _ := store.Has(meeting.ID)
	if has {
		errors = append(errors,"Meeting with ID " + meeting.ID + " already exists.")
	}

	return errors
}

func validateProducts(productList []product.Product) ([]string) {
	var errorList []string
	for i, p := range productList {
		if(p.Name == "") {
			errorList = append(errorList, "Element " + strconv.Itoa(i) + ": Missing field name")
		}
		if(p.Price == 0.0){
			errorList = append(errorList, "Element " + strconv.Itoa(i) + ": Missing field price")
		}
	}
	return errorList
}
