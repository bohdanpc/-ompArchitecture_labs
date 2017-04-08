def menu():
    print("\n---Record fuel consumption---")
    print
    print("1) Show the whole table")
    print("2) Show general summary")
    print("3) Show table for a certain period")
    print("4) Show summary for a certain period")
    print("5) Show info by date")
    print("6) Add a new trip")
    print("7) Exit")
    print
    return input("Make your choice: ")


def enter_date():
    return input("Enter the date(dd mm yyyy): ")


def enter_period():
    date_beg = input("Enter beginning date(dd mm yyyy): ")
    date_end = input("Enter ending date(dd mm yyyy): ")
    return date_beg, date_end


def enter_trip_details():
    date = str(input("Enter the date of your trip(dd-mm-yyyy): "))
    length = int(input("Enter the length of your trip: "))
    coefficient = int(input("Enter the fuel consumption: "))
    return [date, length, coefficient]


def print_record(record, used_fuel = -1):
    print(record.date, "\t", record.length, "\t", record.coefficient, "\t", end="")
    if used_fuel != -1:
        print(used_fuel)
    print


def record_names():
    print("date\tlength(km)\tconsumption(100km)\tfuel used\n")


def print_summary(length, fuel_used):
    print("general length: ", length, "\n")
    print("fuel used: ", fuel_used, "\n")