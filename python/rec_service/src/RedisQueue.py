class RedisQueue(object):
    def __init__(self, name, redis_connection=None):
        self.name = name
        self.redis_connection = redis_connection

    def TestConnection(self):
        if self.redis_connection is None:
            raise Exception("No connection to Redis")
        else:
            print("Connected to Redis")

    def ReadQueue(self):
        if self.redis_connection is None:
            raise Exception("No connection to Redis")
        else:
            while True:
                key, value = self.redis_connection.connection.blpop(self.name)
                yield value.decode("utf-8")

    def WriteQueue(self, value):
        if self.redis_connection is None:
            raise Exception("No connection to Redis")
        else:
            self.redis_connection.connection.rpush(self.name, value)
