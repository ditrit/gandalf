package com.ditrit.gandalf.bus.kafkajava.sample;

import com.ditrit.gandalf.bus.kafkajava.functions.FunctionSynchronizeTopicKafka;
import com.ditrit.gandalf.java.grpc.EventMessage;
import com.ditrit.gandalf.java.grpc.IteratorMessage;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import com.ditrit.gandalf.java.ClientGandalf;

@Component
public class KafkaJavaCLR {
    private ClientGandalf clientGandalf;
    private FunctionSynchronizeTopicKafka functionSynchronizeTopicKafka;

    @Autowired
    public KafkaJavaCLR(ClientGandalf clientGandalf, FunctionSynchronizeTopicKafka functionSynchronizeTopicKafka) {
        this.clientGandalf = clientGandalf;
        this.functionSynchronizeTopicKafka = functionSynchronizeTopicKafka;
    }

    public void sample(String... args) throws Exception {
        IteratorMessage iteratorMessage = this.clientGandalf.CreateIteratorEvent();
        while (true) {
            EventMessage eventMessage = this.clientGandalf.WaitEvent("synchronize", "kafka", iteratorMessage.getId());
            this.functionSynchronizeTopicKafka.executeEvent(eventMessage);
        }
    }
}
