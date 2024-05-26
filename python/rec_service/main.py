import threading
from flask import Flask, request, jsonify, Response
import os

from src.Process import Process

app = Flask(__name__)


def process_queue():
    print("Starting queue processing")
    pq = Process(name="page_queue")
    pq.run()
    pq.ProcessQueue()


proc = Process(name="page_queue")
proc.run()


@app.route("/search", methods=["POST"])
def search():
    data = request.get_json()

    if data is None or "text" not in data.keys():
        print("Missing text in request")
        return Response(status=400)

    vector = proc.Search(data["text"])

    return jsonify(vector)


if __name__ == "__main__":
    queue_thread = threading.Thread(target=process_queue)
    queue_thread.start()

    print("Starting API server")

    api_thread = threading.Thread(
        target=lambda: app.run(
            host="0.0.0.0", port=7999, debug=False, use_reloader=False
        )
    )
    api_thread.start()

    queue_thread.join()
    api_thread.join()
