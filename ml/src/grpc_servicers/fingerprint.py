import ml_pb2 as pb2
import ml_pb2_grpc as pb2_grpc
from services.trainer import Trainer
from services.parser import Parser
from services.predictor import Predictor
from google.protobuf.empty_pb2 import Empty


class FingerprintServicer(pb2_grpc.FingperintServicer):
    def __init__(self, trainer: Trainer, parser: Parser, predictor: Predictor) -> None:
        self.trainer = trainer
        self.parser = parser
        self.predictor = predictor
        print("FingerprintServicer created")

    def Train(self, request, context):
        data_dict = self.parser.parse(request)
        self.trainer.train(data_dict)
        return pb2.TrainRes(completed=True)

    def Predict(self, request, context):
        print("Predict called")
        print(self.predictor.getCurrentModel())

        return pb2.PredictRes(point_id="point1")

    def LoadModel(self, request, context):
        self.predictor.loadModel(request.path)
        print(f"Model loaded: {request.path}")
        return Empty()

    def CheckModel(self, request, context):

        return pb2.CheckModelRes(model_name=self.predictor.getCurrentModel())
