# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from savourrpc import market_pb2 as savourrpc_dot_market__pb2


class PriceServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.getSymbols = channel.unary_unary(
                '/savourrpc.market.PriceService/getSymbols',
                request_serializer=savourrpc_dot_market__pb2.SymbolRequest.SerializeToString,
                response_deserializer=savourrpc_dot_market__pb2.SymbolResponse.FromString,
                )
        self.getSymbolPrices = channel.unary_unary(
                '/savourrpc.market.PriceService/getSymbolPrices',
                request_serializer=savourrpc_dot_market__pb2.SymbolPriceRequest.SerializeToString,
                response_deserializer=savourrpc_dot_market__pb2.SymbolPriceResponse.FromString,
                )
        self.getStableCoins = channel.unary_unary(
                '/savourrpc.market.PriceService/getStableCoins',
                request_serializer=savourrpc_dot_market__pb2.StableCoinRequest.SerializeToString,
                response_deserializer=savourrpc_dot_market__pb2.StableCoinResponse.FromString,
                )
        self.getStableCoinPrice = channel.unary_unary(
                '/savourrpc.market.PriceService/getStableCoinPrice',
                request_serializer=savourrpc_dot_market__pb2.StableCoinPriceRequest.SerializeToString,
                response_deserializer=savourrpc_dot_market__pb2.StableCoinPriceResponse.FromString,
                )


class PriceServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def getSymbols(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def getSymbolPrices(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def getStableCoins(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def getStableCoinPrice(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_PriceServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'getSymbols': grpc.unary_unary_rpc_method_handler(
                    servicer.getSymbols,
                    request_deserializer=savourrpc_dot_market__pb2.SymbolRequest.FromString,
                    response_serializer=savourrpc_dot_market__pb2.SymbolResponse.SerializeToString,
            ),
            'getSymbolPrices': grpc.unary_unary_rpc_method_handler(
                    servicer.getSymbolPrices,
                    request_deserializer=savourrpc_dot_market__pb2.SymbolPriceRequest.FromString,
                    response_serializer=savourrpc_dot_market__pb2.SymbolPriceResponse.SerializeToString,
            ),
            'getStableCoins': grpc.unary_unary_rpc_method_handler(
                    servicer.getStableCoins,
                    request_deserializer=savourrpc_dot_market__pb2.StableCoinRequest.FromString,
                    response_serializer=savourrpc_dot_market__pb2.StableCoinResponse.SerializeToString,
            ),
            'getStableCoinPrice': grpc.unary_unary_rpc_method_handler(
                    servicer.getStableCoinPrice,
                    request_deserializer=savourrpc_dot_market__pb2.StableCoinPriceRequest.FromString,
                    response_serializer=savourrpc_dot_market__pb2.StableCoinPriceResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'savourrpc.market.PriceService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class PriceService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def getSymbols(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/savourrpc.market.PriceService/getSymbols',
            savourrpc_dot_market__pb2.SymbolRequest.SerializeToString,
            savourrpc_dot_market__pb2.SymbolResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def getSymbolPrices(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/savourrpc.market.PriceService/getSymbolPrices',
            savourrpc_dot_market__pb2.SymbolPriceRequest.SerializeToString,
            savourrpc_dot_market__pb2.SymbolPriceResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def getStableCoins(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/savourrpc.market.PriceService/getStableCoins',
            savourrpc_dot_market__pb2.StableCoinRequest.SerializeToString,
            savourrpc_dot_market__pb2.StableCoinResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def getStableCoinPrice(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/savourrpc.market.PriceService/getStableCoinPrice',
            savourrpc_dot_market__pb2.StableCoinPriceRequest.SerializeToString,
            savourrpc_dot_market__pb2.StableCoinPriceResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
