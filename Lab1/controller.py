import view
from model import *


def show_all(records):
    """
    >>> records = [Record("12-12-2012", 253, 172)]
    >>> show_all(records)
    <BLANKLINE>
          Date|Length(km)|Consumption(100km)| Fuel used|
    12-12-2012|    253.00|            172.00|    435.16|
    """
    view.record_names()
    for item in records:
        view.print_record(item, get_used_fuel(item))


def show_summary(records):
    """
    >>> show_summary([Record("12-12-2012", 253, 172), Record("20-11-2015", 128, 123)])
    <BLANKLINE>
    General length: 381
    Fuel used: 592.6
    >>> show_summary([])
    <BLANKLINE>
    General length: 0
    Fuel used: 0
    """
    view.print_summary(get_general_length(records), get_general_fuel_used(records))


def show_summary_period(records, input_func1=input, input_func2=input):
    """
    >>> show_summary_period([Record("12-12-2012", 253, 172), Record("20-11-2015", 128, 123)], \
    lambda:"12-10-2012", lambda:"25-01-2016")
    <BLANKLINE>
    Enter beginning date(dd-mm-yyyy):Enter ending date(dd-mm-yyyy):
    General length: 381
    Fuel used: 592.6

    >>> show_summary_period([Record("12-12-2012", 253, 172), Record("20-11-2015", 128, 123)], \
    lambda:"12-10-2009", lambda:"25-01-2013")
    <BLANKLINE>
    Enter beginning date(dd-mm-yyyy):Enter ending date(dd-mm-yyyy):
    General length: 253
    Fuel used: 435.16
    """
    left, right = view.enter_period(input_func1, input_func2)
    if check_validity_of_date(left) and check_validity_of_date(right):
        records_period = find_by_date_range(records, left, right)
        show_summary(records_period)
    else:
        view.invalid_value()


def show_by_date(records, input_func=input):
    """
    >>> show_by_date([Record("12-12-2012", 253, 172), Record("20-11-2015", 128, 123)], lambda:"12-12-2012")
    <BLANKLINE>
    Enter the date(dd-mm-yyyy):
          Date|Length(km)|Consumption(100km)| Fuel used|
    12-12-2012|    253.00|            172.00|    435.16|
    >>> show_by_date([Record("12-12-2012", 253, 172), Record("20-11-2015", 128, 123)], lambda:"32-12-2017")
    <BLANKLINE>
    Enter the date(dd-mm-yyyy):Invalid values entered
    """
    date = view.enter_date(input_func)

    if check_validity_of_date(date):
        records_daily = find_by_date(records, date)
        show_all(records_daily)
    else:
        view.invalid_value()


def show_by_period(records, input_func1=input, input_func2=input):
    """
    >>> show_by_period([Record("12-12-2012", 253, 172), Record("20-11-2015", 128, 123)], \
    lambda:"12-10-2012", lambda:"25-01-2016")
    <BLANKLINE>
    Enter beginning date(dd-mm-yyyy):Enter ending date(dd-mm-yyyy):
          Date|Length(km)|Consumption(100km)| Fuel used|
    12-12-2012|    253.00|            172.00|    435.16|
    20-11-2015|    128.00|            123.00|    157.44|

    >>> show_by_period([Record("12-12-2012", 253, 172), Record("20-11-2015", 128, 123)], \
    lambda:"12-10-2012", lambda:"25-01-2014")
    <BLANKLINE>
    Enter beginning date(dd-mm-yyyy):Enter ending date(dd-mm-yyyy):
          Date|Length(km)|Consumption(100km)| Fuel used|
    12-12-2012|    253.00|            172.00|    435.16|
    """

    left, right = view.enter_period(input_func1, input_func2)
    if check_validity_of_date(left) and check_validity_of_date(right):
        records_period = find_by_date_range(records, left, right)
        show_all(records_period)
    else:
        view.invalid_value()


def add_record(records):
    item_list = view.enter_trip_details()
    if check_validity_of_date(item_list[0]) and \
            check_validity_of_length(item_list[1]) and check_validity_of_coefficient(item_list[2]):
        record = Record(item_list[0], item_list[1], item_list[2])
        records.append(record)
    else:
        view.invalid_value()


def main_func():
    """ readiness level: 1, 4, 5, 7"""
    choice = ''
    records = initialise("fuel_consumption.pickle")
    while choice != "7":
        choice = str(view.menu())
        if choice == "1":
            show_all(records)
        elif choice == "2":
            show_summary(records)
        elif choice == "3":
            show_by_period(records)
        elif choice == "4":
            show_summary_period(records)
        elif choice == "5":
            show_by_date(records)
        elif choice == "6":
            add_record(records)
        elif choice == "7":
            save_all(records, "fuel_consumption.pickle")


#main_func()

#if __name__ == "__main__":
#    import doctest
#    doctest.testmod()
