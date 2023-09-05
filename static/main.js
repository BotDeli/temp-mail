var socket = new WebSocket("ws://localhost:8080/getMessages")
socket.onopen = function(event) {
    console.log("connection open");
};

socket.onmessage = function(event) {
    let data = event.data.slice(1, -2).split(",");
    if (data[0] === ""){
        return;
    }
    let messages = decodeMessages(data);
    let ids = getHaveId();
    for (let i = 0; i < messages.length; i+=3){
        if (!isIds(ids, messages[i])){
            createNewMessage(messages[i], messages[i+1], messages[i+2], "#")
        } else {
            break
        }
    }
};

function decodeMessages(data) {
    let messages = [];
    let limit = data.length;
    for (let i = 0; i < limit; i+=3) {
        messages.push(data[i].toString().slice(6));
        messages.push(data[i+1].toString().slice(8, -1));
        messages.push(data[i+2].toString().slice(11, -2));
    }
    return messages;
}

windowMessages = document.getElementById("down-window");

function getHaveId(){
    let nodes = windowMessages.children;
    let ids = [];
    for (let i = 1; i < nodes.length; i++){
        ids.push(nodes.item(i).id);
    }
    return ids
}

function isIds(ids, targetId){
    for (let i = 0; i < ids.length; i++){
        if (ids[i].toString() === targetId.toString()){
            return true;
        }
    }
    return false;
}

function createNewMessage(id, sender, theme){
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
    refresh.href = "http://localhost:8080/message?mail="+mailText.innerText+"&id="+id;
    refresh.className = "message-link";
    refresh.appendChild(msg);

    refresh.id = id;

    windowMessages.appendChild(refresh);
}

socket.onclose = function(event ){
    console.log("connection close");
};

window.addEventListener('beforeunload', function(){
    socket.close();
})
