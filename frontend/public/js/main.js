const basicAuthenticationButton = document.getElementById('basic_authentication_button');

basicAuthenticationButton.addEventListener('click', async function(_) {
    const res = await fetch('http://localhost:1991/hello', {mode: 'cors'});
    const message = await res.json();
    const basicAuthenticationmessage = document.getElementById('basic_authentication_message');
    basicAuthenticationmessage.innerHTML = message['message'];
});