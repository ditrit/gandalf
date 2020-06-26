package com.ditrit.gandalf.bus.kafkajava.core.producer;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Component;

@Component
public class KafkaJavaProducer {

    private KafkaTemplate<String, Object> kafkaTemplate;

    @Autowired
    public KafkaJavaProducer(KafkaTemplate kafkaTemplate) {
        this.kafkaTemplate = kafkaTemplate;
    }

    public void sendKafka(String topic, Object message) {
        kafkaTemplate.send(topic, message);
    }
}
