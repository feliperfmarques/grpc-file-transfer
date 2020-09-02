import 'dart:async';
import 'dart:io';

import 'package:grpc/grpc.dart';

import 'proto/grpc_file_transfer.pb.dart' as pb;
import 'proto/grpc_file_transfer.pbgrpc.dart';

Future<void> main(List<String> args) async {
  ClientChannel channel;
  GrpcFileTransferServiceClient stub;

  // Setup

  // final List<int> trustedRoot = File('ca.crt').readAsBytesSync();
  // final ChannelCredentials creds = ChannelCredentials.secure(
  //   certificates: trustedRoot,
  //   authority: 'localhost',
  // );
  channel = ClientChannel('localhost',
      port: 5000,
      // options: ChannelOptions(credentials: creds),
      options: ChannelOptions(credentials: ChannelCredentials.insecure()));
  stub = GrpcFileTransferServiceClient(
    channel,
    options: CallOptions(timeout: const Duration(seconds: 60)),
  );

  // DO
  try {
    List<int> bytes = [];
    final stream =
        await stub.download(pb.Request()..fileName = "video_example.mp4");
    await for (var response in stream) {
      if (response.hasFileChunk()) {
        bytes = [...bytes, ...response.fileChunk];
      }
    }
    new File('video_example.mp4')..writeAsBytesSync(bytes);
  } catch (e, s) {
    print('Caught error: $e. Stacktrace:\n$s');
  }

  // Shutdown
  await channel.shutdown();
}
