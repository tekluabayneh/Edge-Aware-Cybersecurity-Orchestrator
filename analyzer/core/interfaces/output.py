import json
import os

from requests import request
def send_output(event, token): 
  try:
    BASEURL = os.getenv("BASE_URL")
    if BASEURL is None: 
        return 

        response = requests.post(f"{BASEURL}/telementary/report", json=event,  headers={
            "Content-Type": "application/json",
            "Authorization":"Bearer " + token
            })

        response.raise_for_status()
  except Exception as e: 
      print(e)


# print("Clean payload going to pipeline:\n" + json.dumps(event, indent=2))
