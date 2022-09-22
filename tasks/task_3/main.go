package main

// Task: create Processors for HttpMessage and for WebsocketMessage
// Processors must have method ProcessMessage(msg T) with call of fmt.Println("Message processed")
func main() {
	httpProcessor := HttpProcessor[string, HttpMessage[string]]{}
	msg := HttpMessage[string]{
		content: "qwerty",
	}
	httpProcessor.ProcessMessage(msg)

	websocketProcessor := WebsocketProcessor[[]byte, WebsocketMessage[[]byte]]{}
	msg2 := WebsocketMessage[[]byte]{
		content: []byte("asdfgh"),
	}
	websocketProcessor.ProcessMessage(msg2)
}
