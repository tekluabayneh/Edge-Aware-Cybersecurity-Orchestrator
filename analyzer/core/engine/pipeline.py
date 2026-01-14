from core.interfaces.input import Input
from core.interfaces.output import send_output
from core.interfaces.rule import apply_rule
from plugins.integrity.enrich import intgerity_enrich
from plugins.network.enrich import network_enrich 
from plugins.security.enrich import security_enrich 
from plugins.system.enrich import system_enrich 
from fastapi.encoders import jsonable_encoder
import json



def piplinejob(paload): 
     inputReturnValue = Input(paload) 
     json_paylod = json.dumps(jsonable_encoder(inputReturnValue))
     # netwrok_enriched_event = network_enrich(json_paylod) 
     print(json_paylod)
     # print(json.dumps(jsonable_encoder(inputReturnValue), indent=2))

     #
     # system_enriched_event = system_enrich(json_paylod[1]) 
     # integrity_enriched_event = intgerity_enrich(json_paylod[2]) 
     # security_enriched_event = security_enrich(json_paylod[3]) 
     # proccess_enriched_event = proccess_enrich(json_paylod[4]) 
     #
     # list_of_enriches = [netwrok_enriched_event, system_enriched_event, integrity_enriched_event, security_enriched_event, proccess_enriched_event]
     #
     # for en in list_of_enriches: 
     #    print(en)
     #    # ruled_event = apply_rule(en) 
 

     # send_output(ruled_event) 

     # update_state(ruled_event) 



     #
     # print(json.dumps(jsonable_encoder(inputReturnValue[0]), indent=2))
     # print(json.dumps(jsonable_encoder(inputReturnValue[1]), indent=2))
     # print(json.dumps(jsonable_encoder(inputReturnValue[2]), indent=2))
     # print(json.dumps(jsonable_encoder(inputReturnValue[3]), indent=2))
     # print(json.dumps(jsonable_encoder(inputReturnValue[4]), indent=2))
     #


