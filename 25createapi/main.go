package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Model for courses - file
type Course struct {
	CourseId string     `json:"courseid"`
	CourseName string   `json:"coursename"`
	CoursePrice int 	`json:"price"`
	Author *Author    	`json:"author"`
}
type Author struct {
	Fullname string  `json:"fullname"`
	Website string 	 `json:"website"`
}
//fake DB

var courses []Course

//middleware , helper - file
// Define IsEmpty as a method of Course
func (c Course) IsEmpty() bool {
    return c.CourseName == ""
}

func main() {
	fmt.Println("Api - Learn golang online !!")
	r := mux.NewRouter()
	//seeding
	courses = append(courses, Course{
		CourseId: "2",
		CourseName: "Reactjs",
		CoursePrice: 299,
		Author: &Author{Fullname: "Pratim K Hazarika",Website: "prratim.com"},
	})
	courses = append(courses, Course{
		CourseId: "4",
		CourseName: "Mern",
		CoursePrice: 199,
		Author: &Author{Fullname: "Mark zuck",Website: "go.com"},
	})
	///routing
	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/courses",getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}",getOneCourse).Methods("GET")
	r.HandleFunc("/course",createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}",updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}",deleteOneCourse).Methods("DELETE")
	//listen to a port
	log.Fatal(http.ListenAndServe(":4000",r))
}

//controllers - file

//serve home route
///writer - where you write the response
//reader - where we get the value from who ever is sending the request
func serveHome( w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcpme to api by Aws</h1>"))
}

func getAllCourses( w http.ResponseWriter, r *http.Request){
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request){
	 fmt.Println("Get one course")
	 w.Header().Set("Content-Type","application/json")
	 //grab id from request
	 params := mux.Vars(r)
	 courseID := params["id"]
	 //loop through courses , find matching id and return the response

	 for _, course := range courses {
		if course.CourseId == courseID {
			json.NewEncoder(w).Encode(course)
			return
		}
	 }
	  json.NewEncoder(w).Encode(map[string]string{
		"error":"No course found with given id",
		"courseId":courseID,
	  })
	 
}

func createOneCourse(w http.ResponseWriter, r *http.Request){
    fmt.Println("Create one course")
	w.Header().Set("Content-Type","application/json")

	//what id: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	//what about  - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No data inside Json")
		return
	}
	for _,val := range courses {
		if val.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Same course name bro")
			return;
	
		}
	}
	///generate unique id,convert to string

	rand.Seed(time.Now().UnixNano())
	course.CourseId =strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Update one course")
	w.Header().Set("Content-Type","application/json")
	//first -grap id from request
	params := mux.Vars(r)
	courseId := params["id"]

	//loop ,id , remove , add with my id

	for index, course := range courses{
		if course.CourseId == courseId {
			courses = append(courses[:index],courses[index+1:]... )
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = courseId;
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}else{
			json.NewEncoder(w).Encode(map[string]string{
				"error":"No course found with given id",
				"courseId":courseId,
			  })
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type","application/json")
	//first -grap id from request
	params := mux.Vars(r)
	courseId := params["id"]

	for index , course := range courses {
		if course.CourseId == courseId{
			courses = append(courses[:index],courses[index+1:]... )
			json.NewEncoder(w).Encode("Successfully Deleted")
			break;
		}
	}
}
