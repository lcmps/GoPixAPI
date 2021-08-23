var btn_sbmt = document.querySelector('#sbmt')
btn_sbmt.addEventListener('click', function (e) {
    e.preventDefault()
    var payload = {
        name: document.querySelector('#name').value,
        amount: parseFloat(document.querySelector('#amount').value),
        city: document.querySelector('#city').value,
        description: document.querySelector('#description').value,
        pixKey: document.querySelector('#pixKey').value,
        foregroundColor: '#' + document.querySelector('#foregroundColor').value,
        backgroundColor: '#' + document.querySelector('#backgroundColor').value,
    }

    var xhr = new XMLHttpRequest();
    xhr.withCredentials = true;
    xhr.addEventListener("readystatechange", function () {
        if (this.readyState === this.DONE) {
            if (this.status == 200) {
                // var data = JSON.parse(this.responseText)
                document.querySelector('.copyPasta').value = this.responseText
                btn_sbmt.removeAttribute('disabled')
            }
            else{
                console.log(this.responseText)
                window.alert(this.responseText)
                btn_sbmt.removeAttribute('disabled')
            }
        }
    });
    xhr.open("POST", cfg.pasteAPI);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.send(JSON.stringify(payload));
})

function clipboard(){
    var cptxt = document.querySelector('#copyInput')
    cptxt.select()
    cptxt.setSelectionRange(0, 99999)
    navigator.clipboard.writeText(cptxt.value)
    alert('Paste code copied to clipboard')
}