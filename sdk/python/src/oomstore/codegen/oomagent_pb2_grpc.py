# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from . import oomagent_pb2 as oomagent__pb2


class OomAgentStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.OnlineGet = channel.unary_unary(
                '/oomagent.OomAgent/OnlineGet',
                request_serializer=oomagent__pb2.OnlineGetRequest.SerializeToString,
                response_deserializer=oomagent__pb2.OnlineGetResponse.FromString,
                )
        self.OnlineMultiGet = channel.unary_unary(
                '/oomagent.OomAgent/OnlineMultiGet',
                request_serializer=oomagent__pb2.OnlineMultiGetRequest.SerializeToString,
                response_deserializer=oomagent__pb2.OnlineMultiGetResponse.FromString,
                )
        self.Sync = channel.unary_unary(
                '/oomagent.OomAgent/Sync',
                request_serializer=oomagent__pb2.SyncRequest.SerializeToString,
                response_deserializer=oomagent__pb2.SyncResponse.FromString,
                )
        self.ChannelImport = channel.stream_unary(
                '/oomagent.OomAgent/ChannelImport',
                request_serializer=oomagent__pb2.ChannelImportRequest.SerializeToString,
                response_deserializer=oomagent__pb2.ImportResponse.FromString,
                )
        self.Import = channel.unary_unary(
                '/oomagent.OomAgent/Import',
                request_serializer=oomagent__pb2.ImportRequest.SerializeToString,
                response_deserializer=oomagent__pb2.ImportResponse.FromString,
                )
        self.ChannelJoin = channel.stream_stream(
                '/oomagent.OomAgent/ChannelJoin',
                request_serializer=oomagent__pb2.ChannelJoinRequest.SerializeToString,
                response_deserializer=oomagent__pb2.ChannelJoinResponse.FromString,
                )
        self.Join = channel.unary_unary(
                '/oomagent.OomAgent/Join',
                request_serializer=oomagent__pb2.JoinRequest.SerializeToString,
                response_deserializer=oomagent__pb2.JoinResponse.FromString,
                )
        self.ChannelExport = channel.unary_stream(
                '/oomagent.OomAgent/ChannelExport',
                request_serializer=oomagent__pb2.ChannelExportRequest.SerializeToString,
                response_deserializer=oomagent__pb2.ChannelExportResponse.FromString,
                )
        self.Export = channel.unary_unary(
                '/oomagent.OomAgent/Export',
                request_serializer=oomagent__pb2.ExportRequest.SerializeToString,
                response_deserializer=oomagent__pb2.ExportResponse.FromString,
                )
        self.HealthCheck = channel.unary_unary(
                '/oomagent.OomAgent/HealthCheck',
                request_serializer=oomagent__pb2.HealthCheckRequest.SerializeToString,
                response_deserializer=oomagent__pb2.HealthCheckResponse.FromString,
                )


class OomAgentServicer(object):
    """Missing associated documentation comment in .proto file."""

    def OnlineGet(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def OnlineMultiGet(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Sync(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ChannelImport(self, request_iterator, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Import(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ChannelJoin(self, request_iterator, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Join(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ChannelExport(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Export(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def HealthCheck(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_OomAgentServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'OnlineGet': grpc.unary_unary_rpc_method_handler(
                    servicer.OnlineGet,
                    request_deserializer=oomagent__pb2.OnlineGetRequest.FromString,
                    response_serializer=oomagent__pb2.OnlineGetResponse.SerializeToString,
            ),
            'OnlineMultiGet': grpc.unary_unary_rpc_method_handler(
                    servicer.OnlineMultiGet,
                    request_deserializer=oomagent__pb2.OnlineMultiGetRequest.FromString,
                    response_serializer=oomagent__pb2.OnlineMultiGetResponse.SerializeToString,
            ),
            'Sync': grpc.unary_unary_rpc_method_handler(
                    servicer.Sync,
                    request_deserializer=oomagent__pb2.SyncRequest.FromString,
                    response_serializer=oomagent__pb2.SyncResponse.SerializeToString,
            ),
            'ChannelImport': grpc.stream_unary_rpc_method_handler(
                    servicer.ChannelImport,
                    request_deserializer=oomagent__pb2.ChannelImportRequest.FromString,
                    response_serializer=oomagent__pb2.ImportResponse.SerializeToString,
            ),
            'Import': grpc.unary_unary_rpc_method_handler(
                    servicer.Import,
                    request_deserializer=oomagent__pb2.ImportRequest.FromString,
                    response_serializer=oomagent__pb2.ImportResponse.SerializeToString,
            ),
            'ChannelJoin': grpc.stream_stream_rpc_method_handler(
                    servicer.ChannelJoin,
                    request_deserializer=oomagent__pb2.ChannelJoinRequest.FromString,
                    response_serializer=oomagent__pb2.ChannelJoinResponse.SerializeToString,
            ),
            'Join': grpc.unary_unary_rpc_method_handler(
                    servicer.Join,
                    request_deserializer=oomagent__pb2.JoinRequest.FromString,
                    response_serializer=oomagent__pb2.JoinResponse.SerializeToString,
            ),
            'ChannelExport': grpc.unary_stream_rpc_method_handler(
                    servicer.ChannelExport,
                    request_deserializer=oomagent__pb2.ChannelExportRequest.FromString,
                    response_serializer=oomagent__pb2.ChannelExportResponse.SerializeToString,
            ),
            'Export': grpc.unary_unary_rpc_method_handler(
                    servicer.Export,
                    request_deserializer=oomagent__pb2.ExportRequest.FromString,
                    response_serializer=oomagent__pb2.ExportResponse.SerializeToString,
            ),
            'HealthCheck': grpc.unary_unary_rpc_method_handler(
                    servicer.HealthCheck,
                    request_deserializer=oomagent__pb2.HealthCheckRequest.FromString,
                    response_serializer=oomagent__pb2.HealthCheckResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'oomagent.OomAgent', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class OomAgent(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def OnlineGet(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/oomagent.OomAgent/OnlineGet',
            oomagent__pb2.OnlineGetRequest.SerializeToString,
            oomagent__pb2.OnlineGetResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def OnlineMultiGet(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/oomagent.OomAgent/OnlineMultiGet',
            oomagent__pb2.OnlineMultiGetRequest.SerializeToString,
            oomagent__pb2.OnlineMultiGetResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Sync(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/oomagent.OomAgent/Sync',
            oomagent__pb2.SyncRequest.SerializeToString,
            oomagent__pb2.SyncResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ChannelImport(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_unary(request_iterator, target, '/oomagent.OomAgent/ChannelImport',
            oomagent__pb2.ChannelImportRequest.SerializeToString,
            oomagent__pb2.ImportResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Import(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/oomagent.OomAgent/Import',
            oomagent__pb2.ImportRequest.SerializeToString,
            oomagent__pb2.ImportResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ChannelJoin(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_stream(request_iterator, target, '/oomagent.OomAgent/ChannelJoin',
            oomagent__pb2.ChannelJoinRequest.SerializeToString,
            oomagent__pb2.ChannelJoinResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Join(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/oomagent.OomAgent/Join',
            oomagent__pb2.JoinRequest.SerializeToString,
            oomagent__pb2.JoinResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ChannelExport(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/oomagent.OomAgent/ChannelExport',
            oomagent__pb2.ChannelExportRequest.SerializeToString,
            oomagent__pb2.ChannelExportResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Export(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/oomagent.OomAgent/Export',
            oomagent__pb2.ExportRequest.SerializeToString,
            oomagent__pb2.ExportResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def HealthCheck(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/oomagent.OomAgent/HealthCheck',
            oomagent__pb2.HealthCheckRequest.SerializeToString,
            oomagent__pb2.HealthCheckResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
