import _pickle
from _datetime import datetime


class Record( object ) :
    """Class that represents container for length of the way, date and fuel coefficient"""

    def __init__( self , _date , _length , _coefficient ) :
        """Class ctor"""
        self.date = _date
        self.length = _length
        self.coefficient = _coefficient

    def get( self ) :
        """Returns list of private members"""
        return [ self.date , self.length , self.coefficient ]


def initialise( file_name ) :
    """Returns list of values we've already added"""
    try :
        with open( file_name , 'rb' ) as f :
            records = _pickle.load( f )
            f.close( )
        return records
    except Exception as error :
        return [ ]


def save_all( records , file_name ) :
    """Saves list of all values to file"""
    with open( file_name , 'wb' ) as f :
        _pickle.dump( records , f , _pickle.HIGHEST_PROTOCOL )
    f.close( )


def check_validity_of_date( date ) :
    """Returns 'True' if date is valid or 'False' otherwise"""
    tmp = date.split( '-' )
    if tmp.__len__( ) != 3 or tmp[ 0 ].__len__( ) != 2 or tmp[ 1 ].__len__( ) != 2 or tmp[ 2 ].__len__( ) != 4 :
        return False
    for letter in date:
        tmp= int(letter)
        if tmp !=45 and (tmp<48 or tmp>57):
            return False
    return True


def compare_date( first_date , second_date ) :
    """Returns '0', '1' or '-1' dependent on equality of parameters"""
    first = first_date.split( '-' )
    second = second_date.split( '-' )
    first = datetime( int( first[ 2 ] ) , int( first[ 1 ] ) , int( first[ 0 ] ) )
    second = datetime( int( second[ 2 ] ) , int( second[ 1 ] ) , int( second[ 0 ] ) )
    if first < second :
        return -1
    elif first == second :
        return 0
    return 1


def find_by_date( records , date ) :
    """Returns list of items by date or 'False' otherwise"""
    items = [ ]
    for item in records :
        if item.date == date :
            items.append( item )
    if items == [ ] :
        return False
    return items


def find_by_date_range( records , first_date , second_date ) :
    """Returns list of items chosen by date in date range or 'False' otherwise"""
    items = [ ]
    for item in records :
        if compare_date( item.date , first_date ) == -1 or compare_date( item.date , second_date ) == 1 :
            continue
        else :
            items.append( item )
    if items == [ ] :
        return False
    return items


def get_used_fuel( record ) :
    """Returns value of used fuel by given record"""
    return (record.coefficient * record.length) / 100


def get_general_length( records ) :
    """Returns all length we've passed through"""
    res = 0
    for item in records :
        res += item.length
    return res


def get_general_fuel_used( records ) :
    """Returns all fuel we've spent"""
    res = 0
    for item in records :
        res += get_used_fuel( item )
    return res
