import view
import model
import serialize


class Controller:
    @staticmethod
    def show_all(records):
        """show all present records in database

        >>> records = [model.Model.Record("12-12-2012", 253, 172)]
        >>> Controller.show_all(records)
        <BLANKLINE>
              Date|Length(km)|Consumption(100km)| Fuel used|
        12-12-2012|    253.00|            172.00|    435.16|
        """
        view.View.record_names()
        for item in records:
            view.View.print_record(item, model.Model.get_used_fuel(item))

    @staticmethod
    def show_summary(records):
        """show the sum of km's ridden and fuel used

        >>> Controller.show_summary([model.Model.Record("12-12-2012", 253, 172),\
        model.Model.Record("20-11-2015", 128, 123)])
        <BLANKLINE>
        General length: 381
        Fuel used: 592.6
        >>> Controller.show_summary([])
        <BLANKLINE>
        General length: 0
        Fuel used: 0
        """
        view.View.print_summary(model.Model.get_general_length(records),
                           model.Model.get_general_fuel_used(records))

    @staticmethod
    def show_summary_period(records, input_func1=input, input_func2=input):
        """show the sum of km's ridden and fuel used
           in date period (from - to)

        >>> Controller.show_summary_period([model.Model.Record("12-12-2012", 253, 172),\
        model.Model.Record("20-11-2015", 128, 123)], \
        lambda:"12-10-2012", lambda:"25-01-2016")
        <BLANKLINE>
        Enter beginning date(dd-mm-yyyy):Enter ending date(dd-mm-yyyy):
        General length: 381
        Fuel used: 592.6

        >>> Controller.show_summary_period([model.Model.Record("12-12-2012", 253, 172),\
        model.Model.Record("20-11-2015", 128, 123)], lambda:"12-10-2009", lambda:"25-01-2013")
        <BLANKLINE>
        Enter beginning date(dd-mm-yyyy):Enter ending date(dd-mm-yyyy):
        General length: 253
        Fuel used: 435.16

        >>> Controller.show_summary_period([model.Model.Record("12-12-2012", 253, 172),\
        model.Model.Record("20-11-2015", 128, 123)], lambda:"12-25-2013", lambda:"25-01-2014")
        <BLANKLINE>
        Enter beginning date(dd-mm-yyyy):Enter ending date(dd-mm-yyyy):\
    Invalid values entered
        """
        left, right = view.View.enter_period(input_func1, input_func2)
        if model.Model.check_validity_of_date(left) and \
                model.Model.check_validity_of_date(right):
            records_period = model.Model.find_by_date_range(records, left, right)
            Controller.show_summary(records_period)
        else:
            view.View.invalid_value()

    @staticmethod
    def show_by_date(records, input_func=input):
        """show length ridden and fuel used by certain date

        >>> Controller.show_by_date([model.Model.Record("12-12-2012", 253, 172),\
        model.Model.Record("20-11-2015", 128, 123)], lambda:"12-12-2012")
        <BLANKLINE>
        Enter the date(dd-mm-yyyy):
              Date|Length(km)|Consumption(100km)| Fuel used|
        12-12-2012|    253.00|            172.00|    435.16|
        >>> Controller.show_by_date([model.Model.Record("12-12-2012", 253, 172),\
        model.Model.Record("20-11-2015", 128, 123)], lambda:"32-12-2017")
        <BLANKLINE>
        Enter the date(dd-mm-yyyy):Invalid values entered
        """
        date = view.View.enter_date(input_func)

        if model.Model.check_validity_of_date(date):
            records_daily = model.Model.find_by_date(records, date)
            Controller.show_all(records_daily)
        else:
            view.View.invalid_value()

    @staticmethod
    def show_by_period(records, input_func1=input, input_func2=input):
        """show length ridden and fuel used by certain period (from - to)

        >>> Controller.show_by_period([model.Model.Record("12-12-2012", 253, 172),\
        model.Model.Record("20-11-2015", 128, 123)], lambda:"12-10-2012", lambda:"25-01-2016")
        <BLANKLINE>
        Enter beginning date(dd-mm-yyyy):Enter ending date(dd-mm-yyyy):
              Date|Length(km)|Consumption(100km)| Fuel used|
        12-12-2012|    253.00|            172.00|    435.16|
        20-11-2015|    128.00|            123.00|    157.44|

        >>> Controller.show_by_period([model.Model.Record("12-12-2012", 253, 172),\
        model.Model.Record("20-11-2015", 128, 123)], lambda:"12-10-2012", lambda:"25-01-2014")
        <BLANKLINE>
        Enter beginning date(dd-mm-yyyy):Enter ending date(dd-mm-yyyy):
              Date|Length(km)|Consumption(100km)| Fuel used|
        12-12-2012|    253.00|            172.00|    435.16|

        >>> Controller.show_by_period([model.Model.Record("12-12-2012", 253, 172),\
         model.Model.Record("20-11-2015", 128, 123)], \
         lambda:"12-10-2012b", lambda:"42-01-2014")
        <BLANKLINE>
        Enter beginning date(dd-mm-yyyy):Enter ending date(dd-mm-yyyy):\
    Invalid values entered
        """

        left, right = view.View.enter_period(input_func1, input_func2)
        if model.Model.check_validity_of_date(left) and \
                model.Model.check_validity_of_date(right):
            records_period = model.Model.find_by_date_range(records, left, right)
            Controller.show_all(records_period)
        else:
            view.View.invalid_value()

    @staticmethod
    def add_record(records, input_func1=input,
                   input_func2=input, input_func3=input):
        """add new record to the database

        >>> Controller.add_record([model.Model.Record("12-12-2012", 253, 172)],\
         lambda:"12-12-2013", lambda:234, lambda:218)
        <BLANKLINE>
        Enter the date of your trip(dd-mm-yyyy):\
    Enter the length of your trip:Enter the fuel consumption:
        >>> Controller.add_record([model.Model.Record("12-13-2012", 253, 172)],\
        lambda:"31-4-2013", lambda:134, lambda:228)
        <BLANKLINE>
        Enter the date of your trip(dd-mm-yyyy):Enter the length of your trip:\
    Enter the fuel consumption:Invalid values entered
        """
        item_list = view.View.enter_trip_details(input_func1, input_func2, input_func3)
        if model.Model.check_validity(item_list):
            records.append(model.Model.Record(item_list[0], item_list[1], item_list[2]))
        else:
            view.View.invalid_value()

    @staticmethod
    def main_func(input_func1=input):
        """show selection menu and wait for the interaction

        >>> Controller.main_func(lambda:"7")
        <BLANKLINE>
        ---Record fuel consumption---
        1) Show the whole table
        2) Show general summary
        3) Show table for a certain period
        4) Show summary for a certain period
        5) Show info by date
        6) Add a new trip
        7) Exit
        Make your choice:
        """

        choice = ''
        records = serialize.Serialize.load("base")

        while choice != "7":
            choice = str(view.View.menu(input_func1))
            if choice == "1":
                Controller.show_all(records)
            elif choice == "2":
                Controller.show_summary(records)
            elif choice == "3":
                Controller.show_by_period(records)
            elif choice == "4":
                Controller.show_summary_period(records)
            elif choice == "5":
                Controller.show_by_date(records)
            elif choice == "6":
                Controller.add_record(records)
            elif choice == "7":
                serialize.Serialize.save("base", records)


Controller.main_func()
