from google.protobuf.json_format import MessageToDict


class Parser:
    def __init__(self):
        pass

    def parse(self, request):

        data_dict = MessageToDict(
            request,
            preserving_proto_field_name=True,
            use_integers_for_enums=False,
            including_default_value_fields=True,
        )

        return data_dict
