from core.interfaces.normalizer import integrity, network, processes, security, system


# send each raw data to their own normalizer
def Input(payload):
    networkValue = network.netwrok_normalizer(payload.network)
    processesValue=    processes.processes_normalizer(payload.processes)
    integrityValue =  integrity.intgerity_normalizer(payload.integrity)
    systemValue =system.system_normalizer(payload.system)
    securityValue = security.security_normalizer(payload.security)

    return ( 
            networkValue, 
            processesValue,
            systemValue, 
            integrityValue,
            securityValue
            )
   


