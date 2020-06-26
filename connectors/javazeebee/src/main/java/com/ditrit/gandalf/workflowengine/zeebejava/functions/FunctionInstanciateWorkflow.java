package com.ditrit.gandalf.workflowengine.zeebejava.functions;

import com.ditrit.gandalf.java.grpc.CommandMessage;
import com.google.gson.Gson;
import com.google.gson.JsonObject;
import io.zeebe.client.ZeebeClient;
import org.springframework.stereotype.Component;

@Component
public class FunctionInstanciateWorkflow  {

    private Gson mapper;
    private ZeebeClient zeebe;

    public FunctionInstanciateWorkflow() {
        super();
        this.mapper = new Gson();
    }

    public String executeCommand(CommandMessage commandMessage) {
        String payload = commandMessage.getPayload();
        JsonObject jsonObject = mapper.fromJson(payload, JsonObject.class);
        zeebe.newCreateInstanceCommand()
                .bpmnProcessId(jsonObject.get("id").getAsString())
                .latestVersion()
                .variables(jsonObject.get("variables").getAsString())
                .send().join();
        return null;
    }
}
