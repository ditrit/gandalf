// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: connectorCommand.proto

package com.ditrit.gandalf.java.grpc;

/**
 * Protobuf type {@code grpc.CommandMessageUUID}
 */
public  final class CommandMessageUUID extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:grpc.CommandMessageUUID)
    CommandMessageUUIDOrBuilder {
  // Use CommandMessageUUID.newBuilder() to construct.
  private CommandMessageUUID(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private CommandMessageUUID() {
    uUID_ = "";
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return com.google.protobuf.UnknownFieldSet.getDefaultInstance();
  }
  private CommandMessageUUID(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    int mutable_bitField0_ = 0;
    try {
      boolean done = false;
      while (!done) {
        int tag = input.readTag();
        switch (tag) {
          case 0:
            done = true;
            break;
          default: {
            if (!input.skipField(tag)) {
              done = true;
            }
            break;
          }
          case 10: {
            java.lang.String s = input.readStringRequireUtf8();

            uUID_ = s;
            break;
          }
        }
      }
    } catch (com.google.protobuf.InvalidProtocolBufferException e) {
      throw e.setUnfinishedMessage(this);
    } catch (java.io.IOException e) {
      throw new com.google.protobuf.InvalidProtocolBufferException(
          e).setUnfinishedMessage(this);
    } finally {
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return com.ditrit.gandalf.java.grpc.ConnectorCommandProto.internal_static_grpc_CommandMessageUUID_descriptor;
  }

  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return com.ditrit.gandalf.java.grpc.ConnectorCommandProto.internal_static_grpc_CommandMessageUUID_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            com.ditrit.gandalf.java.grpc.CommandMessageUUID.class, com.ditrit.gandalf.java.grpc.CommandMessageUUID.Builder.class);
  }

  public static final int UUID_FIELD_NUMBER = 1;
  private volatile java.lang.Object uUID_;
  /**
   * <code>optional string UUID = 1;</code>
   */
  public java.lang.String getUUID() {
    java.lang.Object ref = uUID_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      uUID_ = s;
      return s;
    }
  }
  /**
   * <code>optional string UUID = 1;</code>
   */
  public com.google.protobuf.ByteString
      getUUIDBytes() {
    java.lang.Object ref = uUID_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      uUID_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  private byte memoizedIsInitialized = -1;
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (!getUUIDBytes().isEmpty()) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 1, uUID_);
    }
  }

  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (!getUUIDBytes().isEmpty()) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(1, uUID_);
    }
    memoizedSize = size;
    return size;
  }

  private static final long serialVersionUID = 0L;
  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof com.ditrit.gandalf.java.grpc.CommandMessageUUID)) {
      return super.equals(obj);
    }
    com.ditrit.gandalf.java.grpc.CommandMessageUUID other = (com.ditrit.gandalf.java.grpc.CommandMessageUUID) obj;

    boolean result = true;
    result = result && getUUID()
        .equals(other.getUUID());
    return result;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptorForType().hashCode();
    hash = (37 * hash) + UUID_FIELD_NUMBER;
    hash = (53 * hash) + getUUID().hashCode();
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static com.ditrit.gandalf.java.grpc.CommandMessageUUID parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.ditrit.gandalf.java.grpc.CommandMessageUUID parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.ditrit.gandalf.java.grpc.CommandMessageUUID parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.ditrit.gandalf.java.grpc.CommandMessageUUID parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.ditrit.gandalf.java.grpc.CommandMessageUUID parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.ditrit.gandalf.java.grpc.CommandMessageUUID parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.ditrit.gandalf.java.grpc.CommandMessageUUID parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static com.ditrit.gandalf.java.grpc.CommandMessageUUID parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.ditrit.gandalf.java.grpc.CommandMessageUUID parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.ditrit.gandalf.java.grpc.CommandMessageUUID parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(com.ditrit.gandalf.java.grpc.CommandMessageUUID prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * Protobuf type {@code grpc.CommandMessageUUID}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:grpc.CommandMessageUUID)
      com.ditrit.gandalf.java.grpc.CommandMessageUUIDOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return com.ditrit.gandalf.java.grpc.ConnectorCommandProto.internal_static_grpc_CommandMessageUUID_descriptor;
    }

    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return com.ditrit.gandalf.java.grpc.ConnectorCommandProto.internal_static_grpc_CommandMessageUUID_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              com.ditrit.gandalf.java.grpc.CommandMessageUUID.class, com.ditrit.gandalf.java.grpc.CommandMessageUUID.Builder.class);
    }

    // Construct using com.ditrit.gandalf.java.grpc.CommandMessageUUID.newBuilder()
    private Builder() {
      maybeForceBuilderInitialization();
    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);
      maybeForceBuilderInitialization();
    }
    private void maybeForceBuilderInitialization() {
      if (com.google.protobuf.GeneratedMessageV3
              .alwaysUseFieldBuilders) {
      }
    }
    public Builder clear() {
      super.clear();
      uUID_ = "";

      return this;
    }

    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return com.ditrit.gandalf.java.grpc.ConnectorCommandProto.internal_static_grpc_CommandMessageUUID_descriptor;
    }

    public com.ditrit.gandalf.java.grpc.CommandMessageUUID getDefaultInstanceForType() {
      return com.ditrit.gandalf.java.grpc.CommandMessageUUID.getDefaultInstance();
    }

    public com.ditrit.gandalf.java.grpc.CommandMessageUUID build() {
      com.ditrit.gandalf.java.grpc.CommandMessageUUID result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    public com.ditrit.gandalf.java.grpc.CommandMessageUUID buildPartial() {
      com.ditrit.gandalf.java.grpc.CommandMessageUUID result = new com.ditrit.gandalf.java.grpc.CommandMessageUUID(this);
      result.uUID_ = uUID_;
      onBuilt();
      return result;
    }

    public Builder clone() {
      return (Builder) super.clone();
    }
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        Object value) {
      return (Builder) super.setField(field, value);
    }
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return (Builder) super.clearField(field);
    }
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return (Builder) super.clearOneof(oneof);
    }
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, Object value) {
      return (Builder) super.setRepeatedField(field, index, value);
    }
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        Object value) {
      return (Builder) super.addRepeatedField(field, value);
    }
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof com.ditrit.gandalf.java.grpc.CommandMessageUUID) {
        return mergeFrom((com.ditrit.gandalf.java.grpc.CommandMessageUUID)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(com.ditrit.gandalf.java.grpc.CommandMessageUUID other) {
      if (other == com.ditrit.gandalf.java.grpc.CommandMessageUUID.getDefaultInstance()) return this;
      if (!other.getUUID().isEmpty()) {
        uUID_ = other.uUID_;
        onChanged();
      }
      onChanged();
      return this;
    }

    public final boolean isInitialized() {
      return true;
    }

    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      com.ditrit.gandalf.java.grpc.CommandMessageUUID parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (com.ditrit.gandalf.java.grpc.CommandMessageUUID) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }

    private java.lang.Object uUID_ = "";
    /**
     * <code>optional string UUID = 1;</code>
     */
    public java.lang.String getUUID() {
      java.lang.Object ref = uUID_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        uUID_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>optional string UUID = 1;</code>
     */
    public com.google.protobuf.ByteString
        getUUIDBytes() {
      java.lang.Object ref = uUID_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        uUID_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>optional string UUID = 1;</code>
     */
    public Builder setUUID(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      uUID_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>optional string UUID = 1;</code>
     */
    public Builder clearUUID() {
      
      uUID_ = getDefaultInstance().getUUID();
      onChanged();
      return this;
    }
    /**
     * <code>optional string UUID = 1;</code>
     */
    public Builder setUUIDBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      uUID_ = value;
      onChanged();
      return this;
    }
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return this;
    }

    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return this;
    }


    // @@protoc_insertion_point(builder_scope:grpc.CommandMessageUUID)
  }

  // @@protoc_insertion_point(class_scope:grpc.CommandMessageUUID)
  private static final com.ditrit.gandalf.java.grpc.CommandMessageUUID DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new com.ditrit.gandalf.java.grpc.CommandMessageUUID();
  }

  public static com.ditrit.gandalf.java.grpc.CommandMessageUUID getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<CommandMessageUUID>
      PARSER = new com.google.protobuf.AbstractParser<CommandMessageUUID>() {
    public CommandMessageUUID parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
        return new CommandMessageUUID(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<CommandMessageUUID> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<CommandMessageUUID> getParserForType() {
    return PARSER;
  }

  public com.ditrit.gandalf.java.grpc.CommandMessageUUID getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

