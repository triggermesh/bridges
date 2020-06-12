FROM python:3.7-slim

COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt

COPY trans.py /trans.py

ENTRYPOINT ["python", "-u", "/trans.py"]

