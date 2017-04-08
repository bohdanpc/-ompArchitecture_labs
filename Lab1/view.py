
def print_record(item):
    pass


def menu():
    print("---Record fuel consumption---")
    print
    print("1) Show all")
    print("2) Show by date")
    print("3) Show for a certain period")
    print("4) Add a new trip")
    print("5) Change consumption of fuel per 100 km")
    print("6) Show consumption of fuel per 100 km")
    print("7) Exit")
    print


def enter_date():
    return input("Enter the date(dd mm yyyy): ")


def enter_period():
    date_beg = input("Enter beginning date(dd mm yyyy): ")
    date_end = input("Enter ending date(dd mm yyyy): ")
    return date_beg, date_end


def print_record(record):
    for item in record:
        print(item, '\t')
    print('\n')


