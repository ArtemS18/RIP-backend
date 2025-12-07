from pydantic import BaseModel

class Component(BaseModel):
	id: int
	title: str
	type: str
	mtbf: str
	mttr: str
	available: float


class ComponentsToSystemCalc(BaseModel):
	replication_count: int
	component_data: Component

class ComponentsToSystemCalcList(BaseModel):
	id: int
	components: list[ComponentsToSystemCalc]
	
class BaseReponse(BaseModel):
	status: str
	
class ResultRequest(BaseModel):
	available_calc: float
	token: str 
	sustem_calc_id: int