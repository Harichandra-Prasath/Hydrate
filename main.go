package main

import (
	"github.com/Harichandra-Prasath/Hydrate/hydrate"
)

func main() {

	Cfg, err := hydrate.ParseConfig()
	if err != nil {
		panic("error in parsing config: " + err.Error())
	}

	err = hydrate.ValidateConfig(Cfg)
	if err != nil {
		panic("error in validation: " + err.Error())

	}

	_, err = hydrate.ScheduleEvents(Cfg)
	if err != nil {
		panic("error in scheduling events: " + err.Error())
	}

}
