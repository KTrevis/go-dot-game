from colorama import Fore
import websockets
from typing import Any
from websockets.asyncio.client import ClientConnection

from ...utils import login, read, sendMessage

async def deleteCharacter(socket: ClientConnection):
    await sendMessage(socket, "DELETE_CHARACTER", {"name": "test"})
    return await read(socket)

async def createCharacter(socket: ClientConnection):
    await sendMessage(socket, "CREATE_CHARACTER", {"name": "test", "class": "Mage"})
    return await read(socket)

async def loggedInTest():
    socket = await websockets.connect("ws://localhost:8080/websocket")
    _ = await login(socket)

    data = await deleteCharacter(socket)

    msgType, data = await createCharacter(socket)
    assert msgType == "CREATE_CHARACTER"
    assert "success" in data

    await sendMessage(socket, "GET_CHARACTER_LIST", {})
    msgType, data = await read(socket)
    arr: list[dict[str, Any]] = data["characterList"]
    assert len(arr) == 1
    assert msgType == "GET_CHARACTER_LIST"

    msgType, data = await createCharacter(socket)
    assert msgType == "CREATE_CHARACTER"
    assert "error" in data

    await sendMessage(socket, "GET_CHARACTER_LIST", {})
    msgType, data = await read(socket)
    arr: list[dict[str, Any]] = data["characterList"]
    assert len(arr) == 1
    assert msgType == "GET_CHARACTER_LIST"

    msgType, data = await deleteCharacter(socket)
    assert msgType == "DELETE_CHARACTER"
    assert "success" in data

    await sendMessage(socket, "GET_CHARACTER_LIST", {})
    msgType, data = await read(socket)
    arr: list[dict[str, Any]] = data["characterList"]
    assert len(arr) == 0
    assert msgType == "GET_CHARACTER_LIST"

    msgType, data = await deleteCharacter(socket)
    assert "error" in data
    assert msgType == "DELETE_CHARACTER"

    await sendMessage(socket, "GET_CHARACTER_LIST", {})
    msgType, data = await read(socket)
    arr: list[dict[str, Any]] = data["characterList"]
    assert len(arr) == 0
    assert msgType == "GET_CHARACTER_LIST"

    msgType, data = await createCharacter(socket)
    assert "success" in data
    assert msgType == "CREATE_CHARACTER"

    await sendMessage(socket, "GET_CHARACTER_LIST", {})
    msgType, data = await read(socket)
    arr: list[dict[str, Any]] = data["characterList"]
    assert len(arr) == 1
    assert msgType == "GET_CHARACTER_LIST"

    msgType, data = await deleteCharacter(socket)
    assert "success" in data
    assert msgType == "DELETE_CHARACTER"

    await sendMessage(socket, "GET_CHARACTER_LIST", {})
    msgType, data = await read(socket)
    arr: list[dict[str, Any]] = data["characterList"]
    assert len(arr) == 0
    assert msgType == "GET_CHARACTER_LIST"

    await socket.close()

async def loggedOutTest():
    socket = await websockets.connect("ws://localhost:8080/websocket")
    _ = await login(socket)

    await sendMessage(socket, "CREATE_CHARACTER", {"name": "test", "class": "Mage"})
    msgType, data = await read(socket)
    assert msgType == "CREATE_CHARACTER"
    assert "success" in data

    await socket.close()
    socket = await websockets.connect("ws://localhost:8080/websocket")

    await sendMessage(socket, "GET_CHARACTER_LIST", {})
    msgType, data = await read(socket)
    assert msgType == "GET_CHARACTER_LIST"
    assert "error" in data

    await sendMessage(socket, "DELETE_CHARACTER", {"name": "test", "class": "Mage"})
    msgType, data = await read(socket)
    assert msgType == "DELETE_CHARACTER"
    assert "error" in data

    await sendMessage(socket, "CREATE_CHARACTER", {"name": "oui", "class": "Mage"})
    msgType, data = await read(socket)
    assert msgType == "CREATE_CHARACTER"
    assert "error" in data

    _ = await login(socket)

    await sendMessage(socket, "DELETE_CHARACTER", {"name": "test", "class": "Mage"})
    msgType, data = await read(socket)
    assert "success" in data
    assert msgType == "DELETE_CHARACTER"

async def testCharacters():
    await loggedInTest()
    await loggedOutTest()
    print(f"{Fore.GREEN}[CHARACTERS OK]\n")
