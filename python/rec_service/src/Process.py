from src.interfaces.Redis_Interface import RedisConnection
from src.interfaces.OpenAI_Interface import OpenAIConnection
from src.RedisQueue import RedisQueue
from src.RedisVector import RedisVector

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
        self.vector = None

    def run(self):
        self.Connect()
        self.TestConnection()
        self.ConnectToQueue()
        self.ConnectToVector()

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

    def ConnectToVector(self):
        if self.redis_connection is None or self.openai_connection is None:
            raise Exception("No connection to Redis or OpenAI")
        else:
            self.vector = RedisVector(
                redis_connection=self.redis_connection,
                index_name="title",
                vector_dim=1536,
                openai_connection=self.openai_connection,
            )
            self.vector.CreateIndex()

    def ProcessQueue(self):
        if self.queue is None or self.vector is None or self.openai_connection is None:
            raise Exception("No connection to queue or vector")
        else:
            for item in self.queue.ReadQueue():
                parsed_item = json.loads(item)

                if (
                    "id" in parsed_item
                    and "body" in parsed_item
                    and "title" in parsed_item
                ):
                    print(f"Processing item: {parsed_item['id']}")

                    parsed_item["page_id"] = parsed_item["id"]
                    parsed_item["body_embedding"] = self.vector.Vectorize(
                        parsed_item["body"]
                    ).tobytes()
                    self.vector.AddDataToIndex(parsed_item)

    def Search(self, text):
        if self.vector is None:
            raise Exception("No connection to vector")
        else:
            print(f"Searching for text: {text}")
            vector = self.vector.Vectorize(text)
            results = self.vector.Search(vector)
            print(f"Found {len(results)} results")
            print(results)
            return results
