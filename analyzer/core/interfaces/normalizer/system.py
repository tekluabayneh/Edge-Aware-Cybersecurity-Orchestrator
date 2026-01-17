from fastapi.encoders import jsonable_encoder
from shared.schemas.schema import Event


def system_normalizer(systemPaylod):
     system_normalizer_result = { 
                                 "payload":systemPaylod,
                                 "type":"system",
                                 "category":"system", 
                                 "source":"agent",
                                 "tags":["system"]
                                 } 
  
     return system_normalizer_result
