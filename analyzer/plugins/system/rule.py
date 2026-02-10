def set_system_rule(event):
        system_rules = []
        cpu_status = event.get("overall_cpu", "")
        if cpu_status in ["Chill", "Light"]:
            system_rules.append("CPU load is low and running normally.")
        elif cpu_status in ["Normal", "Medium"]:
            system_rules.append("CPU load is moderate.")
        elif cpu_status in ["Heavy", "Intense", "Crazy"]:
            system_rules.append("CPU load is high! Potential performance risk.")
        else:
            system_rules.append("CPU load status unknown.")
        # ---- RAM ----
        ram_status = event.get("overall_ram", "")
        if ram_status in ["Chill", "Light"]:
            system_rules.append("RAM usage is low and normal.")
        elif ram_status in ["Normal", "Medium"]:
            system_rules.append("RAM usage is moderate.")
        elif ram_status in ["Heavy", "Intense", "Crazy"]:
            system_rules.append("RAM usage is high! Potential memory risk.")
        else:
            system_rules.append("RAM usage status unknown.")

        # ---- Disk ----
        disk_status = event.get("overall_disk", "")
        if disk_status in ["Chill", "Light"]:
            system_rules.append("Disk usage is low and healthy.")
        elif disk_status in ["Normal", "Medium"]:
            system_rules.append("Disk usage is moderate.")
        elif disk_status in ["Heavy", "Intense", "Crazy"]:
            system_rules.append("Disk usage is high! Potential storage risk.")
        else:
            system_rules.append("Disk usage status unknown.")

        # ---- Network ----
        network_status = event.get("overall_network", "")
        if network_status in ["Chill", "Light"]:
            system_rules.append("Network usage is low and stable.")
        elif network_status in ["Normal", "Medium"]:
            system_rules.append("Network usage is moderate.")
        elif network_status in ["Heavy", "Intense", "Crazy"]:
            system_rules.append("Network usage is very high! Potential network congestion.")
        else:
            system_rules.append("Network usage status unknown.")

        event["system_rules"] = system_rules

