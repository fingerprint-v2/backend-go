from google.protobuf import empty_pb2 as _empty_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Point(_message.Message):
    __slots__ = ("ID", "name", "fingerprints")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    FINGERPRINTS_FIELD_NUMBER: _ClassVar[int]
    ID: str
    name: str
    fingerprints: _containers.RepeatedCompositeFieldContainer[Fingerprint]
    def __init__(self, ID: _Optional[str] = ..., name: _Optional[str] = ..., fingerprints: _Optional[_Iterable[_Union[Fingerprint, _Mapping]]] = ...) -> None: ...

class Fingerprint(_message.Message):
    __slots__ = ("wifis",)
    WIFIS_FIELD_NUMBER: _ClassVar[int]
    wifis: _containers.RepeatedCompositeFieldContainer[Wifi]
    def __init__(self, wifis: _Optional[_Iterable[_Union[Wifi, _Mapping]]] = ...) -> None: ...

class Wifi(_message.Message):
    __slots__ = ("SSID", "BSSID", "frequency", "level")
    SSID_FIELD_NUMBER: _ClassVar[int]
    BSSID_FIELD_NUMBER: _ClassVar[int]
    FREQUENCY_FIELD_NUMBER: _ClassVar[int]
    LEVEL_FIELD_NUMBER: _ClassVar[int]
    SSID: str
    BSSID: str
    frequency: int
    level: int
    def __init__(self, SSID: _Optional[str] = ..., BSSID: _Optional[str] = ..., frequency: _Optional[int] = ..., level: _Optional[int] = ...) -> None: ...

class TrainReq(_message.Message):
    __slots__ = ("name", "points")
    NAME_FIELD_NUMBER: _ClassVar[int]
    POINTS_FIELD_NUMBER: _ClassVar[int]
    name: str
    points: _containers.RepeatedCompositeFieldContainer[Point]
    def __init__(self, name: _Optional[str] = ..., points: _Optional[_Iterable[_Union[Point, _Mapping]]] = ...) -> None: ...

class TrainRes(_message.Message):
    __slots__ = ("completed",)
    COMPLETED_FIELD_NUMBER: _ClassVar[int]
    completed: bool
    def __init__(self, completed: bool = ...) -> None: ...

class PredictReq(_message.Message):
    __slots__ = ("fingerprints",)
    FINGERPRINTS_FIELD_NUMBER: _ClassVar[int]
    fingerprints: _containers.RepeatedCompositeFieldContainer[Fingerprint]
    def __init__(self, fingerprints: _Optional[_Iterable[_Union[Fingerprint, _Mapping]]] = ...) -> None: ...

class PredictRes(_message.Message):
    __slots__ = ("point_id",)
    POINT_ID_FIELD_NUMBER: _ClassVar[int]
    point_id: str
    def __init__(self, point_id: _Optional[str] = ...) -> None: ...

class LoadModelReq(_message.Message):
    __slots__ = ("path",)
    PATH_FIELD_NUMBER: _ClassVar[int]
    path: str
    def __init__(self, path: _Optional[str] = ...) -> None: ...

class CheckModelRes(_message.Message):
    __slots__ = ("model_name",)
    MODEL_NAME_FIELD_NUMBER: _ClassVar[int]
    model_name: str
    def __init__(self, model_name: _Optional[str] = ...) -> None: ...
