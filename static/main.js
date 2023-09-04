const mailText = document.getElementById("mail-text");
mailText.innerText = 'Загрузка...';

let mail = 'error load mail';

newMail();
function newMail(){
    fetch('/newMail').
    then(response => response.json()).
    then(data => {
        mail = data.mail;
    }).
    catch(err => {
        console.log(err);
    });
    mailText.innerText = mail;
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

windowMessages = document.getElementById("window-messages");

createNewMessage("testSender", "testTheme", "#")
createNewMessage("testSender", "testTheme", "#")
createNewMessage("testSender", "testTheme", "#")
createNewMessage("testSender", "testTheme", "#")
createNewMessage("testSender", "testTheme", "#")
createNewMessage("testSender", "testTheme", "#")
createNewMessage("testSender", "testTheme", "#")
createNewMessage("testSender", "testTheme", "#")
createNewMessage("testSender", "testTheme", "#")
createNewMessage("testSender", "testTheme", "#")
createNewMessage("testSender", "testTheme", "#")
createNewMessage("testSender", "testTheme", "#")
createNewMessage("testSender", "testTheme", "#")
createNewMessage("testSender", "testTheme", "#")
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