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
    comments = "Source: connectorEvent.proto")
public final class ConnectorEventGrpc {

  private ConnectorEventGrpc() {}

  public static final String SERVICE_NAME = "grpc.ConnectorEvent";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.ditrit.gandalf.java.grpc.EventMessage,
      com.ditrit.gandalf.java.grpc.Empty> getSendEventMessageMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SendEventMessage",
      requestType = com.ditrit.gandalf.java.grpc.EventMessage.class,
      responseType = com.ditrit.gandalf.java.grpc.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.ditrit.gandalf.java.grpc.EventMessage,
      com.ditrit.gandalf.java.grpc.Empty> getSendEventMessageMethod() {
    io.grpc.MethodDescriptor<com.ditrit.gandalf.java.grpc.EventMessage, com.ditrit.gandalf.java.grpc.Empty> getSendEventMessageMethod;
    if ((getSendEventMessageMethod = ConnectorEventGrpc.getSendEventMessageMethod) == null) {
      synchronized (ConnectorEventGrpc.class) {
        if ((getSendEventMessageMethod = ConnectorEventGrpc.getSendEventMessageMethod) == null) {
          ConnectorEventGrpc.getSendEventMessageMethod = getSendEventMessageMethod =
              io.grpc.MethodDescriptor.<com.ditrit.gandalf.java.grpc.EventMessage, com.ditrit.gandalf.java.grpc.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SendEventMessage"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.ditrit.gandalf.java.grpc.EventMessage.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.ditrit.gandalf.java.grpc.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new ConnectorEventMethodDescriptorSupplier("SendEventMessage"))
              .build();
        }
      }
    }
    return getSendEventMessageMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.ditrit.gandalf.java.grpc.EventMessageWait,
      com.ditrit.gandalf.java.grpc.EventMessage> getWaitEventMessageMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "WaitEventMessage",
      requestType = com.ditrit.gandalf.java.grpc.EventMessageWait.class,
      responseType = com.ditrit.gandalf.java.grpc.EventMessage.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.ditrit.gandalf.java.grpc.EventMessageWait,
      com.ditrit.gandalf.java.grpc.EventMessage> getWaitEventMessageMethod() {
    io.grpc.MethodDescriptor<com.ditrit.gandalf.java.grpc.EventMessageWait, com.ditrit.gandalf.java.grpc.EventMessage> getWaitEventMessageMethod;
    if ((getWaitEventMessageMethod = ConnectorEventGrpc.getWaitEventMessageMethod) == null) {
      synchronized (ConnectorEventGrpc.class) {
        if ((getWaitEventMessageMethod = ConnectorEventGrpc.getWaitEventMessageMethod) == null) {
          ConnectorEventGrpc.getWaitEventMessageMethod = getWaitEventMessageMethod =
              io.grpc.MethodDescriptor.<com.ditrit.gandalf.java.grpc.EventMessageWait, com.ditrit.gandalf.java.grpc.EventMessage>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "WaitEventMessage"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.ditrit.gandalf.java.grpc.EventMessageWait.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.ditrit.gandalf.java.grpc.EventMessage.getDefaultInstance()))
              .setSchemaDescriptor(new ConnectorEventMethodDescriptorSupplier("WaitEventMessage"))
              .build();
        }
      }
    }
    return getWaitEventMessageMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.ditrit.gandalf.java.grpc.Empty,
      com.ditrit.gandalf.java.grpc.IteratorMessage> getCreateIteratorEventMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateIteratorEvent",
      requestType = com.ditrit.gandalf.java.grpc.Empty.class,
      responseType = com.ditrit.gandalf.java.grpc.IteratorMessage.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.ditrit.gandalf.java.grpc.Empty,
      com.ditrit.gandalf.java.grpc.IteratorMessage> getCreateIteratorEventMethod() {
    io.grpc.MethodDescriptor<com.ditrit.gandalf.java.grpc.Empty, com.ditrit.gandalf.java.grpc.IteratorMessage> getCreateIteratorEventMethod;
    if ((getCreateIteratorEventMethod = ConnectorEventGrpc.getCreateIteratorEventMethod) == null) {
      synchronized (ConnectorEventGrpc.class) {
        if ((getCreateIteratorEventMethod = ConnectorEventGrpc.getCreateIteratorEventMethod) == null) {
          ConnectorEventGrpc.getCreateIteratorEventMethod = getCreateIteratorEventMethod =
              io.grpc.MethodDescriptor.<com.ditrit.gandalf.java.grpc.Empty, com.ditrit.gandalf.java.grpc.IteratorMessage>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateIteratorEvent"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.ditrit.gandalf.java.grpc.Empty.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.ditrit.gandalf.java.grpc.IteratorMessage.getDefaultInstance()))
              .setSchemaDescriptor(new ConnectorEventMethodDescriptorSupplier("CreateIteratorEvent"))
              .build();
        }
      }
    }
    return getCreateIteratorEventMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static ConnectorEventStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ConnectorEventStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ConnectorEventStub>() {
        @java.lang.Override
        public ConnectorEventStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ConnectorEventStub(channel, callOptions);
        }
      };
    return ConnectorEventStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static ConnectorEventBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ConnectorEventBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ConnectorEventBlockingStub>() {
        @java.lang.Override
        public ConnectorEventBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ConnectorEventBlockingStub(channel, callOptions);
        }
      };
    return ConnectorEventBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static ConnectorEventFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ConnectorEventFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ConnectorEventFutureStub>() {
        @java.lang.Override
        public ConnectorEventFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ConnectorEventFutureStub(channel, callOptions);
        }
      };
    return ConnectorEventFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class ConnectorEventImplBase implements io.grpc.BindableService {

    /**
     */
    public void sendEventMessage(com.ditrit.gandalf.java.grpc.EventMessage request,
        io.grpc.stub.StreamObserver<com.ditrit.gandalf.java.grpc.Empty> responseObserver) {
      asyncUnimplementedUnaryCall(getSendEventMessageMethod(), responseObserver);
    }

    /**
     */
    public void waitEventMessage(com.ditrit.gandalf.java.grpc.EventMessageWait request,
        io.grpc.stub.StreamObserver<com.ditrit.gandalf.java.grpc.EventMessage> responseObserver) {
      asyncUnimplementedUnaryCall(getWaitEventMessageMethod(), responseObserver);
    }

    /**
     */
    public void createIteratorEvent(com.ditrit.gandalf.java.grpc.Empty request,
        io.grpc.stub.StreamObserver<com.ditrit.gandalf.java.grpc.IteratorMessage> responseObserver) {
      asyncUnimplementedUnaryCall(getCreateIteratorEventMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getSendEventMessageMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                com.ditrit.gandalf.java.grpc.EventMessage,
                com.ditrit.gandalf.java.grpc.Empty>(
                  this, METHODID_SEND_EVENT_MESSAGE)))
          .addMethod(
            getWaitEventMessageMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                com.ditrit.gandalf.java.grpc.EventMessageWait,
                com.ditrit.gandalf.java.grpc.EventMessage>(
                  this, METHODID_WAIT_EVENT_MESSAGE)))
          .addMethod(
            getCreateIteratorEventMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                com.ditrit.gandalf.java.grpc.Empty,
                com.ditrit.gandalf.java.grpc.IteratorMessage>(
                  this, METHODID_CREATE_ITERATOR_EVENT)))
          .build();
    }
  }

  /**
   */
  public static final class ConnectorEventStub extends io.grpc.stub.AbstractAsyncStub<ConnectorEventStub> {
    private ConnectorEventStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ConnectorEventStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ConnectorEventStub(channel, callOptions);
    }

    /**
     */
    public void sendEventMessage(com.ditrit.gandalf.java.grpc.EventMessage request,
        io.grpc.stub.StreamObserver<com.ditrit.gandalf.java.grpc.Empty> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getSendEventMessageMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void waitEventMessage(com.ditrit.gandalf.java.grpc.EventMessageWait request,
        io.grpc.stub.StreamObserver<com.ditrit.gandalf.java.grpc.EventMessage> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getWaitEventMessageMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void createIteratorEvent(com.ditrit.gandalf.java.grpc.Empty request,
        io.grpc.stub.StreamObserver<com.ditrit.gandalf.java.grpc.IteratorMessage> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getCreateIteratorEventMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class ConnectorEventBlockingStub extends io.grpc.stub.AbstractBlockingStub<ConnectorEventBlockingStub> {
    private ConnectorEventBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ConnectorEventBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ConnectorEventBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.ditrit.gandalf.java.grpc.Empty sendEventMessage(com.ditrit.gandalf.java.grpc.EventMessage request) {
      return blockingUnaryCall(
          getChannel(), getSendEventMessageMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.ditrit.gandalf.java.grpc.EventMessage waitEventMessage(com.ditrit.gandalf.java.grpc.EventMessageWait request) {
      return blockingUnaryCall(
          getChannel(), getWaitEventMessageMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.ditrit.gandalf.java.grpc.IteratorMessage createIteratorEvent(com.ditrit.gandalf.java.grpc.Empty request) {
      return blockingUnaryCall(
          getChannel(), getCreateIteratorEventMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class ConnectorEventFutureStub extends io.grpc.stub.AbstractFutureStub<ConnectorEventFutureStub> {
    private ConnectorEventFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ConnectorEventFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ConnectorEventFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.ditrit.gandalf.java.grpc.Empty> sendEventMessage(
        com.ditrit.gandalf.java.grpc.EventMessage request) {
      return futureUnaryCall(
          getChannel().newCall(getSendEventMessageMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.ditrit.gandalf.java.grpc.EventMessage> waitEventMessage(
        com.ditrit.gandalf.java.grpc.EventMessageWait request) {
      return futureUnaryCall(
          getChannel().newCall(getWaitEventMessageMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.ditrit.gandalf.java.grpc.IteratorMessage> createIteratorEvent(
        com.ditrit.gandalf.java.grpc.Empty request) {
      return futureUnaryCall(
          getChannel().newCall(getCreateIteratorEventMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_SEND_EVENT_MESSAGE = 0;
  private static final int METHODID_WAIT_EVENT_MESSAGE = 1;
  private static final int METHODID_CREATE_ITERATOR_EVENT = 2;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final ConnectorEventImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(ConnectorEventImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_SEND_EVENT_MESSAGE:
          serviceImpl.sendEventMessage((com.ditrit.gandalf.java.grpc.EventMessage) request,
              (io.grpc.stub.StreamObserver<com.ditrit.gandalf.java.grpc.Empty>) responseObserver);
          break;
        case METHODID_WAIT_EVENT_MESSAGE:
          serviceImpl.waitEventMessage((com.ditrit.gandalf.java.grpc.EventMessageWait) request,
              (io.grpc.stub.StreamObserver<com.ditrit.gandalf.java.grpc.EventMessage>) responseObserver);
          break;
        case METHODID_CREATE_ITERATOR_EVENT:
          serviceImpl.createIteratorEvent((com.ditrit.gandalf.java.grpc.Empty) request,
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

  private static abstract class ConnectorEventBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    ConnectorEventBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.ditrit.gandalf.java.grpc.ConnectorEventProto.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("ConnectorEvent");
    }
  }

  private static final class ConnectorEventFileDescriptorSupplier
      extends ConnectorEventBaseDescriptorSupplier {
    ConnectorEventFileDescriptorSupplier() {}
  }

  private static final class ConnectorEventMethodDescriptorSupplier
      extends ConnectorEventBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    ConnectorEventMethodDescriptorSupplier(String methodName) {
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
      synchronized (ConnectorEventGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new ConnectorEventFileDescriptorSupplier())
              .addMethod(getSendEventMessageMethod())
              .addMethod(getWaitEventMessageMethod())
              .addMethod(getCreateIteratorEventMethod())
              .build();
        }
      }
    }
    return result;
  }
}
