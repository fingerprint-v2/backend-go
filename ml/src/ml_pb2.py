# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ml.proto
# Protobuf Python Version: 4.25.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x08ml.proto\x12\x02ml\x1a\x1bgoogle/protobuf/empty.proto\"H\n\x05Point\x12\n\n\x02ID\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12%\n\x0c\x66ingerprints\x18\x03 \x03(\x0b\x32\x0f.ml.Fingerprint\"&\n\x0b\x46ingerprint\x12\x17\n\x05wifis\x18\x01 \x03(\x0b\x32\x08.ml.Wifi\"E\n\x04Wifi\x12\x0c\n\x04SSID\x18\x01 \x01(\t\x12\r\n\x05\x42SSID\x18\x02 \x01(\t\x12\x11\n\tfrequency\x18\x04 \x01(\x05\x12\r\n\x05level\x18\x05 \x01(\x05\"3\n\x08TrainReq\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x19\n\x06points\x18\x02 \x03(\x0b\x32\t.ml.Point\"\x1d\n\x08TrainRes\x12\x11\n\tcompleted\x18\x01 \x01(\x08\"3\n\nPredictReq\x12%\n\x0c\x66ingerprints\x18\x01 \x03(\x0b\x32\x0f.ml.Fingerprint\"\x1e\n\nPredictRes\x12\x10\n\x08point_id\x18\x01 \x01(\t\"\x1c\n\x0cLoadModelReq\x12\x0c\n\x04path\x18\x01 \x01(\t\"#\n\rCheckModelRes\x12\x12\n\nmodel_name\x18\x01 \x01(\t2\xd4\x01\n\nFingperint\x12%\n\x05Train\x12\x0c.ml.TrainReq\x1a\x0c.ml.TrainRes\"\x00\x12+\n\x07Predict\x12\x0e.ml.PredictReq\x1a\x0e.ml.PredictRes\"\x00\x12\x37\n\tLoadModel\x12\x10.ml.LoadModelReq\x1a\x16.google.protobuf.Empty\"\x00\x12\x39\n\nCheckModel\x12\x16.google.protobuf.Empty\x1a\x11.ml.CheckModelRes\"\x00\x42\x1bZ\x19github.com/fingerprint/mlb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'ml_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\031github.com/fingerprint/ml'
  _globals['_POINT']._serialized_start=45
  _globals['_POINT']._serialized_end=117
  _globals['_FINGERPRINT']._serialized_start=119
  _globals['_FINGERPRINT']._serialized_end=157
  _globals['_WIFI']._serialized_start=159
  _globals['_WIFI']._serialized_end=228
  _globals['_TRAINREQ']._serialized_start=230
  _globals['_TRAINREQ']._serialized_end=281
  _globals['_TRAINRES']._serialized_start=283
  _globals['_TRAINRES']._serialized_end=312
  _globals['_PREDICTREQ']._serialized_start=314
  _globals['_PREDICTREQ']._serialized_end=365
  _globals['_PREDICTRES']._serialized_start=367
  _globals['_PREDICTRES']._serialized_end=397
  _globals['_LOADMODELREQ']._serialized_start=399
  _globals['_LOADMODELREQ']._serialized_end=427
  _globals['_CHECKMODELRES']._serialized_start=429
  _globals['_CHECKMODELRES']._serialized_end=464
  _globals['_FINGPERINT']._serialized_start=467
  _globals['_FINGPERINT']._serialized_end=679
# @@protoc_insertion_point(module_scope)
