FROM python:3.11-slim

WORKDIR /app

# install uv
RUN pip install uv

# copy project files
COPY . .

# install dependencies using uv
RUN uv sync

# run app
CMD ["uv", "run", "uvicorn", "main:app", "--host", "0.0.0.0", "--port", "9090"]
