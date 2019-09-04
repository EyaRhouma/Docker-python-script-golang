from sanic import Sanic
from sanic.response import json

app = Sanic()

@app.route("/solve")
async def test(request):
    return json({"hello": "Eya"})

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8000)