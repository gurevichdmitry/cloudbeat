"""
This class provides wrapper for Fleet Integrations (packages) api
"""
import json

from .api_client import ApiClient


class Integrations(ApiClient):
    """
    This class provides Integrations (EPM packages) functionality
    """

    def __init__(self, host, username, password):
        super().__init__(host=host,
                                           username=username,
                                           password=password)
        self.url_epm_packages = self.get_fleet_url() + '/epm/packages'
        self.url_epm_packages_name_version = self.get_fleet_url() + "/epm/packages/{}/{}"

    def list_integrations(self, is_experimental: bool = True) -> list:
        """
        This function retrieves list of all kibana integrations (packages)
        @param is_experimental: Boolean if experimental list is included or not
        @return: Integrations list
        """
        url = self.url_epm_packages
        params = {
            'experimental': is_experimental
        }

        response = self.send_request(method='GET',
                                     url=url,
                                     params=params)
        result = ApiClient.handle_response(response=response)
        return result

    def get_integration_context(self, integration_name: str) -> dict:
        """
        This function retrieves integration context by provided name
        @param integration_name: Name of integration the context shall be retrieved from
        @return: Integration context dictionary
        """
        integrations = self.list_integrations()
        for integration in integrations:
            name = integration.get('name', '')
            if name == integration_name:
                return integration
        return {}

    def get_version_by_name(self, integration_name: str):
        """
        This function retrieves Integration's version
        @param integration_name: Integration name
        @return: Version field value
        """
        integration = self.get_integration_context(integration_name=integration_name)
        return integration.get('version', '')

    def is_integration_installed(self, integration_name: str):
        """
        This function retrieves Integration's status
        @param integration_name: Integration's name
        @return: Boolean (True if 'installed', otherwise - False)
        """
        integration = self.get_integration_context(integration_name=integration_name)
        status = integration.get('status', '')
        if status == 'installed':
            return True
        return False

    def install_integration_assets(self, integration_name: str) -> list:
        """
        This function install Integration's assets
        @param integration_name: Integration name
        @return: List of installed assets
        """

        status = self.is_integration_installed(integration_name=integration_name)
        if status:
            return []

        version = self.get_version_by_name(integration_name=integration_name)
        url = self.url_epm_packages_name_version.format(integration_name, version)

        body = {
            'force': False,
            'ignore_constraints': True
        }

        response = self.send_request('POST',
                                     url=url,
                                     data=json.dumps(body))
        result = ApiClient.handle_response(response=response)
        return result

    def uninstall_integration_assets(self, integration_name: str) -> list:
        """
        This function uninstalls integration's assets
        @param integration_name: Integration name the assets should be uninstalled from
        @return: List of uninstalled assets
        """
        status = self.is_integration_installed(integration_name=integration_name)
        if not status:
            return []

        version = self.get_version_by_name(integration_name=integration_name)

        url = self.url_epm_packages_name_version.format(integration_name, version)
        body = {
            'force': True
        }
        response = self.send_request('DELETE',
                                     url=url,
                                     data=json.dumps(body))
        result = ApiClient.handle_response(response=response)
        return result
