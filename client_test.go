package deploygate

func setUpConfigForTest() ClientConfig {
	endpoint := "https://deploygate.com/api"
	apiKey := "dummy_api_key"

	return ClientConfig{
		apiKey:      apiKey,
		apiEndpoint: &endpoint,
	}
}
