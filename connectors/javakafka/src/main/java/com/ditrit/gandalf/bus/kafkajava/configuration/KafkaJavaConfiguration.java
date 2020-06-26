package com.ditrit.gandalf.bus.kafkajava.configuration;

import com.ditrit.gandalf.bus.kafkajava.properties.KafkaJavaProperties;
import com.ditrit.gandalf.java.ClientGandalf;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.core.annotation.Order;
import org.springframework.scheduling.concurrent.ThreadPoolTaskExecutor;

@Configuration
@Order
public class KafkaJavaConfiguration {

    private ApplicationContext context;
    private KafkaJavaProperties kafkaJavaProperties;

    @Autowired
    public KafkaJavaConfiguration(ApplicationContext context, KafkaJavaProperties kafkaJavaProperties) {
        this.context = context;
        this.kafkaJavaProperties = kafkaJavaProperties;
    }

    @Bean
    public ThreadPoolTaskExecutor taskExecutor() {
        ThreadPoolTaskExecutor pool = new ThreadPoolTaskExecutor();
        pool.setCorePoolSize(10);
        pool.setMaxPoolSize(20);
        pool.setWaitForTasksToCompleteOnShutdown(true);
        return pool;
    }

    @Bean
    public ClientGandalf clientGandalf() {
        ClientGandalf clientGandalf = new ClientGandalf(this.kafkaJavaProperties.getIdentity(), this.kafkaJavaProperties.getConnectorCommandConnection(), this.kafkaJavaProperties.getConnectorEventConnection());
        return clientGandalf;
    }

}
