FROM python:3.7-slim

COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt

COPY oci-datadog-trans.py /oci-datadog-trans.py

ENTRYPOINT ["python", "-u", "/oci-datadog-trans.py"]
