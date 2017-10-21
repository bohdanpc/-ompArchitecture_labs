import yaml


class YamlSerializer:
    def serialize(self, data, file):
        """
        Encode obj to yaml format and write it into file.
        """

        yaml.dump(data, file)

    def deserialize(self, file):
        """
        Decode from yaml file to Python-object.
        """

        res = yaml.load(file)
        return res
