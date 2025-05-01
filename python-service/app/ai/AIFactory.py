from typing import Any
from zoldics_service_utils.clients.bedrock_client import LLMOperation
from zoldics_service_utils.clients.bedrock_client.foundation_models import (
    FoundationModels,
)
from app.ai.config import AIEvents, AIEventsConfig, AIEventsConfigFactory
from app.interfaces.decoder import ContextChat_TH, GeneralChat_TH


class AIContextChat(LLMOperation[ContextChat_TH]):
    def __init__(self) -> None:
        self.__agent_config = AIEventsConfigFactory.get_event_config(
            event=AIEvents.CONTEXT_CHAT
        )

    def execute_llm_operation(
        self, payload: ContextChat_TH, modelId: FoundationModels
    ) -> str:
        prompt_config = self.__agent_config.MODEL_PROMPTS(
            user_query=payload["query"], code_snippet=payload["context"]
        )[modelId]
        model_hyper_parameters = self.__agent_config.MODEL_HYPERPARAMETERS()[modelId]
        llmclientpayload = self.construct_llmclient_payload(
            system_prompt=prompt_config.get("system_prompt"),
            user_prompt=prompt_config.get("user_prompt"),
            model_hyperparameters=model_hyper_parameters,
            modelId=modelId,
        )
        response = self.call_llmclient(payload=llmclientpayload)
        return response


class AIGeneralChat(LLMOperation[GeneralChat_TH]):
    def __init__(self) -> None:
        self.__agent_config = AIEventsConfigFactory.get_event_config(
            event=AIEvents.GENERAL_CHAT
        )

    def execute_llm_operation(
        self, payload: GeneralChat_TH, modelId: FoundationModels
    ) -> Any:
        prompt_config = self.__agent_config.MODEL_PROMPTS(user_query=payload["query"])[
            modelId
        ]
        model_hyper_parameters = self.__agent_config.MODEL_HYPERPARAMETERS()[modelId]
        llmclientpayload = self.construct_llmclient_payload(
            system_prompt=prompt_config.get("system_prompt"),
            user_prompt=prompt_config.get("user_prompt"),
            model_hyperparameters=model_hyper_parameters,
            modelId=modelId,
        )
        response = self.call_llmclient(payload=llmclientpayload)
        return response
