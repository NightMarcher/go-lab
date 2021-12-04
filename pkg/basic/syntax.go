package basic

import (
	"errors"
	"fmt"
)

type statusCode struct {
	code   int
	reason string
}

func fakeRequest(sc statusCode) (string, error) {
	if sc.code >= 400 {
		return sc.reason, errors.New(fmt.Sprintf("Fake Request Error => code: %d, reason: %s", sc.code, sc.reason))
	}
	return sc.reason, nil
}

// error
func DemoError() {
	statusCodes := []statusCode{
		{200, "OK"},
		{204, "No Content"},
		{302, "Found"},
		{400, "Bad Request"},
		{403, "Forbidden"},
		{404, "Not Found"},
		{500, "Internal Server Error"},
		{502, "Bad Gateway"},
		{503, "Service Unavailable"},
		{504, "Gateway Timeout"},
	}
	fmt.Println()
	for _, sc := range statusCodes {
		if reason, error := fakeRequest(sc); error == nil {
			fmt.Println(reason)
		} else {
			fmt.Println(error)
		}
	}
}
