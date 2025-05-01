from typing import List
from app.encoders.EncoderHandlers import EmbeddingGenerator
from app.interfaces.encoder import CodeContextWithSimilarity, QueryPayload
import numpy as np


class EncoderService:
    @staticmethod
    def cosine_similarity(vec1: List[float], vec2: List[float]) -> float:
        v1 = np.array(vec1)
        v2 = np.array(vec2)
        if np.linalg.norm(v1) == 0 or np.linalg.norm(v2) == 0:
            return 0.0
        return float(np.dot(v1, v2) / (np.linalg.norm(v1) * np.linalg.norm(v2)))

    @classmethod
    def populate_cosine_similarity(
        cls,
        payload: QueryPayload,
    ) -> List[CodeContextWithSimilarity]:
        query_embedding = EmbeddingGenerator().generate_embedding(payload.query)
        results = []

        for idx, context in enumerate(payload.context):
            similarity_to_query = cls.cosine_similarity(
                query_embedding, context.embedding
            )
            if idx == 0:
                similarity_to_prev = similarity_to_query
            else:
                similarity_to_prev = cls.cosine_similarity(
                    payload.context[idx - 1].embedding, payload.context[idx].embedding
                )

            context_with_similarity = CodeContextWithSimilarity(
                **context.model_dump(),
                similarity_to_query=similarity_to_query,
                similarity_to_prev_query=similarity_to_prev
            )
            results.append(context_with_similarity)

        return results
