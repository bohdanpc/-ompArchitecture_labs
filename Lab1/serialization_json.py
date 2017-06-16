import json


class JsonSerializer:
    def serialize(self, data, file):
        """
        Encode obj to yaml format and write it into file.
        """

        json.dump(data.__dict__, file)

    def deserialize(self, file):
        """
        Decode from yaml file to Python-object.
        """

        return json.load(file)
