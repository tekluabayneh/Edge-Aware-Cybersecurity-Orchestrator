from core.interfaces.normalizer import integrity, network, processes, security, system


# send each raw data to their own normalizer
def Input(payload):
    print(payload.email) 
    print(payload.agent_id) 
    print(payload.agent_token)
    network.netwrok_normalizer(payload.network)
    processes.processes_normalizer(payload.processes)
    integrity.intgerity_normalizer(payload.integrity)
    system.system_normalizer(payload.system)
    security.security_normalizer(payload.security)
