#!/usr/bin/env python3

from flask import Flask
from flask import request

import os
import json

app = Flask(__name__)

@app.route('/', methods=['POST', 'GET'])
def trans():
    try:
        ce = request.get_json(force=True)
        print(request.data)

        # respond with a cloudevent
        headers = {}
        headers['Ce-Specversion']='1.0'
        headers['Ce-Time']=request.headers['Ce-Time']
        headers['Ce-Id']=request.headers['Ce-Id']
        headers['Ce-Type']='io.triggermesh.datadog.metric.aggregated'
        headers['Ce-Source']=request.headers['Ce-Source'] + "/translated"
        headers['Ce-Subject']=request.headers['Ce-Subject']

        # Modify the Oracle OCI Monitoring data to be usable for Datadog
        data = []
        for e in ce:
            displayName = e['name']
            if "resourceDisplayName" in e['dimensions']:
                resourceGroup = e['dimensions']['resourceDisplayName']
            else:
                resourceGroup = e['dimensions']['resourceId']

            metricType = 'distribution'
            for s in e['aggregatedDatapoints']:
                datum = {}
                datum['displayName'] = displayName
                datum['resourceGroup'] = resourceGroup
                datum['timeStamp'] = s['timestamp']
                datum['value'] = s['value']
                datum['metricType'] = metricType
                data.append(datum)

        body = {}
        body['data'] = data

        return app.response_class(
            response=json.dumps(body),
            headers=headers,
            status=200,
            mimetype='application/json'
        )
    except:
        print("Invalid event")
        body = {}
        body['msg'] = "Unable to read event"
        return app.response_class(
            response=json.dumps(body),
            status=400,
            mimetype='application/json')


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)
