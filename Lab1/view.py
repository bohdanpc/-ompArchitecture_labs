def menu():
    """Shows menu on the screen"""
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
    """Asked the customer to enter the date"""
    return input("Enter the date(dd-mm-yyyy): ")


def enter_period():
    """Asked the customer to enter the period"""
    date_beg = input("Enter beginning date(dd-mm-yyyy): ")
    date_end = input("Enter ending date(dd-mm-yyyy): ")
    return date_beg, date_end


def enter_trip_details():
    """Asked the customer to enter a new trip"""
    date = str(input("Enter the date of your trip(dd-mm-yyyy): "))
    length = int(input("Enter the length of your trip: "))
    coefficient = int(input("Enter the fuel consumption: "))
    return [date, length, coefficient]


def print_record(record, used_fuel = -1):
    """Print record on the screen"""
    print('%10s|%10f|%18f|' % (record.date, record.length, record.coefficient), end = "")
    if used_fuel != -1:
        print('%10f|' % (used_fuel))


def record_names():
    """Print names of table columns"""
    print('\n%10s|%10s|%18s|%10s|' % ("Date", "Length(km)", "Consumption(100km)", "Fuel used"))


def print_summary(length, fuel_used):
    """Print summary info"""
    print("\nGeneral length: ", length)
    print("Fuel used: ", fuel_used)


def invalid_value():
    """Print a message about incorrect value"""
    print("Invalid values entered\n")

