package wasm

import (
	"context"
	"fmt"
	"syscall/js"

	"github.com/hown3d/meme-des-tages-cli/reddit"
)

func topPicture(client reddit.Client, waitChan chan struct{}) js.Func {
	f := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		defer func() {
			waitChan <- struct{}{}
		}()
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		subredditVal := args[0].String()
		uri, err := client.GetTopPost(context.Background(), subredditVal)
		if err != nil {
			fmt.Printf("unable to get top post: %w", err)
			return err.Error()
		}
		return uri
	})
	return f

}

func Register(c chan struct{}) {
	js.Global().Set("topPicture", topPicture(reddit.NewClient(), c))
}
