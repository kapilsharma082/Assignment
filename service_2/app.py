from flask import Flask, jsonify

app = Flask(__name__)

@app.route("/ping", methods=["GET"])
def ping():
    return jsonify(status="ok", service="2"), 200

@app.route("/health", methods=["GET"])
def health():
    # Return JSON for easier automated health checks
    return jsonify(status="OK"), 200

@app.route("/hello", methods=["GET"])
def hello():
    return jsonify(message="Hello from Service 2"), 200

if __name__ == "__main__":
    # Listen on 0.0.0.0 so Docker can expose it externally
    app.run(host="0.0.0.0", port=5002)
