package main

import "fmt"

type HttpProcessor[T any, H HttpMessage[T]] struct {
}

type WebsocketProcessor[T any, W WebsocketMessage[T]] struct {
}

func (h *HttpProcessor[T, H]) ProcessMessage(msg HttpMessage[T]) {
	fmt.Printf("Message processed %s\n", msg)
}

func (w *WebsocketProcessor[T, W]) ProcessMessage(msg WebsocketMessage[T]) {
	fmt.Printf("Message processed %s\n", msg)
}

type Message[T any] interface {
	GetContent() T
}

type HttpMessage[T any] struct {
	content T
}

type WebsocketMessage[T any] struct {
	content T
}

func (hm HttpMessage[T]) GetContent() T {
	return hm.content
}

func (wm WebsocketMessage[T]) GetContent() T {
	return wm.content
}
