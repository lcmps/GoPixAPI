var appUrl = window.location.origin;

function isValid() {
    var fields = ['name', 'city', 'pixKey']

    for (let i = 0; i < fields.length; i++) {
        var ele = document.querySelector('#' + fields[i])
        if (ele.value == "") {
            var inv = ele.getAttribute('name')
            return inv
        }
        console.log('ok')
    }
}

var cfg = {
    qrAPI: appUrl + '/qr',
    linkAPI: appUrl + '/link',
    pasteAPI: appUrl + '/paste'
}