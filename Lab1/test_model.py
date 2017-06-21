import model
import unittest


class TestModel(unittest.TestCase):
    def setUp(self):
        self.model = model.Model("none")
        self.model.records.append(model.Model.Record("12-01-2017", 125.50, 3.14198))

    def test_init(self):
        records = [model.Model.Record("12-01-2017", 125.50, 3.14198)]
        self.assertEqual(self.model.records, records)

    def test_check_validity_of_date(self):
        self.assertTrue(self.model.check_validity_of_date(self.model.records[0].date))
        self.assertFalse(self.model.check_validity_of_date("wrong date"))

    def test_check_validity(self):
        self.assertTrue(self.model.check_validity(
            [self.model.records[0].date, self.model.records[0].length, self.model.records[0].coefficient]))
        self.assertFalse(self.model.check_validity(["10-29-2015", "12", 10]))

    def test_compare_date(self):
        self.assertEqual(self.model.compare_date(self.model.records[0].date, "10-09-2015"), 2)
        self.assertEqual(self.model.compare_date(self.model.records[0].date, "12-01-2017"), 1)
        self.assertEqual(self.model.compare_date(self.model.records[0].date, "10-09-2040"), 0)

    def test_find_by_date(self):
        self.model.records.append(model.Model.Record("12-01-2017", 125, 3.14))
        self.model.records.append(model.Model.Record("12-03-2017", 250, 14.20))
        self.model.records.append(model.Model.Record("12-03-2017", 456, 55.11))
        result = self.model.find_by_date("12-03-2017")
        self.assertEqual(result[0].length, 250)

    def test_find_by_date_range(self):
        elements = [model.Model.Record("12-01-2017", 125, 3.14), model.Model.Record("12-02-2017", 250, 14.20), \
                    model.Model.Record("12-03-2017", 456, 55.11), model.Model.Record("12-04-2017", 887, 15), \
                    model.Model.Record("12-05-2017", 337, 1.08), model.Model.Record("12-06-2017", 225, 0.75)]
        self.assertEqual(self.model.find_by_date_range(elements, "01-03-2017", "22-05-2017")[0].length, 456)

    def test_get_used_fuel(self):
        self.assertEqual(self.model.get_used_fuel(model.Model.Record("12-03-2017", 200, 10)), 20.0)

    def test_get_general_length(self):
        self.assertEqual(self.model.get_general_length([model.Model.Record("12-03-2017", 300, 10), \
                                                        model.Model.Record("12-03-2017", 100, 20)]), 400)

    def test_get_general_fuel_used(self):
        self.assertEqual(self.model.get_general_fuel_used([model.Model.Record("12-03-2017", 300, 10), \
                                                           model.Model.Record("12-03-2017", 100, 20)]), 50.0)


if __name__ == '__main__':
    unittest.main()
