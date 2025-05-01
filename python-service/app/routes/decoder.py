from fastapi import APIRouter, Depends, Query
from fastapi.responses import JSONResponse
from app.controllers.decoder import DecoderController
from app.interfaces.decoder import LLMResponseContext_PM
from app.settings.enums import ServicePaths
from zoldics_service_utils.middlewares import ContextSetter


router = APIRouter(
    prefix=ServicePaths.CONTEXT_PATH.value + "/decoder",
    tags=["Decoder Path"],
    responses={"404": {"description": "Not found"}},
    dependencies=[Depends(ContextSetter())],
)


@router.post("/get-llmresponse")
def get_llmresponse(payload: LLMResponseContext_PM):
    llm_response = DecoderController().handle_context_chat(payload=payload)
    return JSONResponse(status_code=200, content=llm_response)


@router.get("/general-chat")
def general_chat(user_query: str = Query(...)):
    llm_response = DecoderController().handle_general_chat(user_query)
    return JSONResponse(status_code=200, content=llm_response)
