package command

import (
	"encoding/json"
	"fmt"

	"github.com/WeisswurstSystems/WWM-BB/user/usecase/activate"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/changePassword"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/deleteAccount"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/register"
)

func ExampleCommandHandler_Activate() {
	// Send a json Request in this form
	var request activate.Request
	data, _ := json.MarshalIndent(request, "", "  ")
	fmt.Printf("%s", data)

	// Output:
	// {
	//   "registrationID": ""
	// }
}

func ExampleCommandHandler_ChangePassword() {
	// Send a json Request in this form
	var request changePassword.Request
	data, _ := json.MarshalIndent(request, "", "  ")
	fmt.Printf("%s", data)

	// Output:
	// {
	//   "login": {
	//     "mail": "",
	//     "password": ""
	//   },
	//   "password": ""
	// }
}

func ExampleCommandHandler_DeleteAccount() {
	// Send a json Request in this form
	var request deleteAccount.Request
	data, _ := json.MarshalIndent(request, "", "  ")
	fmt.Printf("%s", data)

	// Output:
	// {
	//   "login": {
	//     "mail": "",
	//     "password": ""
	//   }
	// }
}

func ExampleCommandHandler_Register() {
	// Send a json Request in this form
	var request register.Request
	data, _ := json.MarshalIndent(request, "", "  ")
	fmt.Printf("%s", data)

	// Output:
	// {
	//   "mail": "",
	//   "password": "",
	//   "mailEnabled": false
	// }
}
