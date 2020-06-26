package com.ditrit.gandalf.bus.kafkajava.functions;

import com.ditrit.gandalf.java.grpc.CommandMessage;
import com.google.gson.Gson;
import com.google.gson.JsonObject;
import org.apache.kafka.clients.admin.AdminClient;
import org.apache.kafka.clients.admin.ListTopicsResult;
import org.apache.kafka.clients.admin.NewTopic;
import org.springframework.kafka.core.KafkaAdmin;
import org.springframework.stereotype.Component;

import java.util.ArrayList;
import java.util.List;

@Component
public class FunctionCreateTopic {

    public static final String COMMAND = "CREATE_TOPIC";
    private KafkaAdmin kafkaAdmin;
    private Gson mapper;

    public FunctionCreateTopic() {
        this.mapper = new Gson();
    }

    public String executeCommand(CommandMessage commandMessage) {
        String payload = commandMessage.getPayload();
        JsonObject jsonObject = mapper.fromJson(payload, JsonObject.class);
        String topic = jsonObject.get("topic").getAsString();

        AdminClient adminClient = AdminClient.create(this.kafkaAdmin.getConfig());
        if(!this.isTopicExist(topic, adminClient)) {
            NewTopic newTopic = new NewTopic(topic, 1, (short)1);
            List<NewTopic> createTopics = new ArrayList<>();
            createTopics.add(newTopic);
            adminClient.createTopics(createTopics);
        }
        adminClient.close();

        return null;
    }

    private boolean isTopicExist(String topic, AdminClient adminClient) {
        ListTopicsResult listTopics = adminClient.listTopics();
        try {
            return  listTopics.names().get().contains(topic);
        } catch (Exception e) {
            e.printStackTrace();
            return false;
        }
    }
}
