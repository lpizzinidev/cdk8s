package jenkinsio


// SecretKeySelector selects a key of a Secret.
type JenkinsSpecNotificationsMailgunApiKeySecretKeySelector struct {
	// The key of the secret to select from.
	//
	// Must be a valid secret key.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The name of the secret in the pod's namespace to select from.
	Secret *JenkinsSpecNotificationsMailgunApiKeySecretKeySelectorSecret `field:"required" json:"secret" yaml:"secret"`
}

