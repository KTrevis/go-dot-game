import json
import websockets
from typing import Any
from websockets.asyncio.client import ClientConnection

from tests.utils import login, sendMessage, sendMessageAndRead

async def deleteCharacter(socket: ClientConnection):
    return await sendMessageAndRead(socket, "DELETE_CHARACTER", {"name": "test", "class": "test"})

async def createCharacter(socket: ClientConnection):
    return await sendMessageAndRead(socket, "CREATE_CHARACTER", {"name": "test", "class": "test"})

async def testCharacters():
    socket = await websockets.connect("ws://localhost:8080/websocket")
    await login(socket)

    data = await createCharacter(socket)
    assert "success" in data

    data = await createCharacter(socket)
    assert "error" in data

    data = await deleteCharacter(socket)
    assert "success" in data

    data = await deleteCharacter(socket)
    assert "error" in data

    data = await createCharacter(socket)
    assert "success" in data

    data = await deleteCharacter(socket)
    assert "success" in data

    await socket.close()
