from fastapi import APIRouter, status, BackgroundTasks
from app.service import send_results
from app.schema import ComponentsToSystemCalcList, BaseReponse

router = APIRouter(prefix="")


@router.post("/calc", status_code=status.HTTP_200_OK, response_model=BaseReponse)
async def create_calulate_task(items: ComponentsToSystemCalcList,  bg: BackgroundTasks):
    bg.add_task(send_results, items)
    return BaseReponse(status="ok")
