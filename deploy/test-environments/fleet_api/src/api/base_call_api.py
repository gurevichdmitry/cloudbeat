"""
This module provides functionality for performing API calls and handling API call exceptions.
It utilizes the 'requests' library for making HTTP requests.

Module contents:
    - APICallException: Exception class raised for API call failures.
    - perform_api_call: Function for making API calls and handling the response.

Dependencies:
    - requests: Library for making HTTP requests.

Note: Make sure to install the 'requests' library before using this module.
You can install it by running 'pip install requests' in your Python environment.
"""
import requests

class APICallException(Exception):
    """
    Exception raised for API call failures.

    Attributes:
        status_code (int): The HTTP status code of the failed API call.
        response_text (str): The response text of the failed API call.
    """

    def __init__(self, status_code, response_text):
        """
        Initialize the APICallException.

        Args:
            status_code (int): The HTTP status code of the failed API call.
            response_text (str): The response text of the failed API call.
        """
        self.status_code = status_code
        self.response_text = response_text


def perform_api_call(method, url, headers, auth, params):
    """
    Perform an API call using the provided parameters.

    Args:
        method (str): The HTTP method for the API call (e.g., 'GET', 'POST', 'PUT', 'DELETE').
        url (str): The URL of the API endpoint.
        headers (dict): The headers to be included in the API request.
        auth (tuple or None): The authentication tuple (username, password)
                        for basic authentication. Set to None for no authentication.
        params (dict): The parameters to be included in the API request.

    Returns:
        dict: The JSON response from the API call.

    Raises:
        APICallException: If the API call returns a non-200 status code.
    """
    response = requests.request(method=method, url=url, headers=headers, auth=auth, **params)
    if response.status_code != 200:
        raise APICallException(response.status_code, response.text)

    return response.json()
