package com.ditrit.gandalf.workflowengine.zeebejava.core.domain;

import java.util.HashMap;

public class ZeebeJavaMessage {

    private String name;
    private String correlationKey;
    private HashMap<String, String> variables;
    private Long duration;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getCorrelationKey() {
        return correlationKey;
    }

    public void setCorrelationKey(String correlationKey) {
        this.correlationKey = correlationKey;
    }

    public HashMap<String, String> getVariables() {
        return variables;
    }

    public void setVariables(HashMap<String, String> variables) {
        this.variables = variables;
    }

    public Long getDuration() {
        return duration;
    }

    public void setDuration(Long duration) {
        this.duration = duration;
    }

    public ZeebeJavaMessage(String name, String correlationKey, HashMap<String, String> variables, Long duration) {
        this.name = name;
        this.correlationKey = correlationKey;
        this.variables = variables;
        this.duration = duration;
    }

    @Override
    public String toString() {
        return "ZeebeMessage{" +
                "name='" + name + '\'' +
                ", correlationKey='" + correlationKey + '\'' +
                ", variables=" + variables +
                ", duration=" + duration +
                '}';
    }
}
