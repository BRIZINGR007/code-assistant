from app.deocders.DecoderHandlers import DecoderHandler
from app.interfaces.decoder import LLMResponseContext_PM
from app.services.decoders import DecoderService
from app.settings.prompts import GeneralChatPromptEnums, ContextChatPromptEnums
from decouple import config


class DecoderController:
    def __init__(self) -> None:
        self.__decoder_service = DecoderService()

    def handle_context_chat(self, payload: LLMResponseContext_PM) -> str:
        context = self.__decoder_service.create_payload_with_context(payload.contexts)

        use_bedrock = config("USE_BEDROCK")
        if use_bedrock == "true":
            print("USING  BEDROCK  TO GENERATE  LLM RESPONSE  .")
            return self.__decoder_service.handle_context_chat_with_bedrock(
                payload.query, context
            )
        print("USING LOCAL LLM TO GENERATE LLM RESPONSE .")
        return self.__decoder_service.handle_context_chat_with_local_llm(
            payload.query, context
        )

    def handle_general_chat(self, query: str) -> str:
        use_bedrock = config("USE_BEDROCK")
        if use_bedrock == "true":
            print("USING  BEDROCK  TO GENERATE  LLM RESPONSE  .")
            return self.__decoder_service.handle_general_chat_with_bedrock(query=query)
        print("USING LOCAL LLM TO GENERATE LLM RESPONSE .")
        return self.__decoder_service.handle_general_chat_with_local_llm(query=query)
