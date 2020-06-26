package com.ditrit.gandalf.bus.kafkajava.core.consumer;

import com.ditrit.gandalf.bus.kafkajava.properties.KafkaJavaProperties;
import com.ditrit.gandalf.java.ClientGandalf;
import com.google.gson.Gson;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;


@Component
public class KafkaJavaJavaConsumer extends RunnableKafkaJavaConsumer {

    private ClientGandalf clientGandalf;
    private KafkaJavaProperties kafkaJavaProperties;
    protected Gson mapper;

    @Autowired
    public KafkaJavaJavaConsumer(ClientGandalf clientGandalf, KafkaJavaProperties kafkaJavaProperties) {
        super();
        this.clientGandalf = clientGandalf;
        this.kafkaJavaProperties = kafkaJavaProperties;
        this.mapper = new Gson();
        this.initRunnable(this.kafkaJavaProperties.getEndPointConnection(), this.kafkaJavaProperties.getGroup(), this.kafkaJavaProperties.getSynchronizeTopics());
    }

    @Override
    protected void publish(Object message) {
        KafkaJavaMessage event = (KafkaJavaMessage) message;
        this.clientGandalf.SendEvent(event.getTopic(), event.getTimeout(), event.getUuid(), event.getEvent() , event.getPayload());
    }

    @Override
    protected Object parse(String value) {
        return this.mapper.fromJson(value, KafkaJavaMessage.class);
    }
}
