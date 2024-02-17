package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "hello")
	} else {
		fmt.Fprintf(w, "hello %s", name)
	}

}

func TestQueryParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Joko", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleParam(w http.ResponseWriter, r *http.Request) {

	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")
	age := r.URL.Query().Get("age")

	fmt.Fprintf(w, "Hello %s %s %s", firstName, lastName, age)

}

func TestQueryMultiParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?last_name=firdaus&first_name=nur&&age=20", nil)
	recorder := httptest.NewRecorder()

	MultipleParam(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleParameterValue(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]

	fmt.Fprintln(w, strings.Join(names, " "))
}

func TestQueryMultiParamValue(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=nur&name=nadiva&name=jasmine", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValue(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
