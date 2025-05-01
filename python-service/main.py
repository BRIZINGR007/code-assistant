from fastapi import FastAPI
import uvicorn
from fastapi.middleware.cors import CORSMiddleware
from decouple import config
from app.routes import encoder, decoder, cross_encoder
from app.settings.LifeSpans import (
    Lifespans,
    bootstrap_lifespan,
)

app = FastAPI(
    title="AI Service",
    description="A  User-Service for User Related MetaData.",
    version="0.0.1",
    lifespan=Lifespans([bootstrap_lifespan]),
)
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

app.include_router(encoder.router)
app.include_router(decoder.router)
app.include_router(cross_encoder.router)

if __name__ == "__main__":
    workers = 1
    uvicorn.run(
        "main:app",
        port=4282,
        host="0.0.0.0",
        reload=False,
        workers=workers,
        lifespan="on",
    )
