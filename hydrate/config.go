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
	StartTime   int     `json:"start_time"`
	EndTime     int     `json:"end_time"`
	Liters      float32 `json:"liters"`
	milliLiters int
	Step        int `json:"step"`
}

// Parse the Config, Assumed .hydrate.json on home
func ParseConfig() (*Config, error) {

	var Cfg Config

	configPath := os.Getenv("HOME") + "/" + ConfigFileName

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

	// Convert to MilliLiters
	mL := Cfg.Liters * 1000

	Cfg.milliLiters = int(mL)

	if Cfg.Step > Cfg.milliLiters || (Cfg.milliLiters)/Cfg.Step < 5.0 {
		return fmt.Errorf("step is too high")
	}

	return nil

}
