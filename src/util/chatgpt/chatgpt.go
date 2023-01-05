package chatgpt

import (
	"fatsharkbot/src/config"
	"github.com/chatgp/chatgpt-go"
	"log"
	"net/http"
	"time"
)

func IndependentConversation(message string) string {
	// new chatgpt client
	token := config.C.ChatGpt.Token
	cfValue := config.C.ChatGpt.Cfvalue

	cookies := []*http.Cookie{
		{
			Name:  "__Secure-next-auth.session-token",
			Value: token,
		},
		{
			Name:  "cf_clearance",
			Value: cfValue,
		},
	}

	cli := chatgpt.NewClient(
		chatgpt.WithDebug(true),
		chatgpt.WithTimeout(60*time.Second),
		chatgpt.WithCookies(cookies),
	)

	// chat in independent conversation
	text, err := cli.GetChatText(message)
	if err != nil {
		log.Println(err)
	}
	return text.Content
}
