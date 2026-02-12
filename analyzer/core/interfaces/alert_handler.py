import requests

def send_alert(alert, context):
    """
    Add user/agent info and send to backend
    """
    payload = {
        "rule_id": alert["rule_id"],
        "severity": alert["severity"],
        "title": alert["title"],
        "event": alert["event"],
        "user_id": context.user_id,
        "email": context.email,
        "agent_id": context.agent_id,
        "agent_token": context.agent_token
    }

    try:
        response = requests.post("https://backend/api/alert", json=payload)
        response.raise_for_status()
        print(f"Alert sent successfully: {payload['rule_id']}")
    except Exception as e:
        print(f"Failed to send alert: {e}")
