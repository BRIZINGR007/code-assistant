FROM python:3.11-bullseye

ARG USE_BEDROCK=true

# Install minimal dependencies
RUN apt-get update && apt-get install -y \
    wget \
    && rm -rf /var/lib/apt/lists/*

# Set working directory
WORKDIR /app

# Install llama-cpp-python using the pre-built wheel
# This uses the CPU-only pre-built wheel from the additional index URL
RUN pip install --upgrade pip && \
    if [ "$USE_BEDROCK" = "false" ]; then \
        pip install llama-cpp-python --extra-index-url https://abetlen.github.io/llama-cpp-python/whl/cpu; \
    fi

    
COPY requirements.txt .
RUN pip install --no-cache-dir --upgrade -r requirements.txt
COPY . .


# Run the models download script
RUN python modelsdownload.py
EXPOSE 4282
ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1
CMD ["python", "main.py"]
