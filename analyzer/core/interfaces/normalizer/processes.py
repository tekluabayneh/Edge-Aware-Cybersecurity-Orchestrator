from fastapi.encoders import jsonable_encoder
from shared.schemas.schema import Event


def processes_normalizer(processesPaylod): 
     processes_normalizer_result = { 
                                    "payload":processesPaylod,
                                    "type":"processes",
                                    "category":"processes", 
                                    "source":"agent",
                                    "tags":["processes"]
                                    }

     return processes_normalizer_result
  
