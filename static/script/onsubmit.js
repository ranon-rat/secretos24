

document.addEventListener('DOMContentLoaded', (event) => {

    document.getElementById("submit").addEventListener("submit", function (e) {
        e.preventDefault() // Cancel the default action
        let formData = new FormData(e.target)
        /**contenido */
        let content = formData.get("content")
        let title = formData.get("title")
        let name = formData.get("name")

        //captcha

        let captchaID = document.getElementById("captcha-id").innerHTML
        let captcha = formData.get("captcha")

        console.log(captcha, captchaID)
        let path = "/thread"+window.location.search
        if (window.location.pathname === "/") {
            path = "/new-thread"
        }
        console.log(path)
        fetch(path, {
            method: "POST",
            body: JSON.stringify({
                content: content,
                title: title,
                name: name,
                captchaID: captchaID,
                captcha: captcha
            })
        }).then(r => r.text()).then(d => console.log(d))
    });
    document.getElementById("get-captcha").onclick= (e) => {
        fetch("/get-captcha")
            .then(r => r.json())
            .then(d => { 
            
    
                document.getElementById("captcha-id").innerHTML=d.id;
                document.getElementById("captcha-image").innerHTML="<img width=200 src='"+d.image+"'>";
                document.getElementById("captcha-image").style.backgroundColor="white"
            })
    }
});
