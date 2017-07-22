package wwm

import (
	"github.com/WeisswurstSystems/WWM-BB/security"
	user "github.com/WeisswurstSystems/WWM-BB/user/driver"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	userRouter := router.PathPrefix("/users").Subrouter()
	user.AddUserRoutes(userRouter)

	// unsecured endpoints
	router.HandleFunc("/meetings", meetingService.ReadAll).Methods("GET")
	router.HandleFunc("/meetings/{meetingId}", meetingService.ReadSingle).Methods("GET")

	router.HandleFunc("/meetings", security.DefAuth(meetingService.Create)).Methods("POST")

	// secured and only meeting owner endpoints
	router.HandleFunc("/meetings/{meetingId}/setPlace", security.DefAuth(security.MeetingAuthenticationHandler(meetingService.SetPlace))).Methods("POST")
	router.HandleFunc("/meetings/{meetingId}/setDate", security.DefAuth(security.MeetingAuthenticationHandler(meetingService.SetDate))).Methods("POST")
	router.HandleFunc("/meetings/{meetingId}/setBuyer", security.DefAuth(security.MeetingAuthenticationHandler(meetingService.SetBuyer))).Methods("POST")
	router.HandleFunc("/meetings/{meetingId}/addProduct", security.DefAuth(security.MeetingAuthenticationHandler(meetingService.AddProduct))).Methods("POST")
	router.HandleFunc("/meetings/{meetingId}/changeProduct", security.DefAuth(security.MeetingAuthenticationHandler(meetingService.ChangeProduct))).Methods("POST")
	router.HandleFunc("/meetings/{meetingId}/closeMeeting", security.DefAuth(security.MeetingAuthenticationHandler(meetingService.CloseMeeting))).Methods("POST")
	router.HandleFunc("/meetings/{meetingId}/openMeeting", security.DefAuth(security.MeetingAuthenticationHandler(meetingService.OpenMeeting))).Methods("POST")

	// Let's go!
	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Printf("Starting on port 8080")
		log.Fatal(http.ListenAndServe(":8080", router))
	} else {
		log.Printf("Starting on port %v", port)
		log.Fatal(http.ListenAndServe(":"+port, router))
	}
}
