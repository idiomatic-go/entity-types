package client

// Configuration - stores client configuration, such as :
//   retention periods for log and other data
//   notifications for entity data changes, one for each entity
//   notification services for streaming
//   contact information
type Configuration struct {
	Version string
}
