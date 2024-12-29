import asyncio
from typing import Any
import websockets
from colored_print import log
from websockets.asyncio.client import ClientConnection
import json

from .ddos.test import testDDOS
from .login.test import testLogin


async def main():
    await testLogin()
    await testDDOS()
    log.success("[ALL TESTS OK]\n")

asyncio.run(main())
