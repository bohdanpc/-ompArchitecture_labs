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
    return input("Enter the date(dd-mm-yyyy): ")


def enter_period():
    date_beg = input("Enter beginning date(dd-mm-yyyy): ")
    date_end = input("Enter ending date(dd-mm-yyyy): ")
    return date_beg, date_end


def enter_trip_details():
    date = str(input("Enter the date of your trip(dd-mm-yyyy): "))
    length = int(input("Enter the length of your trip: "))
    coefficient = int(input("Enter the fuel consumption: "))
    return [date, length, coefficient]


def print_record(record, used_fuel = -1):
    print('%10s|%10d|%18d|' % (record.date, record.length, record.coefficient), end = "")
    if used_fuel != -1:
        print('%10d|' % used_fuel)


def record_names():
    print('\n%10s|%10s|%18s|%10s|' % ("Date", "Length(km)", "Consumption(100km)", "Fuel used"))


def print_summary(length, fuel_used):
    print("\nGeneral length: ", length)
    print("Fuel used: ", fuel_used)
