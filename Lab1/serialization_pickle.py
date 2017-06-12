import pickle


class PickleSerializer:
    def serialize(self, data, file):
        """
        Encode obj to binary row using pickle write it into file.
        """
        pickle.dump(data, file, pickle.HIGHEST_PROTOCOL)

    def deserialize(self, file):
        """
        Decode from binary row to Python-object using pickle.
        """

        return pickle.load(file)
