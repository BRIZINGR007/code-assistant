docker build --progress=plain --no-cache --build-arg USE_BEDROCK=false -t python-service .
docker run -p 4282:4282 python-service