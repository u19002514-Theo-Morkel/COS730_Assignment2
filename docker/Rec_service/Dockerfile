FROM python:3.12.3

WORKDIR /app

COPY ./python/rec_service/requirements.txt requirements.txt

RUN pip install -r requirements.txt

COPY ./python/rec_service/main.py main.py
COPY ./python/rec_service/src ./src

CMD ["python", "main.py"]