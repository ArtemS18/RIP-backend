import asyncio
from httpx import AsyncClient
from app.schema import ComponentsToSystemCalcList, ResultRequest


BACKEND_URL = "http://localhost:8080/api/components"
TOKEN = "token"

async def calculate_avalilable(items: ComponentsToSystemCalcList) -> float:
    system_available: float = 1
    for i in items.components:
        component = i.component_data
        available = component.available
        if available == 0:
            if component.mtbf == 0 or component.mtbf+component.mttr == 0:
                raise 
            available = float(component.mtbf) / float(component.mtbf+component.mttr)
        component_available = float (1 - (float(1-available) ** float(i.replication_count)))
        system_available *= component_available
    return system_available

async def send_results(items: ComponentsToSystemCalcList):
    system_available = await calculate_avalilable(items)
    await asyncio.sleep(5)

    async with AsyncClient() as client:
        resp = await client.get(BACKEND_URL, json=ResultRequest(
            available_calc=system_available, 
            token=TOKEN, 
            sustem_calc_id=items.id
            ).model_dump()
        )
        print(resp)
