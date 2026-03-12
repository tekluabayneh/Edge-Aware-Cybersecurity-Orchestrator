import json
import os

import requests 
def send_output(event, token): 
   try:
     BASEURL = os.getenv("BACKEND_BASE_URL")
     print("base url",BASEURL)
     if BASEURL is None: 
        return 

     print(event.get("email"))
     print(event.get("agent_id"))
     print(event.get("agent_token"))

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


   # print("Clean payload going to pipeline:\n" + json.dumps(event.get("integ"), indent=2))
