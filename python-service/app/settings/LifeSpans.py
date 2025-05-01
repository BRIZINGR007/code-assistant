import os
import contextlib
from contextlib import asynccontextmanager
from typing import cast
from fastapi import FastAPI

from app.deocders.DecoderHandlers import DecoderHandler
from app.encoders.EncoderHandlers import EmbeddingGenerator
from bootstrap import BootStrap


@asynccontextmanager
async def bootstrap_lifespan(app: FastAPI):
    BootStrap()()
    yield


class Lifespans:
    def __init__(self, lifespans) -> None:
        self.lifespans = lifespans

    @asynccontextmanager
    async def __manage_lifespan(self, app: FastAPI):
        async with contextlib.AsyncExitStack() as exit_stack:
            for lifespan in self.lifespans:
                await exit_stack.enter_async_context(lifespan(app))
            yield

    def __call__(self, app: FastAPI):
        self.app = app
        return self.__manage_lifespan(app)
