package cd

type Application string

type CDProvider interface {
	GetApplications() []Application
}
