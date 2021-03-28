package api_registration

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tkanos/gonfig"
)

func Register(file string) {
	fmt.Println("api registration init")
	err := gonfig.GetConf(file, &A)

	if err != nil {
		fmt.Println("error loading api registration")
		fmt.Println(err.Error())
	}
	fmt.Println(A.Name)
	A.HostName = os.Getenv("HOSTNAME")

	if A.HostName == "" {
		fmt.Println("Error getting hostname")
		return
	}

	jsonStr, _ := json.Marshal(A)

	fmt.Println(string(jsonStr))

	ret, _, _ := HTTPRequest(
		"POST",
		A.RegisterAPI,
		jsonStr,
		[]string{})

	fmt.Println(ret)
	fmt.Println("api-registration register")

}
