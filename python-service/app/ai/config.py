from enum import StrEnum
from typing import Dict, Literal, overload
from zoldics_service_utils.clients.bedrock_client import BaseAIEventsConfig
from zoldics_service_utils.clients.bedrock_client.foundation_models import (
    FoundationModels,
)
from zoldics_service_utils.interfaces.interfaces_th import (
    LLM_HyperParameters_TH,
    LLM_PromptTemplates_TH,
)
from app.settings.prompts import ContextChatPromptEnums, GeneralChatPromptEnums


class AIEvents(StrEnum):
    CONTEXT_CHAT = "CONTEXT_CHAT"
    GENERAL_CHAT = "GENERAL_CHAT"


class AIEventsConfig:
    class CONTEXT_CHAT(BaseAIEventsConfig):
        @staticmethod
        def BASE_OUTPUT_TOKENS() -> int:
            return 2048

        @staticmethod
        def MODEL_PROMPTS(
            user_query: str, code_snippet: str
        ) -> Dict[FoundationModels, LLM_PromptTemplates_TH]:
            return {
                FoundationModels.LLAMA_3_1_8B_INSTRUCT: LLM_PromptTemplates_TH(
                    system_prompt=ContextChatPromptEnums.SYSTEM_PROMPT.value.format(
                        code_snippet=code_snippet
                    ),
                    user_prompt=ContextChatPromptEnums.USER_PROMPT.format(
                        user_query=user_query
                    ),
                    base_token_count=200,
                ),
                FoundationModels.LLAMA_3_1_70B_INSTRUCT: LLM_PromptTemplates_TH(
                    system_prompt=ContextChatPromptEnums.SYSTEM_PROMPT.value.format(
                        code_snippet=code_snippet
                    ),
                    user_prompt=ContextChatPromptEnums.USER_PROMPT.format(
                        user_query=user_query
                    ),
                    base_token_count=200,
                ),
                FoundationModels.LLAMA_3_2_1B_INSTRUCT: LLM_PromptTemplates_TH(
                    system_prompt=ContextChatPromptEnums.SYSTEM_PROMPT.value.format(
                        code_snippet=code_snippet
                    ),
                    user_prompt=ContextChatPromptEnums.USER_PROMPT.format(
                        user_query=user_query
                    ),
                    base_token_count=200,
                ),
            }

        @staticmethod
        def MODEL_HYPERPARAMETERS() -> Dict[FoundationModels, LLM_HyperParameters_TH]:
            return {
                FoundationModels.LLAMA_3_1_8B_INSTRUCT: LLM_HyperParameters_TH(
                    max_gen_len=AIEventsConfig.CONTEXT_CHAT.BASE_OUTPUT_TOKENS(),
                    temperature=0.6,
                    top_p=0.6,
                ),
                FoundationModels.LLAMA_3_1_70B_INSTRUCT: LLM_HyperParameters_TH(
                    max_gen_len=AIEventsConfig.CONTEXT_CHAT.BASE_OUTPUT_TOKENS(),
                    temperature=0.6,
                    top_p=0.6,
                ),
                FoundationModels.LLAMA_3_2_1B_INSTRUCT: LLM_HyperParameters_TH(
                    max_gen_len=AIEventsConfig.CONTEXT_CHAT.BASE_OUTPUT_TOKENS(),
                    temperature=0.6,
                    top_p=0.6,
                ),
            }

    class GENERAL_CHAT(BaseAIEventsConfig):
        @staticmethod
        def BASE_OUTPUT_TOKENS() -> int:
            return 2048

        @staticmethod
        def MODEL_PROMPTS(
            user_query: str,
        ) -> Dict[FoundationModels, LLM_PromptTemplates_TH]:
            return {
                FoundationModels.LLAMA_3_1_8B_INSTRUCT: LLM_PromptTemplates_TH(
                    system_prompt=GeneralChatPromptEnums.SYSTEM_PROMPT.value,
                    user_prompt=GeneralChatPromptEnums.USER_PROMPT.format(
                        user_query=user_query
                    ),
                    base_token_count=200,
                ),
                FoundationModels.LLAMA_3_1_70B_INSTRUCT: LLM_PromptTemplates_TH(
                    system_prompt=GeneralChatPromptEnums.SYSTEM_PROMPT.value,
                    user_prompt=GeneralChatPromptEnums.USER_PROMPT.format(
                        user_query=user_query
                    ),
                    base_token_count=200,
                ),
                FoundationModels.LLAMA_3_2_1B_INSTRUCT: LLM_PromptTemplates_TH(
                    system_prompt=GeneralChatPromptEnums.SYSTEM_PROMPT.value,
                    user_prompt=GeneralChatPromptEnums.USER_PROMPT.format(
                        user_query=user_query
                    ),
                    base_token_count=200,
                ),
            }

        @staticmethod
        def MODEL_HYPERPARAMETERS() -> Dict[FoundationModels, LLM_HyperParameters_TH]:
            return {
                FoundationModels.LLAMA_3_1_8B_INSTRUCT: LLM_HyperParameters_TH(
                    max_gen_len=AIEventsConfig.GENERAL_CHAT.BASE_OUTPUT_TOKENS(),
                    temperature=0.6,
                    top_p=0.6,
                ),
                FoundationModels.LLAMA_3_1_70B_INSTRUCT: LLM_HyperParameters_TH(
                    max_gen_len=AIEventsConfig.GENERAL_CHAT.BASE_OUTPUT_TOKENS(),
                    temperature=0.6,
                    top_p=0.6,
                ),
                FoundationModels.LLAMA_3_2_1B_INSTRUCT: LLM_HyperParameters_TH(
                    max_gen_len=AIEventsConfig.GENERAL_CHAT.BASE_OUTPUT_TOKENS(),
                    temperature=0.6,
                    top_p=0.6,
                ),
            }


class AIEventsConfigFactory:
    @staticmethod
    @overload
    def get_event_config(
        event: Literal[AIEvents.GENERAL_CHAT],
    ) -> AIEventsConfig.GENERAL_CHAT:
        return AIEventsConfig.GENERAL_CHAT()

    @staticmethod
    @overload
    def get_event_config(
        event: Literal[AIEvents.CONTEXT_CHAT],
    ) -> AIEventsConfig.CONTEXT_CHAT:
        return AIEventsConfig.CONTEXT_CHAT()

    @staticmethod
    def get_event_config(event: AIEvents):
        match event:
            case AIEvents.GENERAL_CHAT:
                return AIEventsConfig.GENERAL_CHAT()
            case AIEvents.CONTEXT_CHAT:
                return AIEventsConfig.CONTEXT_CHAT()
            case _:
                raise ValueError("Unsupported event type")
