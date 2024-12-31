from colorama import Fore
import websockets
from typing import Any
from websockets.asyncio.client import ClientConnection

from ...utils import login, sendMessageAndRead

async def deleteCharacter(socket: ClientConnection):
    return await sendMessageAndRead(socket, "DELETE_CHARACTER", {"name": "test"})

async def createCharacter(socket: ClientConnection):
    return await sendMessageAndRead(socket, "CREATE_CHARACTER", {"name": "test", "class": "Mage"})

async def loggedInTest():
    socket = await websockets.connect("ws://localhost:8080/websocket")
    await login(socket)

    data = await deleteCharacter(socket)

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

async def loggedOutTest():
    socket = await websockets.connect("ws://localhost:8080/websocket")
    await login(socket)

    data = await sendMessageAndRead(socket, "CREATE_CHARACTER", {"name": "test", "class": "Mage"})
    assert "success" in data

    await socket.close()
    socket = await websockets.connect("ws://localhost:8080/websocket")

    data = await sendMessageAndRead(socket, "GET_CHARACTER_LIST", {})
    assert "error" in data

    data = await sendMessageAndRead(socket, "DELETE_CHARACTER", {"name": "test", "class": "Mage"})
    assert "error" in data

    data = await sendMessageAndRead(socket, "CREATE_CHARACTER", {"name": "oui", "class": "Mage"})
    assert "error" in data

    await login(socket)

    data = await sendMessageAndRead(socket, "DELETE_CHARACTER", {"name": "test", "class": "Mage"})
    assert "success" in data

async def testCharacters():
    await loggedInTest()
    await loggedOutTest()
    print(f"{Fore.GREEN}[CHARACTERS OK]\n")
