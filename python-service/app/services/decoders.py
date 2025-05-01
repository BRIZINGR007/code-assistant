from typing import List

from app.ai.AIFactory import AIContextChat, AIGeneralChat
from app.deocders.DecoderHandlers import DecoderHandler
from app.interfaces.decoder import ContextChat_TH, Contexts_PM, GeneralChat_TH
from zoldics_service_utils.clients.bedrock_client.foundation_models import (
    FoundationModels,
)
from decouple import config

from app.settings.prompts import ContextChatPromptEnums, GeneralChatPromptEnums


class DecoderService:
    @staticmethod
    def create_payload_with_context(contexts: List[Contexts_PM]) -> str:
        code_chunks = [
            {"file_path": ctx.FilePath, "code_chunk": ctx.Code} for ctx in contexts
        ]
        return str(code_chunks)

    @staticmethod
    def handle_context_chat_with_bedrock(query: str, context: str) -> str:
        modelId = FoundationModels.LLAMA_3_2_1B_INSTRUCT
        payload = ContextChat_TH(query=query, context=context)
        llm_response = AIContextChat().execute_llm_operation(
            payload=payload, modelId=modelId
        )
        return llm_response

    @staticmethod
    def handle_general_chat_with_bedrock(query: str) -> str:
        modelId = FoundationModels.LLAMA_3_2_1B_INSTRUCT
        payload = GeneralChat_TH(query=query)
        llm_response = AIGeneralChat().execute_llm_operation(
            payload=payload, modelId=modelId
        )
        return llm_response

    @staticmethod
    def handle_context_chat_with_local_llm(query: str, context: str):
        system_prompt = ContextChatPromptEnums.SYSTEM_PROMPT.value.format(
            code_snippet=str(context)
        )
        user_prompt = ContextChatPromptEnums.USER_PROMPT.value.format(user_query=query)
        llm_response = DecoderHandler().generate_llm_reponse(
            system_prompt=system_prompt, user_prompt=user_prompt
        )
        return llm_response

    @staticmethod
    def handle_general_chat_with_local_llm(query: str) -> str:
        system_prompt = GeneralChatPromptEnums.SYSTEM_PROMPT.value
        user_prompt = GeneralChatPromptEnums.USER_PROMPT.value.format(user_query=query)

        llm_response = DecoderHandler().generate_llm_reponse(
            system_prompt=system_prompt, user_prompt=user_prompt
        )
        return llm_response
