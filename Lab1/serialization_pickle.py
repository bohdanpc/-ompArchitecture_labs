import pickle


class PickleSerializer:
    def serialize(self, data, file):
        """
        Encode obj to binary row using pickle write it into file.
        """
        pickle.dump(data, file)

    def deserialize(self, file):
        """
        Decode from binary row to Python-object using pickle.
        """
        data = pickle.load(file)
        return data
