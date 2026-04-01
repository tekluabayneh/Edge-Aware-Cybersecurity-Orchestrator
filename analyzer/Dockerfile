FROM python:3.12

WORKDIR /app


RUN apt-get update && apt-get install -y curl


RUN curl -Ls https://astral.sh/uv/install.sh | sh


ENV PATH="/root/.local/bin:$PATH"


COPY . .


RUN uv sync

# run app
CMD ["uv", "run", "uvicorn", "main:app", "--host", "0.0.0.0", "--port", "9090"]
