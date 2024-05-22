package drbot

import "fmt"

var ErrNotInitialized = fmt.Errorf("not initialized")
var ErrNeedChatDataExceptStartingTheChat = fmt.Errorf("need 'ChatData' except starting the chat")
var ErrInvalidMakeRequestChatResponseData = fmt.Errorf("invalid 'MakeRequestChat()' response data")
