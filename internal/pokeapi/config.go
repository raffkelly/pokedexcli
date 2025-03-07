package pokeapi

type Config struct {
	Next *string
	Prev *string
}

func GetConfig() *Config {
	nextURL := "https://pokeapi.co/api/v2/location-area/"
	return &Config{
		Next: &nextURL,
		Prev: nil,
	}
}
