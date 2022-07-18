/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "txlabs.blocklesschain.market";

export interface MsgSubmitOrder {
  creator: string;
  functionId: string;
  fuel: string;
}

export interface MsgSubmitOrderResponse {}

export interface MsgSubmitCompletedOrder {
  creator: string;
  orderIndex: string;
}

export interface MsgSubmitCompletedOrderResponse {}

const baseMsgSubmitOrder: object = { creator: "", functionId: "", fuel: "" };

export const MsgSubmitOrder = {
  encode(message: MsgSubmitOrder, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.functionId !== "") {
      writer.uint32(18).string(message.functionId);
    }
    if (message.fuel !== "") {
      writer.uint32(26).string(message.fuel);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSubmitOrder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSubmitOrder } as MsgSubmitOrder;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.functionId = reader.string();
          break;
        case 3:
          message.fuel = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSubmitOrder {
    const message = { ...baseMsgSubmitOrder } as MsgSubmitOrder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
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
    return message;
  },

  toJSON(message: MsgSubmitOrder): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.functionId !== undefined && (obj.functionId = message.functionId);
    message.fuel !== undefined && (obj.fuel = message.fuel);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSubmitOrder>): MsgSubmitOrder {
    const message = { ...baseMsgSubmitOrder } as MsgSubmitOrder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
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
    return message;
  },
};

const baseMsgSubmitOrderResponse: object = {};

export const MsgSubmitOrderResponse = {
  encode(_: MsgSubmitOrderResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSubmitOrderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSubmitOrderResponse } as MsgSubmitOrderResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgSubmitOrderResponse {
    const message = { ...baseMsgSubmitOrderResponse } as MsgSubmitOrderResponse;
    return message;
  },

  toJSON(_: MsgSubmitOrderResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgSubmitOrderResponse>): MsgSubmitOrderResponse {
    const message = { ...baseMsgSubmitOrderResponse } as MsgSubmitOrderResponse;
    return message;
  },
};

const baseMsgSubmitCompletedOrder: object = { creator: "", orderIndex: "" };

export const MsgSubmitCompletedOrder = {
  encode(
    message: MsgSubmitCompletedOrder,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.orderIndex !== "") {
      writer.uint32(18).string(message.orderIndex);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSubmitCompletedOrder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSubmitCompletedOrder,
    } as MsgSubmitCompletedOrder;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.orderIndex = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSubmitCompletedOrder {
    const message = {
      ...baseMsgSubmitCompletedOrder,
    } as MsgSubmitCompletedOrder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.orderIndex !== undefined && object.orderIndex !== null) {
      message.orderIndex = String(object.orderIndex);
    } else {
      message.orderIndex = "";
    }
    return message;
  },

  toJSON(message: MsgSubmitCompletedOrder): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.orderIndex !== undefined && (obj.orderIndex = message.orderIndex);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSubmitCompletedOrder>
  ): MsgSubmitCompletedOrder {
    const message = {
      ...baseMsgSubmitCompletedOrder,
    } as MsgSubmitCompletedOrder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.orderIndex !== undefined && object.orderIndex !== null) {
      message.orderIndex = object.orderIndex;
    } else {
      message.orderIndex = "";
    }
    return message;
  },
};

const baseMsgSubmitCompletedOrderResponse: object = {};

export const MsgSubmitCompletedOrderResponse = {
  encode(
    _: MsgSubmitCompletedOrderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSubmitCompletedOrderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSubmitCompletedOrderResponse,
    } as MsgSubmitCompletedOrderResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgSubmitCompletedOrderResponse {
    const message = {
      ...baseMsgSubmitCompletedOrderResponse,
    } as MsgSubmitCompletedOrderResponse;
    return message;
  },

  toJSON(_: MsgSubmitCompletedOrderResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgSubmitCompletedOrderResponse>
  ): MsgSubmitCompletedOrderResponse {
    const message = {
      ...baseMsgSubmitCompletedOrderResponse,
    } as MsgSubmitCompletedOrderResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  SubmitOrder(request: MsgSubmitOrder): Promise<MsgSubmitOrderResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SubmitCompletedOrder(
    request: MsgSubmitCompletedOrder
  ): Promise<MsgSubmitCompletedOrderResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  SubmitOrder(request: MsgSubmitOrder): Promise<MsgSubmitOrderResponse> {
    const data = MsgSubmitOrder.encode(request).finish();
    const promise = this.rpc.request(
      "txlabs.blocklesschain.market.Msg",
      "SubmitOrder",
      data
    );
    return promise.then((data) =>
      MsgSubmitOrderResponse.decode(new Reader(data))
    );
  }

  SubmitCompletedOrder(
    request: MsgSubmitCompletedOrder
  ): Promise<MsgSubmitCompletedOrderResponse> {
    const data = MsgSubmitCompletedOrder.encode(request).finish();
    const promise = this.rpc.request(
      "txlabs.blocklesschain.market.Msg",
      "SubmitCompletedOrder",
      data
    );
    return promise.then((data) =>
      MsgSubmitCompletedOrderResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
