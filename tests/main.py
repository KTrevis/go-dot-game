import asyncio
from typing import Any
import websockets
from colored_print import log
from websockets.asyncio.client import ClientConnection
import json

from .login.test import testLogin

async def main():
    socket = await websockets.connect("ws://localhost:8080/websocket")
    await testLogin(socket)
    # socket = await websockets.connect("ws://localhost:8080/websocket")
    log.success("[ALL TESTS OK]\n")

asyncio.run(main())
