FROM python:2.7-slim

ARG TARGET

COPY requirements.txt .
RUN pip install -r requirements.txt

COPY trans-${TARGET}.py /trans.py

ENTRYPOINT ["python", "-u", "/trans.py"]
