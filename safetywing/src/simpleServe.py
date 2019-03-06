from flask import Flask, request, render_template, jsonify

app = Flask(__name__)

@app.route("/", methods=['get'])
def home():
    return render_template("index.html")

@app.route("/api/login", methods=['post'])
def login():
    username = request.get_json()["username"]
    password = request.get_json()["password"]
    print(username, password)
    
    if username == "loveuer" and password == "1314loveuer":
        return jsonify({
            "StatusCode": 200,
            "Msg": "200",
            "Val": "new session",
        })
    else:
        return jsonify({
            "StatusCode": 400,
            "Msg": 400,
            "Val": "",
        })

@app.route("/api/session", methods=['get'])
def session():
    print(request.url)
    return jsonify({
        
    })


if __name__ == "__main__":
    app.run(debug=True, port=9119)