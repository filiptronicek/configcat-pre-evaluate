package configcatpreevaluate

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"testing"

	"github.com/sirupsen/logrus"
)

func TestEvaluateRule(t *testing.T) {
	url := "https://gitpod.io/configcat/configuration-files/gitpod/config_v5.json"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	logger := logrus.New()

	evaluator := NewRolloutEvaluator(logger)

	sampleUser := NewUserWithAdditionalAttributes("test", "test@example.com", "CZ", map[string]string{})
	result := evaluator.PreEvaluateJson(body, sampleUser)

	fmt.Println(string(result))
}
