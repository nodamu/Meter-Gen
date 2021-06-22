#!/bin/bash

exec ./run.sh &
exec python receiver.py  &
exec uvicorn main:app --host=0.0.0.0

