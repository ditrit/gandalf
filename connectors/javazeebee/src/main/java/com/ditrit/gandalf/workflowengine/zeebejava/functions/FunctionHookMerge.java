package com.ditrit.gandalf.workflowengine.zeebejava.functions;/*
package com.ditrit.gandalf.demoworkerzeebe.functions;

import com.ditrit.gandalf.gandalfjava.core.zeromqcore.worker.domain.EventFunction;
import com.google.gson.Gson;
import com.google.gson.JsonObject;
import io.zeebe.client.ZeebeClient;
import org.zeromq.ZMsg;

import java.time.Duration;
import java.util.HashMap;

public class FunctionHookMerge  {

    private Gson mapper;
    private ZeebeClient zeebe;

    public FunctionHookMerge() {
        super();
        this.mapper = new Gson();
    }

    public void executeCommand(grpc.CommandMessage commandMessage) {
        Object[] eventArray = event.toArray();
        String payload = eventArray[4].toString();
        String topic = eventArray[2].toString();

        JsonObject jsonObject = this.mapper.fromJson(payload, JsonObject.class);
        System.out.println(jsonObject.get("project_url").getAsString());
        HashMap<String, String> variables = new HashMap<>();
        variables.put("project_name", jsonObject.get("project_name").getAsString());
        variables.put("project_url", jsonObject.get("project_url").getAsString());

        zeebe.newPublishMessageCommand() //
                .messageName("message")
                .correlationKey(topic)
                .variables(variables)
                .timeToLive(Duration.ofMinutes(100L))
                .send().join();
    }
}
*/
