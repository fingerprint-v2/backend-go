from dependency_injector import containers, providers

from grpc_servicers.fingerprint import FingerprintServicer
from services.trainer import Trainer
from services.parser import Parser
from services.predictor import Predictor


class Container(containers.DeclarativeContainer):
    trainer = providers.Factory(Trainer)
    parser = providers.Factory(Parser)
    predictor = providers.Factory(Predictor)
    fingerprintServicer = providers.Factory(
        FingerprintServicer, trainer=trainer, parser=parser, predictor=predictor
    )
