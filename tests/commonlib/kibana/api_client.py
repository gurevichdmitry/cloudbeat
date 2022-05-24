"""
Base Fleet http client for all requests
"""
import json
from uuid import uuid4
import requests
from requests.models import Response
from requests.exceptions import HTTPError
from requests.auth import HTTPBasicAuth

METHODS = ['GET', 'POST', 'DELETE', 'UPDATE']


class ApiClient:
    """
    Base API client for Fleet http requests
    """

    def __init__(self, host: str, username: str, password: str):
        self.host = host
        self.username = username
        self.password = password
        self.fleet_base = "/api/fleet"
        self.kbn_xsrf = str(uuid4())

    def send_request(self,
                     method: str,
                     url: str,
                     data=None,
                     params=None,
                     headers=None) -> Response:
        """
        Wrapper method for http request, that handles predefine params.
        @param method: REST API method string
        @param url: Full endpoint url
        @param data: Data to be passed to http request
        @param params: Params to be passed to http request
        @param headers: Custom headers to be passed, by default base headers are provided
        @return: Response object
        """

        if method.upper() not in METHODS:
            raise Exception(f"The method '{method}' is not supported.")

        base_headers = {
            'content-type': 'application/json',
            'kbn-xsrf': self.kbn_xsrf
        }

        if headers is not None:
            base_headers.update(headers)

        if data is None:
            data = {}

        if params is None:
            params = {}

        kwargs = dict(data=json.dumps(data),
                      params=params,
                      headers=base_headers,
                      auth=self.set_basic_auth())

        return requests.request(method=method.upper(),
                                url=url,
                                **kwargs)

    @staticmethod
    def handle_response(response: Response):
        """
        Base wrapper for all responses:
        - handling status code
        - getting response content
        @param response: Response object
        @return: List result
        """

        try:
            response.raise_for_status()
            result = response.json().get('items', [])
            return result
        except (HTTPError, ConnectionError) as error:
            print(f"Status code is {response.status_code}, message: {error}")
            return []

    def set_basic_auth(self):
        """
        Initiates requests.HTTPBasicAuth authentication
        @return: Basic Authenticator
        """
        return HTTPBasicAuth(self.username, self.password)

    def get_fleet_url(self):
        """
        Creates Fleet base url
        @return: Base URL to Fleet
        """
        return self.host + self.fleet_base
