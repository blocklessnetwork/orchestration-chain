/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "txlabs.blocklesschain.market";

export interface Order {
  index: string;
  functionId: string;
  fuel: string;
  customer: string;
  height: string;
  date: string;
  state: string;
}

export interface OrderFilter {
  customer: string;
}

const baseOrder: object = {
  index: "",
  functionId: "",
  fuel: "",
  customer: "",
  height: "",
  date: "",
  state: "",
};

export const Order = {
  encode(message: Order, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.functionId !== "") {
      writer.uint32(18).string(message.functionId);
    }
    if (message.fuel !== "") {
      writer.uint32(26).string(message.fuel);
    }
    if (message.customer !== "") {
      writer.uint32(34).string(message.customer);
    }
    if (message.height !== "") {
      writer.uint32(42).string(message.height);
    }
    if (message.date !== "") {
      writer.uint32(50).string(message.date);
    }
    if (message.state !== "") {
      writer.uint32(58).string(message.state);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Order {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOrder } as Order;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.functionId = reader.string();
          break;
        case 3:
          message.fuel = reader.string();
          break;
        case 4:
          message.customer = reader.string();
          break;
        case 5:
          message.height = reader.string();
          break;
        case 6:
          message.date = reader.string();
          break;
        case 7:
          message.state = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Order {
    const message = { ...baseOrder } as Order;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.functionId !== undefined && object.functionId !== null) {
      message.functionId = String(object.functionId);
    } else {
      message.functionId = "";
    }
    if (object.fuel !== undefined && object.fuel !== null) {
      message.fuel = String(object.fuel);
    } else {
      message.fuel = "";
    }
    if (object.customer !== undefined && object.customer !== null) {
      message.customer = String(object.customer);
    } else {
      message.customer = "";
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
    if (object.state !== undefined && object.state !== null) {
      message.state = String(object.state);
    } else {
      message.state = "";
    }
    return message;
  },

  toJSON(message: Order): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.functionId !== undefined && (obj.functionId = message.functionId);
    message.fuel !== undefined && (obj.fuel = message.fuel);
    message.customer !== undefined && (obj.customer = message.customer);
    message.height !== undefined && (obj.height = message.height);
    message.date !== undefined && (obj.date = message.date);
    message.state !== undefined && (obj.state = message.state);
    return obj;
  },

  fromPartial(object: DeepPartial<Order>): Order {
    const message = { ...baseOrder } as Order;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.functionId !== undefined && object.functionId !== null) {
      message.functionId = object.functionId;
    } else {
      message.functionId = "";
    }
    if (object.fuel !== undefined && object.fuel !== null) {
      message.fuel = object.fuel;
    } else {
      message.fuel = "";
    }
    if (object.customer !== undefined && object.customer !== null) {
      message.customer = object.customer;
    } else {
      message.customer = "";
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
    if (object.state !== undefined && object.state !== null) {
      message.state = object.state;
    } else {
      message.state = "";
    }
    return message;
  },
};

const baseOrderFilter: object = { customer: "" };

export const OrderFilter = {
  encode(message: OrderFilter, writer: Writer = Writer.create()): Writer {
    if (message.customer !== "") {
      writer.uint32(10).string(message.customer);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): OrderFilter {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOrderFilter } as OrderFilter;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.customer = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OrderFilter {
    const message = { ...baseOrderFilter } as OrderFilter;
    if (object.customer !== undefined && object.customer !== null) {
      message.customer = String(object.customer);
    } else {
      message.customer = "";
    }
    return message;
  },

  toJSON(message: OrderFilter): unknown {
    const obj: any = {};
    message.customer !== undefined && (obj.customer = message.customer);
    return obj;
  },

  fromPartial(object: DeepPartial<OrderFilter>): OrderFilter {
    const message = { ...baseOrderFilter } as OrderFilter;
    if (object.customer !== undefined && object.customer !== null) {
      message.customer = object.customer;
    } else {
      message.customer = "";
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
