package definition

// AppType Application Type
type AppType string

//App Types
const (
	//AppTypeBackground Background Application
	AppTypeBackground AppType = "kubex.application.background"

	// AppTypeProject Standard Application
	AppTypeProject AppType = "kubex.application.project"

	// AppTypeUser Standard User Application (Sidebar)
	AppTypeUser AppType = "kubex.application.user"

	//AppTypeGroup App Group
	AppTypeGroup AppType = "kubex.group"

	//AppTypeService Service Only - internal only
	AppTypeService AppType = "kubex.service"
)
