/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../market/params";
import { Order, OrderFilter } from "../market/order";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { CompletedOrder } from "../market/completed_order";
import { ClaimedOrder } from "../market/claimed_order";

export const protobufPackage = "txlabs.blocklesschain.market";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetOrderRequest {
  index: string;
}

export interface QueryGetOrderResponse {
  order: Order | undefined;
}

export interface QueryAllOrderRequest {
  filter: OrderFilter | undefined;
  pagination: PageRequest | undefined;
}

export interface QueryAllOrderResponse {
  order: Order[];
  pagination: PageResponse | undefined;
}

export interface QueryGetCompletedOrderRequest {
  index: string;
}

export interface QueryGetCompletedOrderResponse {
  completedOrder: CompletedOrder | undefined;
}

export interface QueryAllCompletedOrderRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllCompletedOrderResponse {
  completedOrder: CompletedOrder[];
  pagination: PageResponse | undefined;
}

export interface QueryGetClaimedOrderRequest {
  index: string;
}

export interface QueryGetClaimedOrderResponse {
  claimedOrder: ClaimedOrder | undefined;
}

export interface QueryAllClaimedOrderRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllClaimedOrderResponse {
  claimedOrder: ClaimedOrder[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetOrderRequest: object = { index: "" };

export const QueryGetOrderRequest = {
  encode(
    message: QueryGetOrderRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetOrderRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetOrderRequest } as QueryGetOrderRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetOrderRequest {
    const message = { ...baseQueryGetOrderRequest } as QueryGetOrderRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetOrderRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetOrderRequest>): QueryGetOrderRequest {
    const message = { ...baseQueryGetOrderRequest } as QueryGetOrderRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetOrderResponse: object = {};

export const QueryGetOrderResponse = {
  encode(
    message: QueryGetOrderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.order !== undefined) {
      Order.encode(message.order, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetOrderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetOrderResponse } as QueryGetOrderResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.order = Order.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetOrderResponse {
    const message = { ...baseQueryGetOrderResponse } as QueryGetOrderResponse;
    if (object.order !== undefined && object.order !== null) {
      message.order = Order.fromJSON(object.order);
    } else {
      message.order = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetOrderResponse): unknown {
    const obj: any = {};
    message.order !== undefined &&
      (obj.order = message.order ? Order.toJSON(message.order) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetOrderResponse>
  ): QueryGetOrderResponse {
    const message = { ...baseQueryGetOrderResponse } as QueryGetOrderResponse;
    if (object.order !== undefined && object.order !== null) {
      message.order = Order.fromPartial(object.order);
    } else {
      message.order = undefined;
    }
    return message;
  },
};

const baseQueryAllOrderRequest: object = {};

export const QueryAllOrderRequest = {
  encode(
    message: QueryAllOrderRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.filter !== undefined) {
      OrderFilter.encode(message.filter, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllOrderRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllOrderRequest } as QueryAllOrderRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.filter = OrderFilter.decode(reader, reader.uint32());
          break;
        case 2:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllOrderRequest {
    const message = { ...baseQueryAllOrderRequest } as QueryAllOrderRequest;
    if (object.filter !== undefined && object.filter !== null) {
      message.filter = OrderFilter.fromJSON(object.filter);
    } else {
      message.filter = undefined;
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllOrderRequest): unknown {
    const obj: any = {};
    message.filter !== undefined &&
      (obj.filter = message.filter
        ? OrderFilter.toJSON(message.filter)
        : undefined);
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllOrderRequest>): QueryAllOrderRequest {
    const message = { ...baseQueryAllOrderRequest } as QueryAllOrderRequest;
    if (object.filter !== undefined && object.filter !== null) {
      message.filter = OrderFilter.fromPartial(object.filter);
    } else {
      message.filter = undefined;
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllOrderResponse: object = {};

export const QueryAllOrderResponse = {
  encode(
    message: QueryAllOrderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.order) {
      Order.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllOrderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllOrderResponse } as QueryAllOrderResponse;
    message.order = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.order.push(Order.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllOrderResponse {
    const message = { ...baseQueryAllOrderResponse } as QueryAllOrderResponse;
    message.order = [];
    if (object.order !== undefined && object.order !== null) {
      for (const e of object.order) {
        message.order.push(Order.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllOrderResponse): unknown {
    const obj: any = {};
    if (message.order) {
      obj.order = message.order.map((e) => (e ? Order.toJSON(e) : undefined));
    } else {
      obj.order = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllOrderResponse>
  ): QueryAllOrderResponse {
    const message = { ...baseQueryAllOrderResponse } as QueryAllOrderResponse;
    message.order = [];
    if (object.order !== undefined && object.order !== null) {
      for (const e of object.order) {
        message.order.push(Order.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetCompletedOrderRequest: object = { index: "" };

export const QueryGetCompletedOrderRequest = {
  encode(
    message: QueryGetCompletedOrderRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetCompletedOrderRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetCompletedOrderRequest,
    } as QueryGetCompletedOrderRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetCompletedOrderRequest {
    const message = {
      ...baseQueryGetCompletedOrderRequest,
    } as QueryGetCompletedOrderRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetCompletedOrderRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetCompletedOrderRequest>
  ): QueryGetCompletedOrderRequest {
    const message = {
      ...baseQueryGetCompletedOrderRequest,
    } as QueryGetCompletedOrderRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetCompletedOrderResponse: object = {};

export const QueryGetCompletedOrderResponse = {
  encode(
    message: QueryGetCompletedOrderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.completedOrder !== undefined) {
      CompletedOrder.encode(
        message.completedOrder,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetCompletedOrderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetCompletedOrderResponse,
    } as QueryGetCompletedOrderResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.completedOrder = CompletedOrder.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetCompletedOrderResponse {
    const message = {
      ...baseQueryGetCompletedOrderResponse,
    } as QueryGetCompletedOrderResponse;
    if (object.completedOrder !== undefined && object.completedOrder !== null) {
      message.completedOrder = CompletedOrder.fromJSON(object.completedOrder);
    } else {
      message.completedOrder = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetCompletedOrderResponse): unknown {
    const obj: any = {};
    message.completedOrder !== undefined &&
      (obj.completedOrder = message.completedOrder
        ? CompletedOrder.toJSON(message.completedOrder)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetCompletedOrderResponse>
  ): QueryGetCompletedOrderResponse {
    const message = {
      ...baseQueryGetCompletedOrderResponse,
    } as QueryGetCompletedOrderResponse;
    if (object.completedOrder !== undefined && object.completedOrder !== null) {
      message.completedOrder = CompletedOrder.fromPartial(
        object.completedOrder
      );
    } else {
      message.completedOrder = undefined;
    }
    return message;
  },
};

const baseQueryAllCompletedOrderRequest: object = {};

export const QueryAllCompletedOrderRequest = {
  encode(
    message: QueryAllCompletedOrderRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllCompletedOrderRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllCompletedOrderRequest,
    } as QueryAllCompletedOrderRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllCompletedOrderRequest {
    const message = {
      ...baseQueryAllCompletedOrderRequest,
    } as QueryAllCompletedOrderRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllCompletedOrderRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllCompletedOrderRequest>
  ): QueryAllCompletedOrderRequest {
    const message = {
      ...baseQueryAllCompletedOrderRequest,
    } as QueryAllCompletedOrderRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllCompletedOrderResponse: object = {};

export const QueryAllCompletedOrderResponse = {
  encode(
    message: QueryAllCompletedOrderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.completedOrder) {
      CompletedOrder.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllCompletedOrderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllCompletedOrderResponse,
    } as QueryAllCompletedOrderResponse;
    message.completedOrder = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.completedOrder.push(
            CompletedOrder.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllCompletedOrderResponse {
    const message = {
      ...baseQueryAllCompletedOrderResponse,
    } as QueryAllCompletedOrderResponse;
    message.completedOrder = [];
    if (object.completedOrder !== undefined && object.completedOrder !== null) {
      for (const e of object.completedOrder) {
        message.completedOrder.push(CompletedOrder.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllCompletedOrderResponse): unknown {
    const obj: any = {};
    if (message.completedOrder) {
      obj.completedOrder = message.completedOrder.map((e) =>
        e ? CompletedOrder.toJSON(e) : undefined
      );
    } else {
      obj.completedOrder = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllCompletedOrderResponse>
  ): QueryAllCompletedOrderResponse {
    const message = {
      ...baseQueryAllCompletedOrderResponse,
    } as QueryAllCompletedOrderResponse;
    message.completedOrder = [];
    if (object.completedOrder !== undefined && object.completedOrder !== null) {
      for (const e of object.completedOrder) {
        message.completedOrder.push(CompletedOrder.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetClaimedOrderRequest: object = { index: "" };

export const QueryGetClaimedOrderRequest = {
  encode(
    message: QueryGetClaimedOrderRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetClaimedOrderRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetClaimedOrderRequest,
    } as QueryGetClaimedOrderRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetClaimedOrderRequest {
    const message = {
      ...baseQueryGetClaimedOrderRequest,
    } as QueryGetClaimedOrderRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetClaimedOrderRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetClaimedOrderRequest>
  ): QueryGetClaimedOrderRequest {
    const message = {
      ...baseQueryGetClaimedOrderRequest,
    } as QueryGetClaimedOrderRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetClaimedOrderResponse: object = {};

export const QueryGetClaimedOrderResponse = {
  encode(
    message: QueryGetClaimedOrderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.claimedOrder !== undefined) {
      ClaimedOrder.encode(
        message.claimedOrder,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetClaimedOrderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetClaimedOrderResponse,
    } as QueryGetClaimedOrderResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.claimedOrder = ClaimedOrder.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetClaimedOrderResponse {
    const message = {
      ...baseQueryGetClaimedOrderResponse,
    } as QueryGetClaimedOrderResponse;
    if (object.claimedOrder !== undefined && object.claimedOrder !== null) {
      message.claimedOrder = ClaimedOrder.fromJSON(object.claimedOrder);
    } else {
      message.claimedOrder = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetClaimedOrderResponse): unknown {
    const obj: any = {};
    message.claimedOrder !== undefined &&
      (obj.claimedOrder = message.claimedOrder
        ? ClaimedOrder.toJSON(message.claimedOrder)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetClaimedOrderResponse>
  ): QueryGetClaimedOrderResponse {
    const message = {
      ...baseQueryGetClaimedOrderResponse,
    } as QueryGetClaimedOrderResponse;
    if (object.claimedOrder !== undefined && object.claimedOrder !== null) {
      message.claimedOrder = ClaimedOrder.fromPartial(object.claimedOrder);
    } else {
      message.claimedOrder = undefined;
    }
    return message;
  },
};

const baseQueryAllClaimedOrderRequest: object = {};

export const QueryAllClaimedOrderRequest = {
  encode(
    message: QueryAllClaimedOrderRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllClaimedOrderRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllClaimedOrderRequest,
    } as QueryAllClaimedOrderRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllClaimedOrderRequest {
    const message = {
      ...baseQueryAllClaimedOrderRequest,
    } as QueryAllClaimedOrderRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllClaimedOrderRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllClaimedOrderRequest>
  ): QueryAllClaimedOrderRequest {
    const message = {
      ...baseQueryAllClaimedOrderRequest,
    } as QueryAllClaimedOrderRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllClaimedOrderResponse: object = {};

export const QueryAllClaimedOrderResponse = {
  encode(
    message: QueryAllClaimedOrderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.claimedOrder) {
      ClaimedOrder.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllClaimedOrderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllClaimedOrderResponse,
    } as QueryAllClaimedOrderResponse;
    message.claimedOrder = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.claimedOrder.push(
            ClaimedOrder.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllClaimedOrderResponse {
    const message = {
      ...baseQueryAllClaimedOrderResponse,
    } as QueryAllClaimedOrderResponse;
    message.claimedOrder = [];
    if (object.claimedOrder !== undefined && object.claimedOrder !== null) {
      for (const e of object.claimedOrder) {
        message.claimedOrder.push(ClaimedOrder.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllClaimedOrderResponse): unknown {
    const obj: any = {};
    if (message.claimedOrder) {
      obj.claimedOrder = message.claimedOrder.map((e) =>
        e ? ClaimedOrder.toJSON(e) : undefined
      );
    } else {
      obj.claimedOrder = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllClaimedOrderResponse>
  ): QueryAllClaimedOrderResponse {
    const message = {
      ...baseQueryAllClaimedOrderResponse,
    } as QueryAllClaimedOrderResponse;
    message.claimedOrder = [];
    if (object.claimedOrder !== undefined && object.claimedOrder !== null) {
      for (const e of object.claimedOrder) {
        message.claimedOrder.push(ClaimedOrder.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a Order by index. */
  Order(request: QueryGetOrderRequest): Promise<QueryGetOrderResponse>;
  /** Queries a list of Order items. */
  OrderAll(request: QueryAllOrderRequest): Promise<QueryAllOrderResponse>;
  /** Queries a CompletedOrder by index. */
  CompletedOrder(
    request: QueryGetCompletedOrderRequest
  ): Promise<QueryGetCompletedOrderResponse>;
  /** Queries a list of CompletedOrder items. */
  CompletedOrderAll(
    request: QueryAllCompletedOrderRequest
  ): Promise<QueryAllCompletedOrderResponse>;
  /** Queries a ClaimedOrder by index. */
  ClaimedOrder(
    request: QueryGetClaimedOrderRequest
  ): Promise<QueryGetClaimedOrderResponse>;
  /** Queries a list of ClaimedOrder items. */
  ClaimedOrderAll(
    request: QueryAllClaimedOrderRequest
  ): Promise<QueryAllClaimedOrderResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "txlabs.blocklesschain.market.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Order(request: QueryGetOrderRequest): Promise<QueryGetOrderResponse> {
    const data = QueryGetOrderRequest.encode(request).finish();
    const promise = this.rpc.request(
      "txlabs.blocklesschain.market.Query",
      "Order",
      data
    );
    return promise.then((data) =>
      QueryGetOrderResponse.decode(new Reader(data))
    );
  }

  OrderAll(request: QueryAllOrderRequest): Promise<QueryAllOrderResponse> {
    const data = QueryAllOrderRequest.encode(request).finish();
    const promise = this.rpc.request(
      "txlabs.blocklesschain.market.Query",
      "OrderAll",
      data
    );
    return promise.then((data) =>
      QueryAllOrderResponse.decode(new Reader(data))
    );
  }

  CompletedOrder(
    request: QueryGetCompletedOrderRequest
  ): Promise<QueryGetCompletedOrderResponse> {
    const data = QueryGetCompletedOrderRequest.encode(request).finish();
    const promise = this.rpc.request(
      "txlabs.blocklesschain.market.Query",
      "CompletedOrder",
      data
    );
    return promise.then((data) =>
      QueryGetCompletedOrderResponse.decode(new Reader(data))
    );
  }

  CompletedOrderAll(
    request: QueryAllCompletedOrderRequest
  ): Promise<QueryAllCompletedOrderResponse> {
    const data = QueryAllCompletedOrderRequest.encode(request).finish();
    const promise = this.rpc.request(
      "txlabs.blocklesschain.market.Query",
      "CompletedOrderAll",
      data
    );
    return promise.then((data) =>
      QueryAllCompletedOrderResponse.decode(new Reader(data))
    );
  }

  ClaimedOrder(
    request: QueryGetClaimedOrderRequest
  ): Promise<QueryGetClaimedOrderResponse> {
    const data = QueryGetClaimedOrderRequest.encode(request).finish();
    const promise = this.rpc.request(
      "txlabs.blocklesschain.market.Query",
      "ClaimedOrder",
      data
    );
    return promise.then((data) =>
      QueryGetClaimedOrderResponse.decode(new Reader(data))
    );
  }

  ClaimedOrderAll(
    request: QueryAllClaimedOrderRequest
  ): Promise<QueryAllClaimedOrderResponse> {
    const data = QueryAllClaimedOrderRequest.encode(request).finish();
    const promise = this.rpc.request(
      "txlabs.blocklesschain.market.Query",
      "ClaimedOrderAll",
      data
    );
    return promise.then((data) =>
      QueryAllClaimedOrderResponse.decode(new Reader(data))
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
