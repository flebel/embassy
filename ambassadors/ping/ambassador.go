package ambassador

const Name = "ping"

func Perform() (int, string, []byte, error) {
	return 200, "plain/text", []byte("Pong!"), nil
}
