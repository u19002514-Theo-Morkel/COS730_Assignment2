import threading

import os

from src.Process import Process


def process_queue():
    print("Starting queue processing")
    pq = Process(name="comment_queue")
    pq.run()
    pq.ProcessQueue()


if __name__ == "__main__":
    process_queue()
