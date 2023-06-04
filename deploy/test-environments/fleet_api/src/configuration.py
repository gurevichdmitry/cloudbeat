"""
This module provides configuration settings and paths
for the ELK (Elasticsearch, Logstash, Kibana) integration.

Module contents:
    - elk_config: Munch object containing the ELK configuration settings.
    - state_data_file: Path object representing the path to the state data file.

Dependencies:
    - os: Module for accessing environment variables.
    - pathlib: Module for working with file paths.
    - munch: Module for creating convenient data containers.

Note: This module assumes that environment variables for 
the ELK configuration (ES_USER, ES_PASSWORD, KIBANA_URL)
have been set in the system environment.
"""
import os
from pathlib import Path
from munch import Munch

elk_config = Munch()
elk_config.user = os.getenv("ES_USER", "NA")
elk_config.password = os.getenv("ES_PASSWORD", "NA")
elk_config.kibana_url = os.getenv("KIBANA_URL", "NA")
elk_config.auth = (elk_config.user, elk_config.password)

aws_config = Munch()
aws_config.access_key_id = os.getenv("AWS_ACCESS_KEY_ID", "NA")
aws_config.secret_access_key = os.getenv("AWS_SECRET_ACCESS_KEY", "NA")

state_data_file = Path(__file__).parent / "state_data.json"
