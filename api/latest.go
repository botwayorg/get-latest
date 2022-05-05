package api

import (
	"io/ioutil"
	"net/http"
	"strings"

	httpClient "github.com/abdfnx/resto/client"
	"github.com/gorilla/mux"
	"github.com/tidwall/gjson"
)

func Latest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		user := vars["user"]
		repo := vars["repo"]
		token := vars["token"]
		no_v := vars["no-v"]

		url := "https://api.github.com/repos/" + user + "/" + repo + "/releases/latest"

		req, err := http.NewRequest("GET", url, nil)

		if token != "" {
			req.Header.Add("Authorization", "token " + token)
		}

		if err != nil {
			w.Write([]byte("Error creating request: " + err.Error()))
		}

		client := httpClient.HttpClient()
		res, err := client.Do(req)

		if err != nil {
			w.Write([]byte("Error sending request: " + err.Error()))
		}

		defer res.Body.Close()

		b, err := ioutil.ReadAll(res.Body)

		if err != nil {
			w.Write([]byte("Error sending request: " + err.Error()))
		}

		tag := gjson.Get(string(b), "tag_name")

		if tag.Exists() {
			if no_v != "" {
				// remove the 'v' char from the tag
				tag_without_v := strings.Replace(tag.String(), "v", "", -1)
				w.Write([]byte(tag_without_v))
			} else {
				w.Write([]byte(tag.String()))
			}
		} else {
			w.Write([]byte("no releases found"))
		}
	}
}
