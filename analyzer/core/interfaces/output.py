import json
import os

import requests 
def send_output(event, token): 
   print("Clean payload going to pipeline:\n" + json.dumps(event.get("security"), indent=2))
   try:
     BASEURL = os.getenv("BACKEND_BASE_URL")
     if BASEURL is None: 
        return 

     response = requests.post(f"{BASEURL}/api/telementary/report", json=event,  headers={
      "Content-Type": "application/json",
            "Authorization":"Bearer " + token
      })
     print(response.status_code)
     print(response.text)
     # response.raise_for_status()
     # print("Success:", response.json().get("message", "No message key found"))

   except Exception as e: 
      print("error",e)


