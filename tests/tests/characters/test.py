import json
from colorama import Fore
import websockets
from typing import Any
from websockets.asyncio.client import ClientConnection

from ...utils import login, sendMessageAndRead

async def deleteCharacter(socket: ClientConnection):
    return await sendMessageAndRead(socket, "DELETE_CHARACTER", {"name": "test"})

async def createCharacter(socket: ClientConnection):
    return await sendMessageAndRead(socket, "CREATE_CHARACTER", {"name": "test", "class": "Mage"})

async def testCreateDelete():
    socket = await websockets.connect("ws://localhost:8080/websocket")
    await login(socket)

    data = await createCharacter(socket)
    assert "success" in data
    data = await sendMessageAndRead(socket, "GET_CHARACTERS", {})
    print(data)

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
    print(f"{Fore.GREEN}[CHARACTER CREATE DELETE OK]\n")

async def testCharacters():
    print(f"{Fore.GREEN}[CHARACTERS OK]\n")
