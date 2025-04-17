package hydrate

import (
	"time"
)

func EndlessRun(Cfg *Config, Events []Event) {

	for {
		now := time.Now()
		start := time.Date(now.Year(), now.Month(), now.Day(), Cfg.StartTime-1, 60-Cfg.stepDuration, 0, 0, now.Location())

		// If you are already past your startTime, wait for tomorrow
		if now.After(start) {
			start = start.Add(24 * time.Hour)
		}

		// Sleep until then
		time.Sleep(time.Until(start))

		iterateEvents(Cfg.stepDuration, Events)
	}

}

func iterateEvents(stepDuration int, Events []Event) {

	ticker := time.NewTicker(time.Duration(stepDuration) * time.Minute)
	i := 0

	for {

		select {

		case <-ticker.C:
			ProcessEvent(Events[i])
			i += 1
		}

	}

}
