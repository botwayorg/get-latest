package api

import (
	"io/ioutil"
	"net/http"

	httpClient "github.com/abdfnx/resto/client"
	"github.com/tidwall/gjson"
)

func Latest(repo string, token string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := "https://api.github.com/repos/" + repo + "/releases/latest"

		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			w.Write([]byte("Error creating request: " + err.Error()))
		}

		if token != "" {
			req.Header.Add("Authorization", "Bearer " + token)
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
			w.Write([]byte(tag.String()))
		} else {
			w.Write([]byte("no releases found"))
		}
	}
}
