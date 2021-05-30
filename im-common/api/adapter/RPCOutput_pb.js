/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.adapter.RPCOutput', null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.adapter.RPCOutput = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.adapter.RPCOutput, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.adapter.RPCOutput.displayName = 'proto.adapter.RPCOutput';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.adapter.RPCOutput.prototype.toObject = function(opt_includeInstance) {
  return proto.adapter.RPCOutput.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.adapter.RPCOutput} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.adapter.RPCOutput.toObject = function(includeInstance, msg) {
  var f, obj = {
    ret: jspb.Message.getFieldWithDefault(msg, 1, 0),
    rsp: msg.getRsp_asB64(),
    optMap: (f = msg.getOptMap()) ? f.toObject(includeInstance, undefined) : [],
    desc: jspb.Message.getFieldWithDefault(msg, 4, ""),
    servername: jspb.Message.getFieldWithDefault(msg, 5, ""),
    func: jspb.Message.getFieldWithDefault(msg, 6, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.adapter.RPCOutput}
 */
proto.adapter.RPCOutput.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.adapter.RPCOutput;
  return proto.adapter.RPCOutput.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.adapter.RPCOutput} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.adapter.RPCOutput}
 */
proto.adapter.RPCOutput.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setRet(value);
      break;
    case 2:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setRsp(value);
      break;
    case 3:
      var value = msg.getOptMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readInt32, jspb.BinaryReader.prototype.readString);
         });
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setDesc(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setServername(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setFunc(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.adapter.RPCOutput.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.adapter.RPCOutput.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.adapter.RPCOutput} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.adapter.RPCOutput.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRet();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = message.getRsp_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      2,
      f
    );
  }
  f = message.getOptMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(3, writer, jspb.BinaryWriter.prototype.writeInt32, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getDesc();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getServername();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getFunc();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
};


/**
 * optional int32 Ret = 1;
 * @return {number}
 */
proto.adapter.RPCOutput.prototype.getRet = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.adapter.RPCOutput.prototype.setRet = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional bytes Rsp = 2;
 * @return {!(string|Uint8Array)}
 */
proto.adapter.RPCOutput.prototype.getRsp = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * optional bytes Rsp = 2;
 * This is a type-conversion wrapper around `getRsp()`
 * @return {string}
 */
proto.adapter.RPCOutput.prototype.getRsp_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getRsp()));
};


/**
 * optional bytes Rsp = 2;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getRsp()`
 * @return {!Uint8Array}
 */
proto.adapter.RPCOutput.prototype.getRsp_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getRsp()));
};


/** @param {!(string|Uint8Array)} value */
proto.adapter.RPCOutput.prototype.setRsp = function(value) {
  jspb.Message.setProto3BytesField(this, 2, value);
};


/**
 * map<int32, string> Opt = 3;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<number,string>}
 */
proto.adapter.RPCOutput.prototype.getOptMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<number,string>} */ (
      jspb.Message.getMapField(this, 3, opt_noLazyCreate,
      null));
};


proto.adapter.RPCOutput.prototype.clearOptMap = function() {
  this.getOptMap().clear();
};


/**
 * optional string Desc = 4;
 * @return {string}
 */
proto.adapter.RPCOutput.prototype.getDesc = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.adapter.RPCOutput.prototype.setDesc = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string ServerName = 5;
 * @return {string}
 */
proto.adapter.RPCOutput.prototype.getServername = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/** @param {string} value */
proto.adapter.RPCOutput.prototype.setServername = function(value) {
  jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string Func = 6;
 * @return {string}
 */
proto.adapter.RPCOutput.prototype.getFunc = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/** @param {string} value */
proto.adapter.RPCOutput.prototype.setFunc = function(value) {
  jspb.Message.setProto3StringField(this, 6, value);
};


goog.object.extend(exports, proto.adapter);