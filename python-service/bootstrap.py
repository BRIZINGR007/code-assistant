from typing import Any, cast
from decouple import config
from zoldics_service_utils.utils.env_initlializer import EnvStore

from app.cross_encoders.CrossEncoderHandlers import CrossEncoderHandler
from app.deocders.DecoderHandlers import DecoderHandler
from app.encoders.EncoderHandlers import EmbeddingGenerator


class BootStrap:
    def __call__(self, *args: Any, **kwds: Any) -> Any:
        use_bedrock = config("USE_BEDROCK")
        if use_bedrock == "true":
            EnvStore().aws_access_key_id = cast(str, config("AWS_ACCESS_KEY_ID"))
            EnvStore().aws_secret_access_key = cast(
                str, config("AWS_SECRET_ACCESS_KEY")
            )
            EnvStore().aws_region_name = cast(str, config("AWS_REGION_NAME"))
        elif use_bedrock == "false":
            DecoderHandler()
            print("Decoder Handler  Initialized.")
        else:
            raise ValueError("Invalid Value for Use Bedrock env variable .")
        EmbeddingGenerator()
        print("Embedding Generator Initilized.")
        CrossEncoderHandler()
        print("Cross  Encoder Initialized.")
