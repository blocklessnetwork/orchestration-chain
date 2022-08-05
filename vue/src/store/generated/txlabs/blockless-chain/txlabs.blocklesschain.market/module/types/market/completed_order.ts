/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "txlabs.blocklesschain.market";

export interface CompletedOrder {
  index: string;
  orderIndex: string;
  completedBy: string;
  height: string;
  date: string;
  fuelUsed: string;
  responseId: string;
}

const baseCompletedOrder: object = {
  index: "",
  orderIndex: "",
  completedBy: "",
  height: "",
  date: "",
  fuelUsed: "",
  responseId: "",
};

export const CompletedOrder = {
  encode(message: CompletedOrder, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.orderIndex !== "") {
      writer.uint32(18).string(message.orderIndex);
    }
    if (message.completedBy !== "") {
      writer.uint32(26).string(message.completedBy);
    }
    if (message.height !== "") {
      writer.uint32(34).string(message.height);
    }
    if (message.date !== "") {
      writer.uint32(42).string(message.date);
    }
    if (message.fuelUsed !== "") {
      writer.uint32(50).string(message.fuelUsed);
    }
    if (message.responseId !== "") {
      writer.uint32(58).string(message.responseId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): CompletedOrder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseCompletedOrder } as CompletedOrder;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.orderIndex = reader.string();
          break;
        case 3:
          message.completedBy = reader.string();
          break;
        case 4:
          message.height = reader.string();
          break;
        case 5:
          message.date = reader.string();
          break;
        case 6:
          message.fuelUsed = reader.string();
          break;
        case 7:
          message.responseId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CompletedOrder {
    const message = { ...baseCompletedOrder } as CompletedOrder;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.orderIndex !== undefined && object.orderIndex !== null) {
      message.orderIndex = String(object.orderIndex);
    } else {
      message.orderIndex = "";
    }
    if (object.completedBy !== undefined && object.completedBy !== null) {
      message.completedBy = String(object.completedBy);
    } else {
      message.completedBy = "";
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
    if (object.fuelUsed !== undefined && object.fuelUsed !== null) {
      message.fuelUsed = String(object.fuelUsed);
    } else {
      message.fuelUsed = "";
    }
    if (object.responseId !== undefined && object.responseId !== null) {
      message.responseId = String(object.responseId);
    } else {
      message.responseId = "";
    }
    return message;
  },

  toJSON(message: CompletedOrder): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.orderIndex !== undefined && (obj.orderIndex = message.orderIndex);
    message.completedBy !== undefined &&
      (obj.completedBy = message.completedBy);
    message.height !== undefined && (obj.height = message.height);
    message.date !== undefined && (obj.date = message.date);
    message.fuelUsed !== undefined && (obj.fuelUsed = message.fuelUsed);
    message.responseId !== undefined && (obj.responseId = message.responseId);
    return obj;
  },

  fromPartial(object: DeepPartial<CompletedOrder>): CompletedOrder {
    const message = { ...baseCompletedOrder } as CompletedOrder;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.orderIndex !== undefined && object.orderIndex !== null) {
      message.orderIndex = object.orderIndex;
    } else {
      message.orderIndex = "";
    }
    if (object.completedBy !== undefined && object.completedBy !== null) {
      message.completedBy = object.completedBy;
    } else {
      message.completedBy = "";
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
    if (object.fuelUsed !== undefined && object.fuelUsed !== null) {
      message.fuelUsed = object.fuelUsed;
    } else {
      message.fuelUsed = "";
    }
    if (object.responseId !== undefined && object.responseId !== null) {
      message.responseId = object.responseId;
    } else {
      message.responseId = "";
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
