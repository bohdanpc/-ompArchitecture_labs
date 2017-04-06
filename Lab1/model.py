# TODO NESHTA: add general structure and maybe pickle-file reading/writing

import _pickle


class Record( object ) :
    def __init__( self , _data , _length , _coefficient ) :
        self.data = _data
        self.length = _length
        self.coefficient = _coefficient


def initialise( file_name ) :
    try :
        with open( file_name , 'rb' ) as f :
            list = _pickle.load( f )
            f.close( )
        return list
    except Exception as error :
        return [ ]


def save_all( list , file_name ) :
    with open( file_name , 'wb' ) as f :
        _pickle.dump( list , f , _pickle.HIGHEST_PROTOCOL )
    f.close( )
