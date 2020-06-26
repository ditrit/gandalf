package com.ditrit.gandalf.bus.kafkajava.functions;

import com.ditrit.gandalf.bus.kafkajava.core.producer.KafkaJavaProducer;
import com.ditrit.gandalf.java.grpc.EventMessage;
import com.google.gson.Gson;
import com.google.gson.JsonObject;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.core.KafkaAdmin;
import org.springframework.stereotype.Component;

@Component
public class FunctionSynchronizeTopicKafka {

    public static final String COMMAND = "SYNCHRONIZE_TOPIC_KAFKA";
    private KafkaAdmin kafkaAdmin;
    private Gson mapper;
    private KafkaJavaProducer kafkaJavaProducer;

    @Autowired
    public FunctionSynchronizeTopicKafka(KafkaJavaProducer kafkaJavaProducer) {
        this.kafkaJavaProducer = kafkaJavaProducer;
    }

    public String executeEvent(EventMessage eventMessage) {

        JsonObject jsonObject = new JsonObject();
        jsonObject.addProperty("event", eventMessage.getEvent());
        jsonObject.addProperty("payload", eventMessage.getPayload());

        this.kafkaJavaProducer.sendKafka(eventMessage.getTopic(), jsonObject.toString());

        return null;
    }
}
