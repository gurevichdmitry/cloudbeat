import pytest


@pytest.mark.cis_benchmark
def test_cis_benchmark_install_assets(fleet_client, cis_k8s_benchmark):

    status = fleet_client.is_integration_installed(cis_k8s_benchmark.name)
    if status:
        fleet_client.uninstall_integration_assets(integration_name=cis_k8s_benchmark.name)

    result = fleet_client.install_integration_assets(integration_name=cis_k8s_benchmark.name)

    status = fleet_client.is_integration_installed(cis_k8s_benchmark.name) #
    assert status, f"Install integration assets failed"


@pytest.mark.cis_benchmark
def test_cis_benchmark_delete_assets(fleet_client, cis_k8s_benchmark):

    result = fleet_client.is_integration_installed(integration_name=cis_k8s_benchmark.name)
    if result:
        fleet_client.uninstall_integration_assets(integration_name=cis_k8s_benchmark.name)

    is_installed = fleet_client.is_integration_installed(integration_name=cis_k8s_benchmark.name)
    assert not is_installed, f"Uninstalling integration assets failed."
