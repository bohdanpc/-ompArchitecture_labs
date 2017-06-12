import unittest
from io import StringIO
import serialization_pickle
import serialization_yaml
import serialization_json
import model


class PickleTest(unittest.TestCase):
    def setUp(self):
        self.data = model.Record("12-01-2017", 125.50, 3.14198)

    def test_pickle_save(self):
        outfile = StringIO()
        serialization_pickle.PickleSerializer().serialize(self.data, outfile)
        content = outfile.getvalue()
        outfile.close()
        self.assertEqual(content, self.data)

    def test_pickle_read(self):
        outfile = StringIO()
        serialization_pickle.PickleSerializer().serialize(self.data, outfile)
        outfile.seek(0)
        content = serialization_pickle.PickleSerializer().deserialize(outfile)
        outfile.close()
        self.assertEqual(content, self.data)


class YamlTest(unittest.TestCase):
    def setUp(self):
        self.data = model.Record("12-01-2017", 125.50, 3.14198)

    def test_yaml_save(self):
        outfile = StringIO()
        serialization_yaml.YAMLSerializer().serialize(self.data, outfile)
        outfile.seek(0)
        content = outfile.getvalue()
        outfile.close()
        self.assertEqual(content, self.data)

    def test_yaml_read(self):
        outfile = StringIO()
        serialization_yaml.YAMLSerializer().serialize(self.data, outfile)
        outfile.seek(0)
        content = serialization_yaml.YAMLSerializer().deserialize(outfile)
        self.assertEqual(content, self.data)


class JsonTest(unittest.TestCase):
    def setUp(self):
        self.data = model.Record("12-01-2017", 125.50, 3.14198)

    def test_json_save(self):
        outfile = StringIO()
        serialization_json.JsonSerializer().serialize(self.data, outfile)
        content = outfile.getvalue()
        outfile.close()
        self.assertEqual(content, [{"coefficient": 3.14198, "date": "12-01-2017", "length": 125.50}])

    def test_json_read(self):
        outfile = StringIO()
        serialization_json.JsonSerializer().serialize(self.data, outfile)
        outfile.seek(0)
        content = serialization_json.JsonSerializer().deserialize(outfile)
        outfile.close()
        self.assertEqual(content, self.data)


if __name__ == '__main__':
    unittest.main()
