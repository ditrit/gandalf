package com.ditrit.gandalf.workflowengine.zeebejava.sample;

import com.ditrit.gandalf.java.ClientGandalf;
import com.ditrit.gandalf.java.grpc.CommandMessage;
import com.ditrit.gandalf.java.grpc.IteratorMessage;
import com.ditrit.gandalf.workflowengine.zeebejava.functions.FunctionInstanciateWorkflow;
import com.ditrit.gandalf.workflowengine.zeebejava.functions.FunctionSendMessage;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class ZeebeJavaCLR {

    private FunctionInstanciateWorkflow functionInstanciateWorkflow;
    private FunctionSendMessage functionSendMessage;
    private ClientGandalf clientGandalf;

    @Autowired
    public ZeebeJavaCLR(ClientGandalf clientGandalf, FunctionInstanciateWorkflow functionInstanciateWorkflow, FunctionSendMessage functionSendMessage) {
        this.clientGandalf = clientGandalf;
        this.functionInstanciateWorkflow = functionInstanciateWorkflow;
        this.functionSendMessage = functionSendMessage;
    }

    public void sample(String... args) throws Exception {
        while (true) {
            IteratorMessage iteratorMessage = this.clientGandalf.CreateIteratorCommand();
            System.out.println("toto");
            System.out.println(iteratorMessage);
            CommandMessage commandMessage = this.clientGandalf.WaitCommand("instanciate_workflow", iteratorMessage.getId());
            System.out.println("toto1");
            this.functionInstanciateWorkflow.executeCommand(commandMessage);
            System.out.println("toto2");
            commandMessage = this.clientGandalf.WaitCommand("send_message_workflow", iteratorMessage.getId());
            System.out.println("toto3");
            this.functionSendMessage.executeCommand(commandMessage);
        }
    }
}
