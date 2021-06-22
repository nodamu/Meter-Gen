from db import retrieve_student
from fastapi import FastAPI

app = FastAPI()


@app.get("/{meterId}")
async def getMeterEvent(meterId: str,startDate : str, endDate: str):
    return await retrieve_student(meterId)




