from concurrent import futures
import logging
import math
import time

import grpc
import todo_pb2
import todo_pb2_grpc


class TodoServicer(todo_pb2_grpc.TodoServiceServicer):
    def GetTodos(self, request, context):
        # todos = [todo_pb2.Todos(
        #     id=1,
        #     title='test',
        #     completed=False
        # )]
        todo = todo_pb2.Todo(id=1, title="test", completed=False)

        todos = todo_pb2.Todos(todos=[todo, todo])
        return todos

        # return super().GetTodos(request, context)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    todo_pb2_grpc.add_TodoServiceServicer_to_server(TodoServicer(), server)
    server.add_insecure_port("[::]:50051")
    server.start()
    print("Server started at localhost:50051")
    server.wait_for_termination()


if __name__ == "__main__":
    logging.basicConfig()
    serve()
