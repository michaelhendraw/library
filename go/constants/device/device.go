package device

const (
	All       = 0
	Web       = 1
	MobileWeb = 2
	Android   = 3
	Ios       = 4
)

var (
	XDeviceList = map[int]string{
		All:       "all",
		Web:       "web",
		MobileWeb: "mobile_web",
		Android:   "android",
		Ios:       "ios",
	}

	XDeviceMapping = map[string]int{
		"all":        All,
		"web":        Web,
		"mobile_web": MobileWeb,
		"android":    Android,
		"ios":        Ios,
	}
)
