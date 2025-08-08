package constants

const (
	// # source
	SOURCE_WEB_APPLICATION    string = "WEB_APPLICATION"
	SOURCE_MOBILE_APPLICATION string = "MOBILE_APPLICATION"
	SOURCE_WEB_MANAGEMENT     string = "WEB_MANAGEMENT"
)

var AllowedSources = map[string]bool{
	SOURCE_WEB_APPLICATION:    true,
	SOURCE_MOBILE_APPLICATION: true,
	SOURCE_WEB_MANAGEMENT:     true,
}
