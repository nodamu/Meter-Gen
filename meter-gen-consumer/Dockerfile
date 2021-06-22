
# pull the official docker image
FROM python:3.9.4-slim

# set work directory
WORKDIR /app

# set env variables
ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

# install dependencies
COPY requirements.txt .
RUN pip install -r requirements.txt

EXPOSE 8000

# copy project
COPY . /app/

CMD ["uvicorn", "meter-gen-consumer.main:app", "--host=0.0.0.0", "--reload"]
