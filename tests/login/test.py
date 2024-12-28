import asyncio
from typing import Any
import websockets
from tests.utils import sendMessage
from .valid_login import validLogin
from .invalid_login import invalidLogin
from websockets.asyncio.client import ClientConnection
import json
from colored_print import log

async def loginDifferentSocket():
    socket = await websockets.connect("ws://localhost:8080/websocket")
    msg = {
        "username": "test",
        "password": "test"
            }
    await sendMessage(socket, "LOGIN", msg)
    data: dict[str, Any] = json.loads(await socket.recv())

    assert "error" in data
    assert data["error"] == "this account is already logged in"
    log.success("[LOGIN TO SAME ACCOUNT DIFFERENT SOCKET OK]\n")


async def doubleLogin(socket: ClientConnection):
    msg = {
        "username": "test",
        "password": "test"
            }
    await sendMessage(socket, "LOGIN", msg)

    data: dict[str, Any] = json.loads(await socket.recv())
    assert "error" in data
    assert data["error"] == "you are already authenticated"
    log.success("[DOUBLE LOGIN OK]\n")

async def testLogin(socket: ClientConnection):
    print("[LOGIN TESTS]\n")
    await asyncio.gather(
            invalidLogin(),
            validLogin(socket),
            )
    await loginDifferentSocket()
    await doubleLogin(socket)
