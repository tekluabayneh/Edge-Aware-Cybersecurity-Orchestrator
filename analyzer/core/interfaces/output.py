import json
def send_output(event): 
    print("Clean payload going to pipeline:\n" + json.dumps(event, indent=2))
