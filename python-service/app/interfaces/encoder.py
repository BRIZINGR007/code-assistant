from pydantic import BaseModel
from typing import List


class CodeContext(BaseModel):
    vector_id: str
    codebase_id: str
    codebase_name: str
    hashId: str
    filePath: str
    code: str
    embedding: List[float]


class QueryPayload(BaseModel):
    query: str
    context: List[CodeContext]


class CodeContextWithSimilarity(CodeContext):
    similarity_to_query: float
    similarity_to_prev_query: float
