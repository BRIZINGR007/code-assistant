# MicroService Chat Assistant

A multi-service chat application with Angular frontend and microservice backend architecture.

## Demo Video
![DemoVideo](./assets/demo.gif)

## Architecture Overview

- **Frontend**: Angular  
- **Backend Services**:
  - Service 1: Gin (Golang)  
  - Service 2: Python (FastAPI)

## Features

- RAG (Retrieval-Augmented Generation) implementation
- Dual LLM mode (AWS Bedrock or local execution)
- Vector database integration
- Customizable reference limit for context

## Configuration

Create a `.env` file with the following variables:

```env
# MongoDB Configuration (Required)
MONGO_URI=mongodb+srv://username:password@cluster.mongodb.net/

# LLM Configuration
USE_BEDROCK=true                     # Set to false to use local LLM
REFERENCES_LIMIT=1                   # Number of references sent to LLM

# AWS Configuration (Required when USE_BEDROCK=true)
AWS_ACCESS_KEY_ID=AKIXXXXXXXXXXXXXXXXXXXXXX
AWS_SECRET_ACCESS_KEY=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
AWS_REGION_NAME=us-east-1
```


## Configuration Notes

**USE_BEDROCK**:
- `true`: Uses AWS Bedrock service (requires AWS credentials)
- `false`: Runs LLM locally (AWS credentials optional)

**REFERENCES_LIMIT**: Controls how many reference documents are sent to the LLM for context

## Models Used

| Model Type   | Model Name |
|--------------|------------|
| Embedding    | `intfloat/multilingual-e5-small` |
| ReRanker     | `cross-encoder/mmarco-mMiniLMv2-L12-H384-v1` |
| Bedrock LLM  | `us.meta.llama3-2-1b-instruct-v1:0` |
| Local LLM    | `bartowski/Llama-3.2-1B-Instruct-GGUF/Llama-3.2-1B-Instruct-Q4_K_L.gguf` |

## Getting Started

1. Clone the repository
2. Configure your `.env` file
3. Run the application:

```bash
chmod +x run-all.sh
./run-all.sh
```