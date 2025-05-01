from fastapi.responses import JSONResponse
from app.cross_encoders.CrossEncoderHandlers import CrossEncoderHandler
from app.interfaces.cross_encoder import Reranking_PM


class CrossEncoderController:
    def __init__(self) -> None:
        pass

    def rerank(self, payload: Reranking_PM):
        context_payload = [
            [payload.query, each_context] for each_context in payload.contexts
        ]
        scores = CrossEncoderHandler().rerank(pairs=context_payload)
        max_index = scores.index(max(scores))
        return JSONResponse(status_code=200, content=max_index)
