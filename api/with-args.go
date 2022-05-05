package api

import (
	"io/ioutil"
	"net/http"
	"strings"

	httpClient "github.com/abdfnx/resto/client"
	"github.com/tidwall/gjson"
)

func LatestWithArgs(repo, token string, no_v bool) string {
	url := "https://api.github.com/repos/" + repo + "/releases/latest"

	req, err := http.NewRequest("GET", url, nil)

	if token != "" {
		req.Header.Add("Authorization", "token " + token)
	}

	if err != nil {
		return "Error creating request: " + err.Error()
	}

	client := httpClient.HttpClient()
	res, err := client.Do(req)

	if err != nil {
		return "Error sending request: " + err.Error()
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "Error sending request: " + err.Error()
	}

	tag := gjson.Get(string(b), "tag_name")

	if tag.Exists() {
		if no_v {
			// remove the 'v' char from the tag
			tag_without_v := strings.Replace(tag.String(), "v", "", -1)
			return tag_without_v
		} else {
			return tag.String()
		}
	} else {
		return "no releases found"
	}
}
