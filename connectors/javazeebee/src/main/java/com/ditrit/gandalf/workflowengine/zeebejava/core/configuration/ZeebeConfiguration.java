package com.ditrit.gandalf.workflowengine.zeebejava.core.configuration;

import com.ditrit.gandalf.workflowengine.zeebejava.properties.ZeebeJavaProperties;
import io.zeebe.client.ZeebeClient;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class ZeebeConfiguration {

    private ZeebeJavaProperties zeebeJavaProperties;

    @Autowired
    public ZeebeConfiguration(ZeebeJavaProperties zeebeJavaProperties) {
        this.zeebeJavaProperties = zeebeJavaProperties;
    }

    @Bean
    public ZeebeClient zeebe() {
        //TODO
        System.out.println("Connector connect to endpoint: " /*+ this.connectorZeebeProperties.getEndPointConnection()*/);
        ZeebeClient zeebeClient = ZeebeClient.newClientBuilder()
                .brokerContactPoint(this.zeebeJavaProperties.getEndPointConnection())
                .build();
        return zeebeClient;
    }
}
