from typing import Any, Type
from pymongo import MongoClient
import os, logging, datetime

client = MongoClient(os.environ.get("MONGO_URL","localhost"), 27017)

logging.basicConfig(format='%(asctime)s - %(message)s', level=logging.INFO)
# logging.basicConfig()
# MONGO_DETAILS = os.environ.get("MONGO_URL","localhost")

# client = motor.motor_asyncio.AsyncIOMotorClient(MONGO_DETAILS)

database = client.metergen

meter_events = database.get_collection("meter_events")

# save message
def saveMessage(msg):
    database = client['metergen']
    meter_events = database['meter_events']
    meter_events.insert(msg)

# Retrieve a meter with a matching ID and date time
def retrieve_meter_data(meterId: str, startDate : str, endDate: str, limit : int) -> Any:
    try:
        from_date = datetime.datetime.strptime(startDate, '%Y-%m-%dT%H:%M:%S%z')
        to_date = datetime.datetime.strptime(endDate, '%Y-%m-%dT%H:%M:%S%z')
        event =  meter_events.find({'smart_meter_id':meterId,'timestamp': {'$gte':from_date,'$lt':to_date}})
        data = []
        for e in event:
            data.append(e)
        return data
    except TypeError:
        logging.info("Query failed")

print(retrieve_meter_data("hello","2009-01-06T09:19:18Z","2009-01-06T09:20:28Z",5))