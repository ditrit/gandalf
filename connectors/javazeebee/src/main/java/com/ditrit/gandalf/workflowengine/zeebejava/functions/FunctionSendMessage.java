package com.ditrit.gandalf.workflowengine.zeebejava.functions;

import com.ditrit.gandalf.java.grpc.CommandMessage;
import com.ditrit.gandalf.workflowengine.zeebejava.core.domain.ZeebeJavaMessage;
import com.google.gson.Gson;
import io.zeebe.client.ZeebeClient;
import org.springframework.stereotype.Component;

import java.time.Duration;

@Component
public class FunctionSendMessage {

    private Gson mapper;
    private ZeebeClient zeebe;

    public FunctionSendMessage() {
        super();
        this.mapper = new Gson();
    }

    public String executeCommand(CommandMessage commandMessage) {
        String payload = commandMessage.getPayload();
        ZeebeJavaMessage connectorZeebeMessage = this.mapper.fromJson(payload, ZeebeJavaMessage.class);
        zeebe.newPublishMessageCommand()
                .messageName(connectorZeebeMessage.getName())
                .correlationKey(connectorZeebeMessage.getCorrelationKey())
                .variables(connectorZeebeMessage.getVariables())
                .timeToLive(Duration.ofMinutes(connectorZeebeMessage.getDuration()))
                .send().join();
        return null;
    }
}
