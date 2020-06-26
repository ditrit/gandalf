package com.ditrit.gandalf.workflowengine.zeebejava.properties;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Configuration;

@Configuration
public class ZeebeJavaProperties {

    @Value("${identity}")
    private String identity;
    @Value("${endPointConnection}")
    private String endPointConnection;
    @Value("${connectorCommandConnection}")
    private String connectorCommandConnection;
    @Value("${connectorEventConnection}")
    private String connectorEventConnection;

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
}
