from openai import OpenAI
import tiktoken


class OpenAIConnection:
    def __init__(self, api_key):
        self.api_key = api_key
        self.openai = None

    def connect(self):
        self.openai = OpenAI(api_key=self.api_key)

    def get_connection(self):
        return self.openai

    def get_api_key(self):
        return self.api_key

    def get_embedding(self, text):
        if self.openai is None:
            raise Exception("No connection to OpenAI")

        try:
            self.num_tokens_from_string(text)
            text = text.strip().replace("\n", " ")
            return (
                self.openai.embeddings.create(
                    input=[text], model="text-embedding-3-small"
                )
                .data[0]
                .embedding
            )
        except Exception as e:
            print(f"Error: {e}")
            return None

    def num_tokens_from_string(self, text):
        text = text.strip().replace("\n", " ")
        encoding = tiktoken.get_encoding("cl100k_base")
        tokens = encoding.encode(text)
        return len(tokens)
