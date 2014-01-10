package polls

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

var question_id = ""

var question = `{"text":"Which is the best IPL team?","options":[{"option_id":"1","name":"RCB"},{"option_id":"2","name":"KKR"},{"option_id":"3","name":"CSK"},{"option_id":"4","name":"MI"}]}`

func TestCreateQuestion(t *testing.T) {
	testServer(func(s *Server) {
		//CQINIT OMIT
		res, err := testHttpRequest("POST", "/api/v1/q/", question)
		if err != nil {
			t.Fatalf("Unable to create question: %v", err)

		} else {
			//CQEND OMIT

			body, _ := ioutil.ReadAll(res.Body)
			defer res.Body.Close()

			if res.StatusCode == 200 {
				fmt.Println(string(body))
			}

			var response map[string]interface{}
			err := json.Unmarshal(body, &response)
			if err != nil {
				t.Fatalf("fail to parse body: %v", string(body))
			}

			question_id = response["id"].(string)

			fmt.Printf("question_id: %v\n", question_id)

		}

	})

}

func TestVote(t *testing.T) {
	testServer(func(s *Server) {
		//VINIT OMIT
		api := "/api/v1/q/" + question_id

		res, err := testHttpRequest("PUT", api, `{"option":"1"}`)
		if err != nil {
			t.Fatalf("Unable to vote: %v", err)

		} else {
			//VEND OMIT

			body, _ := ioutil.ReadAll(res.Body)
			defer res.Body.Close()

			if res.StatusCode == 200 {
				fmt.Println(string(body))
			}

		}

	})

}

func TestGetQuestion(t *testing.T) {
	testServer(func(s *Server) {
		//GINIT OMIT
		api := "/api/v1/q/" + question_id

		res, err := testHttpRequest("GET", api, ``)
		if err != nil {
			t.Fatalf("Unable to get: %v", err)

		} else {
			//GEND OMIT

			body, _ := ioutil.ReadAll(res.Body)
			defer res.Body.Close()

			if res.StatusCode == 200 {
				fmt.Println(string(body))
			}

		}

	})

}

func TestDeleteQuestion(t *testing.T) {
	testServer(func(s *Server) {
		//DINIT OMIT
		api := "/api/v1/q/" + question_id

		res, err := testHttpRequest("DELETE", api, ``)
		if err != nil {
			t.Fatalf("Unable to delete: %v", err)

		} else {
			//DEND OMIT

			body, _ := ioutil.ReadAll(res.Body)
			defer res.Body.Close()

			if res.StatusCode == 200 {
				fmt.Println(string(body))
			}

		}

	})

}
