from db import retrieve_meter_data
from pymongo import MongoClient
import os
import asyncio
from bson.objectid import ObjectId
import motor.motor_asyncio

client = MongoClient(os.environ.get("MONGO_URL","localhost"), 27017)

# MONGO_DETAILS = os.environ.get("MONGO_URL","localhost")

# client = motor.motor_asyncio.AsyncIOMotorClient(MONGO_DETAILS)

database = client.metergen

meter_events = database.get_collection("meter_events")


