//go:build no_runtime_type_checking

package jenkinsio

// Building without runtime type checking enabled, so all the below just return nil

func validateJenkinsSpecMasterContainersLivenessProbeTcpSocketPort_FromNumberParameters(value *float64) error {
	return nil
}

func validateJenkinsSpecMasterContainersLivenessProbeTcpSocketPort_FromStringParameters(value *string) error {
	return nil
}

