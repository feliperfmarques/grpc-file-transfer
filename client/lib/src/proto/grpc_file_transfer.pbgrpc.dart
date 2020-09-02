///
//  Generated code. Do not modify.
//  source: proto/grpc_file_transfer.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'grpc_file_transfer.pb.dart' as $0;
export 'grpc_file_transfer.pb.dart';

class GrpcFileTransferServiceClient extends $grpc.Client {
  static final _$download = $grpc.ClientMethod<$0.Request, $0.Response>(
      '/grpcfiletransfer.GrpcFileTransferService/Download',
      ($0.Request value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Response.fromBuffer(value));

  GrpcFileTransferServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseStream<$0.Response> download($0.Request request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$download, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseStream(call);
  }
}

abstract class GrpcFileTransferServiceBase extends $grpc.Service {
  $core.String get $name => 'grpcfiletransfer.GrpcFileTransferService';

  GrpcFileTransferServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.Request, $0.Response>(
        'Download',
        download_Pre,
        false,
        true,
        ($core.List<$core.int> value) => $0.Request.fromBuffer(value),
        ($0.Response value) => value.writeToBuffer()));
  }

  $async.Stream<$0.Response> download_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Request> request) async* {
    yield* download(call, await request);
  }

  $async.Stream<$0.Response> download(
      $grpc.ServiceCall call, $0.Request request);
}
