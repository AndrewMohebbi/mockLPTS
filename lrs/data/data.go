package data

import (
	"lrs/structs"
)

// HardCoded returns some hard-coded structs
func HardCoded() (structs.Message, structs.Message) {
	return sheila, george
}

// Hard coded data
var sheilaEvents1 = []structs.Event{
	structs.Event{Section: 1, Completed: true},
	structs.Event{Section: 2, Completed: true},
	structs.Event{Section: 3, Completed: true},
	structs.Event{Section: 4, Completed: false},
}

var sheilaEvents2 = []structs.Event{
	structs.Event{Section: 1, Completed: true},
	structs.Event{Section: 2, Completed: true},
	structs.Event{Section: 3, Completed: false},
	structs.Event{Section: 4, Completed: true},
	structs.Event{Section: 5, Completed: true},
	structs.Event{Section: 6, Completed: true},
}

var sheilaCourses = []structs.Course{
	structs.Course{
		CourseCode: "gettingstarted", CourseName: "MATLAB Onramp",
		Version: "R2019b", Language: "en", TotalSections: 37, Events: sheilaEvents1,
	},
	structs.Course{
		CourseCode: "mlpr", CourseName: "MATLAB Programming Techniques",
		Version: "R2019b", Language: "en", TotalSections: 23, Events: sheilaEvents2,
	},
}

var sheila = structs.Message{ProfileID: "Sheila", Courses: sheilaCourses}

var georgeEvents1 = []structs.Event{
	structs.Event{Section: 1, Completed: true},
	structs.Event{Section: 2, Completed: true},
	structs.Event{Section: 3, Completed: true},
	structs.Event{Section: 4, Completed: true},
	structs.Event{Section: 5, Completed: true},
	structs.Event{Section: 6, Completed: true},
	structs.Event{Section: 7, Completed: true},
}

var georgeEvents2 = []structs.Event{
	structs.Event{Section: 1, Completed: true},
	structs.Event{Section: 2, Completed: true},
	structs.Event{Section: 3, Completed: true},
	structs.Event{Section: 4, Completed: true},
	structs.Event{Section: 5, Completed: true},
	structs.Event{Section: 6, Completed: true},
	structs.Event{Section: 7, Completed: false},
	structs.Event{Section: 8, Completed: true},
}

var georgeCourses = []structs.Course{
	structs.Course{
		CourseCode: "gettingstarted", CourseName: "MATLAB Onramp",
		Version: "R2019b", Language: "en", TotalSections: 37, Events: georgeEvents1,
	},
	structs.Course{
		CourseCode: "simulink", CourseName: "Simulink Onramp",
		Version: "R2019b", Language: "en", TotalSections: 17, Events: georgeEvents2,
	},
}

var george = structs.Message{ProfileID: "George", Courses: georgeCourses}
