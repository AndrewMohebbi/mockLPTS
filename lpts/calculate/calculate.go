// Package calculate contains the functions which calculate progress
package calculate

import (
	"encoding/json"
	"log"

	"lpts/structs"
)

//IDCalculate calculates progress for Profile ID request
func IDCalculate(body []byte) structs.AllMessage {

	var data map[string]interface{}

	if err := json.Unmarshal(body, &data); err != nil {
		log.Println("IDCalculate: unmarshalling failed")
	}

	////////Hard coded return
	var hardProgress = []structs.Course{
		structs.Course{
			CourseCode: "gettingstarted", CourseName: "MATLAB Onramp",
			Version: "R2019b", Language: "en", Percentage: 6.2,
		},
		structs.Course{
			CourseCode: "mlpr", CourseName: "MATLAB Programming Techniques",
			Version: "R2019b", Language: "en", Percentage: 11.7,
		},
	}

	var hard = structs.AllMessage{CourseProgress: hardProgress}

	return hard
}
