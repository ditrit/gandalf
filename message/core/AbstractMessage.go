package connector

import "nanomsg.org/go/mangos/v2"

type AbstractMessage struct {
	uuid  string
	topic string

	//TODO SUITE
	/* 	private Map<String, String> routing;
	private Map<String, String> access;
	private Map<String, String> info;
	protected ObjectMapper objectMapper; */
}

func (a AbstractMessage) sendWith(socket mangos.Socket, routingInfo string)
