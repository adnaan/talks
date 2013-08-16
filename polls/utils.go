package lokyantra

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var (
	host        = "http://localhost:9947"
	contentType = "application/json"
)

//TESTSERVERINIT OMIT
func testServer(f func(s *Server)) {

	server := NewServer()
	server.ListenAndServe()
	defer server.Shutdown()
	f(server)
}

//TESTSERVEREND OMIT

//TESTHTTPINIT OMIT
func testHttpRequest(verb string, resource string, body string) (*http.Response, error) {
	client := &http.Client{Transport: &http.Transport{DisableKeepAlives: true}}
	r, _ := http.NewRequest(verb, fmt.Sprintf("%s%s", host, resource), strings.NewReader(body))
	r.Header.Add("Content-Type", contentType)
	return client.Do(r)
}

//TESTHTTPEND OMIT

//READJSONINIT OMIT
func readJson(d interface{}, r *http.Request, w http.ResponseWriter) error {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {

		http.Error(w, "Bad Data!", http.StatusBadRequest)
		return err
	}

	return json.Unmarshal(body, &d)

}

//READJSONEND OMIT

//SERVEJSONINIT OMIT

func serveJson(w http.ResponseWriter, v interface{}) {
	content, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Length", strconv.Itoa(len(content)))
	w.Header().Set("Content-Type", "application/json")
	w.Write(content)

}

//SERVEJSONEND OMIT
