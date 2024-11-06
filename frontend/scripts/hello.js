const submit_button = document.getElementById("submit_button");

function SubmitUserData()
{
    var name = document.getElementById("input_name").value;
    var email = document.getElementById("input_email").value;
    var password = document.getElementById("input_password").value;
    
    var payload = {
        name: name,
        email: email,
        password: password
    };

    fetch("http://localhost:8080/ping",
        {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(payload)
        }).then(console.log(payload))
}

submit_button.addEventListener("click", SubmitUserData);