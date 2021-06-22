import datetime
from db import saveMessage
import pika, os, time, logging, json
import asyncio

logging.basicConfig(format='%(asctime)s - %(message)s', level=logging.INFO)
logging.basicConfig()



def process_meter_event(event):
  # print(" [x] Received " + str(msg))
  # logging.info(event)

  payload = json.loads(event)

  logging.info(payload)
  payload["timestamp"] = datetime.datetime.strptime( payload["timestamp"], '%Y-%m-%dT%H:%M:%S%z')

  saveMessage(payload)

  # delays for 5 seconds
  time.sleep(5) 
  return #payload;

# Access the CLODUAMQP_URL environment variable and parse it (fallback to localhost)
url = os.environ.get('MQADD', 'amqp://guest:guest@localhost:5672/%2f')
params = pika.URLParameters(url)
connection = pika.BlockingConnection(params)
channel = connection.channel() # start a channel
channel.queue_declare(queue='queue') # Declare a queue

# create a function which is called on incoming messages
def callback(ch, method, properties, body):
  process_meter_event(body)



# set up subscription on the queue
channel.basic_consume('meter-queue',
  callback,
  auto_ack=True)

# start consuming (blocks)
channel.start_consuming()
connection.close()