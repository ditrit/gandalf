package com.ditrit.gandalf.java;

public class ClientConfiguration {

    private String SenderCommandConnection;
    private String SenderEventConnection;
    private String WaiterCommandConnection;
    private String WaiterEventConnection;
    private String IteratorConnection;
    private String Identity;

    public String getSenderCommandConnection() {
        return SenderCommandConnection;
    }

    public void setSenderCommandConnection(String senderCommandConnection) {
        SenderCommandConnection = senderCommandConnection;
    }

    public String getSenderEventConnection() {
        return SenderEventConnection;
    }

    public void setSenderEventConnection(String senderEventConnection) {
        SenderEventConnection = senderEventConnection;
    }

    public String getWaiterCommandConnection() {
        return WaiterCommandConnection;
    }

    public void setWaiterCommandConnection(String waiterCommandConnection) {
        WaiterCommandConnection = waiterCommandConnection;
    }

    public String getWaiterEventConnection() {
        return WaiterEventConnection;
    }

    public void setWaiterEventConnection(String waiterEventConnection) {
        WaiterEventConnection = waiterEventConnection;
    }

    public String getIteratorConnection() {
        return IteratorConnection;
    }

    public void setIteratorConnection(String iteratorConnection) {
        IteratorConnection = iteratorConnection;
    }

    public String getIdentity() {
        return Identity;
    }

    public void setIdentity(String identity) {
        Identity = identity;
    }

    //TODO LOAD
}
