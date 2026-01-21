import json


def network_enrich(event):
    print("Clean payload going to pipeline:\n" + json.dumps(event, indent=2))
