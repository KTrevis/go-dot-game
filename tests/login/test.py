import asyncio
from typing import Any
import websockets
from tests.ddos.test import testDDOS
from tests.utils import sendMessage
from .valid_login import validLogin
from .invalid_login import invalidLogin
from websockets.asyncio.client import ClientConnection
import json
from colored_print import log


async def doubleLogin():
    socket = await websockets.connect("ws://localhost:8080/websocket")
    msg = {
        "username": "test",
        "password": "test"
            }
    await sendMessage(socket, "LOGIN", msg)
    await sendMessage(socket, "LOGIN", msg)

    data: dict[str, Any] = json.loads(await socket.recv())
    data: dict[str, Any] = json.loads(await socket.recv())
    assert "error" in data
    assert data["error"] == "you are already authenticated"
    await socket.close()
    log.success("[DOUBLE LOGIN OK]\n")

async def testLogin():
    print("[LOGIN TESTS]\n")
    await asyncio.gather(
            invalidLogin(),
            validLogin(),
            )
    await doubleLogin()
