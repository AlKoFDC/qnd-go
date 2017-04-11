package icndb

const (
	sfw           = "?exclude=[explicit]"
	randomJokeURL = "https://api.icndb.com/jokes/random" + sfw
)

func GetRandomJoke() (string, error) {
	// You can exchange the order of these operations to use the one or the other mock.
	return getRandomJokeWithInterface(icndbServer{})
	return getRandomJokeWithFunction()
}
