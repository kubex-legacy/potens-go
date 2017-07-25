package definition

// AppRelease Application Version
type AppRelease string

//App Types
const (
	//AppReleaseStable stable version
	AppReleaseStable AppRelease = "stable"
	//AppReleaseBeta beta version
	AppReleaseBeta AppRelease = "beta"
	//AppReleaseAlpha alpha version
	AppReleaseAlpha AppRelease = "alpha"
	//AppReleasePreAlpha pre-alpha version
	AppReleasePreAlpha AppRelease = "pre-alpha"
)
