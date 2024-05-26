import redis


class RedisConnection:
    def __init__(self, host, port):
        self.host = host
        self.port = port
        self.connection = None

    def connect(self):
        self.connection = redis.Redis(host=self.host, port=self.port)
        respc = self.connection.ping()
        if respc:
            print("Connected to Redis")
        else:
            print("Failed to connect to Redis")

    def get_connection(self):
        return self.connection

    def read_queue(self, queue_name):
        if self.connection is None:
            raise Exception("No connection to Redis")

        while True:
            key, value = self.connection.blpop([queue_name])  # type: ignore
            yield value.decode("utf-8")

    def write_queue(self, queue_name, value):
        if self.connection is None:
            raise Exception("No connection to Redis")

        self.connection.rpush(queue_name, value)
