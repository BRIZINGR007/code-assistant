from fastapi import APIRouter, Depends, Query
from fastapi.responses import JSONResponse
from app.controllers.encoder import EncoderController
from app.encoders.EncoderHandlers import EmbeddingGenerator
from app.interfaces.encoder import QueryPayload
from app.settings.enums import ServicePaths
from zoldics_service_utils.middlewares import ContextSetter


router = APIRouter(
    prefix=ServicePaths.CONTEXT_PATH.value + "/encoder",
    tags=["Enocder Path"],
    responses={"404": {"description": "Not found"}},
    dependencies=[Depends(ContextSetter())],
)


@router.get("/get-embedding")
def get_embeddings(query: str = Query(...)):
    embedding = EmbeddingGenerator().generate_embedding(query)
    return JSONResponse(status_code=200, content=embedding)


@router.post("/get-cosine-similarity-scores")
def get_similarity(payload: QueryPayload):
    return EncoderController().populate_cosine_similarity(payload=payload)
