from pymongo import MongoClient
import os
import asyncio
from bson.objectid import ObjectId
import motor.motor_asyncio

client = MongoClient(os.environ.get("MONGO_URL","localhost"), 27017)

MONGO_DETAILS = os.environ.get("MONGO_URL","localhost")

client = motor.motor_asyncio.AsyncIOMotorClient(MONGO_DETAILS)

database = client.metergen

meter_events = database.get_collection("meter_events")

# save message
def saveMessage(msg):
    database = client['metergen']
    meter_events = database['meter_events']
    meter_events.insert(msg)

# Retrieve a meter with a matching ID
async def retrieve_student(meterId: str) -> dict:
    event = await meter_events.find_one({"smart_meter_id": meterId})
    if event:
        return event

