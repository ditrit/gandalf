package com.ditrit.gandalf.java;

import com.ditrit.gandalf.java.command.ClientCommand;
import com.ditrit.gandalf.java.event.ClientEvent;
import com.ditrit.gandalf.java.grpc.CommandMessage;
import com.ditrit.gandalf.java.grpc.CommandMessageReply;
import com.ditrit.gandalf.java.grpc.EventMessage;
import com.ditrit.gandalf.java.grpc.IteratorMessage;

public class ClientGandalf {

   private String identity;
   private String clientCommandConnection;
   private String clientEventConnection;
   private ClientCommand clientCommand;
   private ClientEvent clientEvent;


    public String getIdentity() {
        return identity;
    }

    public void setIdentity(String identity) {
        this.identity = identity;
    }

    public String getClientCommandConnection() {
        return clientCommandConnection;
    }

    public void setClientCommandConnection(String clientCommandConnection) {
        this.clientCommandConnection = clientCommandConnection;
    }

    public String getClientEventConnection() {
        return clientEventConnection;
    }

    public void setClientEventConnection(String clientEventConnection) {
        this.clientEventConnection = clientEventConnection;
    }

    public ClientCommand getClientCommand() {
        return clientCommand;
    }

    public void setClientCommand(ClientCommand clientCommand) {
        this.clientCommand = clientCommand;
    }

    public ClientEvent getClientEvent() {
        return clientEvent;
    }

    public void setClientEvent(ClientEvent clientEvent) {
        this.clientEvent = clientEvent;
    }

    public ClientGandalf(String identity, String clientCommandConnection, String clientEventConnection) {
        this.identity = identity;
        this.clientCommandConnection = clientCommandConnection;
        this.clientEventConnection = clientEventConnection;
        this.clientCommand = new ClientCommand(this.getClientCommandConnection(), this.getIdentity());
        this.clientEvent = new ClientEvent(this.getClientEventConnection(), this.getIdentity());
    }

    public String SendCommand(String context, String timeout, String uuid, String connectorType, String commandType, String command, String payload) {
        return this.getClientCommand().SendCommand(context, timeout, uuid, connectorType, commandType, command,  payload);
    }

    public void SendCommandReply(CommandMessage commandMessage, String reply, String payload) {
        this.getClientCommand().SendCommandReply(commandMessage, reply, payload);
    }

    public void SendEvent(String topic, String timeout, String uuid, String event, String payload) {
        this.getClientEvent().SendEvent(topic, timeout, uuid, event, payload);
    }

    public CommandMessage WaitCommand(String command, String idIterator) {
        return this.getClientCommand().WaitCommand(command, idIterator);
    }

    public CommandMessageReply WaitCommandreply(String uuid, String idIterator) {
        return this.getClientCommand().WaitCommandReply(uuid, idIterator);
    }

    public EventMessage WaitEvent(String event, String topic, String idIterator) {
        return this.getClientEvent().WaitEvent(event, topic, idIterator);
    }

    public IteratorMessage CreateIteratorCommand() {
        return this.getClientCommand().CreateIteratorCommand();
    }

    public IteratorMessage CreateIteratorEvent() {
        return this.getClientEvent().CreateIteratorEvent();
    }
}
