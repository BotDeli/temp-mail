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