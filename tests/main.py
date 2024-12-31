import asyncio
from colorama import Fore

from .tests.characters.test import testCharacters
from .tests.ddos.test import testDDOS
from .tests.login.test import testLogin

async def main():
    await testLogin()
    _ = await asyncio.gather(
        testCharacters(),
    )
    # await testDDOS()
    print(f"{Fore.GREEN}[ALL TESTS OK]\n")

asyncio.run(main())
