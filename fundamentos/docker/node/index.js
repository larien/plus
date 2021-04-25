// docker run --rm -it -v $(pwd)/:/usr/src/app -p 3000:3000 node:15 bash
// cd /usr/src/app
// npm init
// npm install express --save
// node index.js

const express = require('express')
const app = express()
const port = 3000

app.get('/', (req, res) => {
    res.send('<h1>Olá, Lauren!</h1>')
})

app.listen(port, () => {
    console.log("Running at" + port)
})