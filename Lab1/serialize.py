import json
import pickle
import configparser
import sys
import os
import yaml
import model


class Serialize:
    """
    Choose file to serialize
    """

    @staticmethod
    def save(filePath, table):
        """
        Choose file to save data
        """
        config = configparser.RawConfigParser()
        config.read('config.cfg')
        cls = config.get('SerializationSettings', 'serialize')
        _serializer = sys.modules["serialize"].__dict__.get(cls)
        return _serializer.save(filePath, table)

    @staticmethod
    def load(filePath):
        """
        Choose file to load data
        """
        config = configparser.RawConfigParser()
        config.read('config.cfg')
        cls = config.get('SerializationSettings', 'serialize')
        _serializer = sys.modules["serialize"].__dict__.get(cls)
        return _serializer.load(filePath)


class JsonSerialize:
    """
    Serialize database to json file
    """

    @staticmethod
    def load(fileName):
        """
        load data from file :filePath and return as list. If file doesn't
        exist, return empty list
        """
        if not os.path.isfile(fileName + ".json"):
            return []
        with open(fileName + ".json", "r") as f:
            data = json.load(f)
            res_db = []
            for elem in data:
                res_db.append(model.Record(elem["date"],
                                           elem["length"],
                                           elem["coefficient"]))
            return res_db

    @staticmethod
    def save(fileName, data):
        """
            serialize to file :filePath list :data
        """
        with open(fileName + ".json", "w") as f:
            json.dump([record.__dict__ for record in data], f)
            f.close()


class PickleSerialize:
    """
    Serialize database to pickle file
    """

    @staticmethod
    def save(fileName, data):
        """
        serialize to file :filePath list :data
        """
        with open(fileName + ".pickle", "wb") as f:
            pickle.dump(data, f)
            f.close()

    @staticmethod
    def load(fileName):
        """
        load data from file :filePath and return as list. If file doesn't
        exist, return empty list
        """
        if not os.path.isfile(fileName + ".pickle"):
            return []
        with open(fileName + ".pickle", "rb") as f:
            res = pickle.load(f)
            return res


class YamlSerialize:
    """
    Serialize database to yaml file
    """

    @staticmethod
    def save(fileName, data):
        """
        serialize to file :filePath list :data
        """
        with open(fileName + ".yaml", "w") as f:
            yaml.dump(data, f)

    @staticmethod
    def load(fileName):
        """
        load data from file :filePath and return as list. If file doesn't
        exist, return empty list
        """
        if not os.path.isfile(fileName + ".yaml"):
            return []
        with open(fileName + ".yaml", "r") as f:
            res = yaml.load(f)
            return res
