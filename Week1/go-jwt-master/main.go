package main 

import (
	"fmt"
	"net/http"
	"go-jwt/handler"
	"go-jwt/config"
	"go-jwt/driver"
)

func main() {
	driver.ConnectMongoDB(config.DB_USER, config.DB_PASS)
	
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/user", handler.GetUser)

	fmt.Println("Server running [:8000]")
	http.ListenAndServe(":8000", nil)
}


// TestFindAPersonWrongTypeId test case id exists
func TestFindAPersonExists(t *testing.T){
	t.Parallel()

	r, _ := http.NewRequest("GET", "/api/person", nil)
	id := "60d555b6ba8b0ead3bb0596d"
	//Hack to try to fake gorilla/mux vars
	vars := map[string]string{
		"id": id,
	}
	r = mux.SetURLVars(r, vars)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.FindAPerson)
	handler.ServeHTTP(rr, r)
	// check httpStatus
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	docId, errCheckId := primitive.ObjectIDFromHex(id)

	if errCheckId != nil{
		t.Error("handler returned wrong ")
	}
	expected := model.Person{
		ID: docId,
		Name: "quang",
	}
	responseData := model.Person{}

	//parse repsonse body to struct Error
	if err := json.NewDecoder(rr.Body).Decode(&responseData); err != nil {
		log.Fatalln(err)
	}

	//compare expected and reponse data
	if expected.ID != responseData.ID || expected.Name != responseData.Name{
		t.Errorf("handler returned unexpected status: got %v want %v",
			responseData, expected)
	}
}


// TestFindAPersonWrongTypeId test case id is not ObjectId
func TestFindAPersonWrongTypeId(t *testing.T){
	t.Parallel()

	r, _ := http.NewRequest("GET", "/api/person", nil)

	//Hack to try to fake gorilla/mux vars
	vars := map[string]string{
		"id": "1",
	}
	r = mux.SetURLVars(r, vars)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.FindAPerson)
	handler.ServeHTTP(rr, r)
	// check httpStatus
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := model.Error{
		Status: 400,
		Message: "Bad Request",
	}
	responseData := model.Error{}

	//parse repsonse body to struct Error
	if err := json.NewDecoder(rr.Body).Decode(&responseData); err != nil {
		log.Fatalln(err)
	}

	//compare expected and reponse data
	if expected.Status != responseData.Status{
		t.Errorf("handler returned unexpected status: got %v want %v",
			responseData, expected.Status)
	}
}