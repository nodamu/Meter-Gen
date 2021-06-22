from typing import Any, List
from db import retrieve_meter_data
from fastapi import FastAPI

app = FastAPI()


@app.get("/{meterId}")
def getMeterEvent(meterId: str,startDate : str, endDate: str, limit: int) -> Any:
    meter_data =  retrieve_meter_data(meterId,startDate,endDate,limit)

    return meter_data


