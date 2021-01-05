let index = {
    init: function() {
        // Init
        asticode.loader.init();
        asticode.modaler.init();
        asticode.notifier.init();

        // Wait for astilectron to be ready
        document.addEventListener('astilectron-ready', function() {
            // Listen
            index.listen();
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
            elem.appendChild(text)
            document.getElementById("history").appendChild(elem)
        });
        // display bookmarks
        data.knjizneoznakedata.map((bk) => {
            const elem = document.createElement('div');
            const text = document.createTextNode(bk.ime+" - "+bk.url+", kategorija: "+bk.kategorija)
            elem.setAttribute('data-id', bk.id)
            elem.appendChild(text)
            document.getElementById("bookmarks").appendChild(elem)
        });
        // display extensions
        data.prosirenjadata.map((ex) => {
            const elem = document.createElement('div');
            const text = document.createTextNode(ex.ime+" - "+ex.opis)
            elem.setAttribute('data-id', ex.id)
            elem.appendChild(text)
            document.getElementById("extensions").appendChild(elem)
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
    }
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