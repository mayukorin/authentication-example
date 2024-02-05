function showMessage(elementId, message) {
    const element = document.getElementById(elementId);
    element.innerHTML = message;
}

const hello_with_basic_auth_url = 'http://localhost:1991/basic_auth/hello';
const token_auth_jwt_token_generate_url = 'http://localhost:1991/token_auth/jwt_token';
const basic_auth_button_id = 'basic_authentication_button';
const basic_auth_message_id = 'basic_authentication_message';
const token_auth_jwt_generate_message_id = 'token_authentication_jwt_generate_message';
const token_auth_jwt_generate_form_id = 'token_authentication_jwt_generate_form';
const token_auth_jwt_generate_button_id = 'token_authentication_jwt_generate_button';

const basicAuthenticationButton = document.getElementById(basic_auth_button_id);
basicAuthenticationButton.addEventListener('click', async function(_) {
    const res = await fetch(hello_with_basic_auth_url, { mode: 'cors'});
    if (res.status === 401) {
        const username = prompt('username を入力してください', '');
        const password = prompt('password を入力してください', '');
        if (username !== null && password !== null) {
            const res2 = await fetch(hello_with_basic_auth_url, {headers: {
                'Authorization': 'Basic '+btoa(username+':'+password)
            }, mode: 'cors'});
            if (res2.status !==  200) {
                showMessage(basic_auth_message_id, 'username か password が間違っています');
            } else {
                const message = await res2.json();
                showMessage(basic_auth_message_id, message['message']);
            }
        } else {
            alert('basic auth をキャンセルしました');
        }
    } else {
        const message = await res.json();
        showMessage(basic_auth_message_id, message['message']);
    }
});

const tokenAuthenticationJWTGenerateButton = document.getElementById(token_auth_jwt_generate_button_id);
tokenAuthenticationJWTGenerateButton.addEventListener('click', async function(_) {
    const form = document.getElementById(token_auth_jwt_generate_form_id);
    const formData = new FormData(form);
    const res = await fetch(token_auth_jwt_token_generate_url, { 
        method: 'POST', 
        body: formData,
        mode: 'cors'});

    if (res.status === 200) {
        const message = await res.json();
        const jwtToken = message['jwt_token'];
        localStorage.setItem('jwt_token', jwtToken);
        showMessage(token_auth_jwt_generate_message_id, 'jwt Token を取得しました');
    } else {
        showMessage(token_auth_jwt_generate_message_id, 'jwt Token の取得に失敗しました');
    }


});