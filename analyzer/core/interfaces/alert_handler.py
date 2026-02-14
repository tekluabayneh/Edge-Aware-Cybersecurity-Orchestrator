import requests
import json
import os
def send_alert(alert):
    """
    Add user/agent info and send to backend
    """
    print("Clean payload going to pipeline in send:\n" + json.dumps(alert, indent=2))
    
    try:
        BASE_URL = os.getenv("BACKEND_BASE_URL") 
        if BASE_URL is None: 
            print("BACKEND_BASE_URL not found")
            return
        response = requests.post(f"{BASE_URL}/newAlert/create", json=alert,  headers={"Content-Type": "application/json"})
        print("api response",response)
        response.raise_for_status()
        print(f"Alert sent successfully")
    except Exception as e:
        print(f"Failed to send alert: {e}")
