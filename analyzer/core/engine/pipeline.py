from core.interfaces.normalizer.integrity import intgerity_normalizer
from core.interfaces.normalizer.network import netwrok_normalizer
from core.interfaces.normalizer.processes import processes_normalizer
from core.interfaces.normalizer.security import security_normalizer
from core.interfaces.normalizer.system import system_normalizer


def piplinejob(): 
    networkValue  =  netwrok_normalizer
    securityValue =  security_normalizer
    systemValue =    system_normalizer
    integrityValue = intgerity_normalizer
    proccessValue =   processes_normalizer


         

   
