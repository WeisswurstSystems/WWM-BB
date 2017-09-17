package command

import (
	"encoding/json"
	"fmt"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/closemeeting"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/createmeeting"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/putproduct"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/removeproduct"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/setbuyer"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/setplace"
	"github.com/WeisswurstSystems/WWM-BB/meeting/usecase/invite"
)

func ExampleCommandHandler_CreateMeeting() {
	// Send a json Request in this form
	var request createmeeting.Request
	data, _ := json.MarshalIndent(request, "", "  ")
	fmt.Printf("%s", data)

	// Output:
	// {
	//   "meeting": {
	//     "id": "",
	//     "place": "",
	//     "creator": "",
	//     "buyer": "",
	//     "date": "0001-01-01T00:00:00Z",
	//     "closeDate": "0001-01-01T00:00:00Z",
	//     "closed": false,
	//     "orders": null,
	//     "offer": null
	//   },
	//   "login": {
	//     "mail": "",
	//     "password": ""
	//   }
	// }
}

func ExampleCommandHandler_CloseMeeting() {
	// Send a json Request in this form
	var request closemeeting.Request
	data, _ := json.MarshalIndent(request, "", "  ")
	fmt.Printf("%s", data)

	// Output:
	// {
	//   "meetingID": "",
	//   "login": {
	//     "mail": "",
	//     "password": ""
	//   }
	// }
}

func ExampleCommandHandler_PutProduct() {
	// Send a json Request in this form
	var request putproduct.Request
	data, _ := json.MarshalIndent(request, "", "  ")
	fmt.Printf("%s", data)

	// Output:
	// {
	//   "meetingID": "",
	//   "product": {
	//     "name": "",
	//     "price": 0
	//   },
	//   "login": {
	//     "mail": "",
	//     "password": ""
	//   }
	// }
}

func ExampleCommandHandler_RemoveProduct() {
	// Send a json Request in this form
	var request removeproduct.Request
	data, _ := json.MarshalIndent(request, "", "  ")
	fmt.Printf("%s", data)

	// Output:
	// {
	//   "productName": "",
	//   "meetingID": "",
	//   "login": {
	//     "mail": "",
	//     "password": ""
	//   }
	// }
}

func ExampleCommandHandler_SetBuyer() {
	// Send a json Request in this form
	var request setbuyer.Request
	data, _ := json.MarshalIndent(request, "", "  ")
	fmt.Printf("%s", data)

	// Output:
	// {
	//   "buyer": "",
	//   "meetingID": "",
	//   "login": {
	//     "mail": "",
	//     "password": ""
	//   }
	// }
}

func ExampleCommandHandler_SetPlace() {
	// Send a json Request in this form
	var request setplace.Request
	data, _ := json.MarshalIndent(request, "", "  ")
	fmt.Printf("%s", data)

	// Output:
	// {
	//   "place": "",
	//   "meetingID": "",
	//   "login": {
	//     "mail": "",
	//     "password": ""
	//   }
	// }
}

func ExampleCommandHandler_Invite() {
	var request invite.Request
	data, _ := json.MarshalIndent(request, "", "  ")
	fmt.Printf("%s", data)

	// Output:
	// {
	//   "meetingID": "",
	//   "userMails": null,
	//   "login": {
	//     "mail": "",
	//     "password": ""
	//   }
	// }
}