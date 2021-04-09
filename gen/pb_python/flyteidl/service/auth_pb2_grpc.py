# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from flyteidl.service import auth_pb2 as flyteidl_dot_service_dot_auth__pb2


class AuthServiceStub(object):
  """The following defines an RPC service that is also served over HTTP via grpc-gateway.
  Standard response codes for both are defined here: https://github.com/grpc-ecosystem/grpc-gateway/blob/master/runtime/errors.go
  """

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.OAuth2Metadata = channel.unary_unary(
        '/flyteidl.service.AuthService/OAuth2Metadata',
        request_serializer=flyteidl_dot_service_dot_auth__pb2.OAuth2MetadataRequest.SerializeToString,
        response_deserializer=flyteidl_dot_service_dot_auth__pb2.OAuth2MetadataResponse.FromString,
        )
    self.FlyteClient = channel.unary_unary(
        '/flyteidl.service.AuthService/FlyteClient',
        request_serializer=flyteidl_dot_service_dot_auth__pb2.FlyteClientRequest.SerializeToString,
        response_deserializer=flyteidl_dot_service_dot_auth__pb2.FlyteClientResponse.FromString,
        )
    self.UserInfo = channel.unary_unary(
        '/flyteidl.service.AuthService/UserInfo',
        request_serializer=flyteidl_dot_service_dot_auth__pb2.UserInfoRequest.SerializeToString,
        response_deserializer=flyteidl_dot_service_dot_auth__pb2.UserInfoResponse.FromString,
        )


class AuthServiceServicer(object):
  """The following defines an RPC service that is also served over HTTP via grpc-gateway.
  Standard response codes for both are defined here: https://github.com/grpc-ecosystem/grpc-gateway/blob/master/runtime/errors.go
  """

  def OAuth2Metadata(self, request, context):
    """Anonymously accessible
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def FlyteClient(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def UserInfo(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_AuthServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'OAuth2Metadata': grpc.unary_unary_rpc_method_handler(
          servicer.OAuth2Metadata,
          request_deserializer=flyteidl_dot_service_dot_auth__pb2.OAuth2MetadataRequest.FromString,
          response_serializer=flyteidl_dot_service_dot_auth__pb2.OAuth2MetadataResponse.SerializeToString,
      ),
      'FlyteClient': grpc.unary_unary_rpc_method_handler(
          servicer.FlyteClient,
          request_deserializer=flyteidl_dot_service_dot_auth__pb2.FlyteClientRequest.FromString,
          response_serializer=flyteidl_dot_service_dot_auth__pb2.FlyteClientResponse.SerializeToString,
      ),
      'UserInfo': grpc.unary_unary_rpc_method_handler(
          servicer.UserInfo,
          request_deserializer=flyteidl_dot_service_dot_auth__pb2.UserInfoRequest.FromString,
          response_serializer=flyteidl_dot_service_dot_auth__pb2.UserInfoResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'flyteidl.service.AuthService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
