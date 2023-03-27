# ConfigCat pre-evaluate

This Go package provides a simple way to pre-evaluate your feature flags when using [ConfigCat](https://configcat.com/). It transforms your `config_v5.json` file into one that is stripped of all the rules and only contains the values for the flags.

> **Note**: this is heavily based on [configcat/go-sdk](https://github.com/configcat/go-sdk)

## Usage

```go

import (
    "io/ioutil"
    "github.com/sirupsen/logrus"
    "github.com/filiptronicek/configcat-pre-evaluate"
)

func main() {
    // Read the config_v5.json file
    config, err := ioutil.ReadFile("config_v5.json")
    if err != nil {
        panic(err)
    }

	logger := logrus.New()
	evaluator := NewRolloutEvaluator(logger)

    // Pre-evaluate the config
    preEvaluatedConfig, err := evaluator.PreEvaluateJson(config)
    if err != nil {
        panic(err)
    }
}
```
