import redis
from redisvl.index import SearchIndex
from redisvl.schema import IndexSchema
from redisvl.query import VectorQuery
import numpy as np


class RedisVector:
    def __init__(
        self, redis_connection, index_name=None, vector_dim=1536, openai_connection=None
    ):
        self.redis_connection = redis_connection
        self.index_name = index_name
        self.vector_dim = vector_dim
        self.index = None
        self.openai_connection = openai_connection

    def TestConnection(self):
        if self.redis_connection is None:
            raise Exception("No connection to Redis")
        else:
            print("Connected to Redis")

    def CreateIndex(self):
        schema_dict = {
            "index": {"name": "title", "type": "text"},
            "fields": [
                {"name": "page_id", "type": "text"},
                {"name": "title", "type": "text"},
                {"name": "body", "type": "text"},
                {
                    "name": "body_embedding",
                    "type": "vector",
                    "attrs": {
                        "dims": self.vector_dim,
                        "algorithm": "flat",
                        "datatype": "float32",
                        "distance_metric": "cosine",
                    },
                },
            ],
        }

        index = SearchIndex.from_dict(schema_dict)
        index.set_client(self.redis_connection.connection)
        index.create(overwrite=False)
        self.index = index

    def AddDataToIndex(self, data):
        if self.index is None or type(data) is not dict:
            raise Exception("No index created")
        else:
            print(f"Adding data to index: {data['id']}")
            self.index.load([data])

    def Search(self, vector, top_k=15):
        if self.index is None:
            raise Exception("No index created")
        else:
            query = VectorQuery(
                vector=[vector],
                vector_field_name="body_embedding",
                return_fields=["id", "page_id", "title", "body"],
                num_results=top_k,
            )
            results = self.index.query(query)
            return results

    def Vectorize(self, text):
        if self.index is None or self.openai_connection is None:
            raise Exception("No index created")
        else:
            length = self.openai_connection.num_tokens_from_string(text)
            print(f"Number of tokens: {length}")
            if length > 1024:
                text = text[:1024]

            vector = self.openai_connection.get_embedding(text)
            return np.array(vector, dtype=np.float32)

    def SearchText(self, text, top_k=15):
        if self.index is None or self.openai_connection is None:
            raise Exception("No index created")
        else:
            vector = self.Vectorize(text).tobytes()

            return self.Search(vector, top_k)
