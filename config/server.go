package config

type ServerConfig struct {
	clientID  string
	apiSecret string
	senderID  string
}

func Make(clientID, apiSecret, senderID string) *ServerConfig {
	config := ServerConfig{clientID, apiSecret, senderID}
	return &config
}
