package com.ditrit.gandalf.bus.kafkajava.core.consumer;

import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.apache.kafka.clients.consumer.ConsumerRecords;
import org.apache.kafka.clients.consumer.KafkaConsumer;

import java.time.Duration;
import java.util.List;
import java.util.Properties;

public abstract class RunnableKafkaJavaConsumer implements Runnable {

    private String busConnection;
    private String group;
    private List<String> topics;
    protected  KafkaConsumer<String, String> consumer;

    protected void initRunnable(String busConnection, String group, List<String> topics) {
        this.busConnection = busConnection;
        this.group = group;
        this.topics = topics;

        this.consumer = new KafkaConsumer<>(this.consumerConfig(this.busConnection, this.group));
        consumer.subscribe(this.topics);
    }

    public void run() {
        while (!Thread.currentThread().isInterrupted()) {
            ConsumerRecords<String, String> records = consumer.poll(Duration.ofMillis(100));
            for (ConsumerRecord<String, String> record : records) {
                this.publish(record.value());
            }
        }
    }

    protected abstract void publish(Object message);

    protected abstract Object parse(String value);

    private Properties consumerConfig(String busConnection, String group) {
        Properties props = new Properties();
        props.setProperty("bootstrap.servers", busConnection);
        props.setProperty("group.id", group);
        props.setProperty("enable.auto.commit", "true");
        props.setProperty("auto.commit.interval.ms", "1000");
        props.setProperty("key.deserializer", "org.apache.kafka.common.serialization.StringDeserializer");
        props.setProperty("value.deserializer", "org.apache.kafka.common.serialization.StringDeserializer");
        return props;
    }

    public void addTopic(String topic) {
        this.topics.add(topic);
        this.consumer.subscribe(this.topics);
    }
}
