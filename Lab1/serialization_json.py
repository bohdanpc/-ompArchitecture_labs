import json
import model


class JsonSerializer:
    def serialize(self, data, file):
        """
        Encode obj to yaml format and write it into file.
        """

        json.dump([record.__dict__ for record in data], file)

    def deserialize(self, file):
        """
        Decode from yaml file to Python-object.
        """

        data = json.load(file)
        result = []
        for elem in data:
            result.append(model.Record(elem["date"],
                                       elem["length"],
                                       elem["coefficient"]))
        return result
