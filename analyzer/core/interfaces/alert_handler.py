import requests
import json
import os
def send_alert(alert, token):
    """
    Add user/agent info and send to backend
    """
    try:
        BASE_URL = os.getenv("BACKEND_BASE_URL") 
        if BASE_URL is None: 
            print("BACKEND_BASE_URL not found", BASE_URL)
            return


        response = requests.post(f"{BASE_URL}/newAlert/create", json=alert,  headers={
            "Content-Type": "application/json",
            "Authorization":"Bearer " + token
            })

        response.raise_for_status()
    except Exception as e:
        print(f"Failed to send alert: {e}")
