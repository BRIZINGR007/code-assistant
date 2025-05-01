import os
import requests
from pathlib import Path
import requests
from pathlib import Path
from decouple import config


class HuggingFaceModelDownload:

    def __download_hf_file(self, model_repo: str, filename: str, output_dir: str):
        base_url = f"https://huggingface.co/{model_repo}/resolve/main/{filename}"
        output_path = Path(output_dir) / filename
        output_path.parent.mkdir(parents=True, exist_ok=True)

        if output_path.exists():
            print(f"FILE ALREADY EXISTS: {output_path}", flush=True)
            return

        response = requests.get(base_url, stream=True)
        if response.status_code == 200:
            with open(output_path, "wb") as f:
                for chunk in response.iter_content(chunk_size=8192):
                    f.write(chunk)
            print(f"DOWNLOADED: {filename}", flush=True)
        else:
            print(f"FAILED TO DOWNLOAD {filename}: {response.status_code}", flush=True)

    def __download_model(self, url, output_path):
        response = requests.get(url, stream=True)
        output_path = Path(output_path)
        output_path.parent.mkdir(parents=True, exist_ok=True)

        if response.status_code == 200:
            with open(output_path, "wb") as f:
                for chunk in response.iter_content(chunk_size=8192):
                    f.write(chunk)
            print(f"MODEL DOWNLOADED TO {output_path}", flush=True)
        else:
            print(f"FAILED TO DOWNLOAD MODEL: {response.status_code}", flush=True)

    def download_intfloat_embedding_model(self):
        print("DOWNLOADING EMBEDDING MODELS", flush=True)
        files = [
            "config.json",
            "model.safetensors",
            "sentencepiece.bpe.model",
            "special_tokens_map.json",
            "tokenizer.json",
            "tokenizer_config.json",
        ]
        model_repo = "intfloat/multilingual-e5-small"
        output_dir = "intfloat/multilingual-e5-small"

        for file in files:
            self.__download_hf_file(model_repo, file, output_dir)

    def download_baai_reranking_model(self):
        files = [
            "config.json",
            "model.safetensors",
            "special_tokens_map.json",
            "tokenizer_config.json",
            "tokenizer.json",
        ]
        model_repo = "BAAI/bge-reranker-base"
        output_dir = "BAAI/bge-reranker-base"
        for file in files:
            self.__download_hf_file(model_repo, file, output_dir)

    def download_marco_minilm_reranking_model(self):
        print("DOWNLOADING RERANKING MODELS ...", flush=True)
        files = [
            "config.json",
            "model.safetensors",
            "special_tokens_map.json",
            "tokenizer_config.json",
            "tokenizer.json",
        ]
        model_repo = "cross-encoder/mmarco-mMiniLMv2-L12-H384-v1"
        output_dir = "cross-encoder/mmarco-mMiniLMv2-L12-H384-v1"
        for file in files:
            self.__download_hf_file(model_repo, file, output_dir)

    def download_gguf_llm(self):
        model_url = "https://huggingface.co/bartowski/Llama-3.2-1B-Instruct-GGUF/resolve/main/Llama-3.2-1B-Instruct-Q4_K_L.gguf"
        output_file = (
            "bartowski/Llama-3.2-1B-Instruct-GGUF/Llama-3.2-1B-Instruct-Q4_K_L.gguf"
        )
        print(f"DOWNLOADING LLM {output_file}", flush=True)
        self.__download_model(model_url, output_file)


if __name__ == "__main__":
    _hfmd = HuggingFaceModelDownload()
    if os.environ.get("USE_BEDROCK", "false").lower() == "false":
        _hfmd.download_gguf_llm()
    elif os.environ.get("USE_BEDROCK", "false").lower() == "true":
        print("SETTING  TO BEDROCK INSTEAD ...")
    _hfmd.download_intfloat_embedding_model()
    _hfmd.download_marco_minilm_reranking_model()
