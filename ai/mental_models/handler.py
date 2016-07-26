import json
import psycopg2
import sys

class Strum(object):
    def __init__(self, id):
        self.ID = id

        self.config = json.load(open('config.json', 'r'))
        self.conn = psycopg2.connect(self.config['database'])

        cur = self.conn.cursor()

        cur.execute("SELECT data FROM strum_stats where id = %s", (id,))

        print cur.fetchone()

        self.conn.commit()

        cur.close()


    def close(self):
        self.conn.close()

    def json(self):
        pass

def handler(event, context):
    """
    The handler works to so that it can also run on AWS Lambda.
    """

    strum = Strum(event['ID'])

    strum.close()

    return {
        'Message': "OK"
    }

if __name__ == "__main__":
    print handler({
        'ID': sys.argv[1]
    }, None)
