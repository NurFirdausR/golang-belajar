package golang_web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-NFR-name"
	cookie.Value = r.URL.Query().Get("name")

	http.SetCookie(w, cookie)
	fmt.Fprintf(w, "Success set cookie")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("X-NFR-name")
	if cookie.Value == "" {
		fmt.Fprintf(w, "No Cookie")
	} else {
		name := cookie.Value
		fmt.Fprintf(w, "Cookie %s", name)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/?name=Nur", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)
	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookie %s:%s \n", cookie.Name, cookie.Value)
	}
}
