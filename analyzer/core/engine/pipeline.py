from core.interfaces.input import Input
from fastapi.encoders import jsonable_encoder
import json

def piplinejob(paload): 
     inputReturnValue = Input(paload) 
     print(json.dumps(jsonable_encoder(inputReturnValue[0]), indent=2))
     print(json.dumps(jsonable_encoder(inputReturnValue[1]), indent=2))
     print(json.dumps(jsonable_encoder(inputReturnValue[2]), indent=2))
     print(json.dumps(jsonable_encoder(inputReturnValue[3]), indent=2))
     print(json.dumps(jsonable_encoder(inputReturnValue[4]), indent=2))
        
