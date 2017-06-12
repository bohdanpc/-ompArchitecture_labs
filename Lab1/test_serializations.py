import unittest
from io import StringIO, BytesIO
import serialization_pickle
import serialization_yaml
import serialization_json
import model


class PickleTest(unittest.TestCase):
    def setUp(self):
        self.data = model.Record("12-01-2017", 125.50, 3.14198)
        self.bytes = b'\x80\x04\x95Y\x00\x00\x00\x00\x00\x00\x00\x8c\x05model\x94\x8c\x06Record\x94\x93\x94)\x81\x94}\x94(\x8c\x04date\x94\x8c\n12-01-2017\x94\x8c\x06length\x94G@_`\x00\x00\x00\x00\x00\x8c\x0bcoefficient\x94G@\t"\xc6i\x05}\x18ub.'

    def test_pickle_save(self):
        outfile = BytesIO()
        serialization_pickle.PickleSerializer().serialize(self.data, outfile)
        content = outfile.getvalue()
        outfile.close()
        self.assertEqual(content, self.bytes)

    def test_pickle_read(self):
        outfile = BytesIO()
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
        print(content)
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
        # self.expected = '"{\\"date\\": \\"12-01-2017\\", \\"length\\": 125.5, \\"coefficient\\": 3.14198}"'
        self.expected = {"date": "12-01-2017", "length": 125.5, "coefficient": 3.14198}

    def test_json_save(self):
        outfile = StringIO()
        serialization_json.JsonSerializer().serialize(self.data, outfile)
        content = outfile.getvalue()
        outfile.close()
        self.assertEqual(content, '{"date": "12-01-2017", "length": 125.5, "coefficient": 3.14198}')

    def test_json_read(self):
        outfile = StringIO()
        serialization_json.JsonSerializer().serialize(self.data, outfile)
        outfile.seek(0)
        content = serialization_json.JsonSerializer().deserialize(outfile)
        outfile.close()
        self.assertEqual(content, self.expected)


if __name__ == '__main__':
    unittest.main()
