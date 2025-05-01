from typing import List, TypedDict
from pydantic import BaseModel


class Contexts_PM(BaseModel):
    Code: str
    FilePath: str


class LLMResponseContext_PM(BaseModel):
    query: str
    contexts: List[Contexts_PM]


class ContextChat_TH(TypedDict):
    query: str
    context: str


class GeneralChat_TH(TypedDict):
    query: str
