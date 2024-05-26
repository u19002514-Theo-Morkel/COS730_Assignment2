from src.interfaces.Redis_Interface import RedisConnection
from src.interfaces.OpenAI_Interface import OpenAIConnection
from src.RedisQueue import RedisQueue

from dotenv import load_dotenv
import json
import numpy as np
import os


class Process:
    def __init__(self, name, redis_connection=None, openai_connection=None):
        self.name = name
        self.redis_connection = redis_connection
        self.openai_connection = openai_connection
        self.queue = None

    def run(self):
        self.Connect()
        self.TestConnection()
        self.ConnectToQueue()

    def Connect(self):
        if self.redis_connection is None:
            self.redis_connection = RedisConnection(host="192.168.3.105", port=6379)
        self.redis_connection.connect()

        if self.openai_connection is None:
            load_dotenv()

            api_key = os.getenv("OPENAI_API_KEY")
            if api_key is None:
                raise Exception("No OpenAI API key found")
            self.openai_connection = OpenAIConnection(api_key=api_key)
        self.openai_connection.connect()

    def TestConnection(self):
        if self.redis_connection is None or self.openai_connection is None:
            raise Exception("No connection to Redis or OpenAI")
        else:
            print("Connected to Redis and OpenAI")

    def ConnectToQueue(self):
        if self.redis_connection is None:
            raise Exception("No connection to Redis")
        else:
            self.queue = RedisQueue(
                name=self.name, redis_connection=self.redis_connection
            )

    def ProcessQueue(self):
        if self.queue is None or self.openai_connection is None:
            raise Exception("No connection to queue or vector")
        else:
            for item in self.queue.ReadQueue():
                parsed_item = json.loads(item)

                if "id" in parsed_item and "text" in parsed_item:
                    print(f"Processing item: {parsed_item['id']}")
                    moderation = self.openai_connection.moderation(
                        text=parsed_item["text"]
                    )
                    parsed_item["commentId"] = parsed_item["id"]
                    parsed_item["moderation"] = moderation.model_dump()["results"][0]  # type: ignore
                    parsed_item["flagged"] = parsed_item["moderation"]["flagged"]
                    print(f"Processed item: {parsed_item['id']}")
                    print(parsed_item)

                    self.queue.WriteQueue(value=parsed_item)
