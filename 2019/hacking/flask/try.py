from flask import Flask
from flask import render_template
app = Flask(__name__)

@app.route("/")
def main(name=None):
    return render_template('index.html', name=name) 

@app.route("/apoorva")
def hello():
    return "<h2>jQuery and AJAX is FUN!</h2>"

app.run(debug=True,threaded=True)
