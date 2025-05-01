from typing import final, List

from transformers.models.auto.tokenization_auto import AutoTokenizer
from transformers.models.auto.modeling_auto import AutoModel
from transformers.tokenization_utils_base import BatchEncoding
import torch
from zoldics_service_utils.ioc import SingletonMeta


@final
class EmbeddingGenerator(metaclass=SingletonMeta):
    def __init__(self) -> None:
        self.tokenizer = AutoTokenizer.from_pretrained("intfloat/multilingual-e5-small")
        self.model = AutoModel.from_pretrained("intfloat/multilingual-e5-small")
        self.model.eval()

    def generate_embedding(self, text: str) -> List[float]:
        encoded_input: BatchEncoding = self.tokenizer(
            text, padding=True, truncation=True, return_tensors="pt"
        )
        with torch.no_grad():
            model_output = self.model(**encoded_input)
            sentence_embeddings = model_output[0][:, 0]
        sentence_embeddings: torch.Tensor = torch.nn.functional.normalize(
            sentence_embeddings, p=2, dim=1
        )
        return sentence_embeddings.tolist()[0]
