#!/usr/bin/env python

from flask import Flask
from flask import request
from restful_lib import Connection

import os
import json

def sink():
    headers = {}
    headers['Ce-Specversion']='1.0'
    headers['Ce-Time']=request.headers['Ce-Time']
    headers['Ce-Id']=request.headers['Ce-Id']
    headers['Ce-Source']='translators.triggermesh.io/partsunlimited-demo-translator'
    headers['Ce-Type']='io.triggermesh.targets.sink'

    return app.response_class(
            headers=headers,
            status=204,
    )


app = Flask(__name__)
@app.route('/', methods=['POST', 'GET'])
def trans():
    base_url = "https://partunlimited.demo.triggermesh.io:8080/api" # The API endpoint for the parts store
    conn = Connection(base_url)

    ce = request.get_json(force=True)
    print(request.data)
    print(request.headers)
    ceSource = request.headers['Ce-Source']

    headers = {}
    headers['Ce-Specversion']='1.0'
    headers['Ce-Time']=request.headers['Ce-Time']
    headers['Ce-Id']=request.headers['Ce-Id']
    headers['Ce-Source']='translators.triggermesh.io/partsunlimited-demo-translator'

    # For events we don't care about, just return
    if ceSource is not None and not ceSource.startswith('tmtestdb.demo.triggermesh.com/'):
        print("invalid source: " + ceSource)
        return sink()

    # Handle the replenishment event by posting a message to Zendesk
    if ceSource == "tmtestdb.demo.triggermesh.com/replenish":
        headers['Ce-Type']='com.zendesk.ticket.create'
        # Need to extract the manufacturer details
        resp = conn.request_get("/product/" + str(ce["new"]["ID"]))
        respBody = json.loads(resp[u'body'])

        if ce["op"] == "UPDATE" and ce["new"]["QUANTITY"] == 1:
            body = {
                "subject": "Parts Unlimited Replenishment Request",
                "body": "It is time to reorder " + respBody["name"] + " from " + respBody["manufacturer"]["manufacturer"]
            }

            return app.response_class(
                response=json.dumps(body),
                headers=headers,
                status=200,
                mimetype='application/json'
            )
        else:
            print("invalid replenish")
            return sink()

    # Handle the new order event by sending it to an Oracle Cloud function
    if ceSource == "tmtestdb.demo.triggermesh.com/neworder":
        headers['Ce-Type']='com.triggermesh.targets.oracle.function.partsunlimited-neworder'
        # Need to extract the order details
        resp = conn.request_get("/order/" + str(ce["new"]["ID"]))
        respBody = json.loads(resp[u'body'])

        if ce["op"] == "INSERT":
            body = {
                "name": respBody["user"]["name"],
                "address": respBody["user"]["address"],
                "totalCost": respBody["totalCost"],
                "paymentMethod": respBody["paymentType"],
                "ordered": respBody["dateOrdered"]
            }

            return app.response_class(
                response=json.dumps(body),
                headers=headers,
                status=200,
                mimetype='application/json'
            )
        else:
            print("invalid neworder")
            return sink()

    else:
        print("unknown source" + ceSource)
        return sink()

if __name__ == '__main__':
	app.run(host='0.0.0.0', port=8080)
