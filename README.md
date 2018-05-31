Clone the repo
Resolve dependency using  go get
Server will run at localhost:3000
Run the main

To find nearby coffee shop
curl "http://localhost:3000/needcoffee?x=47.6&y=-122.4"

To add  a new coffeeshop
curl -X "POST" "http://localhost:3000/addcoffees" \
    -H "Content-Type: application/json; charset=utf-8" \ -d $'[
    {
        "name": "Slight Blue Bottle",
        "x": -123,
        "y": 32.123
    }, {
        "name": "Near Sight Coffee",
        "x": -122.221,
        "y": 32.011
} ]'