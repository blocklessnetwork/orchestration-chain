/* eslint-disable */
import { Params } from "../market/params";
import { Order } from "../market/order";
import { CompletedOrder } from "../market/completed_order";
import { ClaimedOrder } from "../market/claimed_order";
import { NodeRegistration } from "../market/node_registration";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "txlabs.blocklesschain.market";

/** GenesisState defines the market module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  orderList: Order[];
  completedOrderList: CompletedOrder[];
  claimedOrderList: ClaimedOrder[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  nodeRegistrationList: NodeRegistration[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.orderList) {
      Order.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.completedOrderList) {
      CompletedOrder.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.claimedOrderList) {
      ClaimedOrder.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.nodeRegistrationList) {
      NodeRegistration.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.orderList = [];
    message.completedOrderList = [];
    message.claimedOrderList = [];
    message.nodeRegistrationList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.orderList.push(Order.decode(reader, reader.uint32()));
          break;
        case 3:
          message.completedOrderList.push(
            CompletedOrder.decode(reader, reader.uint32())
          );
          break;
        case 4:
          message.claimedOrderList.push(
            ClaimedOrder.decode(reader, reader.uint32())
          );
          break;
        case 5:
          message.nodeRegistrationList.push(
            NodeRegistration.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.orderList = [];
    message.completedOrderList = [];
    message.claimedOrderList = [];
    message.nodeRegistrationList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.orderList !== undefined && object.orderList !== null) {
      for (const e of object.orderList) {
        message.orderList.push(Order.fromJSON(e));
      }
    }
    if (
      object.completedOrderList !== undefined &&
      object.completedOrderList !== null
    ) {
      for (const e of object.completedOrderList) {
        message.completedOrderList.push(CompletedOrder.fromJSON(e));
      }
    }
    if (
      object.claimedOrderList !== undefined &&
      object.claimedOrderList !== null
    ) {
      for (const e of object.claimedOrderList) {
        message.claimedOrderList.push(ClaimedOrder.fromJSON(e));
      }
    }
    if (
      object.nodeRegistrationList !== undefined &&
      object.nodeRegistrationList !== null
    ) {
      for (const e of object.nodeRegistrationList) {
        message.nodeRegistrationList.push(NodeRegistration.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.orderList) {
      obj.orderList = message.orderList.map((e) =>
        e ? Order.toJSON(e) : undefined
      );
    } else {
      obj.orderList = [];
    }
    if (message.completedOrderList) {
      obj.completedOrderList = message.completedOrderList.map((e) =>
        e ? CompletedOrder.toJSON(e) : undefined
      );
    } else {
      obj.completedOrderList = [];
    }
    if (message.claimedOrderList) {
      obj.claimedOrderList = message.claimedOrderList.map((e) =>
        e ? ClaimedOrder.toJSON(e) : undefined
      );
    } else {
      obj.claimedOrderList = [];
    }
    if (message.nodeRegistrationList) {
      obj.nodeRegistrationList = message.nodeRegistrationList.map((e) =>
        e ? NodeRegistration.toJSON(e) : undefined
      );
    } else {
      obj.nodeRegistrationList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.orderList = [];
    message.completedOrderList = [];
    message.claimedOrderList = [];
    message.nodeRegistrationList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.orderList !== undefined && object.orderList !== null) {
      for (const e of object.orderList) {
        message.orderList.push(Order.fromPartial(e));
      }
    }
    if (
      object.completedOrderList !== undefined &&
      object.completedOrderList !== null
    ) {
      for (const e of object.completedOrderList) {
        message.completedOrderList.push(CompletedOrder.fromPartial(e));
      }
    }
    if (
      object.claimedOrderList !== undefined &&
      object.claimedOrderList !== null
    ) {
      for (const e of object.claimedOrderList) {
        message.claimedOrderList.push(ClaimedOrder.fromPartial(e));
      }
    }
    if (
      object.nodeRegistrationList !== undefined &&
      object.nodeRegistrationList !== null
    ) {
      for (const e of object.nodeRegistrationList) {
        message.nodeRegistrationList.push(NodeRegistration.fromPartial(e));
      }
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
