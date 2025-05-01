from fastapi import APIRouter, Depends
from app.controllers.cross_encoder import CrossEncoderController
from app.interfaces.cross_encoder import Reranking_PM
from app.settings.enums import ServicePaths
from zoldics_service_utils.middlewares import ContextSetter

router = APIRouter(
    prefix=ServicePaths.CONTEXT_PATH.value + "/cross-encoder",
    tags=["Cross Encoder  Path"],
    responses={"404": {"description": "Not Found"}},
    dependencies=[Depends(ContextSetter())],
)


@router.post("/reranking")
def initiate_reranking(payload: Reranking_PM):
    return CrossEncoderController().rerank(payload)
