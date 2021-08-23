var btn_sbmt = document.querySelector('#sbmt')
btn_sbmt.addEventListener('click', function (e) {
    e.preventDefault()
    btn_sbmt.setAttribute('disabled', true)

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
                var data = JSON.parse(this.responseText)
                createImgHolder();
                document.querySelector('.img-content img').src = appUrl + data.path
                document.querySelector('#dl-img').href= appUrl + data.path
                btn_sbmt.removeAttribute('disabled')
            }
            else{
                console.log(this.responseText)
                window.alert(this.responseText)
                btn_sbmt.removeAttribute('disabled')
            }
        }
    });
    xhr.open("POST", cfg.linkAPI);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.send(JSON.stringify(payload));

})

function createImgHolder() {
    if (document.querySelector('.img-holder') != undefined) {
        return
    }

    var ct = document.querySelector('#qrcode .row')
    var html = `<div class="col-xs-12 col-sm-12 col-md-12 col-lg-6 col-xl-6 mx-auto img-holder">
    <div class="card">
        <div class="card-header">
        <div class="row">
            <div class="col-sm-6 text-left">
            </div>
        </div>
    </div>
    <div class="card-body img-content"><img class="img-fluid qr-img" src="">
        <div class="dl-btn"><a href="" id="dl-img" target="_blank" download="" class="btn btn-primary">Download</a>
        </div>
    </div>
    </div>
    </div>`
    var dc = new DOMParser().parseFromString(html, "text/html")
    ct.appendChild(dc.querySelector('body div'))
}