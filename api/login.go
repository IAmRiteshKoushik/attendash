package api

import "time"

func SubmitLoginForm(email, pwd string) (bool, error) {
	// structure your API call here
	time.Sleep(5 * time.Second)

	return true, nil
}
