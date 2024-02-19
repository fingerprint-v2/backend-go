import logging
from concurrent import futures

import grpc
from dependency_injector.wiring import Provide, inject

import ml_pb2_grpc as pb2_grpc
from container import Container


@inject
def main(
    fingerprintServicer: pb2_grpc.FingperintServicer = Provide[
        Container.fingerprintServicer
    ],
):

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    pb2_grpc.add_FingperintServicer_to_server(fingerprintServicer, server)
    server.add_insecure_port("[::]:50051")
    server.start()
    print("Server started at localhost:50051")
    server.wait_for_termination()


if __name__ == "__main__":
    logging.basicConfig()
    container = Container()
    container.init_resources()
    container.wire(modules=[__name__])
    main()
