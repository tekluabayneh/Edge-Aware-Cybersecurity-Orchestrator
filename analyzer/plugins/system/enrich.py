import json
import math

def get_level(value):
    if value <= 10:
        return "Chill"
    elif value <= 25:
        return "Light"
    elif value <= 40:
        return "Normal"
    elif value <= 55:
        return "Medium"
    elif value <= 70:
        return "Heavy"
    elif value <= 90:
        return "Intense"
    else:
        return "Crazy"


def system_enrich(event):
    payload = event.get("payload", {})
    for key, value in payload.items():

        if key == "cpu":
            cpu = math.floor(value[0])
            event["cpu_percent"] = cpu
            event["overall_cpu"] = get_level(cpu)

        elif key == "ram":
            ram = math.floor(value)
            event["ram_percent"] = ram
            event["overall_ram"] = get_level(ram)

        elif key == "disk":
            disk = math.floor(value)
            event["disk_percent"] = disk
            event["overall_disk"] = get_level(disk)

        elif key == "network":
            network = math.floor(value)
            event["network_percent"] = network
            event["overall_network"] = get_level(network)

    return event 
