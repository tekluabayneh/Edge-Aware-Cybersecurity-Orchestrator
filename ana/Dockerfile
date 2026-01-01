# 1️⃣ Use an official Python base image
FROM python:3.11-slim


WORKDIR /app


COPY . .


RUN pip install uv


RUN uv sync


CMD ["uv", "run", "python", "main.py"]
