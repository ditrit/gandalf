package com.ditrit.gandalf.java.grpc;

import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.27.1)",
    comments = "Source: connectorCommand.proto")
public final class ConnectorCommandGrpc {

  private ConnectorCommandGrpc() {}

  public static final String SERVICE_NAME = "grpc.ConnectorCommand";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.ditrit.gandalf.java.grpc.CommandMessage,
      com.ditrit.gandalf.java.grpc.CommandMessageUUID> getSendCommandMessageMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SendCommandMessage",
      requestType = com.ditrit.gandalf.java.grpc.CommandMessage.class,
      responseType = com.ditrit.gandalf.java.grpc.CommandMessageUUID.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.ditrit.gandalf.java.grpc.CommandMessage,
      com.ditrit.gandalf.java.grpc.CommandMessageUUID> getSendCommandMessageMethod() {
    io.grpc.MethodDescriptor<com.ditrit.gandalf.java.grpc.CommandMessage, com.ditrit.gandalf.java.grpc.CommandMessageUUID> getSendCommandMessageMethod;
    if ((getSendCommandMessageMethod = ConnectorCommandGrpc.getSendCommandMessageMethod) == null) {
      synchronized (ConnectorCommandGrpc.class) {
        if ((getSendCommandMessageMethod = ConnectorCommandGrpc.getSendCommandMessageMethod) == null) {
          ConnectorCommandGrpc.getSendCommandMessageMethod = getSendCommandMessageMethod =
              io.grpc.MethodDescriptor.<com.ditrit.gandalf.java.grpc.CommandMessage, com.ditrit.gandalf.java.grpc.CommandMessageUUID>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SendCommandMessage"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.ditrit.gandalf.java.grpc.CommandMessage.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.ditrit.gandalf.java.grpc.CommandMessageUUID.getDefaultInstance()))
              .setSchemaDescriptor(new ConnectorCommandMethodDescriptorSupplier("SendCommandMessage"))
              .build();
        }
      }
    }
    return getSendCommandMessageMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.ditrit.gandalf.java.grpc.CommandMessageReply,
      com.ditrit.gandalf.java.grpc.Empty> getSendCommandMessageReplyMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SendCommandMessageReply",
      requestType = com.ditrit.gandalf.java.grpc.CommandMessageReply.class,
      responseType = com.ditrit.gandalf.java.grpc.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.ditrit.gandalf.java.grpc.CommandMessageReply,
      com.ditrit.gandalf.java.grpc.Empty> getSendCommandMessageReplyMethod() {
    io.grpc.MethodDescriptor<com.ditrit.gandalf.java.grpc.CommandMessageReply, com.ditrit.gandalf.java.grpc.Empty> getSendCommandMessageReplyMethod;
    if ((getSendCommandMessageReplyMethod = ConnectorCommandGrpc.getSendCommandMessageReplyMethod) == null) {
      synchronized (ConnectorCommandGrpc.class) {
        if ((getSendCommandMessageReplyMethod = ConnectorCommandGrpc.getSendCommandMessageReplyMethod) == null) {
          ConnectorCommandGrpc.getSendCommandMessageReplyMethod = getSendCommandMessageReplyMethod =
              io.grpc.MethodDescriptor.<com.ditrit.gandalf.java.grpc.CommandMessageReply, com.ditrit.gandalf.java.grpc.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SendCommandMessageReply"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.ditrit.gandalf.java.grpc.CommandMessageReply.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.ditrit.gandalf.java.grpc.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new ConnectorCommandMethodDescriptorSupplier("SendCommandMessageReply"))
              .build();
        }
      }
    }
    return getSendCommandMessageReplyMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.ditrit.gandalf.java.grpc.CommandMessageWait,
      com.ditrit.gandalf.java.grpc.CommandMessage> getWaitCommandMessageMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "WaitCommandMessage",
      requestType = com.ditrit.gandalf.java.grpc.CommandMessageWait.class,
      responseType = com.ditrit.gandalf.java.grpc.CommandMessage.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.ditrit.gandalf.java.grpc.CommandMessageWait,
      com.ditrit.gandalf.java.grpc.CommandMessage> getWaitCommandMessageMethod() {
    io.grpc.MethodDescriptor<com.ditrit.gandalf.java.grpc.CommandMessageWait, com.ditrit.gandalf.java.grpc.CommandMessage> getWaitCommandMessageMethod;
    if ((getWaitCommandMessageMethod = ConnectorCommandGrpc.getWaitCommandMessageMethod) == null) {
      synchronized (ConnectorCommandGrpc.class) {
        if ((getWaitCommandMessageMethod = ConnectorCommandGrpc.getWaitCommandMessageMethod) == null) {
          ConnectorCommandGrpc.getWaitCommandMessageMethod = getWaitCommandMessageMethod =
              io.grpc.MethodDescriptor.<com.ditrit.gandalf.java.grpc.CommandMessageWait, com.ditrit.gandalf.java.grpc.CommandMessage>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "WaitCommandMessage"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.ditrit.gandalf.java.grpc.CommandMessageWait.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.ditrit.gandalf.java.grpc.CommandMessage.getDefaultInstance()))
              .setSchemaDescriptor(new ConnectorCommandMethodDescriptorSupplier("WaitCommandMessage"))
              .build();
        }
      }
    }
    return getWaitCommandMessageMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.ditrit.gandalf.java.grpc.CommandMessageWait,
      com.ditrit.gandalf.java.grpc.CommandMessageReply> getWaitCommandMessageReplyMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "WaitCommandMessageReply",
      requestType = com.ditrit.gandalf.java.grpc.CommandMessageWait.class,
      responseType = com.ditrit.gandalf.java.grpc.CommandMessageReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.ditrit.gandalf.java.grpc.CommandMessageWait,
      com.ditrit.gandalf.java.grpc.CommandMessageReply> getWaitCommandMessageReplyMethod() {
    io.grpc.MethodDescriptor<com.ditrit.gandalf.java.grpc.CommandMessageWait, com.ditrit.gandalf.java.grpc.CommandMessageReply> getWaitCommandMessageReplyMethod;
    if ((getWaitCommandMessageReplyMethod = ConnectorCommandGrpc.getWaitCommandMessageReplyMethod) == null) {
      synchronized (ConnectorCommandGrpc.class) {
        if ((getWaitCommandMessageReplyMethod = ConnectorCommandGrpc.getWaitCommandMessageReplyMethod) == null) {
          ConnectorCommandGrpc.getWaitCommandMessageReplyMethod = getWaitCommandMessageReplyMethod =
              io.grpc.MethodDescriptor.<com.ditrit.gandalf.java.grpc.CommandMessageWait, com.ditrit.gandalf.java.grpc.CommandMessageReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "WaitCommandMessageReply"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.ditrit.gandalf.java.grpc.CommandMessageWait.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.ditrit.gandalf.java.grpc.CommandMessageReply.getDefaultInstance()))
              .setSchemaDescriptor(new ConnectorCommandMethodDescriptorSupplier("WaitCommandMessageReply"))
              .build();
        }
      }
    }
    return getWaitCommandMessageReplyMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.ditrit.gandalf.java.grpc.Empty,
      com.ditrit.gandalf.java.grpc.IteratorMessage> getCreateIteratorCommandMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateIteratorCommand",
      requestType = com.ditrit.gandalf.java.grpc.Empty.class,
      responseType = com.ditrit.gandalf.java.grpc.IteratorMessage.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.ditrit.gandalf.java.grpc.Empty,
      com.ditrit.gandalf.java.grpc.IteratorMessage> getCreateIteratorCommandMethod() {
    io.grpc.MethodDescriptor<com.ditrit.gandalf.java.grpc.Empty, com.ditrit.gandalf.java.grpc.IteratorMessage> getCreateIteratorCommandMethod;
    if ((getCreateIteratorCommandMethod = ConnectorCommandGrpc.getCreateIteratorCommandMethod) == null) {
      synchronized (ConnectorCommandGrpc.class) {
        if ((getCreateIteratorCommandMethod = ConnectorCommandGrpc.getCreateIteratorCommandMethod) == null) {
          ConnectorCommandGrpc.getCreateIteratorCommandMethod = getCreateIteratorCommandMethod =
              io.grpc.MethodDescriptor.<com.ditrit.gandalf.java.grpc.Empty, com.ditrit.gandalf.java.grpc.IteratorMessage>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateIteratorCommand"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.ditrit.gandalf.java.grpc.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.ditrit.gandalf.java.grpc.IteratorMessage.getDefaultInstance()))
              .setSchemaDescriptor(new ConnectorCommandMethodDescriptorSupplier("CreateIteratorCommand"))
              .build();
        }
      }
    }
    return getCreateIteratorCommandMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static ConnectorCommandStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ConnectorCommandStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ConnectorCommandStub>() {
        @java.lang.Override
        public ConnectorCommandStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ConnectorCommandStub(channel, callOptions);
        }
      };
    return ConnectorCommandStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static ConnectorCommandBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ConnectorCommandBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ConnectorCommandBlockingStub>() {
        @java.lang.Override
        public ConnectorCommandBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ConnectorCommandBlockingStub(channel, callOptions);
        }
      };
    return ConnectorCommandBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static ConnectorCommandFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ConnectorCommandFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ConnectorCommandFutureStub>() {
        @java.lang.Override
        public ConnectorCommandFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ConnectorCommandFutureStub(channel, callOptions);
        }
      };
    return ConnectorCommandFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class ConnectorCommandImplBase implements io.grpc.BindableService {

    /**
     */
    public void sendCommandMessage(com.ditrit.gandalf.java.grpc.CommandMessage request,
        io.grpc.stub.StreamObserver<com.ditrit.gandalf.java.grpc.CommandMessageUUID> responseObserver) {
      asyncUnimplementedUnaryCall(getSendCommandMessageMethod(), responseObserver);
    }

    /**
     */
    public void sendCommandMessageReply(com.ditrit.gandalf.java.grpc.CommandMessageReply request,
        io.grpc.stub.StreamObserver<com.ditrit.gandalf.java.grpc.Empty> responseObserver) {
      asyncUnimplementedUnaryCall(getSendCommandMessageReplyMethod(), responseObserver);
    }

    /**
     */
    public void waitCommandMessage(com.ditrit.gandalf.java.grpc.CommandMessageWait request,
        io.grpc.stub.StreamObserver<com.ditrit.gandalf.java.grpc.CommandMessage> responseObserver) {
      asyncUnimplementedUnaryCall(getWaitCommandMessageMethod(), responseObserver);
    }

    /**
     */
    public void waitCommandMessageReply(com.ditrit.gandalf.java.grpc.CommandMessageWait request,
        io.grpc.stub.StreamObserver<com.ditrit.gandalf.java.grpc.CommandMessageReply> responseObserver) {
      asyncUnimplementedUnaryCall(getWaitCommandMessageReplyMethod(), responseObserver);
    }

    /**
     */
    public void createIteratorCommand(com.ditrit.gandalf.java.grpc.Empty request,
        io.grpc.stub.StreamObserver<com.ditrit.gandalf.java.grpc.IteratorMessage> responseObserver) {
      asyncUnimplementedUnaryCall(getCreateIteratorCommandMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getSendCommandMessageMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                com.ditrit.gandalf.java.grpc.CommandMessage,
                com.ditrit.gandalf.java.grpc.CommandMessageUUID>(
                  this, METHODID_SEND_COMMAND_MESSAGE)))
          .addMethod(
            getSendCommandMessageReplyMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                com.ditrit.gandalf.java.grpc.CommandMessageReply,
                com.ditrit.gandalf.java.grpc.Empty>(
                  this, METHODID_SEND_COMMAND_MESSAGE_REPLY)))
          .addMethod(
            getWaitCommandMessageMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                com.ditrit.gandalf.java.grpc.CommandMessageWait,
                com.ditrit.gandalf.java.grpc.CommandMessage>(
                  this, METHODID_WAIT_COMMAND_MESSAGE)))
          .addMethod(
            getWaitCommandMessageReplyMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                com.ditrit.gandalf.java.grpc.CommandMessageWait,
                com.ditrit.gandalf.java.grpc.CommandMessageReply>(
                  this, METHODID_WAIT_COMMAND_MESSAGE_REPLY)))
          .addMethod(
            getCreateIteratorCommandMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                com.ditrit.gandalf.java.grpc.Empty,
                com.ditrit.gandalf.java.grpc.IteratorMessage>(
                  this, METHODID_CREATE_ITERATOR_COMMAND)))
          .build();
    }
  }

  /**
   */
  public static final class ConnectorCommandStub extends io.grpc.stub.AbstractAsyncStub<ConnectorCommandStub> {
    private ConnectorCommandStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ConnectorCommandStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ConnectorCommandStub(channel, callOptions);
    }

    /**
     */
    public void sendCommandMessage(com.ditrit.gandalf.java.grpc.CommandMessage request,
        io.grpc.stub.StreamObserver<com.ditrit.gandalf.java.grpc.CommandMessageUUID> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getSendCommandMessageMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void sendCommandMessageReply(com.ditrit.gandalf.java.grpc.CommandMessageReply request,
        io.grpc.stub.StreamObserver<com.ditrit.gandalf.java.grpc.Empty> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getSendCommandMessageReplyMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void waitCommandMessage(com.ditrit.gandalf.java.grpc.CommandMessageWait request,
        io.grpc.stub.StreamObserver<com.ditrit.gandalf.java.grpc.CommandMessage> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getWaitCommandMessageMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void waitCommandMessageReply(com.ditrit.gandalf.java.grpc.CommandMessageWait request,
        io.grpc.stub.StreamObserver<com.ditrit.gandalf.java.grpc.CommandMessageReply> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getWaitCommandMessageReplyMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void createIteratorCommand(com.ditrit.gandalf.java.grpc.Empty request,
        io.grpc.stub.StreamObserver<com.ditrit.gandalf.java.grpc.IteratorMessage> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getCreateIteratorCommandMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class ConnectorCommandBlockingStub extends io.grpc.stub.AbstractBlockingStub<ConnectorCommandBlockingStub> {
    private ConnectorCommandBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ConnectorCommandBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ConnectorCommandBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.ditrit.gandalf.java.grpc.CommandMessageUUID sendCommandMessage(com.ditrit.gandalf.java.grpc.CommandMessage request) {
      return blockingUnaryCall(
          getChannel(), getSendCommandMessageMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.ditrit.gandalf.java.grpc.Empty sendCommandMessageReply(com.ditrit.gandalf.java.grpc.CommandMessageReply request) {
      return blockingUnaryCall(
          getChannel(), getSendCommandMessageReplyMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.ditrit.gandalf.java.grpc.CommandMessage waitCommandMessage(com.ditrit.gandalf.java.grpc.CommandMessageWait request) {
      return blockingUnaryCall(
          getChannel(), getWaitCommandMessageMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.ditrit.gandalf.java.grpc.CommandMessageReply waitCommandMessageReply(com.ditrit.gandalf.java.grpc.CommandMessageWait request) {
      return blockingUnaryCall(
          getChannel(), getWaitCommandMessageReplyMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.ditrit.gandalf.java.grpc.IteratorMessage createIteratorCommand(com.ditrit.gandalf.java.grpc.Empty request) {
      return blockingUnaryCall(
          getChannel(), getCreateIteratorCommandMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class ConnectorCommandFutureStub extends io.grpc.stub.AbstractFutureStub<ConnectorCommandFutureStub> {
    private ConnectorCommandFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ConnectorCommandFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ConnectorCommandFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.ditrit.gandalf.java.grpc.CommandMessageUUID> sendCommandMessage(
        com.ditrit.gandalf.java.grpc.CommandMessage request) {
      return futureUnaryCall(
          getChannel().newCall(getSendCommandMessageMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.ditrit.gandalf.java.grpc.Empty> sendCommandMessageReply(
        com.ditrit.gandalf.java.grpc.CommandMessageReply request) {
      return futureUnaryCall(
          getChannel().newCall(getSendCommandMessageReplyMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.ditrit.gandalf.java.grpc.CommandMessage> waitCommandMessage(
        com.ditrit.gandalf.java.grpc.CommandMessageWait request) {
      return futureUnaryCall(
          getChannel().newCall(getWaitCommandMessageMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.ditrit.gandalf.java.grpc.CommandMessageReply> waitCommandMessageReply(
        com.ditrit.gandalf.java.grpc.CommandMessageWait request) {
      return futureUnaryCall(
          getChannel().newCall(getWaitCommandMessageReplyMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.ditrit.gandalf.java.grpc.IteratorMessage> createIteratorCommand(
        com.ditrit.gandalf.java.grpc.Empty request) {
      return futureUnaryCall(
          getChannel().newCall(getCreateIteratorCommandMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_SEND_COMMAND_MESSAGE = 0;
  private static final int METHODID_SEND_COMMAND_MESSAGE_REPLY = 1;
  private static final int METHODID_WAIT_COMMAND_MESSAGE = 2;
  private static final int METHODID_WAIT_COMMAND_MESSAGE_REPLY = 3;
  private static final int METHODID_CREATE_ITERATOR_COMMAND = 4;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final ConnectorCommandImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(ConnectorCommandImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_SEND_COMMAND_MESSAGE:
          serviceImpl.sendCommandMessage((com.ditrit.gandalf.java.grpc.CommandMessage) request,
              (io.grpc.stub.StreamObserver<com.ditrit.gandalf.java.grpc.CommandMessageUUID>) responseObserver);
          break;
        case METHODID_SEND_COMMAND_MESSAGE_REPLY:
          serviceImpl.sendCommandMessageReply((com.ditrit.gandalf.java.grpc.CommandMessageReply) request,
              (io.grpc.stub.StreamObserver<com.ditrit.gandalf.java.grpc.Empty>) responseObserver);
          break;
        case METHODID_WAIT_COMMAND_MESSAGE:
          serviceImpl.waitCommandMessage((com.ditrit.gandalf.java.grpc.CommandMessageWait) request,
              (io.grpc.stub.StreamObserver<com.ditrit.gandalf.java.grpc.CommandMessage>) responseObserver);
          break;
        case METHODID_WAIT_COMMAND_MESSAGE_REPLY:
          serviceImpl.waitCommandMessageReply((com.ditrit.gandalf.java.grpc.CommandMessageWait) request,
              (io.grpc.stub.StreamObserver<com.ditrit.gandalf.java.grpc.CommandMessageReply>) responseObserver);
          break;
        case METHODID_CREATE_ITERATOR_COMMAND:
          serviceImpl.createIteratorCommand((com.ditrit.gandalf.java.grpc.Empty) request,
              (io.grpc.stub.StreamObserver<com.ditrit.gandalf.java.grpc.IteratorMessage>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class ConnectorCommandBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    ConnectorCommandBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.ditrit.gandalf.java.grpc.ConnectorCommandProto.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("ConnectorCommand");
    }
  }

  private static final class ConnectorCommandFileDescriptorSupplier
      extends ConnectorCommandBaseDescriptorSupplier {
    ConnectorCommandFileDescriptorSupplier() {}
  }

  private static final class ConnectorCommandMethodDescriptorSupplier
      extends ConnectorCommandBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    ConnectorCommandMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (ConnectorCommandGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new ConnectorCommandFileDescriptorSupplier())
              .addMethod(getSendCommandMessageMethod())
              .addMethod(getSendCommandMessageReplyMethod())
              .addMethod(getWaitCommandMessageMethod())
              .addMethod(getWaitCommandMessageReplyMethod())
              .addMethod(getCreateIteratorCommandMethod())
              .build();
        }
      }
    }
    return result;
  }
}
