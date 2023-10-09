
from flask import Flask, jsonify, request
import json

app = Flask(__name__)


@app.route('/')
def hello_world():
    return 'Hello World'


@app.route('/hello')
def products():
    return 'Hello, from the API'

# getting all the shows


@app.route('/shows', methods=['GET'])
def get_shows():
    with open('database.json') as f:
        file_data = json.load(f)
        shows_data = file_data['data']
    return jsonify(shows=shows_data), 200

# adding a new show
# {"name": "NAME OF THE SHOW", "platform": "NAME OF PLATFORM", "rating": "999999"}
@app.route('/addshow', methods=["POST"])
def add_shows():
    new_show = request.json
    # get the currents shows
    with open('database.json') as f:
        file_data = json.load(f)
        current_shows = file_data['data']

    # add the new show from request
    current_shows.append(new_show)

    # write to the file
    with open('database.json', 'w') as f:
        json.dump({'data': current_shows}, f)

    # success
    return "success", 200

# sample data
# {

# "show": "KGF",
# "description" : "nice"

# }
# editing a new show


@app.route('/editshow', methods=["PUT"])
def edit_show():
    show_name = request.json['show']
    description = request.json['description']
    # get the currents shows
    with open('database.json') as f:
        file_data = json.load(f)
        current_shows = file_data['data']

    # update the information about the show
    for show in current_shows:
        if show['name'] == show_name:
            show['description'] = description

    # write to the file
    with open('database.json', 'w') as f:
        json.dump({'data': current_shows}, f)

    # success
    return "success", 200

# Deleting a new show
# sample data
# it should be an dictionary already existing in the database 
# 
# {"name": "Good boys",
# "platform": "Netflix",
# "rating": "7.8",
# "description": "nice"}
# 


@app.route('/deleteshow', methods=["DELETE"])
def delete_show():
    show = request.json
    # get the currents shows
    with open('database.json') as f:
        file_data = json.load(f)
        current_shows = file_data['data']

    # update the information about the show
    current_shows.remove(show)

    # write to the file
    with open('database.json', 'w') as f:
        json.dump({'data': current_shows}, f)

    # success
    return "success", 200


if __name__ == '__main__':
    app.run(debug=True)

# For better naming convention, refer https://restfulapi.net/resource-naming/
