#!/usr/bin/env python3

from flask import Flask
from flask import request

import os
import json

app = Flask(__name__)

@app.route('/', methods=['POST', 'GET'])
def trans():
    ce = request.get_json(force=True)
    print(request.data)

    # respond with a cloudevent
    headers = {}
    headers['Ce-Specversion']='1.0'
    headers['Ce-Time']=request.headers['Ce-Time']
    headers['Ce-Id']=request.headers['Ce-Id']
    headers['Content-Type']='application/json'
    headers['Ce-Type']='com.triggermesh.targets.oracledb.fn.dumper'
    headers['Ce-Source']='translators.triggermesh.io/oracle-source-translator'

    # FOR SQS
    body = { "message": "missing oracle db change data"}

    if ce is not None and "op" in ce:
        if ce["op"] == "INSERT":
            body = { "message": "A new user wants to say something: " + ce["new"]["DATA"] }
        elif ce["op"] == "DELETE":
            body = { "message": "A user wants to say goodbye: " + ce["old"]["DATA"] }
        else:
            body = { "message": "Don't know what to do with this: " + json.dumps(ce) }

        return app.response_class(
            response=json.dumps(body),
            headers=headers,
            status=200,
            mimetype='application/json'
    )
    else:
        print("Unkown event, skipping...")

        return app.response_class(
            response=json.dumps(body),
            headers=headers,
            status=500,
            mimetype='application/json'
        )

    return response 

if __name__ == '__main__':
	app.run(host='0.0.0.0', port=8080)

