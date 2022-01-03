# Gandalf - Messages

## Type

Communication is done through six types of messages:

+ Command: Used for sending commands to be executed by Workers
+ Event: Used to send information events or responses to Command.
+ Config: Used to send Connector configurations.
+ Configuration: Used to send logical configurations of gandalf components
+ Configuration Database: Used to retrieve the given database configuration for aggregaters.
+ Secret: Used for sending and validating the secret of gandalf components

## Form

The messages are all from the same database.

## Base

+ UUID: the identifier of the message
+ Tenant: the tenant of the message
+ Token:
+ Timeout: the duration of the message
+ Timestamp: the date the message was created
+ Payload: the payload of the message
+ Next:
+ Major: the major version
+ Minor: the minor version

## Command

+ Target: The target connector of the control
+ Command: The name of the command
+ Context: the application context of the command

## Event

+ Topic: The topic of the message
+ Event: The name of the event
+ ReferenceUUID: the reference to the UUID of a command

## Config

+ Target: The target connector of the control
+ Command: The name of the command
+ Context: the application context of the command

## Configuration

+ Target: The control target connector (if needed)
+ Command: The name of the command
+ Context: the application context of the command

## Configuration Database

+ Command: The name of the command
+ Context: the application context of the command

## Secret

+ Target: The control target connector (if needed)
+ Command: The name of the command
+ Context: the application context of the command
