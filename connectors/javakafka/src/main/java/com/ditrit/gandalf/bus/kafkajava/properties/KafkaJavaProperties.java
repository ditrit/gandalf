package com.ditrit.gandalf.bus.kafkajava.properties;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.autoconfigure.condition.ConditionalOnExpression;
import org.springframework.context.annotation.Configuration;

import java.util.List;

@Configuration
public class KafkaJavaProperties {

    @Value("${identity}")
    private String identity;
    @Value("${endPointConnection}")
    private String endPointConnection;
    @Value("${connectorCommandConnection}")
    private String connectorCommandConnection;
    @Value("${connectorEventConnection}")
    private String connectorEventConnection;
    @Value("${synchronize.topics}")
    private List<String> synchronizeTopics;
    @Value("${group}")
    private String group;

    public String getIdentity() {
        return identity;
    }

    public void setIdentity(String identity) {
        this.identity = identity;
    }

    public String getConnectorCommandConnection() {
        return connectorCommandConnection;
    }

    public void setConnectorCommandConnection(String connectorCommandConnection) {
        this.connectorCommandConnection = connectorCommandConnection;
    }

    public String getConnectorEventConnection() {
        return connectorEventConnection;
    }

    public void setConnectorEventConnection(String connectorEventConnection) {
        this.connectorEventConnection = connectorEventConnection;
    }

    public String getEndPointConnection() {
        return endPointConnection;
    }

    public void setEndPointConnection(String endPointConnection) {
        this.endPointConnection = endPointConnection;
    }

    public String getGroup() {
        return group;
    }

    public void setGroup(String group) {
        this.group = group;
    }

    public List<String> getSynchronizeTopics() {
        return synchronizeTopics;
    }

    public void setSynchronizeTopics(List<String> synchronizeTopics) {
        this.synchronizeTopics = synchronizeTopics;
    }

}
