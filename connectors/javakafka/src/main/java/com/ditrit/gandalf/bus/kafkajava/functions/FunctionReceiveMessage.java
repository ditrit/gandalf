package com.ditrit.gandalf.bus.kafkajava.functions;

import com.ditrit.gandalf.java.grpc.CommandMessage;
import com.google.gson.Gson;

public class FunctionReceiveMessage {

    public static final String COMMAND = "RECEIVE_MESSAGE";
    private Gson mapper;

    public FunctionReceiveMessage() {
    }

    public String executeCommand(CommandMessage commandMessage) {
        return null;
    }
}
