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
	"log"
)

func SetPlace(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	result, err := store.FindOne(vars["meetingId"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newPlace := req.URL.Query().Get("price")

	if newPlace == "" {
		http.Error(w, "Missing query parameter: place", http.StatusBadRequest)
		return
	}

	result.Place = newPlace
	err = store.Update(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func SetDate(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	result, err := store.FindOne(vars["meetingId"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	param := req.URL.Query().Get("date")
	newDate, err := time.Parse("2006-01-02T15:04:05Z", param)

	if err != nil || newDate.IsZero() {
		http.Error(w, "Missing parameter: date in form <2006-01-02T15:04:05Z>", http.StatusBadRequest)
		return
	}

	if newDate.Before(time.Now()) {
		http.Error(w, "The date has to be in the future", http.StatusBadRequest)
	}

	result.Date = newDate
	err = store.Update(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func SetBuyer(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	result, err := store.FindOne(vars["meetingId"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newBuyer := req.URL.Query().Get("buyer")

	if newBuyer == "" {
		http.Error(w, "Missing query parameter: buyer", http.StatusBadRequest)
		return
	}

	result.Buyer = newBuyer
	err = store.Update(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func AddProduct(w http.ResponseWriter, req *http.Request) {
	if req.Body == nil {
		http.Error(w, "No body received", http.StatusBadRequest)
		return
	}

	var newProduct product.Product

	err := json.NewDecoder(req.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	errorList := validateProduct(newProduct, 0)

	if len(errorList) != 0 {
		http.Error(w, strings.Join(errorList[:], "\n"), http.StatusBadRequest)
		return
	}

	vars := mux.Vars(req)
	result, err := store.FindOne(vars["meetingId"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	index := util.SliceIndex(len(result.Products), func(i int) bool {
		return result.Products[i].Name == newProduct.Name
	})

	if index != -1 {
		http.Error(w, "Product with name "+newProduct.Name+" already exists in meeting with ID "+result.ID, http.StatusInternalServerError)
		return
	}

	result.Products = append(result.Products, newProduct)
	err = store.Update(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func ChangeProduct(w http.ResponseWriter, req *http.Request) {
	if req.Body == nil {
		http.Error(w, "No body received", http.StatusBadRequest)
		return
	}

	var newProduct product.Product

	err := json.NewDecoder(req.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	errorList := validateProduct(newProduct, 0)

	if len(errorList) != 0 {
		http.Error(w, strings.Join(errorList[:], "\n"), http.StatusBadRequest)
		return
	}

	vars := mux.Vars(req)
	result, err := store.FindOne(vars["meetingId"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	index := util.SliceIndex(len(result.Products), func(i int) bool {
		return result.Products[i].Name == newProduct.Name
	})

	if index == -1 {
		http.Error(w, "No product with name "+newProduct.Name+" is in meeting with ID "+result.ID, http.StatusInternalServerError)
		return
	}

	result.Products[index] = newProduct
	err = store.Update(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func CloseMeeting(w http.ResponseWriter, req *http.Request) {
	setIsMeetingClosed(true, w, req)
}

func OpenMeeting(w http.ResponseWriter, req *http.Request) {
	setIsMeetingClosed(false, w, req)
}

func setIsMeetingClosed(isClosed bool, w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	result, err := store.FindOne(vars["meetingId"])

	log.Printf("Closing meeting %v", vars["meetingId"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result.Closed = isClosed
	err = store.Update(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

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
		http.Error(w, strings.Join(meetingErrors[:], "\n"), http.StatusBadRequest)
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
		errors = append(errors, "Meeting with ID "+meeting.ID+" already exists.")
	}

	return errors
}

func validateProducts(productList []product.Product) ([]string) {
	var errorList []string
	for i, p := range productList {
		errorList = append(errorList, validateProduct(p, i)...)
	}
	return errorList
}

func validateProduct(product product.Product, i int) []string {
	var errorList []string
	if product.Name == "" {
		errorList = append(errorList, "Product "+strconv.Itoa(i)+": Missing field name")
	}
	if product.Price == 0.0 {
		errorList = append(errorList, "Product "+strconv.Itoa(i)+": Missing field price")
	}
	return errorList;
}
