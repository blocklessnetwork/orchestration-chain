/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "txlabs.blocklesschain.market";

export interface NodeRegistration {
  index: string;
  nodeId: string;
  nodePort: string;
  nodeIp: string;
  nodeOwner: string;
  height: string;
  date: string;
}

const baseNodeRegistration: object = {
  index: "",
  nodeId: "",
  nodePort: "",
  nodeIp: "",
  nodeOwner: "",
  height: "",
  date: "",
};

export const NodeRegistration = {
  encode(message: NodeRegistration, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.nodeId !== "") {
      writer.uint32(18).string(message.nodeId);
    }
    if (message.nodePort !== "") {
      writer.uint32(26).string(message.nodePort);
    }
    if (message.nodeIp !== "") {
      writer.uint32(34).string(message.nodeIp);
    }
    if (message.nodeOwner !== "") {
      writer.uint32(42).string(message.nodeOwner);
    }
    if (message.height !== "") {
      writer.uint32(50).string(message.height);
    }
    if (message.date !== "") {
      writer.uint32(58).string(message.date);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): NodeRegistration {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNodeRegistration } as NodeRegistration;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.nodeId = reader.string();
          break;
        case 3:
          message.nodePort = reader.string();
          break;
        case 4:
          message.nodeIp = reader.string();
          break;
        case 5:
          message.nodeOwner = reader.string();
          break;
        case 6:
          message.height = reader.string();
          break;
        case 7:
          message.date = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): NodeRegistration {
    const message = { ...baseNodeRegistration } as NodeRegistration;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.nodeId !== undefined && object.nodeId !== null) {
      message.nodeId = String(object.nodeId);
    } else {
      message.nodeId = "";
    }
    if (object.nodePort !== undefined && object.nodePort !== null) {
      message.nodePort = String(object.nodePort);
    } else {
      message.nodePort = "";
    }
    if (object.nodeIp !== undefined && object.nodeIp !== null) {
      message.nodeIp = String(object.nodeIp);
    } else {
      message.nodeIp = "";
    }
    if (object.nodeOwner !== undefined && object.nodeOwner !== null) {
      message.nodeOwner = String(object.nodeOwner);
    } else {
      message.nodeOwner = "";
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = String(object.height);
    } else {
      message.height = "";
    }
    if (object.date !== undefined && object.date !== null) {
      message.date = String(object.date);
    } else {
      message.date = "";
    }
    return message;
  },

  toJSON(message: NodeRegistration): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.nodeId !== undefined && (obj.nodeId = message.nodeId);
    message.nodePort !== undefined && (obj.nodePort = message.nodePort);
    message.nodeIp !== undefined && (obj.nodeIp = message.nodeIp);
    message.nodeOwner !== undefined && (obj.nodeOwner = message.nodeOwner);
    message.height !== undefined && (obj.height = message.height);
    message.date !== undefined && (obj.date = message.date);
    return obj;
  },

  fromPartial(object: DeepPartial<NodeRegistration>): NodeRegistration {
    const message = { ...baseNodeRegistration } as NodeRegistration;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.nodeId !== undefined && object.nodeId !== null) {
      message.nodeId = object.nodeId;
    } else {
      message.nodeId = "";
    }
    if (object.nodePort !== undefined && object.nodePort !== null) {
      message.nodePort = object.nodePort;
    } else {
      message.nodePort = "";
    }
    if (object.nodeIp !== undefined && object.nodeIp !== null) {
      message.nodeIp = object.nodeIp;
    } else {
      message.nodeIp = "";
    }
    if (object.nodeOwner !== undefined && object.nodeOwner !== null) {
      message.nodeOwner = object.nodeOwner;
    } else {
      message.nodeOwner = "";
    }
    if (object.height !== undefined && object.height !== null) {
      message.height = object.height;
    } else {
      message.height = "";
    }
    if (object.date !== undefined && object.date !== null) {
      message.date = object.date;
    } else {
      message.date = "";
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
