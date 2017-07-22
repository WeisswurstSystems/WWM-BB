package main

import (
	"github.com/WeisswurstSystems/WWM-BB/mail"
	"github.com/WeisswurstSystems/WWM-BB/meeting/store"
	"github.com/WeisswurstSystems/WWM-BB/security"
	userMiddleware "github.com/WeisswurstSystems/WWM-BB/user/middleware"
	userStore "github.com/WeisswurstSystems/WWM-BB/user/store"
	userUsecase "github.com/WeisswurstSystems/WWM-BB/user/usecase"
	userHandler "github.com/WeisswurstSystems/WWM-BB/user/webhandler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	//entities
	mailService := mail.NewSMTPService()
	userStore := userStore.NewMongoStore()

	//usecase
	userInteractor := userUsecase.Interactor{
		UserStore:   userStore,
		MailService: mailService,
	}

	//handler
	userCommand := userHandler.CommandHandler{
		UserInteractor: userInteractor,
	}
	userQuery := userHandler.QueryHandler{
		UserStore: userStore,
	}
	userMiddleware := userMiddleware.MiddlewareHandler{
		UserStore: userStore,
	}

	router := mux.NewRouter()

	// unsecured endpoints
	router.HandleFunc("/users/do/register", userCommand.Register).Methods("POST")
	router.HandleFunc("/users/do/activate", userCommand.Activate).Methods("POST")
	router.HandleFunc("/meetings", meetingService.ReadAll).Methods("GET")
	router.HandleFunc("/meetings/{meetingId}", meetingService.ReadSingle).Methods("GET")

	// secured endpoints
	router.HandleFunc("/users", userMiddleware.Authenticated(userQuery.Read)).Methods("GET")
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
