import unittest
from io import StringIO, BytesIO
import serialization_pickle
import serialization_yaml
import serialization_json
import model


class PickleTest(unittest.TestCase):
    def setUp(self):
        self.data = model.Model.Record("12-01-2017", 125.50, 3.14198)

    def test_pickle_save(self):
        outfile = BytesIO()
        serialization_pickle.PickleSerializer().serialize(self.data, outfile)
        outfile.seek(0)
        self.assertEqual(
            serialization_pickle.PickleSerializer().deserialize(outfile),
            self.data)

    def test_pickle_read(self):
        outfile = BytesIO()
        serialization_pickle.PickleSerializer().serialize(self.data, outfile)
        outfile.seek(0)
        content = serialization_pickle.PickleSerializer().deserialize(outfile)
        outfile.close()
        self.assertEqual(content, self.data)


class YamlTest(unittest.TestCase):
    def setUp(self):
        self.data = [{"date": "12-01-2017", "length": 125.5, "coefficient": 3.14198}]

    def test_yaml_save(self):
        outfile = StringIO()
        serialization_yaml.YamlSerializer().serialize(self.data, outfile)
        outfile.seek(0)
        self.assertEqual(serialization_yaml.YamlSerializer().deserialize(outfile), self.data)

    def test_yaml_read(self):
        outfile = StringIO()
        serialization_yaml.YamlSerializer().serialize(self.data, outfile)
        outfile.seek(0)
        content = serialization_yaml.YamlSerializer().deserialize(outfile)
        self.assertEqual(content, self.data)


class JsonTest(unittest.TestCase):
    def setUp(self):
        self.data = model.Model.Record("12-01-2017", 125.50, 3.14198)
        self.expected = {"date": "12-01-2017", "length": 125.5, "coefficient": 3.14198}

    def test_json_save(self):
        outfile = StringIO()
        serialization_json.JsonSerializer().serialize(self.data, outfile)
        outfile.seek(0)
        resdict = serialization_json.JsonSerializer().deserialize(outfile)
        result=model.Model.Record(resdict["date"],resdict["length"],resdict["coefficient"])
        self.assertEqual(result, self.data)

    def test_json_read(self):
        outfile = StringIO()
        serialization_json.JsonSerializer().serialize(self.data, outfile)
        outfile.seek(0)
        content = serialization_json.JsonSerializer().deserialize(outfile)
        outfile.close()
        self.assertEqual(content, self.expected)


if __name__ == '__main__':
    unittest.main()
