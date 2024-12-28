import asyncio
import websockets
from colored_print import log
from websockets.asyncio.client import ClientConnection

from tests.utils import sendMessage

async def createSocket():
    socket = await websockets.connect("ws://localhost:8080/websocket")
    await sendMessage(socket, "LOGIN", {"username": "test", "password": "test"})
    await sendMessage(socket, "LOGIN", {"username": "test", "password": "test"})
    await socket.close()

async def testDDOS():
    tasks = []

    for _ in range(1000):
        task = asyncio.create_task(createSocket())
        tasks.append(task)
    await asyncio.gather(*tasks)
    log.success("[TEST DDOS OK]\n")
