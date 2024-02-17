package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

type Data struct {
	Title string
	Name  string
}
type Nilai struct {
	Title      string
	FinalValue int16
}

func TemplateIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))

	t.ExecuteTemplate(w, "if.gohtml", Data{
		Title: "bangsat",
		Name:  "Nur FIrdaus ",
	})
}

func TestTemplateActionStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000", nil)
	recorder := httptest.NewRecorder()

	TemplateIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateIfOperator(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))

	t.ExecuteTemplate(w, "comparator.gohtml", Nilai{
		Title:      "bangsat",
		FinalValue: 100,
	})
}

func TestTemplateActionOperatorStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000", nil)
	recorder := httptest.NewRecorder()

	TemplateIfOperator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
