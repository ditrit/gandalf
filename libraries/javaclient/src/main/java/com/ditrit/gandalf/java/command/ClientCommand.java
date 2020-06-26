package com.ditrit.gandalf.java.command;

import com.ditrit.gandalf.java.grpc.*;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

public class ClientCommand {

    private String clientCommandConnection;
    private String identity;
    private ManagedChannel channel;

    public String getClientCommandConnection() {
        return clientCommandConnection;
    }

    public void setClientCommandConnection(String clientCommandConnection) {
        this.clientCommandConnection = clientCommandConnection;
    }

    public String getIdentity() {
        return identity;
    }

    public void setIdentity(String identity) {
        this.identity = identity;
    }

    public ClientCommand(String clientCommandConnection, String identity) {
        this.clientCommandConnection = clientCommandConnection;
        this.identity = identity;
        String[] connections = this.getClientCommandConnection().split(":");
        this.channel = ManagedChannelBuilder.forAddress(connections[0], Integer.parseInt(connections[1]))
                .usePlaintext()
                .build();
    }

    public String SendCommand(String context, String timeout, String uuid, String connectorType, String commandType, String command, String payload) {
        ConnectorCommandGrpc.ConnectorCommandBlockingStub stub = ConnectorCommandGrpc.newBlockingStub(channel);

        CommandMessageUUID commandMessageUUID = stub.sendCommandMessage(CommandMessage.newBuilder()
                .setTimeout(timeout).setConnectorType(connectorType).setCommandType(commandType).setCommand(command).setPayload(payload)
                .build());

        return commandMessageUUID.getUUID();
    }

    public void SendCommandReply(CommandMessage commandMessage, String reply, String payload) {
        ConnectorCommandGrpc.ConnectorCommandBlockingStub stub = ConnectorCommandGrpc.newBlockingStub(channel);

        Empty empty = stub.sendCommandMessageReply(CommandMessageReply.newBuilder()
                .setSourceAggregator(commandMessage.getSourceAggregator())
                .setSourceConnector(commandMessage.getSourceConnector())
                .setSourceWorker(commandMessage.getSourceWorker())
                .setDestinationAggregator(commandMessage.getDestinationAggregator())
                .setDestinationConnector(commandMessage.getDestinationConnector())
                .setDestinationWorker(commandMessage.getDestinationWorker())
                .setTenant(commandMessage.getTenant())
                .setContext(commandMessage.getContext())
                .setTimeout(commandMessage.getTimeout())
                .setTimestamp(commandMessage.getTimestamp())
                .setUUID(commandMessage.getUUID())
                .setReply(reply)
                .setPayload(payload)
                .build());
    }

    public CommandMessage WaitCommand(String command, String idIterator) {
        ConnectorCommandGrpc.ConnectorCommandBlockingStub stub = ConnectorCommandGrpc.newBlockingStub(channel);

        CommandMessage commandMessage = stub.waitCommandMessage(CommandMessageWait.newBuilder()
                .setValue(command)
                .setIteratorId(idIterator)
                .build());
        return commandMessage;
    }

    public CommandMessageReply WaitCommandReply(String uuid, String idIterator) {
        ConnectorCommandGrpc.ConnectorCommandBlockingStub stub = ConnectorCommandGrpc.newBlockingStub(channel);

        CommandMessageReply commandMessageReply = stub.waitCommandMessageReply(CommandMessageWait.newBuilder()
                .setValue(uuid).setIteratorId(idIterator)
                .build());
        return commandMessageReply;
    }

    public IteratorMessage CreateIteratorCommand() {
        ConnectorCommandGrpc.ConnectorCommandBlockingStub stub = ConnectorCommandGrpc.newBlockingStub(channel);
        IteratorMessage iteratorMessage = stub.createIteratorCommand(Empty.newBuilder()
                .build());
        return iteratorMessage;
    }
}

