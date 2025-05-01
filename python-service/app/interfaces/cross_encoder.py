from typing import List

from pydantic import BaseModel


class Reranking_PM(BaseModel):
    query: str
    contexts: List[str]
