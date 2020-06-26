package com.ditrit.gandalf.bus.kafkajava.core.configuration;

import com.ditrit.gandalf.bus.kafkajava.properties.KafkaJavaProperties;
import org.apache.kafka.clients.admin.AdminClientConfig;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.kafka.core.KafkaAdmin;

import java.util.HashMap;
import java.util.Map;

@Configuration
public class KafkaConfiguration {

    private KafkaJavaProperties kafkaJavaProperties;

    @Autowired
    public KafkaConfiguration(KafkaJavaProperties kafkaJavaProperties) {
        this.kafkaJavaProperties = kafkaJavaProperties;
    }

    @Bean
    public KafkaAdmin kafkaAdmin() {
        Map<String, Object> configs = new HashMap<>();
        System.out.println(this.kafkaJavaProperties.getEndPointConnection());
        configs.put(AdminClientConfig.BOOTSTRAP_SERVERS_CONFIG, this.kafkaJavaProperties.getEndPointConnection());
        return new KafkaAdmin(configs);
    }
}
