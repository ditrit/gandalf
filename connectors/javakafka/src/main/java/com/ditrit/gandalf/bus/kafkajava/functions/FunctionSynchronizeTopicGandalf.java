package com.ditrit.gandalf.bus.kafkajava.functions;


import com.ditrit.gandalf.java.grpc.CommandMessage;
import com.google.gson.Gson;
import org.springframework.kafka.core.KafkaAdmin;

public class FunctionSynchronizeTopicGandalf  {

    public static final String COMMAND = "SYNCHRONIZE_TOPIC_GANDALF";
    private KafkaAdmin kafkaAdmin;
    private Gson mapper;

    public FunctionSynchronizeTopicGandalf() {
    }

    public String executeCommand(CommandMessage commandMessage) {
        return null;
    }
}
