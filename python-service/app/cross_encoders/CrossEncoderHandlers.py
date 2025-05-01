from typing import List, Tuple
from zoldics_service_utils.ioc import SingletonMeta
from transformers.models.auto.tokenization_auto import AutoTokenizer
from transformers.models.auto.modeling_auto import AutoModelForSequenceClassification
import torch
from torch import Tensor


class CrossEncoderHandler(metaclass=SingletonMeta):
    def __init__(self) -> None:
        self.__modelId: str = "cross-encoder/mmarco-mMiniLMv2-L12-H384-v1"
        self.__tokenizer = AutoTokenizer.from_pretrained(self.__modelId)
        self.__model = AutoModelForSequenceClassification.from_pretrained(
            self.__modelId
        )

    def rerank(self, pairs: List[List[str]]) -> List[float]:
        """
        Takes in a list of sentence pairs and returns a list of scores (float) indicating
        the relevance between each pair.

        :param pairs: A list of sentence pairs e.g., [["sentence1", "sentence2"], ...]
        :return: A list of float scores for each sentence pair.
        """
        with torch.no_grad():
            inputs = self.__tokenizer(
                pairs,
                padding=True,
                truncation=True,
                return_tensors="pt",
                max_length=512,
            )
            scores: Tensor = (
                self.__model(**inputs, return_dict=True).logits.view(-1).float()
            )

        return scores.tolist()
