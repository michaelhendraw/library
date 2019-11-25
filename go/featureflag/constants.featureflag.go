package featureflag

// List of instances
const (
	InstanceApp Instance = iota
)

// List of flag statuses
const (
	StatusDisabled Status = iota
	StatusQA
	StatusSpecific
	StatusAll
	StatusInPercentageUser
)

var (
	instanceTypeMap = map[Instance]string{
		InstanceApp: "App",
	}

	// statusvalue is status in multivalue row, 1 is active and 0 is inactive
	statusValue = map[string]bool{
		"1": true,
		"0": false,
	}
)
