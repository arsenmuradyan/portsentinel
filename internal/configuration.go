package internal 

type ArgoCDConfiguration struct {
	Url      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
type Configuration struct {
	Argocd *ArgoCDConfiguration `yaml:"argocd"`
}
