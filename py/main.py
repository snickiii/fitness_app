from flask import Flask, jsonify, request
from flask_cors import CORS
import datetime

class User:
    def __init__(self, id, username, password, name, subname, email) -> None:
        self.id = id
        self.username = username
        self.password = password
        self.name = name
        self.subname = subname
        self.email = email
        self.created_at = None
    
    def toJson(self):
        return {
            "id": self.id,
            "username": str(self.username),
            "password": str(self.password),
            "name": str(self.name),
            "subname": str(self.subname),
            "mail": str(self.email),
            "created_at": str(datetime.datetime.now())
        }
    
user1 = User(1, 'user1', '97&G&$#f9c', 'John', 'Doe', 'johndoe@gmail.com')

app = Flask(__name__)

CORS(app)

@app.route("/user", methods=['GET'])
def get_user():
    return jsonify(user1.toJson())

@app.route("/user/update", methods=['PUT'])
def update_user():
    data = request.json

    user1.username = data['username']
    user1.password = data['password']
    user1.name = data['name']
    user1.subname = data['subname']
    user1.mail = data['email']

    return jsonify(user1.toJson())

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)