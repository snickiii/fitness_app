from flask import Flask, jsonify
import json
import datetime

app = Flask(__name__)

@app.route("/")
def hello_world():
    a = {
        "id": 1,
        "username": "user1",
        "password": "97&G&$#f9c",
        "name": "John",
        "subname": "Doe",
        "mail": "johndoe@gmail.com",
        "role": "user",
        "created_at": datetime.datetime.now()
    }
    return jsonify(a)