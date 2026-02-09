import json
from core.interfaces.input import Input
from core.interfaces.normalizer.network import netwrok_normalizer
from core.interfaces.output import send_output
from core.interfaces.output import send_output
from plugins.integrity.enrich import intgerity_enrich
from plugins.integrity.rule import set_integrity_rule
from plugins.network.rule import set_network_rule
from plugins.processes.enrich import proccess_enrich 
from plugins.network.enrich import network_enrich 
from plugins.processes.rule import set_proccess_rule
from plugins.security.enrich import security_enrich 
from plugins.security.rule import set_security_rule
from plugins.system.enrich import system_enrich 
from plugins.system.rule import set_system_rule 



def piplinejob(payload):
    data = Input(payload)
    network_enriched = network_enrich(data["network"])
    processes_enriched = proccess_enrich(data["processes"])
    system_enriched    = system_enrich(data["system"])
    integrity_enriched = intgerity_enrich(data["integrity"])
    security_enriched  = security_enrich(data["security"])

    list_of_enriches = {
        "network":   network_enriched,
        "processes": processes_enriched,
        "system":    system_enriched,
        "integrity": integrity_enriched,
        "security":  security_enriched,
    }

    
    for key, enriched_value in list_of_enriches.items():
        if key == "network":
            updated_value = set_network_rule(enriched_value)
            list_of_enriches["network"] = updated_value   
        elif key == "system":
            updated_value = set_system_rule(enriched_value)
            list_of_enriches["system"] = updated_value
        elif key == "processes":
            updated_value = set_proccess_rule(enriched_value)
            list_of_enriches["processes"] = updated_value
        elif key == "security":
            updated_value = set_security_rule(enriched_value)
            list_of_enriches["security"] = updated_value
        elif key == "integrity":
            updated_value = set_integrity_rule(enriched_value)
            list_of_enriches["integrity"] = updated_value
        else:
            print(f"Unknown category: {key}") 

        send_output(list_of_enriches)     




