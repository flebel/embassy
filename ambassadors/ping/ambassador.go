package ambassador

import (
	"github.com/flebel/embassy/ambassadors"
)

const Name = "ping"

func Perform(amb config.Ambassador) (int, string, []byte, error) {
	return 200, "plain/text", []byte("Pong!"), nil
}
