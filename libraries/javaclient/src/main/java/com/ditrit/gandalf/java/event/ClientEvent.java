package com.ditrit.gandalf.java.event;

import com.ditrit.gandalf.java.grpc.*;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

public class ClientEvent {

    private String clientEventConnection;
    private String identity;
    private ManagedChannel channel;


    public String getClientEventConnection() {
        return clientEventConnection;
    }

    public void setClientEventConnection(String clientEventConnection) {
        this.clientEventConnection = clientEventConnection;
    }

    public String getIdentity() {
        return identity;
    }

    public void setIdentity(String identity) {
        this.identity = identity;
    }

    public ClientEvent(String senderEventConnection, String identity) {
        this.clientEventConnection = senderEventConnection;
        this.identity = identity;
        String[] connections = this.getClientEventConnection().split(":");
        this.channel = ManagedChannelBuilder.forAddress(connections[0], Integer.parseInt(connections[1]))
                .usePlaintext()
                .build();
    }

    public void SendEvent(String topic, String timeout, String uuid, String event, String payload) {
        ConnectorEventGrpc.ConnectorEventBlockingStub stub = ConnectorEventGrpc.newBlockingStub(channel);

        Empty empty = stub.sendEventMessage(EventMessage.newBuilder()
                .setTopic(topic)
                .setTimeout(timeout)
                .setUUID(uuid)
                .setEvent(event)
                .setPayload(payload)
                .build());
    }


    public EventMessage WaitEvent(String event , String topic, String idIterator) {
        ConnectorEventGrpc.ConnectorEventBlockingStub stub = ConnectorEventGrpc.newBlockingStub(channel);

        EventMessage eventMessage = stub.waitEventMessage(EventMessageWait.newBuilder()
                .setTopic(topic).setEvent(event).setIteratorId(idIterator)
                .build());
        return eventMessage;
    }

    public IteratorMessage CreateIteratorEvent() {
        ConnectorEventGrpc.ConnectorEventBlockingStub stub = ConnectorEventGrpc.newBlockingStub(channel);

        IteratorMessage iteratorMessage = stub.createIteratorEvent(Empty.newBuilder()
                .build());
        return iteratorMessage;
    }
}
