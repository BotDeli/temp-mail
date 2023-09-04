const mailText = document.getElementById("mail-text");
mailText.innerText = 'Загрузка...';

getMail();
function getMail(){
    fetch("/get_mail").
    then(response => response.json()).
    then(data => {
        mailText.innerText = data.mail;
    }).
    catch(err => {
        console.error(err);
        mailText.innerText = 'error load mail';
    });
}

document.getElementById("copy-button").addEventListener("click", copyMail);
function copyMail(event) {
    event.preventDefault();
    navigator.clipboard.writeText(mailText.innerText).
    then(() => {
        alert("successful copied")
    }).
    catch(() => {
        console.log("failed to copy")
    });
}

var socket = new WebSocket("ws://localhost:8080/getMessages")
socket.onopen = function(event) {
    console.log("connection open");
};
socket.onmessage = function(event) {
    console.log("message:", event.data);
};
socket.onclose = function(event ){
    console.log("connection close");
};

window.addEventListener('beforeunload', function(){
    socket.close();
})

// socket.send(msg);

windowMessages = document.getElementById("window-messages");

createNewMessage("testSender", "testTheme", "#")

function createNewMessage(sender, theme, sendURL){
    let msg = document.createElement("div");
    msg.className = "message";

    let div = document.createElement("div");
    div.className = "messages-content";
    div.innerText = sender;
    msg.appendChild(div);

    div = document.createElement("div");
    div.className = "messages-content";
    div.innerText = theme;
    msg.appendChild(div);

    div = document.createElement("div");
    div.className = "arrow";
    msg.appendChild(div);

    let refresh = document.createElement("a");
    refresh.href = sendURL;
    refresh.className = "message-link";
    refresh.appendChild(msg);

    windowMessages.appendChild(refresh);
}