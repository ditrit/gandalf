package com.ditrit.gandalf.bus.kafkajava.functions;

import com.ditrit.gandalf.bus.kafkajava.core.producer.KafkaJavaProducer;
import com.ditrit.gandalf.java.grpc.CommandMessage;
import com.google.gson.Gson;
import com.google.gson.JsonObject;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class FunctionSendMessage {

    public static final String COMMAND = "SEND_MESSAGE";
    private Gson mapper;

    @Autowired
    private KafkaJavaProducer kafkaJavaProducer;

    public FunctionSendMessage() {
    }

    public String executeCommand(CommandMessage commandMessage) {
        String payload = commandMessage.getPayload();
        JsonObject jsonObject = mapper.fromJson(payload, JsonObject.class);
        this.kafkaJavaProducer.sendKafka(jsonObject.get("topic").getAsString(), jsonObject.get("message").getAsString());
        return null;
    }
}
