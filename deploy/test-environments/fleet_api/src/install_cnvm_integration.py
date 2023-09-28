#!/usr/bin/env python
"""
This script installs CNVM AWS integration

The following steps are performed:
1. Create an agent policy.
2. Create a CNVM AWS integration.
3. Create a deploy/cloudformation/config.json file to be used by the just deploy-cloudformation command.
"""
import sys
import json
from pathlib import Path
from munch import Munch
import configuration_fleet as cnfg
from api.agent_policy_api import create_agent_policy
from api.package_policy_api import create_cnvm_integration
from api.common_api import (
    get_enrollment_token,
    get_fleet_server_host,
    get_artifact_server,
    get_package_version,
)
from loguru import logger
from state_file_manager import state_manager, PolicyState
from package_policy import (
    version_compatible,
    VERSION_MAP,
    load_data,
    generate_random_name,
)

CNVM_EXPECTED_AGENTS = 1
CNVM_CLOUDFORMATION_CONFIG = "../../../cloudformation/config.json"
CNVM_AGENT_TAGS = ["cft_version:CFT_VERSION", "cft_arn:arn:aws:cloudformation:.*"]
PKG_DEFAULT_VERSION = VERSION_MAP.get("vuln_mgmt_aws", "")
INTEGRATION_NAME = "CNVM AWS"
INTEGRATION_INPUT = {
    "name": generate_random_name("pkg-cnvm-aws"),
    "input_name": "vuln_mgmt_aws",
    "posture": "vuln_mgmt",
    "deployment": "aws",
}
AGENT_INPUT = {
    "name": generate_random_name("cnvm-aws"),
}
cnvm_cloudformation_config = Path(__file__).parent / CNVM_CLOUDFORMATION_CONFIG


if __name__ == "__main__":
    # pylint: disable=duplicate-code
    package_version = get_package_version(cfg=cnfg.elk_config)
    logger.info(f"Package version: {package_version}")
    if not version_compatible(
        current_version=package_version,
        required_version=PKG_DEFAULT_VERSION,
    ):
        logger.warning(f"{INTEGRATION_NAME} is not supported in version {package_version}")
        sys.exit(0)
    logger.info(f"Starting installation of {INTEGRATION_NAME} integration.")
    agent_data, package_data = load_data(
        cfg=cnfg.elk_config,
        agent_input=AGENT_INPUT,
        package_input=INTEGRATION_INPUT,
    )

    logger.info("Create agent policy")
    agent_policy_id = create_agent_policy(cfg=cnfg.elk_config, json_policy=agent_data)

    logger.info(f"Create {INTEGRATION_NAME} integration for policy {agent_policy_id}")
    package_policy_id = create_cnvm_integration(
        cfg=cnfg.elk_config,
        pkg_policy=package_data,
        agent_policy_id=agent_policy_id,
    )

    state_manager.add_policy(
        PolicyState(
            agent_policy_id,
            package_policy_id,
            CNVM_EXPECTED_AGENTS,
            CNVM_AGENT_TAGS,
        ),
    )

    cloudformation_params = Munch()
    cloudformation_params.ENROLLMENT_TOKEN = get_enrollment_token(
        cfg=cnfg.elk_config,
        policy_id=agent_policy_id,
    )

    cloudformation_params.FLEET_URL = get_fleet_server_host(cfg=cnfg.elk_config)
    cloudformation_params.ELASTIC_AGENT_VERSION = cnfg.elk_config.stack_version
    cloudformation_params.ELASTIC_ARTIFACT_SERVER = get_artifact_server(cnfg.elk_config.stack_version)

    with open(cnvm_cloudformation_config, "w") as file:
        json.dump(cloudformation_params, file)

    logger.info(f"Installation of {INTEGRATION_NAME} integration is done")
