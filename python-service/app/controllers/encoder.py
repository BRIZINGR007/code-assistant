from fastapi.responses import JSONResponse
from app.encoders.EncoderHandlers import EmbeddingGenerator
from app.interfaces.encoder import QueryPayload
from app.services.encoders import EncoderService


class EncoderController:
    def __init__(self) -> None:
        self.__service = EncoderService()

    def populate_cosine_similarity(self, payload: QueryPayload) -> JSONResponse:
        codecontext_with_simialrity = self.__service.populate_cosine_similarity(
            payload=payload
        )
        json_payload = [
            each_payload.model_dump(mode="json")
            for each_payload in codecontext_with_simialrity
        ]
        return JSONResponse(status_code=200, content=json_payload)
