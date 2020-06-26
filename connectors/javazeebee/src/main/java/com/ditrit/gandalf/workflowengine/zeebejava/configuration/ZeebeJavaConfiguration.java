package com.ditrit.gandalf.workflowengine.zeebejava.configuration;

import com.ditrit.gandalf.java.ClientGandalf;
import com.ditrit.gandalf.workflowengine.zeebejava.properties.ZeebeJavaProperties;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.core.annotation.Order;
import org.springframework.scheduling.concurrent.ThreadPoolTaskExecutor;

@Configuration
@Order
public class ZeebeJavaConfiguration {

    private ApplicationContext context;
    private ZeebeJavaProperties zeebeJavaProperties;

    @Autowired
    public ZeebeJavaConfiguration(ApplicationContext context, ZeebeJavaProperties zeebeJavaProperties) {
        this.context = context;
        this.zeebeJavaProperties = zeebeJavaProperties;
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
        return new ClientGandalf(this.zeebeJavaProperties.getIdentity() , this.zeebeJavaProperties.getConnectorCommandConnection(), this.zeebeJavaProperties.getConnectorEventConnection());
    }
}
