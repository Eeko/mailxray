package finding

func createFinding(message string, location [2]int, severity int) Finding {
	finding := Finding{
		Message:  message,
		Location: location,
		Severity: severity,
	}
	return finding
}

// AdjustLocation (adjustment) can be used as a shortcut on modifying the detected location of a finding.
func (finding *Finding) AdjustLocation(adjustment int) {
	finding.Location = [2]int{finding.Location[0] + adjustment, finding.Location[1] + adjustment}
}

// AdjustLocations runs AdjustLocation to a slice of Findings
func AdjustLocations(findings []Finding, adjustment int) []Finding {
	// this is an ugly way of doing this via a copy of the slice, but not sure if pointer inferno is any better
	var newFind []Finding
	for _, find := range findings {
		newFind = append(newFind,
			createFinding(
				find.Message,
				[2]int{find.Location[0] + adjustment, find.Location[1] + adjustment},
				find.Severity,
			),
		)
		//find.Location = [2]int{find.Location[0] + adjustment, find.Location[1] + adjustment}
		//findings[idx].AdjustLocation(adjustment)
	}
	return newFind
}
