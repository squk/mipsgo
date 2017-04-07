var conn;
var editor

$(document).ready(function() {
    $("#run").click(function() {
        src = editor.getValue();
        if (!conn) {
            return false;
        }

        conn.send(JSON.stringify({ Source: src, Command: "run"}));
        return false;
    });

    $("#step").click(function() {
        src = editor.getValue();
        if (!conn) {
            return false;
        }
        conn.send(JSON.stringify({ Source: src, Command: "step"}));
        return false;
    });

    $("#break").click(function() {
        var pos = editor.selection.getCursor();
        editor.gotoLine(pos.row + 1)
        editor.insert("break\n");
    });

    var log = document.getElementById("log");

    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/ws");

        conn.onclose = function (evt) {
        };

        conn.onmessage = function (evt) {
            var response = JSON.parse(evt.data);
            var consoleBox = $("#console");
            var currentText = consoleBox.value;
            logLine(response["output"]);

            for (var reg in response["registers"]) {
                $("#" + reg).text(response["registers"][reg]);
            }
        };

    }
    editor = ace.edit("editor");
    editor.setTheme("ace/theme/solarized_dark");
    editor.getSession().setMode("ace/mode/assembly_x86");
    // set some basic starter code
    editor.setValue(" addi $t0, $0, 10\n sll $t0, $t0, 3\n pbin $t0");
});

function logLine(str) {
    var log = $("#console")

    var item = document.createElement("div");
    item.innerText = str;
    log.append(item);
    log.scrollTop($(document).height());
}

