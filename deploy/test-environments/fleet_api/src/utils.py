"""
This module provides utility functions for working with JSON files and performing file operations.

Functions:
- read_json(json_path: Path) -> dict: Read JSON data from a file.
- save_state(file_path: Path, data: list) -> None: Save data to a JSON file.
- delete_file(file_path: Path): Delete a file.

Module Dependencies:
- json: Standard library module for working with JSON data.
- pathlib.Path: Class representing file system paths.
- loguru.logger: Logger object for logging messages.

Usage:
Import this module to utilize the provided functions for JSON file operations and file deletion.
"""
import json
from typing import Union
from pathlib import Path
from loguru import logger


def read_json(json_path: Path) -> dict:
    """
    Read JSON data from a file.

    Args:
        json_path (Path): Path to the JSON file.

    Returns:
        dict: Dictionary containing the JSON data.

    Raises:
        FileNotFoundError: If the specified JSON file does not exist.
    """

    try:
        with json_path.open("r") as json_file:
            return json.load(json_file)
    except FileNotFoundError:
        logger.error(f"{json_path.name} file not found.")
        return {}


def save_state(file_path: Path, data: list) -> None:
    """
    Save data to a JSON file.

    If the file already exists, the new data is appended to the existing data.
    If the file does not exist, a new file is created with the provided data.

    Args:
        file_path (Path): Path to the JSON file.
        data (list): List of data to be saved.

    Raises:
        Exception: If an error occurs while saving the JSON data.
    """
    try:
        if file_path.exists():
            with file_path.open("r") as exist_file:
                policies_data = json.load(exist_file)
            policies_data["policies"].extend(data)
        else:
            policies_data = {"policies": data}
        with file_path.open("w") as policies_file:
            json.dump(policies_data, policies_file)
        logger.info(f"JSON data saved to {file_path}")
    except FileNotFoundError as ex:
        logger.error(f"Error occurred while saving JSON data: File '{file_path}' not found.")
        raise ex
    except IOError as ex:
        logger.error(f"Error occurred while saving JSON data: {ex}")
        raise ex


def delete_file(file_path: Path):
    """
    Delete a file.

    Args:
        file_path (Path): Path to the file to be deleted.

    Raises:
        FileNotFoundError: If the specified file does not exist.
        Exception: If an error occurs while deleting the file.
    """
    try:
        file_path.unlink()
        logger.info(f"File '{file_path}' deleted successfully.")
    except FileNotFoundError:
        logger.warning(f"File '{file_path}' does not exist.")
    except OSError as ex:
        logger.error(f"Error occurred while deleting file '{file_path}': {ex}")
        raise ex


def update_key(data: Union[dict, list], search_key: str, value_to_apply: str):
    """Update the value of a specific key in the given data.

    If the data is a dictionary, it searches for the specified key and updates its value.
    If the data is a list, it recursively calls the function for each item in the list.

    Args:
        data (Union[dict,list]): The data to be updated, which can be either a dictionary or a list.
        search_key (str): The key to search for in the data.
        value_to_apply (str): The value to be applied to the matching key.

    Returns:
        None
    """
    if isinstance(data, list):
        for item in data:
            update_key(item, search_key, value_to_apply)
    elif isinstance(data, dict):
        for key, value in data.items():
            if key == search_key:
                data[key]["value"] = value_to_apply
            elif isinstance(value, (dict, list)):
                update_key(value, search_key, value_to_apply)


def delete_key(data: Union[dict, list], search_key: str, key_to_delete: str):
    """Delete a specific key from the data if it matches the search key.

    If the data is a dictionary,
    it searches for the specified key and deletes the corresponding key_to_delete.

    If the data is a list, it recursively calls the function for each item in the list.

    Args:
        data (Union[dict, list]): The data from which the key should be deleted,
                                    which can be either a dictionary or a list.
        search_key (str): The key to search for in the data.
        key_to_delete (str): The key to delete if it matches the search key.

    Returns:
        None
    """
    if isinstance(data, list):
        for item in data:
            delete_key(item, search_key, key_to_delete)
    elif isinstance(data, dict):
        for key, value in data.items():
            if key == search_key and isinstance(value, dict) and "value" in value:
                del value[key_to_delete]
            elif isinstance(value, (dict, list)):
                delete_key(value, search_key, key_to_delete)
