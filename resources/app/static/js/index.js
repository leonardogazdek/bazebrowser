let sel_user = 0;
let index = {
    init: function() {
        // Init
        asticode.loader.init();
        asticode.modaler.init();
        asticode.notifier.init();

        // Wait for astilectron to be ready
        document.addEventListener('astilectron-ready', function() {
            // Listen
        })
    },
    addrInit() {
        // Init
        asticode.loader.init();
        asticode.modaler.init();
        asticode.notifier.init();

        // Wait for astilectron to be ready
        document.addEventListener('astilectron-ready', function() {
            index.getUsers();
            document.body.addEventListener("click", (e) => {
                if(e.target.className == "form-update") {
                    var datas = [].filter.call(e.target.attributes, at => /^data-/.test(at.name));
                    datas = datas.map (el => el.localName.substring(5, el.length))
                    const form = e.target.parentElement.parentElement.querySelector("form")
                    datas.map(data => {
                        form.querySelector(`input[name=${data}]`).value = e.target.getAttribute(`data-${data}`)
                    })
                    
                }
            })
        })
    },
    getUsers: function() {
        let message = {
            "name": "getUsers",
        }
        message.payload="idc"
        astilectron.sendMessage(message,function(response){
            const users = response.payload;
            users.map((usr) => {
                const elem = document.createElement('div');
                const text = document.createTextNode(usr.korisnickoime)
                elem.setAttribute('data-id', usr.id)
                elem.appendChild(text)
                document.getElementById("users").appendChild(elem)

            });
            const elems = document.querySelectorAll('#users div')
            elems.forEach(el => {
                el.addEventListener('click', function(e) {
                    const id = el.getAttribute('data-id')
                    sel_user = id
                    // load user
                    let message = {
                        "name": "fetchUserData",
                    }
                    message.payload=id
                    astilectron.sendMessage(message,function(response){
                        index.displayData(response.payload)
                    })
                }, false)
            })
        });
    },
    displayData(payload) {
        const loaddata = document.querySelectorAll(".loaddata")
        loaddata.forEach(el => {
            el.innerHTML = '';
        })
        // display history
        const data = payload;
        data.povijestdata.map((hist) => {
            const elem = document.createElement('div');
            const text = document.createTextNode(hist.url+" - "+hist.vremenskistambilj)
            elem.setAttribute('data-id', hist.id)
            elem.setAttribute('data-url', hist.url)
            elem.setAttribute('data-vremenskistambilj', hist.vremenskistambilj)
            elem.className="form-update"
            elem.appendChild(text)
            document.getElementById("history").appendChild(elem)
        });
        // display bookmarks
        data.knjizneoznakedata.map((bk) => {
            const elem = document.createElement('div');
            const text = document.createTextNode(bk.ime+" - "+bk.url+", kategorija: "+bk.kategorija)
            elem.setAttribute('data-id', bk.id)
            elem.setAttribute('data-ime', bk.ime)
            elem.setAttribute('data-url', bk.url)
            elem.setAttribute('data-kategorije_id', bk.kategorije_id)
            elem.className="form-update"
            elem.appendChild(text)
            document.getElementById("bookmarks").appendChild(elem)
        });
        // display extensions
        data.prosirenjadata.map((ex) => {
            const elem = document.createElement('div');
            const text = document.createTextNode(ex.ime+" - "+ex.opis)
            elem.setAttribute('data-id', ex.id)
            elem.setAttribute('data-ime', ex.ime)
            elem.setAttribute('data-opis', ex.opis)
            elem.className="form-update"
            elem.appendChild(text)
            document.getElementById("extensions").appendChild(elem)
        });
        data.otvorenekarticedata.map((ex) => {
            const elem = document.createElement('div');
            const text = document.createTextNode(ex.url)
            elem.setAttribute('data-id', ex.id)
            elem.setAttribute('data-url', ex.url)
            elem.className="form-update"
            elem.appendChild(text)
            document.getElementById("opentabs").appendChild(elem)
        });

        data.postavkedata.map((ex) => {
            const elem = document.createElement('div');
            const text = document.createTextNode(ex.ime+" "+ex.vrijednost)
            elem.setAttribute('data-id', ex.id)
            elem.setAttribute('data-ime', ex.ime)
            elem.setAttribute('data-vrijednost', ex.vrijednost)
            elem.className="form-update"
            elem.appendChild(text)
            document.getElementById("settings").appendChild(elem)
        });

    },
    changeUrl: function(e) {
        e.preventDefault();
        let path = document.getElementById("url").value;
        console.log("navigating to " + path);
        // Create message
        let message = {"name": "changeUrl"};
        if (typeof path !== "undefined") {
            message.payload = path
        }

        // Send message
        asticode.loader.show();
        astilectron.sendMessage(message, function(message) {
            // Init
            asticode.loader.hide();

            // Check error
            if (message.name === "error") {
                asticode.notifier.error(message.payload);
                return
            }
        })
        return false
    },
    historyNav: function(attr) {
        // Create message
        let message = {"name": "historyNav"};

        message.payload = (attr == 0) ? "back" : "forward";

        // Send message
        astilectron.sendMessage(message, function(message) {
            // Init

            // Check error
            if (message.name === "error") {
                asticode.notifier.error(message.payload);
                return
            }
        })
    },
    createBookmark(action) {
        // Create message
        let message = {
            name: (action == 'insert') ? "insertBookmark" : "updateBookmark",
            payload: {
                id: parseInt(document.querySelector("#form_bookmarks input[name=id]").value, 10),
                ime: document.querySelector("#form_bookmarks input[name=ime]").value,
                url: document.querySelector("#form_bookmarks input[name=url]").value,
                kategorije_id: parseInt(document.querySelector("#form_bookmarks input[name=kategorije_id]").value, 10),
                korisnici_id: parseInt(sel_user, 10),
            },
        };

        console.log(message)

        // Send message
        astilectron.sendMessage(message, function(message) {
            // Check error
            if (message.name === "error") {
                asticode.notifier.error(message.payload);
                return
            }
            else {
                // FEEDBACK OK
            }
        })
    },
    createExtension(action) {
        // Create message
        let message = {
            name: (action == 'insert') ? "insertExtension" : "updateExtensions",
            payload: {
                id: parseInt(document.querySelector("#form_extensions input[name=id]").value, 10),
                ime: document.querySelector("#form_extensions input[name=ime]").value,
                opis: document.querySelector("#form_extensions input[name=opis]").value,
            },
        };

        console.log(message)

        // Send message
        astilectron.sendMessage(message, function(message) {
            // Check error
            if (message.name === "error") {
                asticode.notifier.error(message.payload);
                return
            }
            else {
                // FEEDBACK OK
            }
        })
    },
    createTab(action) {
        // Create message
        let message = {
            name: (action == 'insert') ? "insertTab" : "updateTab",
            payload: {
                id: parseInt(document.querySelector("#form_opentabs input[name=id]").value, 10),
                url: document.querySelector("#form_opentabs input[name=url]").value,
                korisnici_id: parseInt(sel_user, 10),
            },
        };

        console.log(message)

        // Send message
        astilectron.sendMessage(message, function(message) {
            // Check error
            if (message.name === "error") {
                asticode.notifier.error(message.payload);
                return
            }
            else {
                // FEEDBACK OK
            }
        })
    },
    createSettings(action) {
        // Create message
        let message = {
            name: (action == 'insert') ? "insertSettings" : "updateSettings",
            payload: {
                id: parseInt(document.querySelector("#form_settings input[name=id]").value, 10),
                ime: document.querySelector("#form_settings input[name=ime]").value,
                vrijednost: document.querySelector("#form_settings input[name=vrijednost]").value,
            },
        };

        console.log(message)

        // Send message
        astilectron.sendMessage(message, function(message) {
            // Check error
            if (message.name === "error") {
                asticode.notifier.error(message.payload);
                return
            }
            else {
                // FEEDBACK OK
            }
        })
    },
    /*
    listen: function() {
        astilectron.onMessage(function(message) {
            switch (message.name) {
                case "about":
                    index.about(message.payload);
                    return {payload: "payload"};
                    break;
                case "check.out.menu":
                    asticode.notifier.info(message.payload);
                    break;
            }
        });
    }*/
};