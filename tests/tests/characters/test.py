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

async def testCharacters():
    socket = await websockets.connect("ws://localhost:8080/websocket")
    await login(socket)

    data = await sendMessageAndRead(socket, "GET_CHARACTER_LIST", {})
    arr: list[dict[str, Any]] = data["characterList"]
    assert len(arr) == 0

    data = await createCharacter(socket)
    assert "success" in data

    data = await sendMessageAndRead(socket, "GET_CHARACTER_LIST", {})
    arr: list[dict[str, Any]] = data["characterList"]
    assert len(arr) == 1

    data = await createCharacter(socket)
    assert "error" in data

    data = await sendMessageAndRead(socket, "GET_CHARACTER_LIST", {})
    arr: list[dict[str, Any]] = data["characterList"]
    assert len(arr) == 1

    data = await deleteCharacter(socket)
    assert "success" in data

    data = await sendMessageAndRead(socket, "GET_CHARACTER_LIST", {})
    arr: list[dict[str, Any]] = data["characterList"]
    assert len(arr) == 0

    data = await deleteCharacter(socket)
    assert "error" in data

    data = await sendMessageAndRead(socket, "GET_CHARACTER_LIST", {})
    arr: list[dict[str, Any]] = data["characterList"]
    assert len(arr) == 0


    data = await createCharacter(socket)
    assert "success" in data

    data = await sendMessageAndRead(socket, "GET_CHARACTER_LIST", {})
    arr: list[dict[str, Any]] = data["characterList"]
    assert len(arr) == 1

    data = await deleteCharacter(socket)
    assert "success" in data

    data = await sendMessageAndRead(socket, "GET_CHARACTER_LIST", {})
    arr: list[dict[str, Any]] = data["characterList"]
    assert len(arr) == 0

    await socket.close()
    print(f"{Fore.GREEN}[CHARACTERS OK]\n")
