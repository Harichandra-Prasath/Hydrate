package hydrate

import (
	"encoding/json"
	"fmt"
	"os"
)

const ConfigFileName = ".hydrate.json"

// StartTime should be morning and EndTime should be Night
// Assumed StartTime with AM and EndTime with PM
type Config struct {
	StartTime int `json:"start_time"`
	EndTime   int `json:"end_time"`
	Liters    int `json:"liters"`
	Step      int `json:"step"`
	Capacity  int `json:"capacity"`
}

// Parse the Config, Assumed .hydrate.json on home
func ParseConfig() (*Config, error) {

	var Cfg Config

	configPath := os.Getenv("HOME") + ConfigFileName

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &Cfg)
	if err != nil {
		return nil, err
	}

	return &Cfg, nil
}

func ValidateConfig(Cfg *Config) error {

	// Check if the difference between them is less than two hours
	if ((Cfg.EndTime + 12) - Cfg.StartTime) < 2 {
		return fmt.Errorf("cannot schedule for less than 2 hours")
	}

	if Cfg.Liters < 2 {
		return fmt.Errorf("liters is too low")
	}

	if Cfg.Step > Cfg.Liters || (Cfg.Liters)/Cfg.Step < 5.0 {
		return fmt.Errorf("step is too high")
	}

	return nil

}
