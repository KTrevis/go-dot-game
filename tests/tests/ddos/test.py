import asyncio
import websockets
from websockets.asyncio.client import ClientConnection
from colorama import Fore

from ...utils import sendMessage

sockets: list[ClientConnection] = []

async def createSocket():
    socket = await websockets.connect("ws://localhost:8080/websocket")
    sockets.append(socket)
    await sendMessage(socket, "LOGIN", {"username": "test", "password": "test"})
    await sendMessage(socket, "LOGIN", {"username": "test", "password": "test"})

async def testDDOS():
    tasks: list[asyncio.Task[None]] = []

    for _ in range(1000):
        task = asyncio.create_task(createSocket())
        tasks.append(task)
    _ = await asyncio.gather(*tasks)

    for socket in sockets:
        task = asyncio.create_task(socket.close())
        tasks.append(task)
    _ = await asyncio.gather(*tasks)
    print(f"{Fore.GREEN}[TEST DDOS OK]\n")
