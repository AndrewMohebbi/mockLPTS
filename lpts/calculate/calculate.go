package calculate

import (
	"encoding/json"
	"lpts/structs"
)

//IDCalculate calculates progress for Profile ID request
func IDCalculate(body []byte) structs.IDMessage {

	var data map[string]interface{}

	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	////////Hard coded return
	var hardProgress = []structs.Progress{
		structs.Progress{
			CourseCode: "gettingstarted", CourseName: "MATLAB Onramp",
			Version: "R2019b", Language: "en", Percentage: 6.2,
		},
		structs.Progress{
			CourseCode: "mlpr", CourseName: "MATLAB Programming Techniques",
			Version: "R2019b", Language: "en", Percentage: 11.7,
		},
	}

	var hard = structs.IDMessage{CourseProgress: hardProgress}

	return hard
}
