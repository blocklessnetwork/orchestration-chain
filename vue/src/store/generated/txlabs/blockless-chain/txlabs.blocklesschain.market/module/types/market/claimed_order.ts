/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "txlabs.blocklesschain.market";

export interface ClaimedOrder {
  index: string;
  claimedBy: string;
  claimedHeight: string;
  date: string;
}

const baseClaimedOrder: object = {
  index: "",
  claimedBy: "",
  claimedHeight: "",
  date: "",
};

export const ClaimedOrder = {
  encode(message: ClaimedOrder, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.claimedBy !== "") {
      writer.uint32(26).string(message.claimedBy);
    }
    if (message.claimedHeight !== "") {
      writer.uint32(34).string(message.claimedHeight);
    }
    if (message.date !== "") {
      writer.uint32(42).string(message.date);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ClaimedOrder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseClaimedOrder } as ClaimedOrder;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 3:
          message.claimedBy = reader.string();
          break;
        case 4:
          message.claimedHeight = reader.string();
          break;
        case 5:
          message.date = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ClaimedOrder {
    const message = { ...baseClaimedOrder } as ClaimedOrder;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.claimedBy !== undefined && object.claimedBy !== null) {
      message.claimedBy = String(object.claimedBy);
    } else {
      message.claimedBy = "";
    }
    if (object.claimedHeight !== undefined && object.claimedHeight !== null) {
      message.claimedHeight = String(object.claimedHeight);
    } else {
      message.claimedHeight = "";
    }
    if (object.date !== undefined && object.date !== null) {
      message.date = String(object.date);
    } else {
      message.date = "";
    }
    return message;
  },

  toJSON(message: ClaimedOrder): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.claimedBy !== undefined && (obj.claimedBy = message.claimedBy);
    message.claimedHeight !== undefined &&
      (obj.claimedHeight = message.claimedHeight);
    message.date !== undefined && (obj.date = message.date);
    return obj;
  },

  fromPartial(object: DeepPartial<ClaimedOrder>): ClaimedOrder {
    const message = { ...baseClaimedOrder } as ClaimedOrder;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.claimedBy !== undefined && object.claimedBy !== null) {
      message.claimedBy = object.claimedBy;
    } else {
      message.claimedBy = "";
    }
    if (object.claimedHeight !== undefined && object.claimedHeight !== null) {
      message.claimedHeight = object.claimedHeight;
    } else {
      message.claimedHeight = "";
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
