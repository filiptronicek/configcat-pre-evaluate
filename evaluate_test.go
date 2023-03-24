package configcatpreevaluate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
	"testing"
)

func TestEvaluate(t *testing.T) {
	// Download the file and unmarshal it
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

	var config interface{}
	err = json.Unmarshal(body, &config)
	if err != nil {
		t.Fatal(err)
	}

	// Create a logger
	logger := logrus.New()

	// Create a configcatpreevaluate.Evaluator
	evaluator := newRolloutEvaluator(logger)

	// Evaluate the config

	sampleUser := NewUserWithAdditionalAttributes("test", "test@example.com", "CZ", map[string]string{})
	result := evaluator.evaluate(config, "isUsageBasedBillingEnabled", sampleUser)

	// Print the result
	fmt.Println(result)
}
