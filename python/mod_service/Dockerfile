FROM python:3.12.3

WORKDIR /app

COPY requirements.txt requirements.txt

RUN pip install -r requirements.txt

COPY main.py main.py
COPY ./src ./src

CMD ["python", "main.py"]