const express = require("express");
const app = express();
const PORT = 3000;

app.use(express.json());
app.use(express.urlencoded({extended: true}))

app.get("/", (req, res) => {
    res.status(200).send("Welcome to the golang tutorial...");
});

app.get("/get", (req, res) => {
    res.status(200).send("Hey there! Have a good day.");
});

app.post("/post", (req, res) => {
    let myJson = req.body;

    res.status(200).send(myJson);
});

app.post("/postform", (req, res) => {
    res.status(200).send(JSON.stringify(req.body));
});

app.listen(PORT, () => {
    console.log(`Nodejs server running on: http://localhost:${PORT}`);
});

/* 
    run it on either postman or thunder client,
    because in browser sometimes it doesn't work
*/

