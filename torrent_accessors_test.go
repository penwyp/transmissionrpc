package transmissionrpc

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

func Test_xx(t *testing.T) {

	bodys, err := os.ReadFile("/Users/pen.wyp/Downloads/xxx.log")
	if err != nil {
		log.Fatal(err)
	}

	var result torrentGetResults
	var items answerPayload
	items.Arguments = &result

	uErr := json.Unmarshal(bodys, &items)
	if uErr != nil {
		log.Fatal(uErr)
	}
	log.Println(len(items.Result))
}
