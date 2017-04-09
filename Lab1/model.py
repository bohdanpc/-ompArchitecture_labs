import _pickle


class Record(object):
    def __init__(self, _date, _length, _coefficient):
        self.date = _date
        self.length = _length
        self.coefficient = _coefficient

    def get(self):
        return [self.date, self.length, self.coefficient]


def initialise(file_name):
    try:
        with open(file_name, 'rb') as f:
            records = _pickle.load(f)
            f.close()
        return records
    except Exception as error:
        return []


def save_all(records, file_name):
    with open(file_name, 'wb') as f:
        _pickle.dump(records, f, _pickle.HIGHEST_PROTOCOL)
    f.close()


def find_by_date(records, date):
    for item in records:
        if item.date == date:
            return item
    return False


def get_used_fuel(record):
    return (record.coefficient * record.length) / 100


def get_general_length(records):
    res = 0
    for item in records:
        res += item.length
    return res


def get_general_fuel_used(records):
    res = 0
    for item in records:
        res += get_used_fuel(item)
