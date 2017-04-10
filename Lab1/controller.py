import view
from model import *


def show_all(records):
    view.record_names()
    for item in records:
        view.print_record(item, get_used_fuel(item))


def show_summary(records):
    view.print_summary(get_general_length(records), get_general_fuel_used(records))


def show_summary_period(records):
    left, right = view.enter_period()
    if check_validity_of_date(left) and check_validity_of_date(right):
        records_period = find_by_date_range(records, left, right)
        show_summary(records_period)
    else:
        view.invalid_value()


def show_by_date(records):
    date = view.enter_date()
    if check_validity_of_date(date):
        records_daily = find_by_date(records, date)
        show_all(records_daily)
    else:
        view.invalid_value()


def show_by_period(records):
    left, right = view.enter_period()
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
        else:
            pass


main_func()
