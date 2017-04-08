import view
from model import *


def show_all(records):
    for item in records:
        view.print_record(item)


def show_by_date():
    pass


def show_period():
    pass

def add_record():
    pass



def main_func():
    choice = ''
    records = initialise("fuel_consumption.pickle")
    while choice != "9":
        choice = str(view.menu())
        if choice == "1":
            print("first case")
        elif choice == "2":
            pass
        elif choice == "3":
            pass
        elif choice == "4":
            pass
        elif choice == "5":
            pass
        elif choice == "6":
            pass
        elif choice == "7":
            save_all(records, "fuel_consumption.pickle")
        else:
            pass


main_func()

