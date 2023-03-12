package jumpserver

const (
	api_users   = "/api/v1/users/users/"
	api_assets  = "/api/v1/assets/assets/"
	api_nodes   = "/api/v1/assets/nodes/"
	api_metrics = "/api/v1/prometheus/metrics"

	gmtFmt    = "Mon, 02 Jan 2006 15:04:05 GMT"
	userAgent = "go-jumpserver"
)

type Platform string

const (
	Linux      Platform = "Linux"
	MacOS      Platform = "MacOS"
	BSD        Platform = "BSD"
	Windows    Platform = "Windows"
	Other      Platform = "Other"
	WindowsRDP Platform = "WindowsRDP"
	WindowsTLS Platform = "WindowsTLS"
	AIX        Platform = "AIX"
)

//type Platform int
//
//const (
//    Linux Platform = iota
//    MacOS
//    BSD
//    Windows
//    Other
//    WindowsRDP
//    WindowsTLS
//    AIX
//)
//
//var platformMap = map[Platform]string{
//    Linux:      "Linux",
//    MacOS:      "MacOS",
//    BSD:        "BSD",
//    Windows:    "Windows",
//    Other:      "Other",
//    WindowsRDP: "WindowsRDP",
//    WindowsTLS: "WindowsTLS",
//}
//
//func (p Platform) String() string {
//    val, ok := platformMap[p]
//    if ok {
//        return val
//    }
//    return "Unknown"
//}
