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
    headers['ce-specversion']='1.0'
    headers['ce-time']='2018-04-05T03:56:24Z'
    headers['ce-id']='XXX-YYY-ZZZ-WWW'
    headers['content-type']='application/json'
    headers['ce-type']='com.triggermesh.transform'
    headers['ce-source']='oracle-source'
    
    # FOR SQS
    body = { "message": "missing ce data"}

    if ce is not None:
        if ce["op"] == "INSERT":
            body = { "message": "A new user wants to say something: " + ce["new"]["DATA"] }
        elif ce["op"] == "DELETE":
            body = { "message": "A user wants to say goodbye: " + ce["old"]["DATA"] }
        else:
            body = { "message": "Don't know what to do with this: " + json.dumps(ce) }

    print(body)
    response = app.response_class(
            response=json.dumps(body),
            headers=headers,
            status=200,
            mimetype='application/json'
    )

    return response 

if __name__ == '__main__':
	app.run(host='0.0.0.0', port=8080)

