var conn;
var editor;
var mem_range = 0;
var memory = "";

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

  $("#clear").click(function() {
    $("#console").empty();
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
      if (response["memory"]) {
        memory = response["memory"]
        setMemory();
        console.log("Memory loaded");
      }
      var current_line = response["data"]["current_line"];
      editor.gotoLine(current_line);
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
  log.scrollTop(1E10);
}

function setMemory() {
  lastMem = memory;
  $("#middleE").val(memory);
  refreshMemory(false);
}
