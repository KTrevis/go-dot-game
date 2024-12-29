import asyncio
from colored_print import log

from .characters.test import testCharacters
from .ddos.test import testDDOS
from .login.test import testLogin

async def main():
    # await testLogin()
    await asyncio.gather(
        testLogin(),
        # testCharacters(),
    )
    # await testDDOS()
    log.success("[ALL TESTS OK]\n")

asyncio.run(main())
