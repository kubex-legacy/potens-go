package definition

// UIMode UI Mode Provided By The App
type UIMode string

//UI Modes
const (
	//UIModeNone - No UI Provided
	UIModeNone UIMode = "none"
	//UIModeIntegration - Integrated into pages
	UIModeIntegration UIMode = "integration"
	//UIModeFull - Full UI Application
	UIModeFull UIMode = "full"
)
