package hydrate

type Event struct {
	Step      int
	Remaining int
}

func ScheduleEvents(Cfg *Config) ([]Event, error) {

	totalSteps := (Cfg.milliLiters) / (Cfg.Step)

	// Check the extra water
	unallocated := Cfg.milliLiters - totalSteps*Cfg.Step

	// total  hours
	totalHours := Cfg.EndTime + 12 - Cfg.StartTime

	// We have to allocate totalSteps in totalHours
	totalMinutes := totalHours * 60

	stepDuration := totalMinutes / totalSteps

	Cfg.stepDuration = stepDuration

	// For every stepDuration, we have to notify the user starting from the endtime
	var events []Event

	unallocatedStep := unallocated / (totalSteps / 2)

	i := 0
	rem := Cfg.milliLiters

	for step := 0; step < totalSteps; step++ {

		e := Event{
			Step:      Cfg.Step,
			Remaining: rem - Cfg.Step,
		}

		if step%2 == 0 && i < (totalSteps/2) {
			e.Step += unallocatedStep
			e.Remaining -= unallocatedStep
			i += 1
		}

		events = append(events, e)
		rem = rem - e.Step
	}

	return events, nil
}
