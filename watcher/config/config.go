package config

var Config struct {
	GitCheckInterval    int    `env:"BOSUN_GIT_CHECK_INTERVAL,5"`       // git check interval in seconds
	GithubRepositoryUrl string `env:"BOSUN_GITHUB_REPOSITORY_URL"`      // github repository url
	GithubUsername      string `env:"BOSUN_GITHUB_USERNAME,"`           // github username
	GithubPassword      string `env:"BOSUN_GITHUB_PASSWORD,"`           // github password
	KubernetesNamespace string `env:"BOSUN_KUBERNETES_NAMESPACE,bosun"` // target installation namespace of bosun in the cluster
}
